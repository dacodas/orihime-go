build: internal/sql.go
	@echo "Make sure to symlink 'ln -s \$${PWD} \$${HOME}/go/src/orihime'"
	@echo "Make sure to 'go get ./cmd ./internal'"
	go build -o orihime orihime/pkg

internal/sql.go: template/src/function-skeletons.expanded $(shell find template/template -type f)
	( cd template ; ./bin/altogether-now )

.PHONY: build
