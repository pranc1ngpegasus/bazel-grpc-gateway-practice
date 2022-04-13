.PHONY: resolve-dependencies
resolve-dependencies:
	bazel run //:gazelle
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

.PHONY: bazel-build
bazel-build:
	bazel build --sandbox_debug //adapter/... //cmd/... //proto/...
	bazel build --sandbox_debug --define build_platform=darwin_arm64 //:gen-copy-sh
	@./bazel-bin/copy.sh
