package main

import (
	"context"
	"errors"
	"fmt"
	todo_api "github.com/brharrelldev/todo_grpc/api"
	"github.com/urfave/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var DeleteCmd cli.Command

func init() {
	DeleteCmd.Name = "delete"
	DeleteCmd.Description = "command to delete task"
	DeleteCmd.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "delete-task",
			Usage: "todo-client delete --delete-task=<id>",
		},
		cli.BoolFlag{
			Name:  "delete-all",
			Usage: "todo-client delete --delete-all",
		},
	}
	DeleteCmd.Action = DeleteTasks
}

func DeleteTasks(c *cli.Context) error {
	defer client.Close()

	if c.Bool("delete-all") && c.String("delete-task") != "" {
		return errors.New("delete-all and delete-task cannot be set together ")
	}

	if c.Bool("delete-all") {
		if err := deleteAll(); err != nil{
			return fmt.Errorf("error deleting tasks %v", err)
		}

		return nil

	}

	if err := cli.ShowSubcommandHelp(c);err != nil{
		return fmt.Errorf("error showing subcommand %v", err)
	}

	return nil
}

func deleteAll() error {
	var ids []string

	getAllIds := todo_api.NewGetAllTodosClient(client)
	deleteAllIds := todo_api.NewDeleteTodoClient(client)

	resp, err := getAllIds.GetTodos(context.Background(), &todo_api.GetAllTodoRequest{})
	retrieveStatus := status.Convert(err)

	if retrieveStatus.Code() != codes.OK {
		return fmt.Errorf("invalid response received %v", err)
	}

	for _, id := range resp.Todos {
		ids = append(ids, id.Id)
	}

	for _, id := range ids {
		if _, err := deleteAllIds.DeleteTodo(context.Background(), &todo_api.DeleteTodoRequest{
			Id: id,
		}); err != nil {
			if status.Convert(err).Code() != codes.OK {
				return fmt.Errorf("invalid response received %v", err)
			}

		}

		fmt.Printf("deletion of %s successful\n", id)
	}

	return nil

}
