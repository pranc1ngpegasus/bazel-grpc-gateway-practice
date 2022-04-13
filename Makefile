.PHONY: resolve-dependencies
resolve-dependencies:
	bazel run //:gazelle
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

.PHONY: bazel-build
bazel-build:
	bazel build //cmd/... //adapter/...
