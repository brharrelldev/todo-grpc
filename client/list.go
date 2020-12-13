package main

import (
	"context"
	"errors"
	"fmt"
	todo_api "github.com/brharrelldev/todo_grpc/api"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"strconv"
	"time"
)

var ListCmd cli.Command

func init()  {
	ListCmd.Name = "list"
	ListCmd.Description = "listing tasks with task id"
	ListCmd.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "list-task",
			Usage: "todo-client --list-task=<id>",
		},
		cli.BoolFlag{
			Name: "list-all",
			Usage: "todo-client --list-all",
		},
	}
	ListCmd.Action = listTasks

}

func listTasks(c *cli.Context) error {
	defer client.Close()
	if c.Bool("list-all") && c.String("list-task") != "" {
		return errors.New("both list-all and list-task can't be set together")
	}

	if c.Bool("list-all") {
		if err := listAll(); err != nil{
			return fmt.Errorf("error listing tasks %v", err)
		}
		return nil
	}

	if c.String("list-task") != ""{
		showOne := todo_api.NewGetTodoClient(client)
		resp, err := showOne.GetTodo(context.Background(), &todo_api.GetTodoRequest{
			Id:  c.String("list-task"),
		})
		if err != nil{
			return fmt.Errorf("error getting record %v", err)
		}

		log.Println(resp)
		return nil
	}

	if err := cli.ShowSubcommandHelp(c); err != nil{
		return fmt.Errorf("error show subcommand %v", err)
	}
	return nil
}

func listAll() error  {
	var data [][]string
	listAll := todo_api.NewGetAllTodosClient(client)

	start := time.Now()
	resp, err := listAll.GetTodos(context.Background(), &todo_api.GetAllTodoRequest{})
	callStatus := status.Convert(err)
	if callStatus.Code() != codes.OK {
		return fmt.Errorf("could not list all due to %v", err)
	}

	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader([]string{ "id", "task name", "done?"})

	for _, task := range resp.Todos {
		taskEntry := []string{ task.Id, task.TaskName, strconv.FormatBool(task.Done)}
		data = append(data, taskEntry)
	}

	count := 0
	for _, entries := range data {
		tw.Append(entries)
		count++
	}

	tw.Render()
	end := time.Since(start).Milliseconds()

	fmt.Println()
	fmt.Printf("found %d record(s) in %d ms\n", count, end)
	return nil

}