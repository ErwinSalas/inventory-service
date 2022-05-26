module github.com/ErwinSalas/inventory-service

go 1.18

replace github.com/ErwinSalas/inventory-service/proto => ./proto

require (
	github.com/ErwinSalas/inventory-service/proto v0.0.0-00010101000000-000000000000
	github.com/jinzhu/gorm v1.9.10
	github.com/joho/godotenv v1.3.0
	google.golang.org/grpc v1.46.2
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20210121164019-fc48d45331c7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
