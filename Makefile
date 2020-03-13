export PATH := /opt/local/libexec/gnubin:$(PATH)
export ORIHIME_ROOT := $(PWD)

GENERATED_FILES := internal/database/sql.template.go internal/protobuf/orihime.pb.go
COMMON_GO_SOURCE := $(shell find internal -type f -name '*.go') $(GENERATED_FILES)

all: warning build/orihime build/orihime-server build/orihime-test

warning:
	@echo "Make sure to symlink 'ln -s \$${PWD} \$${HOME}/go/src/orihime'"
	@echo "Make sure to 'go get ./cmd ./internal'"

build/orihime-test: $(COMMON_GO_SOURCE) $(shell find cmd/orihime-test -name '*.go')
	go build -o build/orihime-test orihime/cmd/orihime-test

build/orihime-server: $(COMMON_GO_SOURCE) $(shell find cmd/orihime-server -name '*.go')
	go build -o build/orihime-server orihime/cmd/orihime-server

build/orihime: $(COMMON_GO_SOURCE) $(shell find cmd/orihime -name '*.go')
	go build -o build/orihime orihime/cmd/orihime

build:
	mkdir build

internal/database/sql.template.go: $(shell find tools/template/template tools/template/src -type f)
	( cd tools/template ; ./bin/template )

internal/protobuf/orihime.pb.go: tools/protobuf/orihime.proto 
	@echo "We are at $${ORIHIME_ROOT}"
	( cd tools/protobuf ; protoc --go_out=plugins=grpc:${ORIHIME_ROOT}/internal/protobuf orihime.proto )

.PHONY: warning all
