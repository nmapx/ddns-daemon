OS_LIST := linux windows darwin

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
	for os in $(OS_LIST); do \
		bash -c "GOOS=$$os exec go build -a -o ddns-daemon_amd64_$$os -ldflags=\"-X 'github.com/nmapx/ddns-daemon/cmd.version=$(VERSION)'\""; \
	done

.PHONY: get
get:
	exec go get

.PHONY: tidy
tidy:
	exec go mod tidy

.PHONY: fmt
fmt:
	exec gofmt -w -s ./
