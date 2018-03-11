GO ?= go
DOCKER ?= docker

VERSION := 0.1.1

.PHONY:
run:
	$(GO) run -ldflags '-X main.version=$(VERSION)' hello-http-server.go

hello-http-server: hello-http-server.go
	CGO_ENABLED=0 GOOS=linux $(GO) build -a -ldflags '-extldflags "-static" -X main.version=$(VERSION)' -o $@ $<

.PHONY:
clean:
	rm -rf hello-http-server

.PHONY:
docker-image: hello-http-server
	$(DOCKER) build -t hello-http-server:$(VERSION) .

.PHONY:
run-docker: docker-image
	$(DOCKER) run --rm -it -p 8080:8080 hello-http-server:$(VERSION)