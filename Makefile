export PATH := /opt/local/libexec/gnubin:$(PATH)
export ORIHIME_ROOT := $(PWD)

build/orihime: internal/database/sql.template.go internal/protobuf/orihime.pb.go
	@echo "Make sure to symlink 'ln -s \$${PWD} \$${HOME}/go/src/orihime'"
	@echo "Make sure to 'go get ./cmd ./internal'"
	go build -o build/orihime orihime/cmd/orihime

build: 
	mkdir build

internal/database/sql.template.go: $(shell find tools/template/template tools/template/src -type f)
	( cd tools/template ; ./bin/template )

internal/protobuf/orihime.pb.go: tools/protobuf/orihime.proto 
	@echo "We are at $${ORIHIME_ROOT}"
	( cd tools/protobuf ; protoc --go_out=plugins=grpc:${ORIHIME_ROOT}/internal/protobuf orihime.proto )
