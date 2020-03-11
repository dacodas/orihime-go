export PATH := /opt/local/libexec/gnubin:$(PATH)
export ORIHIME_ROOT := $(PWD)

build: internal/database/sql.go internal/protobuf/orihime.pb.go
	@echo "Make sure to symlink 'ln -s \$${PWD} \$${HOME}/go/src/orihime'"
	@echo "Make sure to 'go get ./cmd ./internal'"
	go build -o orihime orihime/cmd/orihime

internal/database/sql.go: $(shell find tools/template/template tools/template/src -type f)
	( cd tools/template ; ./bin/template )

internal/protobuf/orihime.pb.go: tools/protobuf/orihime.proto 
	@echo "We are at $${ORIHIME_ROOT}"
	( cd tools/protobuf ; protoc --go_out=plugins=grpc:${ORIHIME_ROOT}/internal/protobuf orihime.proto )

.PHONY: build

