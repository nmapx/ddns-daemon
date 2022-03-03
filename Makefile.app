.PHONY: daemon
daemon: ./main.go
	exec go run . daemon

.PHONY: test
test:
	exec go test -v ./...

.PHONY: build
build: ./main.go
	exec go build -a -o ddns-daemon .

.PHONY: get
get:
	exec go get

.PHONY: tidy
tidy:
	exec go mod tidy

.PHONY: fmt
fmt:
	exec gofmt -w -s ./
