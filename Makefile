APP_BIN = app/build/app

build:
	go build -o $APP_BIN ./cmd/app/main.go

gen:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.
