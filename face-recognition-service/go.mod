module face-recognition-service

go 1.16

replace grpc-schemas => ../grpc-schemas

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/go-redis/redis/v8 v8.11.0
	github.com/gorilla/mux v1.8.0
	github.com/stretchr/testify v1.6.1
	go.mongodb.org/mongo-driver v1.5.4
	gonum.org/v1/gonum v0.9.3
	google.golang.org/grpc v1.39.0
	grpc-schemas v0.0.0-00010101000000-000000000000
)
