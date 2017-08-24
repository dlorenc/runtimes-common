git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.5.5",
)

load(
    "@io_bazel_rules_go//go:def.bzl",
    "new_go_repository",
    "go_repositories",
)

go_repositories()

load("@io_bazel_rules_go//proto:go_proto_library.bzl", "go_proto_repositories")

go_proto_repositories()

new_go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    remote = "https://github.com/go-yaml/yaml",
    tag = "v2",
    vcs = "git",
)

new_go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    tag = "master",
)

git_repository(
    name = "io_bazel_rules_docker",
    commit = "a114402e4b98aee45ca67917bf7fb728913cf7d3",
    remote = "https://github.com/bazelbuild/rules_docker.git",
)

# For node-build
new_http_archive(
    name = "httplib2",
    build_file_content = """
py_library(
   name = "httplib2",
   srcs = glob(["**/*.py"]),
   data = ["cacerts.txt"],
   visibility = ["//visibility:public"]
)""",
    sha256 = "d1bee28a68cc665c451c83d315e3afdbeb5391f08971dcc91e060d5ba16986f1",
    strip_prefix = "httplib2-0.10.3/python2/httplib2/",
    type = "tar.gz",
    url = "https://codeload.github.com/httplib2/httplib2/tar.gz/v0.10.3",
)

# Used by oauth2client
new_http_archive(
    name = "six",
    build_file_content = """
# Rename six.py to __init__.py
genrule(
    name = "rename",
    srcs = ["six.py"],
    outs = ["__init__.py"],
    cmd = "cat $< >$@",
)
py_library(
   name = "six",
   srcs = [":__init__.py"],
   visibility = ["//visibility:public"],
)""",
    sha256 = "e24052411fc4fbd1f672635537c3fc2330d9481b18c0317695b46259512c91d5",
    strip_prefix = "six-1.9.0/",
    type = "tar.gz",
    url = "https://pypi.python.org/packages/source/s/six/six-1.9.0.tar.gz",
)

# Used for authentication in containerregistry
new_http_archive(
    name = "oauth2client",
    build_file_content = """
py_library(
   name = "oauth2client",
   srcs = glob(["**/*.py"]),
   visibility = ["//visibility:public"],
   deps = [
     "@httplib2//:httplib2",
     "@six//:six",
   ]
)""",
    sha256 = "7230f52f7f1d4566a3f9c3aeb5ffe2ed80302843ce5605853bee1f08098ede46",
    strip_prefix = "oauth2client-4.0.0/oauth2client/",
    type = "tar.gz",
    url = "https://codeload.github.com/google/oauth2client/tar.gz/v4.0.0",
)

# Used for parallel execution in containerregistry
new_http_archive(
    name = "concurrent",
    build_file_content = """
py_library(
   name = "concurrent",
   srcs = glob(["**/*.py"]),
   visibility = ["//visibility:public"]
)""",
    sha256 = "a7086ddf3c36203da7816f7e903ce43d042831f41a9705bc6b4206c574fcb765",
    strip_prefix = "pythonfutures-3.0.5/concurrent/",
    type = "tar.gz",
    url = "https://codeload.github.com/agronholm/pythonfutures/tar.gz/3.0.5",
)

# For packaging python tools.
git_repository(
    name = "subpar",
    commit = "7e12cc130eb8f09c8cb02c3585a91a4043753c56",
    remote = "https://github.com/google/subpar",
)

git_repository(
    name = "containerregistry",
    commit = "b0278a1544238d03648861b6d9395414d4c958e5",
    remote = "https://github.com/google/containerregistry",
)

load(
    "@io_bazel_rules_docker//docker:docker.bzl",
    "docker_repositories",
    "docker_pull",
)

docker_repositories()

new_http_archive(
    name = "mock",
    build_file_content = """
# Rename mock.py to __init__.py
genrule(
    name = "rename",
    srcs = ["mock.py"],
    outs = ["__init__.py"],
    cmd = "cat $< >$@",
)
py_library(
   name = "mock",
   srcs = [":__init__.py"],
   visibility = ["//visibility:public"],
)""",
    sha256 = "b839dd2d9c117c701430c149956918a423a9863b48b09c90e30a6013e7d2f44f",
    strip_prefix = "mock-1.0.1/",
    type = "tar.gz",
    url = "https://pypi.python.org/packages/source/m/mock/mock-1.0.1.tar.gz",
)

docker_pull(
    name = "python_base",
    registry = "gcr.io",
    repository = "google-appengine/python",
    digest = "sha256:163a514abdb54f99ba371125e884c612e30d6944628dd6c73b0feca7d31d2fb3",
)

http_file(
    name = "docker_credential_gcr",
    url ="https://github.com/GoogleCloudPlatform/docker-credential-gcr/releases/download/v1.4.1/docker-credential-gcr_linux_amd64-1.4.1.tar.gz",
    sha256 = "c4f51ff78c25e2bfef38af0f38c6966806e25da7c5e43092c53a4d467fea4743",
)

docker_pull(
    name = "py_base",
    digest = "sha256:beab5f862b7b41e38f89ba4788b34943577eea11b149d3324db57fe3cedc109c",
    registry = "gcr.io",
    repository = "google-appengine/nodejs",
)
