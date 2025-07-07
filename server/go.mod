module pearviewer/server

go 1.23.4

require (
	github.com/louischm/pkg v0.1.2
	golang.org/x/crypto v0.39.0
	google.golang.org/grpc v1.73.0
	gorm.io/driver/sqlite v1.6.0
	gorm.io/gorm v1.30.0
	pearviewer/generated v0.0.0-00010101000000-000000000000
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.28 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250603155806-513f23925822 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace pearviewer/generated => ../generated
