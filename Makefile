# Docker
BASE_DOCKER_COMPOSE = ./docker/docker-compose.yml
COMPOSE_OPTS = -f "$(BASE_DOCKER_COMPOSE)" -p bazel-grpc-gateway-practice

.PHONY: build
build:
	docker-compose $(COMPOSE_OPTS) up -d --build

.PHONY: setup
setup: build

.PHONY: up
up:
	docker-compose $(COMPOSE_OPTS) up -d

.PHONY: down
down:
	docker-compose $(COMPOSE_OPTS) down

.PHONY: start
start:
	docker-compose $(COMPOSE_OPTS) start

.PHONY: restart
restart:
	docker-compose $(COMPOSE_OPTS) restart

.PHONY: stop
stop:
	docker-compose $(COMPOSE_OPTS) stop

.PHONY: logs
logs:
	docker-compose $(COMPOSE_OPTS) logs -f

# Bazel
BAZEL_BUILD_PLATFORM := darwin_arm64

ifdef GITHUB_ACTIONS
	BAZEL_BUILD_PLATFORM = k8
endif

.PHONY: resolve-dependencies
resolve-dependencies: proto-clean
	bazelisk run //:buildifier
	bazelisk run //:gazelle
	bazelisk run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

.PHONY: proto-clean
proto-clean:
	@find ./proto | grep -E ".*.pb(.gw)?.go" | xargs rm -f

.PHONY: bazel-build
bazel-build: proto-clean
	bazelisk build --sandbox_debug //adapter/... //cmd/... //proto/...
	bazelisk build --sandbox_debug --define build_platform=${BAZEL_BUILD_PLATFORM} //:gen-copy-sh
	@./bazel-bin/copy.sh

.PHONY: generate
generate: bazel-build
	go generate ./...

.PHONY: ent-new
ent-new:
	go run entgo.io/ent/cmd/ent init ${PKGNAME}
