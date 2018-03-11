# hello-http-server

A simple "Hello, world" http server written in go and packaged as a Docker container.

Used to function test any Docker based deployment infrastructure.

## Building the server

```
$ make hello-http-server
```

to build the server.

## Running the server

Build before you run

```
$ ./hello-http-server
```

to start the server running on port :8080. Open http://localhost:8080/ in your browser to test.

## Building the Docker image

```
$ make docker-image
```

will build the server and bundle it as a docker image. The image will be tagged `hello-world-server:$VERSION` where `$VERSION` is defined in the [Makefile](./Makefile)
