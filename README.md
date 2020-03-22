# Messaging API
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/prasetyowira/message/CI?style=flat-square)
![CI](https://github.com/prasetyowira/message/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/prasetyowira/message)](https://goreportcard.com/report/github.com/prasetyowira/message)
[![GolangCI](https://golangci.com/badges/github.com/prasetyowira/message.svg)](https://golangci.com/r/github.com/prasetyowira/message)

A Messaging with Golang for Warung Pintar recruitment assignment

## Getting started

Go Version: 1.14

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make up
$ make start
$ make run
```

Or running inside docker
```console
$ make up
$ make start
$ docker-compose -f docker-compose.prod.yml up -d
```


### REST API

```console
$ make start
```

```http
GET/POST http://127.0.0.1:8000/message
GET http://127.0.0.1:8000/message/{id}
```

Open openapi doc on port [127.0.0.1:81](127.0.0.1:81)

### Graphql


```http
GET/POST http://127.0.0.1:8000/graphql
```

To explore more, open graphql playground
```http
GET http://127.0.0.1:8000/playground
```

### Websocket


```http
ws://127.0.0.1:8000/ws
```

To use, open sample client
```http
GET http://127.0.0.1:8000/websocket
```


### GRPC


```http
http://127.0.0.1:8001
```

.proto file
```console
./api/proto/messaging/v1
```


### Testing

``make test``
