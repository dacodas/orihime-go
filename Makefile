build: internal/sql.go
	go build -o orihime orihime/pkg

internal/sql.go: template/src/function-skeletons.expanded $(shell find template/template -type f)
	( cd template ; ./bin/altogether-now )

.PHONY: build
