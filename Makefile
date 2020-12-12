BIN_DIR=bin
CLIENT_BIN=$(BIN_DIR)/todo-client
SERVER_BIN=$(BIN_DIR)/todo-server

generate_grpc:
	protoc -I. -I$(GOPATH)/src   --go_out=plugins=grpc:. api/*.proto

server-build:
	go build   -o $(SERVER_BIN) server/*.go

client-build:
	go build   -o $(CLIENT_BIN) client/*.go

cleanup:
	rm -rf /tmp/todo_db


all: generate_grpc server-build client-build