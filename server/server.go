package main

import (
	"context"
	"fmt"
	todo_api "github.com/brharrelldev/todo_grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type TodoServer struct {
	db         *TodoDatabase
	grpcServer *grpc.Server
	host       string
	port       string
}

func (s *TodoServer) GetTodos(ctx context.Context, request *todo_api.GetAllTodoRequest) (*todo_api.GetAllTodoResponse, error) {
	var todos []*todo_api.Todo
	todoIter, err := s.db.GetAll()
	if err != nil{
		return &todo_api.GetAllTodoResponse{}, status.Error(codes.NotFound, err.Error())
	}

	for _, i := range todoIter{
		t := &todo_api.Todo{
			TaskName: i.TaskName,
			Done: i.Done,
		}

		todos = append(todos, t)
	}

	return &todo_api.GetAllTodoResponse{
		Todos: todos,
	}, status.Error(codes.OK, "retrieved successfully")
}

func NewTodoGRPCServer(host, port string) (*TodoServer, error) {
	log.Println("initializing database")
	db, err := NewDatabase()
	if err != nil {
		log.Println("error initializing database")
	}

	s := grpc.NewServer()

	return &TodoServer{
		grpcServer: s,
		db:         db,
		host:       host,
		port:       port,
	}, nil

}

func (s *TodoServer) GetTodo(ctx context.Context, request *todo_api.GetTodoRequest) (*todo_api.GetTodoResponse, error) {

	id := request.Id

	t, err := s.db.Get(id)
	if err != nil{
		return &todo_api.GetTodoResponse{}, status.Errorf(codes.NotFound, err.Error())
	}

	resp := &todo_api.GetTodoResponse{
		Response:            &todo_api.Todo{
			Id:                   id,
			TaskName:             t.TaskName,
			Done:                 t.Done,
		},
	}
	return resp, status.Error(codes.OK, "record created")
}

func (s *TodoServer) CreateTodo(ctx context.Context, request *todo_api.CreateTodoRequest) (*todo_api.CreateTodoResponse, error) {
	t := &TodoData{
		TaskName: request.TaskName,
	}

	id, err := s.db.Add(t)
	if err != nil{
		return &todo_api.CreateTodoResponse{}, status.Error(codes.Unknown, err.Error())
	}

	return &todo_api.CreateTodoResponse{
		Id:                   id,
		Message:              "successfully created new record",
	}, status.Error(codes.OK, "successfully created new record")
}



func (s *TodoServer) StartServer() error {

	todo_api.RegisterCreateTodoServer(s.grpcServer, s)
	log.Println("registering rpc method CreateTodo")

	todo_api.RegisterGetTodoServer(s.grpcServer, s)
	log.Println("registering rpc method GetTodo")
	todo_api.RegisterGetAllTodosServer(s.grpcServer, s)

	lis, err := net.Listen("tcp", net.JoinHostPort(s.host, s.port))
	if err != nil {
		return fmt.Errorf("error allocating list %v", err)
	}

	log.Printf("start todo-server on host %s and port %s", s.host, s.port)
	if err := s.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("error starting grpc server %v", err)
	}

	return nil

}

func (s *TodoServer) Close() error  {
	defer s.grpcServer.Stop()
	if err := s.db.db.Close(); err != nil{
		return fmt.Errorf("error closing db %v", err)
	}

	return nil
}
