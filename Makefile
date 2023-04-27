gen_go:
	protoc --go_out=. --go-grpc_out=.  api/grpc/feed-service.proto

gen_python:
	python -m grpc_tools.protoc -I api/grpc/ --python_out=./python --pyi_out=./python --grpc_python_out=./python api/grpc/feed-service.proto

build:
	go build -o ./bin/server_app ./cmd/server/main.go
