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

## Preparing dev environment
Use [Make](https://en.wikipedia.org/wiki/Make_(software))

```bash
make prepare_dev
```

## Build
```bash
make build
```
