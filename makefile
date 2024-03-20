build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o rustdesk-api-server .
buildr:
	goreleaser build --rm-dist