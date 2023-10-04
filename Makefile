.DEFAULT_GOAL := help

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build ./cmd/httpRedirect

build-local:
	go build ./cmd/testProxy

.PHONY: docker
docker: build
	docker build --platform=linux/amd64 -t httpredirect:latest .