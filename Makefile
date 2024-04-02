# ==================================================================================== #
# HELPERS
# ==================================================================================== #
## make help OR make: Show the possible commands within the make file
.PHONY: help
help:
	@echo 'List of available commands:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
GIT_COMMIT := $(shell git rev-list -1 HEAD)
## make tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v
	golangci-lint run

## make audit: run quality control checks
.PHONY: audit
audit:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go test -race -vet=off ./...
	go mod verify