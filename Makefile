build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/go-snippets-linux-amd64

build-mac:
	go build -o bin/go-snippets-darwin-amd64

install-mac:
	go install github.com/lwabish/go-snippets@latest
