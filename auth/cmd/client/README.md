## Usage

### Help
```go run main.go -h```

### Create User
```go run main.go -action create -login foo -password bar```

### Get Token
- ```go run main.go -action token -login foo -password bar```
- wrong password: ```go run main.go -action token -login foo -password barbar```

### Change Password
```go run main.go -action change -login foo -password bar -new-pass barbar```

### Get Token Again
```go run main.go -action token -login foo -password barbar```