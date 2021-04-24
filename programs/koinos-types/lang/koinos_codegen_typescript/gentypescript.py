
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

class RenderError(Exception):
    pass

def fq_name(name):
    return "::".join(name)

def cpp_namespace(name):
    u = name.split("::")
    if len(u) <= 1:
        return ""
    return "::".join(u[:-1])

def ts_name(name):
    u = name

    # Fix integer naming
    if u.startswith("ui"): u = "u_i" + u[2:]
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

def path_ts_file(name):
    p = name.split("::")
    filename = ts_name(p[-1])
    folders = p[1:-1]
    if len(folders) == 0:
        folders.insert(0, "common")
    for i in range(len(folders)):
        folders[i] = folders[i].replace("_","")
    return "/".join(folders) + "/" + filename

def typeref(tref):
    if tref["info"]["type"] == "IntLiteral":
        return tref["value"]
    if tref["name"][-1] == "vector":
        return "Vector<" + typeref(tref["targs"][0]) +">"
    if tref["name"][-1] == "fixed_blob":
        return "FixedBlob"# + typeref(tref["targs"][0])
    if tref["name"][-1] == "opaque":
        return "Opaque<" + typeref(tref["targs"][0]) +">"
    return ts_name(tref["name"][-1])

def typereflike(tref):
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
    return "Json" + typeref(tref)

def get_dependencies(decl):
    dep = set()
    for field in decl["fields"]:
        dep.add(typereflike(field["tref"]))
        dep.add(typeref(field["tref"]))
    return dep

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


def generate_typescript(schema):
    env = jinja2.Environment(
            loader=jinja2.PackageLoader(__package__, "templates"),
            keep_trailing_newline=True,
        )
    env.filters["fq_name"] = fq_name
    env.filters["tuple"] = tuple
    decls_by_name = collections.OrderedDict(((fq_name(name), decl) for name, decl in schema["decls"]))
    decl_namespaces = sorted(set(cpp_namespace(name) for name in decls_by_name))

    ctx = {"schema" : schema,
           "decls_by_name" : decls_by_name,
           "decl_namespaces" : decl_namespaces,
           "ts_name" : ts_name,
           "decl_fixed_blob" : decl_fixed_blob,
           "get_fixed_blobs" : get_fixed_blobs,
           "decl_opaque" : decl_opaque,
           "get_opaque" : get_opaque,
           "decl_vector": decl_vector,
           "get_vectors": get_vectors,
           "is_struct_impl" : is_struct,
           "get_bad_bytes_impl" : get_bad_bytes,
           "is_empty_struct_impl" : is_empty_struct
          }

    result = collections.OrderedDict()
    result_files = collections.OrderedDict()
    result["files"] = result_files

    template_names = [
        "koinos.ts.j2",
        # "koinos_test.ts.j2"
        ]
    j2_template_struct = env.get_template("koinos-struct.ts.j2")
    j2_template_typedef = env.get_template("koinos-typedef.ts.j2")

    #for template_name in template_names:
    #    j2_template = env.get_template(template_name)
    for name, decl in decls_by_name.items():
        if decl["info"]["type"] == "Struct":
            dependencies = get_dependencies(decl)
            out_filename = path_ts_file(name) + ".ts"
            result_files[out_filename] = j2_template_struct.render({
                "decl": decl,
                "dependencies" : dependencies, 
                "ts_name" : ts_name,
                "typeref": typeref,
                "typereflike": typereflike,
            })
        elif decl["info"]["type"] == "Typedef":
            out_filename = path_ts_file(name) + ".ts"
            result_files[out_filename] = j2_template_typedef.render({
                "decl": decl,
                "ts_name" : ts_name,
                "typeref": typeref,
                "typereflike": typereflike,
            })
            print(out_filename)
        else:
            print(decl["info"]["type"])

    rt_path = os.path.join(os.path.dirname(__file__), "rt")
    for root, dirs, files in os.walk(rt_path):
        for f in files:
            filepath = os.path.join(root, f)
            relpath = "basetypes/" + os.path.relpath(filepath, rt_path)
            with open(filepath, "r") as f:
                content = f.read()
            result_files[relpath] = content
    return result

def escape_json(obj):
   import json
   return json.dumps(obj).replace('"', '\\"')

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

        test_cases.append({"typename": typename, "json": escape_json(test['json'])})

    ctx = {"test_cases" : test_cases,
           "go_name" : go_name
          }

    return env.get_template("main.ts.j2").render(ctx), "test.ts"

def setup(app):
    app.register_target("typescript", generate_typescript)
    # app.register_target("typescript_test", generate_tests)
