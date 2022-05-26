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