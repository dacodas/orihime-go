export PATH := /opt/local/libexec/gnubin:$(PATH)

build: internal/database/sql.go internal/protobuf/orihime.pb.go
	@echo "Make sure to symlink 'ln -s \$${PWD} \$${HOME}/go/src/orihime'"
	@echo "Make sure to 'go get ./cmd ./internal'"
	go build -o orihime orihime/cmd/orihime

internal/database/sql.go: tools/template/src/function-skeletons.expanded $(shell find tools/template/template -type f)
	( cd tools/template ; ./bin/altogether-now )

internal/protobuf/orihime.pb.go: tools/protobuf/orihime.proto 
	protoc tools/protobuf/orihime.proto --go_out=plugins=grpc:internal/protobuf/

.PHONY: build
