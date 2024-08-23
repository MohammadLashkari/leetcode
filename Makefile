
## help: print help message
.PHONY: help
help:
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':'

## add/solution: create solution file and add it to the readme table
.PHONY: add/solution
add/solution:
	@go run ./cmd/add.go

