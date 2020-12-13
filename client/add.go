package main

import (
	"bytes"
	"context"
	"fmt"
	todo_api "github.com/brharrelldev/todo_grpc/api"
	"github.com/urfave/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

var AddCmd cli.Command

func init() {
	AddCmd.Name = "add"
	AddCmd.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "task-name",
			Usage: "task-client add --task-name=<test>",
		},
		cli.BoolFlag{
			Name:  "task-file",
			Usage: "task-client add --task-file",
		},
	}
	AddCmd.Description = "add task to task server"
	AddCmd.Action = AddTask
}

func AddTask(c *cli.Context) error {
	defer client.Close()

	if c.Bool("task-file") && c.String("task-name") != "" {
		return fmt.Errorf("task-file param and task-name cannot be set together")
	}

	if c.Bool("task-file") {
		if err := fileHandler(); err != nil {
			return fmt.Errorf("error occured for task-file handler %v", err)
		}
		return nil
	}


	if c.String("task-name") != "" {
		resp, err := addTaskById(c.String("task-name"))
		if err != nil {
			return err

		}

		fmt.Println(resp)
		return nil
	}

	if err := cli.ShowSubcommandHelp(c); err != nil{
		return fmt.Errorf("could not show help %v", err)
	}

	return nil
}

func fileHandler() error {
	start := time.Now()
	var tasks Tasks
	addTask := todo_api.NewCreateTodoClient(client)

	buf := make([]byte, 165000)

	f, err := os.Open("task_data/todo_list.yml")
	if err != nil {
		return fmt.Errorf("error getting task data source %v", err)
	}

	defer f.Close()

	offset, err := f.Read(buf)
	if err != nil {
		return fmt.Errorf("error reading fle content %v", err)
	}

	content := buf[:offset]

	if err := yaml.NewDecoder(bytes.NewBuffer(content)).Decode(&tasks); err != nil {
		return fmt.Errorf("error decoding yaml data file %v", err)
	}

	count := 0
	for _, task := range tasks.Tasks {
		req := &todo_api.CreateTodoRequest{
			TaskName: task.TaskName,
		}

		resp, err := addTask.CreateTodo(context.Background(), req)
		respStatus := status.Convert(err)

		if respStatus.Code() != codes.OK {
			return fmt.Errorf("error occured, request is not ok")
		}

		fmt.Println(resp)
		count++
	}

	end := time.Since(start).Milliseconds()

	fmt.Printf("wrote %d record(s) in %d ms \n", count, end)
	return nil

}

func addTaskById(taskName string) (*todo_api.CreateTodoResponse, error) {

	add := todo_api.NewCreateTodoClient(client)
	resp, err := add.CreateTodo(context.Background(), &todo_api.CreateTodoRequest{
		TaskName: taskName,
	})

	if status.Convert(err).Code() != codes.OK {
		return nil, fmt.Errorf("error occured %v", err)
	}

	return resp, nil

}
