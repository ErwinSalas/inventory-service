## Prerequisites
1. Go, any one of the three latest major releases of Go.

For installation instructions, see Go’s Getting Started guide.

2. Protocol buffer compiler, protoc, version 3.

For installation instructions, see Protocol Buffer Compiler Installation.

3. Go plugins for the protocol compiler:

Install the protocol compiler plugins for Go using the following commands:
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
Update your PATH so that the protoc compiler can find the plugins:
```
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Generate Protobuf

```
protoc -I=proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative proto/inventory.proto
```


## Enviroment Variables
```
DB_HOST=fullstack-mysql
DB_DRIVER=mysql 
DB_USER=username
DB_PASSWORD=password
DB_NAME=fullstack_api
DB_PORT=3306
 
API_HOST=inventory-api
```