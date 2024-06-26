"""A simple bazel rule to build a single go_library from a set of zserio schema files.
"""

load("@rules_go//go:def.bzl", _go_library = "go_library")

def go_zserio_srcs(name, srcs, rootpackage, pkg = None, format = True, sql = False):
    """Generate Go source code for the given zserio files.

    Args:
        name: The name of the target.
        srcs: Zserio source files.
        rootpackage: The rootpackage for the generated zserio code.
        pkg: The package name for generation.
        format: Should we format the source code with a code formatter? Default to True.
        sql: Should we generate SQL marshalling and unmarshalling interfaces? Default to False.
    """
    script = "\n".join([
        "#!/bin/bash",
        "./{bin} generate \\",
        "  --out - \\",
        "  --rootpackage {rootpackage} \\",
        "  --only {pkg} \\",
        "  {noformat} \\",
        "  {sql} \\",
        "  {srcs} 2>/dev/null",
    ]).format(
        bin = "$(execpath //cmd/go-zserio)",
        rootpackage = rootpackage,
        pkg = pkg,
        noformat = "" if format else "--noformat",
        sql = "--sql" if sql else "",
        srcs = "$(SRCS)",
    )

    native.genrule(
        name = name + ".script",
        srcs = srcs,
        outs = [name + ".sh"],
        cmd = "echo -e '{}' >$@".format(script),
        tools = [
            "//cmd/go-zserio",
        ],
        executable = True,
    )
    native.genrule(
        name = name,
        # TODO @aignas 2022-02-08: use something more similar to rules_proto,
        # so that we can encode dependencies between different zs files.
        srcs = srcs,
        outs = [name + ".zs.go"],
        cmd = "{bin} > $@".format(
            bin = "$(execpath {}.script)".format(name),
        ),
        tools = [
            name + ".script",
            "//cmd/go-zserio",
        ],
    )

def go_zserio_library(name, srcs, rootpackage, pkg, sql = True, **kwargs):
    """go_zserio_library generates go source code and a go library.

    Args:
        name: The name of the resultant go library.
        srcs: Zserio source files.
        rootpackage: rootpackage for the zserio bindings.
        pkg: The package name to output.
        sql: Should we generate SQL marshalling and unmarshalling interfaces? Default to False.
        **kwargs: Extra keyword arguments to be passed to the underlying go_library.
    """
    go_zserio_srcs(name = name + "_gen", srcs = srcs, rootpackage = rootpackage, pkg = pkg, sql = sql)
    _go_library(
        name = name,
        srcs = [name + "_gen"],
        importpath = "{}/{}".format(rootpackage, pkg.replace(".", "/")),
        deps = [
            "//:go-zserio",
            "//ztype",
            "@com_github_icza_bitio//:bitio",
        ] + kwargs.pop("deps", []),
        **kwargs
    )
