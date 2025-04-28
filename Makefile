DATE    := $(shell date)
HASH    := $(shell git rev-parse --short HEAD)
LDFLAGS	:= -ldflags "-s -X 'main.Version=${HASH}' -X 'main.BuildDate=${DATE}'"
#
# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

.PHONY: audit
audit:
	go mod tidy # -diff
	go mod verify
	test -z "$(shell gofmt -l .)"
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...


# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

.PHONY: build-linux-amd64
build-linux-amd64:
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o growatt-export-limit-control growatt-export-limit-control.go
	upx -5 growatt-export-limit-control
