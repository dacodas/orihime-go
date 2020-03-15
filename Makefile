export PATH := /opt/local/libexec/gnubin:$(PATH)
export ORIHIME_ROOT := $(PWD)

TEMPLATED_GO_SERVER_SOURCE := internal/database/sql-insert.templated.go internal/database/sql-query.templated.go
PROTO_FILES := internal/protobuf/orihime.pb.go internal/protobuf/orihime.pb.json.go
COMMON_GO_SOURCE := $(PROTO_FILES)

all: warning build/orihime build/orihime-server build/orihime-test

graph-make:
	make -Bnd build/orihime | make2graph | dot -Tsvg -o /tmp/orihime.svg
	make -Bnd build/orihime-server| make2graph | dot -Tsvg -o /tmp/orihime-server.svg
	make -Bnd build/orihime-test | make2graph | dot -Tsvg -o /tmp/orihime-test.svg

warning:
	@echo "Make sure to symlink 'ln -s \$${PWD} \$${HOME}/go/src/orihime'"
	@echo "Make sure to 'go get ./cmd ./internal'"

build/orihime-test: $(COMMON_GO_SOURCE) $(TEMPLATED_GO_SERVER_SOURCE) $(shell find cmd/orihime-test -name '*.go')
	go build -o build/orihime-test orihime/cmd/orihime-test

build/orihime-server: $(COMMON_GO_SOURCE) $(TEMPLATED_GO_SERVER_SOURCE) $(shell find cmd/orihime-server internal/server -name '*.go')
	go build -o build/orihime-server orihime/cmd/orihime-server

build/orihime: $(COMMON_GO_SOURCE) $(shell find cmd/orihime internal/client -name '*.go')
	go build -o build/orihime orihime/cmd/orihime

build:
	mkdir build

internal/protobuf/orihime.pb.go internal/protobuf/orihime.pb.json.go: tools/protobuf/orihime.proto
	( cd tools/protobuf ; protoc --go_out=plugins=grpc:${ORIHIME_ROOT}/internal/protobuf --go-json_out=${ORIHIME_ROOT}/internal/protobuf orihime.proto )

internal/database:
	mkdir -p internal/database

internal/database/sql-insert.templated.go internal/database/sql-query.templated.go: internal/database
	( cd tools/template ; ./bin/template )

internal/database/sql-insert.templated.go: \
	tools/template/src/sql-insert.expanded \
	tools/template/src/sql-insert-preamble \
	tools/template/src/sql-insert-skeleton

internal/database/sql-query.templated.go: \
	tools/template/src/sql-query.expanded \
	tools/template/src/sql-query-preamble \
	tools/template/src/sql-query-skeleton

.PHONY: warning all graph-make
