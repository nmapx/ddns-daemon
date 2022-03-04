.PHONY: daemon
daemon: ./main.go
	exec go run . daemon

.PHONY: test
test:
	exec go test -v ./...

.PHONY: build
build: ./main.go
	exec go build -a -o ddns-daemon .

.PHONY: release
release: ./main.go
	GOOS=linux exec go build -a -o ddns-daemon_amd64_linux -ldflags="-X 'github.com/nmapx/ddns-daemon/cmd.version=$(VERSION)'"
	GOOS=windows exec go build -a -o ddns-daemon_amd64_windows -ldflags="-X 'github.com/nmapx/ddns-daemon/cmd.version=$(VERSION)'"
	GOOS=darwin exec go build -a -o ddns-daemon_amd64_darwin -ldflags="-X 'github.com/nmapx/ddns-daemon/cmd.version=$(VERSION)'"

.PHONY: get
get:
	exec go get

.PHONY: tidy
tidy:
	exec go mod tidy

.PHONY: fmt
fmt:
	exec gofmt -w -s ./
