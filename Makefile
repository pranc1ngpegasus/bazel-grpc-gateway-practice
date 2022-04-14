BAZEL_BUILD_PLATFORM := darwin_arm64

ifdef GITHUB_ACTIONS
	BAZEL_BUILD_PLATFORM = k8
endif

.PHONY: resolve-dependencies
resolve-dependencies: proto-clean
	bazel run //:buildifier
	bazel run //:gazelle
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

.PHONY: proto-clean
proto-clean:
	@find ./proto | grep -E ".*.pb(.gw)?.go" | xargs rm -f

.PHONY: bazel-build
bazel-build: proto-clean
	bazel build --sandbox_debug //adapter/... //cmd/... //proto/...
	bazel build --sandbox_debug --define build_platform=${BAZEL_BUILD_PLATFORM} //:gen-copy-sh
	@./bazel-bin/copy.sh

.PHONY: generate
generate: bazel-build
	go generate ./...

.PHONY: prepare-sql-migrate
prepare-sql-migrate:
	envsubst < ./dbconfig.yaml.tpl > ./dbconfig.yaml

.PHONY: migrate-new
migrate-new: prepare-sql-migrate
	go run github.com/rubenv/sql-migrate/sql-migrate new $(FILENAME) -config=dbconfig.yaml

.PHONY: migrate-up
migrate-up: prepare-sql-migrate
	go run github.com/rubenv/sql-migrate/sql-migrate up -env=${ENV} -config=dbconfig.yaml

.PHONY: migrate-down
migrate-down: prepare-sql-migrate
	go run github.com/rubenv/sql-migrate/sql-migrate down -env=${ENV} -config=dbconfig.yaml
