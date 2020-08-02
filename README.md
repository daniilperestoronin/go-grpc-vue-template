# go-grpc-vue-template

*Status: in development*

## Description
Monolithic web application template.

Technology stack:
- [Go](https://golang.org) for backend.
- [Vue.js](https://vuejs.org) and [Vuetify.js](https://vuetifyjs.com/) for frontend.
- [gRPC](https://grpc.io) for code generation and communication between frontend and backend.
- Embedded proxy [grpcweb](https://github.com/improbable-eng/grpc-web/tree/master/go/grpcweb) for gRPC communication between frontend and backend.
- Embedded OAuth2.0 [ory/fosite](https://github.com/ory/fosite/) for security.
- Embedded static server using standard [net/http](https://golang.org/pkg/net/http/).

## Project structure
    .
    ├── api              # Protobuf api definition
    ├── front            # Frontend src
    ├── server           # Server side src

## Preparing dev environment

For local development you need:
- [Go]()
- [protobuf]() with [protoc-gen-go]() and [protoc-gen-grpc-web]()
- [nodejs]()
- [yarn]()

Use [Make](https://en.wikipedia.org/wiki/Make_(software)) for install Protoc (currently support only for macOS)

```bash
make install-protoc
```

## Build
```bash
make app-build
```
