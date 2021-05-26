
try:
    import importlib.resources as pkg_resources
except ImportError:
    # Try backported to PY<37 `importlib_resources`.
    import importlib_resources as pkg_resources

import jinja2

import collections
import os
import re

fixed_blobs = set()
opaque = set()
vectors = set()
indexes = []

class RenderError(Exception):
    pass

def fq_name(name):
    return "::".join(name)

def cpp_namespace(name):
    u = name.split("::")
    if len(u) <= 1:
        return ""
    return "::".join(u[:-1])

def camel_case(name):
    u = name.replace("_", " ").title().replace(" ", "")
    return u[0].lower() + u[1:]

def ts_name(name):
    if name == 'boolean':
        return 'Bool'
    if name == 'string':
        return 'Str'
    if name == 'number':
        return 'Num'
    if name == 'varint':
        return 'VarInt'

    u = name

    # Fix integer naming
    if u.endswith("_t"): u = u[-2:]

    # Convert to pascal case
    u = u.replace("_", " ").title().replace(" ", "")

    # Replace instances of Id with ID
    m = re.search("^([A-Za-z]*[a-z])?Id([A-Z][A-Za-z]*)?$", u)
    if m is not None:
        u = ""
        if m.group(1) is not None:
            u += m.group(1)
        u += "ID"
        if m.group(2) is not None:
            u += m.group(2)

    # Replace instances of Rpc with RPC
    m = re.search("^([A-Za-z]*[a-z])?Rpc([A-Z][A-Za-z]*)?$", u)
    if m is not None:
        u = ""
        if m.group(1) is not None:
            u += m.group(1)
        u += "RPC"
        if m.group(2) is not None:
            u += m.group(2)

    return u

def index_type(index, parts):
    if len(parts) == 1:
        index.append(parts[0])
        return
    inserted = False
    for i in range(len(index)):
        if type(index[i]) == str:
            continue
        if index[i]["folder"] == parts[0]:
            index_type(index[i]["files"], parts[1:])
            inserted = True
    if not inserted:
        index.append({
            "folder": parts[0],
            "files": []
        })
        index_type(index[-1]["files"], parts[1:])

def path(decl, relativeTo = "."):
    relative = relativeTo.split("/")
    i = 0
    for i in range(len(relative)):
        if relative[i] == decl["name"][i+1]:
            i = i + 1
        else:
            break; 
    r = relativeTo.count("/") - i
    
    filename = ts_name(decl["name"][-1])
    folders = decl["name"][1+i:-1]
    if len(folders) == 0 and i == 0:
        if is_basetype(decl["name"][-1]):
            folders.insert(0, "basetypes")
        else:
            folders.insert(0, "common")
    for i in range(len(folders)):
        folders[i] = folders[i].replace("_","")
    pathFromRoot = "/".join(folders)
    if len(folders) == 0:
        pathFromRoot = "."
    
    return ("../" * r) + pathFromRoot + "/" + filename

def path_ts_file(name):
    p = name.split("::")
    filename = ts_name(p[-1])
    folders = p[1:-1]
    if len(folders) == 0:
        if is_basetype(p[-1]):
            folders.insert(0, "basetypes")
        else:
            folders.insert(0, "common")
    for i in range(len(folders)):
        folders[i] = folders[i].replace("_","")
    return "/".join(folders) + "/" + filename

def typeref(tref):
    if tref["info"]["type"] == "IntLiteral":
        return tref["value"]
    if tref["name"][-1] == "vector":
        return "Vector<" + typeref(tref["targs"][0]) +">"
    if tref["name"][-1] == "opaque":
        return "Opaque<" + typeref(tref["targs"][0]) +">"
    if tref["name"][-1] == "optional":
        return "Optional<" + typeref(tref["targs"][0]) +">"
    return ts_name(tref["name"][-1])

def typeref_like(tref):
    if tref["name"][-1] in ['int8', 'uint8', 'int16',
      'uint16', 'int32', 'uint32', 'int64', 'uint64',
      'int128', 'uint128', 'int160', 'uint160',
      'int256', 'uint256', 'block_height_type',
      'timestamp_type']:
        return "NumberLike"
    if tref["name"][-1] == "string":
        return "StringLike"
    if tref["name"][-1] == "boolean":
        return "BooleanLike"
    if tref["name"][-1] in ['variable_blob', 'fixed_blob']:
        return "VariableBlobLike"
    if tref["name"][-1] == "vector":
        return typeref_like(tref["targs"][0]) + "[]"
    if tref["name"][-1] == "opaque":
        return typeref_like(tref["targs"][0])
    if tref["name"][-1] == "optional":
        return typeref_like(tref["targs"][0])
    return ts_name(tref["name"][-1]) + "Like"

def typeref_constructor(tref):
    if tref["name"][-1] == "opaque":
        return "Opaque(" + typeref(tref["targs"][0]) +", "
    if tref["name"][-1] == "optional":
        return "Optional(" + typeref(tref["targs"][0]) +", "
    if tref["name"][-1] == "vector":
        return "Vector(" + typeref(tref["targs"][0]) +", "
    return typeref(tref) + "("

def deserialize_type(tref):
    if tref["name"][-1] == "opaque":
        return "deserializeOpaque<" + typeref(tref["targs"][0]) + ">(" + typeref(tref["targs"][0]) + ").getNative()"
    if tref["name"][-1] == "optional":
        return "deserializeOptional<" + typeref(tref["targs"][0]) + ">(" + typeref(tref["targs"][0]) + ").value"
    if tref["name"][-1] == "vector":
        return "deserializeVector<" + typeref(tref["targs"][0]) + ">(" + typeref(tref["targs"][0]) + ").items"
    return "deserialize(" + typeref(tref) + ")"

def is_basetype(name):
    if name in ['vector', 'variant', 'string', 'boolean',
    'int8', 'uint8', 'int16', 'uint16', 'int32', 'uint32',
    'int64', 'uint64', 'int128', 'uint128', 'int160', 'uint160',
    'int256', 'uint256', 'multihash', 'variable_blob',
    'fixed_blob', 'timestamp_type','block_height_type', 'opaque',
    'optional', 'number', 'varint']:
      return True
    return False

def insert_type_like_dependency(deps, className, tref, nameRef):
    if tref["name"][-1] in ['int8', 'uint8', 'int16',
      'uint16', 'int32', 'uint32', 'int64', 'uint64',
      'int128', 'uint128', 'int160', 'uint160',
      'int256', 'uint256', 'block_height_type',
      'timestamp_type']:
        deps = insert_dependency(deps, "NumberLike", { "name": ["koinos", "number"]}, nameRef, False)
    elif tref["name"][-1] == "string":
        deps = insert_dependency(deps, "StringLike", { "name": ["koinos", "string"]}, nameRef, False)
    elif tref["name"][-1] == "boolean":
        deps = insert_dependency(deps, "BooleanLike", { "name": ["koinos", "boolean"]}, nameRef, False)
    elif tref["name"][-1] in ['variable_blob', 'fixed_blob']:
        deps = insert_dependency(deps, "VariableBlobLike", { "name": ["koinos", "variable_blob"]}, nameRef, False)
    elif tref["name"][-1] == "opaque":
        className = ts_name(tref["targs"][0]["name"][-1])
        deps = insert_type_like_dependency(deps, className, tref["targs"][0], nameRef)
    elif tref["name"][-1] == "optional":
        className = ts_name(tref["targs"][0]["name"][-1])
        deps = insert_type_like_dependency(deps, className, tref["targs"][0], nameRef)
    elif tref["name"][-1] == "vector":
        className = ts_name(tref["targs"][0]["name"][-1])
        deps = insert_type_like_dependency(deps, className, tref["targs"][0], nameRef)
    else:
        deps = insert_dependency(deps, className + "Like", tref, nameRef, False)

    return deps

def insert_dependency(deps, className, tref, nameRef, insertJsonlike = True):
    pathFile = path(tref, nameRef)
    newPathFile = True
    for i in range(len(deps)):
        if pathFile == deps[i][1]:
            deps[i][0].add(className)
            newPathFile = False
            break
    if newPathFile:
        d = set()
        d.add(className)
        deps.append([d, pathFile])
    if tref["name"][-1] in ['vector', 'opaque', 'optional']:
        className = ts_name(tref["targs"][0]["name"][-1])
        deps = insert_dependency(deps, className, tref["targs"][0], nameRef)
    
    # typeref_like
    if insertJsonlike:
        deps = insert_type_like_dependency(deps, className, tref, nameRef)
    return deps

def get_dependencies_struct(decl, nameRef):
    dep = []
    dep = insert_dependency(dep, "VariableBlob", { "name": ["koinos", "variable_blob"]}, nameRef, False)
    for field in decl["fields"]:
        className = ts_name(field["tref"]["name"][-1])
        dep = insert_dependency(dep, className, field["tref"], nameRef)
    return dep

def get_dependencies_parse(decl, nameRef):
    dep = []
    

def get_dependencies_variant(decl, nameRef):
    dep = []
    dep = insert_dependency(dep, "VariableBlob", { "name": ["koinos", "variable_blob"]}, nameRef, False)
    dep = insert_dependency(dep, "VarInt", { "name": ["koinos", "varint"]}, nameRef, False)
    for arg in decl["tref"]["targs"]:
        className = ts_name(arg["name"][-1])
        dep = insert_dependency(dep, className, arg, nameRef)
    return dep

def get_dependencies_typedef(decl, nameRef):
    className = ts_name(decl["tref"]["name"][-1])
    dep = []
    dep = insert_dependency(dep, "VariableBlob", { "name": ["koinos", "variable_blob"]}, nameRef, False)
    dep = insert_dependency(dep, className, decl["tref"], nameRef)
    return dep

def get_dependencies_enum(decl, nameRef):
    return get_dependencies_typedef(decl, nameRef)


def decl_fixed_blob(length):
    fixed_blobs.add(length)
    return ""

def get_fixed_blobs():
    fb_list = list(fixed_blobs)
    fb_list.sort()
    return fb_list

def decl_opaque(o_type, typename):
    opaque.add((o_type, typename))
    return ""

def get_opaque():
    o_list = list(opaque)
    o_list.sort()
    return o_list

def decl_vector(v_type):
    vectors.add(v_type)
    return ""

def get_vectors():
    v_list = list(vectors)
    v_list.sort()
    return v_list

def is_struct(targ, decls_by_name):
    ns = "::"
    type_name = ns.join(targ["name"])

    if targ["targs"] != None:
        type_name = ns.join(targ["targs"][0]["name"])

    type_info = decls_by_name[type_name]["info"]["type"]

    if type_info == "Struct":
        return True;
    elif type_info == "BaseType":
        return False
    else:
        return is_struct(decls_by_name[type_name]["tref"], decls_by_name)

def is_empty_struct(targ, decls_by_name):
    ns = "::"

    if targ["info"]["type"] == "Field":
        type_name = ns.join(targ["name"])
        type_name = ns.join(targ["tref"]["name"])

        field_type = decls_by_name[type_name]
        type_info = field_type["info"]["type"]
    else:
        field_type = targ
        type_info = field_type["info"]["type"]

    if type_info == "Struct":
        if len(field_type["fields"]) == 0:
            return True

        for field in field_type["fields"]:
            if not is_empty_struct(field, decls_by_name):
                return False

        return True
    elif type_info == "BaseType":
        return False
    else:
        type_name = ""
        if "tref" in targ and targ["tref"] != None:
            type_name = ns.join(targ["tref"]["name"])
        else:
            type_name = ns.join(targ["name"])
        return is_empty_struct(decls_by_name[type_name], decls_by_name)

def get_bad_bytes(targ, decls_by_name, field=None):
    ns = "::"
    type_info = targ["info"]["type"]

    if type_info == "Struct":
        b = ""
        if field == None:
            if len(targ["fields"]) > 0:
                field_type_name = ns.join(targ["fields"][0]["tref"]["name"])
                return get_bad_bytes(decls_by_name[field_type_name], decls_by_name)
            else:
                return ""

        for f in targ["fields"]:
            field_type_name = ns.join(f["tref"]["name"])

            if f["name"] == field["name"]:
                b += get_bad_bytes(decls_by_name[field_type_name], decls_by_name)
                b = b[:-2]
                return b

            b += get_good_bytes(decls_by_name[field_type_name], decls_by_name) + ", "

        b = b[:-2]
        return b
    elif type_info == "BaseType":
        type_name = ns.join(targ["name"])
        return ""
    else:
        type_name = ""
        if "tref" in targ and targ["tref"] != None:
            type_name = ns.join(targ["tref"]["name"])
        else:
            type_name = ns.join(targ["name"])
        return get_bad_bytes(decls_by_name[type_name], decls_by_name)


good_bytes = {
    "std::vector" : "0x00",
    "std::variant" : "0x00",
    "std::string" : "0x00",
    "koinos::boolean" : "0x00",
    "koinos::int8" : "0x00",
    "koinos::uint8" : "0x00",
    "koinos::int16" : "0x00, 0x00",
    "koinos::uint16" : "0x00, 0x00",
    "koinos::int32" : "0x00, 0x00, 0x00, 0x00",
    "koinos::uint32" : "0x00, 0x00, 0x00, 0x00",
    "koinos::int64" : "0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00",
    "koinos::uint64" : "0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00",
    "koinos::int128" : "0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00",
    "koinos::uint128" : "0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00",
    "koinos::int160" : "0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00",
    "koinos::uint160" : "0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00",
    "koinos::int256" : "0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00",
    "koinos::uint256" : "0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00",
    "koinos::multihash" : "0x00, 0x00",
    "koinos::multihash_vector" : "0x00, 0x00, 0x00",
    "koinos::variable_blob" : "0x00",
    "koinos::fixed_blob" : "",
    "koinos::timestamp_type" : "0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00",
    "koinos::block_height_type" : "0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00",
    "koinos::opaque" : "0x00",
}

def get_good_bytes(targ, decls_by_name):
    ns = "::"
    type_name = ns.join(targ["name"])

    if "targs" in targ and targ["targs"] != None:
        type_name = ns.join(targ["targs"][0]["name"])
    elif "tref" in targ and targ["tref"] != None:
        type_name = ns.join(targ["tref"]["name"])

    type_info = targ["info"]["type"]

    if type_name == "koinos::fixed_blob":
        n = targ["tref"]["targs"][0]["value"]
        return (n - 1) * (good_bytes["koinos::int8"] + ", ") + good_bytes["koinos::int8"]
    elif type_info == "Struct":
        b = ""
        for f in targ["fields"]:
            field_type_name = ns.join(f["tref"]["name"])
            b += get_good_bytes(decls_by_name[field_type_name], decls_by_name) + ", "
        b = b[:-2]
        return b
    elif type_info == "BaseType":
        return good_bytes[type_name]
    else:
        return get_good_bytes(decls_by_name[type_name], decls_by_name)


def generate_index_template(index, prefix, result_files, template):
    items = []
    for i in range(len(index)):
        if type(index[i]) == str:
            items.append(index[i])
            continue
        folder = index[i]["folder"] + "/"
        items.append(index[i]["folder"])
        generate_index_template(index[i]["files"], prefix + folder, result_files, template)
    result_files["src/" + prefix + "index.ts"] = template.render({"items": items})

def generate_basic_tests(types_typedef, result_files, template):
    result_files["tests/basic-tests.spec.ts"] = template.render({
      "types_typedef": types_typedef
    })

def generate_typescript(schema):
    env = jinja2.Environment(
            loader=jinja2.PackageLoader(__package__, "templates"),
            keep_trailing_newline=True,
        )
    env.filters["fq_name"] = fq_name
    env.filters["tuple"] = tuple
    decls_by_name = collections.OrderedDict(((fq_name(name), decl) for name, decl in schema["decls"]))
    decl_namespaces = sorted(set(cpp_namespace(name) for name in decls_by_name))

    result = collections.OrderedDict()
    result_files = collections.OrderedDict()
    result["files"] = result_files

    j2_template_struct = env.get_template("koinos-struct.ts.j2")
    j2_template_variant = env.get_template("koinos-variant.ts.j2")
    j2_template_typedef = env.get_template("koinos-typedef.ts.j2")
    j2_template_enum = env.get_template("koinos-enum.ts.j2")
    j2_template_index = env.get_template("koinos-index.ts.j2")
    j2_template_parser = env.get_template("koinos-parser.ts.j2")
    j2_template_basic_tests = env.get_template("koinos-basic-tests.ts.j2")

    types_typedef = set()
    dep_parser = []
    decls_parser = []
    parser_file = "parser.ts"

    for name, decl in decls_by_name.items():
        ctx = {
            "decl": decl,
            "dependencies" : [], 
            "ts_name" : ts_name,
            "camel_case" : camel_case,
            "typeref": typeref,
            "typeref_constructor": typeref_constructor,
            "typeref_like": typeref_like,
            "deserialize_type": deserialize_type,
            "len": len,
            "str": str,
        }

        out_filename = path_ts_file(name) + ".ts"
        index_type(indexes, path_ts_file(name).split("/"))
        
        if decl["info"]["type"] == "Struct":
            class_name = ts_name(decl["name"])
            ctx["class_name"] = class_name
            ctx["dependencies"] = get_dependencies_struct(decl, out_filename)
            result_files["src/" + out_filename] = j2_template_struct.render(ctx)
            types_typedef.add(class_name)
            dep_parser = insert_dependency(dep_parser, class_name, { "name": name.split("::")}, parser_file)
            decls_parser.append({ "name": name.split("::")})
        elif decl["info"]["type"] == "Typedef" and decl["tref"]["name"][-1] == "variant":
            class_name = ts_name(decl["name"])
            ctx["class_name"] = class_name
            ctx["dependencies"] = get_dependencies_variant(decl, out_filename)
            result_files["src/" + out_filename] = j2_template_variant.render(ctx)
            types_typedef.add(class_name)
            dep_parser = insert_dependency(dep_parser, class_name, { "name": name.split("::")}, parser_file)
            decls_parser.append({ "name": name.split("::")})
        elif decl["info"]["type"] == "Typedef":
            class_name = ts_name(decl["name"])
            ctx["class_name"] = class_name
            ctx["ref_name"] = typeref(decl["tref"])
            ctx["dependencies"] = get_dependencies_typedef(decl, out_filename)
            result_files["src/" + out_filename] = j2_template_typedef.render(ctx)
            types_typedef.add(class_name)
            dep_parser = insert_dependency(dep_parser, class_name, { "name": name.split("::")}, parser_file)
            decls_parser.append({ "name": name.split("::")})
        elif decl["info"]["type"] == "EnumClass":
            class_name = ts_name(decl["name"])
            ctx["class_name"] = class_name
            ctx["ref_name"] = typeref(decl["tref"])
            ctx["dependencies"] = get_dependencies_enum(decl, out_filename)
            result_files["src/" + out_filename] = j2_template_enum.render(ctx)
            types_typedef.add(class_name)
            dep_parser = insert_dependency(dep_parser, class_name, { "name": name.split("::")}, parser_file)
            decls_parser.append({ "name": name.split("::")})
        else:
            if decl["name"][-1] in ['vector', 'opaque', 'optional', 'variant']:
                continue
            class_name = ts_name(decl["name"][-1])
            dep_parser = insert_dependency(dep_parser, class_name, { "name": name.split("::")}, parser_file)
            decls_parser.append({ "name": name.split("::")})
        print(decl["name"])

    index_type(indexes, ["parser"])
    generate_index_template(indexes, "", result_files, j2_template_index)
    generate_basic_tests(types_typedef, result_files, j2_template_basic_tests)

    ctx = {
        "decls": decls_parser,
        "dependencies" : dep_parser, 
        "ts_name" : ts_name,
        "typeref_like": typeref_like,
    }
    result_files["src/" + parser_file] = j2_template_parser.render(ctx)

    rt_path = os.path.join(os.path.dirname(__file__), "rt")
    for root, dirs, files in os.walk(rt_path):
        for f in files:
            filepath = os.path.join(root, f)
            relpath = "basetypes/" + os.path.relpath(filepath, rt_path)
            with open(filepath, "r") as f:
                content = f.read()
            result_files["src/" + relpath] = content
    return result

def generate_tests(test_data):
    import json
    env = jinja2.Environment(
        loader=jinja2.PackageLoader(__package__, "test"),
        keep_trailing_newline=True
    )
    test_cases = []

    for test in test_data:
        split_ns = test['type'].rsplit('::', 1)
        namespace, typename = '', ''
        if len(split_ns) == 2:
            namespace = split_ns[0]
            typename = split_ns[1]
        else:
            typename = split_ns[0]

        test_cases.append({"typename": typename, "json": json.dumps(test['json'])})

    ctx = {"test_cases" : test_cases,
           "ts_name" : ts_name
          }

    return env.get_template("koinos-test.ts.j2").render(ctx), "tests/test.ts"

def setup(app):
    app.register_target("typescript", generate_typescript)
    app.register_target("typescript_test", generate_tests)
