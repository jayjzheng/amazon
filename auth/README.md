# Authentication Service
verifies user credentials and provides shortlived JWT.

## Ideas to demonstrate
- repo layout, the usage of `internal` package
- the concept of `domain` package, and its usage [here](https://github.com/jayjzheng/amazon/tree/master/auth/internal/domain)
- service types for implementing business logic independent of external deps. [here](https://github.com/jayjzheng/amazon/blob/master/auth/internal/domain/auth_service.go)
- unit testing [here](https://github.com/jayjzheng/amazon/blob/master/auth/internal/domain/auth_service_test.go)
- flexible variadic configuration [here](https://github.com/jayjzheng/amazon/blob/master/auth/internal/jwt/jwt.go)

## Goals
- simple to change business logic
- simple to add/change implementation of UserStore, TokenGenerator
- simple to change from long running server to cli
- simple to change from gRPC to JSON

## Unfinished business
- logging
- instrumentation
- end to end testing

## Technology
- JWT
- datastore of choice
- language of choice

## Domain Object
- User { login, password }

## Caller
- API Gateway

## Actions
- Create User
- Change Password
- Create Token
- Refresh Token

## Publishes
- userCreated
- passwordChanged
- tokenCreated
- tokenRefreshed

## Open Discussions
### datastore
- both NoSQL and RDBMS works fine.

### programming language
depends on the team members, experience, preferences, etc.

## Usage
### Prereq
- ```brew install protobuf```
- ```go get -u github.com/golang/protobuf/protoc-gen-go```
- ```go install google.golang.org/grpc/cmd/protoc-gen-go-grpc```

### Generate Go Protobuffer code
```protoc --go_out=. pb/auth.proto && protoc --go-grpc_out=. ./pb/auth.proto```
