
try:
    import importlib.resources as pkg_resources
except ImportError:
    # Try backported to PY<37 `importlib_resources`.
    import importlib_resources as pkg_resources

#from . import templates  # relative-import the *package* containing the templates

import jinja2

import collections
import os

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

def go_name(name):
    u = name
    # Fix integer naming
    if u.startswith("ui"): u = "u_i" + u[2:]
    if u.endswith("_t"): u = u[-2:]
    # Convert to pascal case
    return u.replace("_", " ").title().replace(" ", "")

def decl_fixed_blob(length):
    fixed_blobs.add(length)
    return ""

def get_fixed_blobs():
    fb_list = list(fixed_blobs)
    fb_list.sort()
    return fb_list

def decl_opaque(o_type):
    opaque.add(o_type)
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

def generate_golang(schema):
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
           "go_name" : go_name,
           "decl_fixed_blob" : decl_fixed_blob,
           "get_fixed_blobs" : get_fixed_blobs,
           "decl_opaque" : decl_opaque,
           "get_opaque" : get_opaque,
           "decl_vector": decl_vector,
           "get_vectors": get_vectors,
           "is_struct_impl" : is_struct
          }
    #for name, val in ctx["decls_by_name"].items():
    #    print(name)
    #    import json
    #    print(json.dumps(val))

    result = collections.OrderedDict()
    result_files = collections.OrderedDict()
    result["files"] = result_files

    template_names = [
        "koinos.go.j2",
        ]

    for template_name in template_names:
        j2_template = env.get_template(template_name)
        out_filename = os.path.splitext(template_name)[0]
        result_files[out_filename] = j2_template.render(ctx)

    rt_path = os.path.join(os.path.dirname(__file__), "rt")
    for root, dirs, files in os.walk(rt_path):
        for f in files:
            filepath = os.path.join(root, f)
            relpath = os.path.relpath(filepath, rt_path)
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

    return env.get_template("test.go.j2").render(ctx), "test.go"

def setup(app):
    app.register_target("golang", generate_golang)
    app.register_target("golang_test", generate_tests)
