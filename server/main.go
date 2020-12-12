package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "todo-server"
	app.Description = "start todo server"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "host",
			Usage:    "todo-server --host <host address>",
			EnvVar:   "TODO_SERVER_HOST",
			Required: true,
		},
		cli.StringFlag{
			Name:     "port",
			Usage:    "todo-server --port <port num>",
			EnvVar:   "TODO_SERVER_PORT",
			Required: true,
		},
	}

	app.Action = serverStart

	if err := app.Run(os.Args); err != nil{
		log.Fatal(err)
	}

}

func serverStart(c *cli.Context) error  {

	errChan := make(chan error, 1)
	sigChan := make(chan os.Signal, 1)

	s, err := NewTodoGRPCServer(c.String("host"), c.String("port"))
	if err != nil{
		return fmt.Errorf("error cannot start server %v",err )

	}

	go func() {
		if err := s.StartServer(); err != nil{
			errChan <- err
			return
		}
	}()

	select {
	case err :=  <- errChan:
		if err != nil{
			return fmt.Errorf("error could not start server due to %v", err)
		}
	case <-sigChan:
		if err := s.Close(); err != nil{
			return fmt.Errorf("error shutting down todo-server")
		}
		log.Println("graceful shutdown received")
	}

	return nil

}
