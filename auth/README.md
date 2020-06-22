# Authentication Service
verifies user credentials and provides shortlived JWT.

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
