sql.go: template/src/function-skeletons.expanded $(shell find template/template -type f)
	( cd template ; ./bin/altogether-now )
