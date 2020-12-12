package main

import (
	"bytes"
	"context"
	"fmt"
	todo_api "github.com/brharrelldev/todo_grpc/api"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"
	"log"
	"net"
	"os"
)

type Tasks struct {
	Tasks []Task `yaml:"tasks"`
}

type Task struct {
	TaskName string `yaml:"task_name"`
}

func main()  {
	app := cli.NewApp()
	app.Name = "todo-client"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "host",
			Usage:       "todo-client --host",
			EnvVar:      "TODO_SERVER_HOST",
			Required: true,
		},
		cli.StringFlag{
			Name:        "port",
			Usage:       "todo-client --port",
			EnvVar:      "TODO_SERVER_PORT",
			Required: true,
		},
		cli.BoolFlag{
			Name:        "list-all",
			Usage:       "todo-client --list-all",
		},
		cli.BoolFlag{
			Name: "add-task",
			Usage: "todo-client --add-task",
		},
	}
	app.Action = func(c *cli.Context) error {
		client, err := grpc.Dial(net.JoinHostPort(c.String("host"),c.String("port")),
			grpc.WithInsecure())
		if err != nil{
			return fmt.Errorf("error dialing grpc server %v", err)
		}

		defer client.Close()

		if c.Bool("list-all"){
			var data [][]string
			listAll := todo_api.NewGetAllTodosClient(client)
			resp, err := listAll.GetTodos(context.Background(),&todo_api.GetAllTodoRequest{})
			callStatus  := status.Convert(err)
			if callStatus.Code() != codes.OK{
				return fmt.Errorf("could not list all due to %v", err)
			}

			tw := tablewriter.NewWriter(os.Stdout)
			tw.SetHeader([]string{"task name"})

			for _, task := range resp.Todos{
				taskEntry := []string{task.TaskName}
				data = append(data, taskEntry)
			}


			for _, entries := range data{
				tw.Append(entries)
			}

			tw.Render()




		}

		if c.Bool("add-task"){
			var tasks Tasks
			addTask := todo_api.NewCreateTodoClient(client)

			buf := make([]byte, 8192)

			f, err := os.Open("task_data/todo_list.yml")
			if err != nil{
				return fmt.Errorf("error getting task data source %v", err)
			}

			defer f.Close()

			offset, err := f.Read(buf)
			if err != nil{
				return fmt.Errorf("error reading fle content %v", err)
			}

			content  := buf[:offset]

			if err := yaml.NewDecoder(bytes.NewBuffer(content)).Decode(&tasks); err != nil{
				return fmt.Errorf("error decoding yaml data file %v", err)
			}

			for _, task := range tasks.Tasks{
				req := &todo_api.CreateTodoRequest{
					TaskName:             task.TaskName,
				}

				resp, err := addTask.CreateTodo(context.Background(), req)
				respStatus := status.Convert(err)

				if respStatus.Code() != codes.OK{
					return fmt.Errorf("error occured, request is not ok")
				}

				fmt.Println(resp)
			}



		}


		return nil
	}

	if err := app.Run(os.Args); err != nil{
		log.Fatal(err)

	}

}
