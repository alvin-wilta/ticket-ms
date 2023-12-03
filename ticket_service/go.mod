module github.com/alvin-wilta/ticket-ms/ticket_service

go 1.21.4

replace github.com/alvin-wilta/ticket-ms/proto => ../proto

require (
	github.com/alvin-wilta/ticket-ms/proto v0.0.0-00010101000000-000000000000
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/joho/godotenv v1.5.1
	google.golang.org/grpc v1.59.0
	gorm.io/driver/postgres v1.5.4
	gorm.io/gorm v1.25.5
)

require github.com/golang/snappy v0.0.1 // indirect

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/nsqio/go-nsq v1.1.0
	github.com/stretchr/testify v1.8.2 // indirect
	golang.org/x/crypto v0.16.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
