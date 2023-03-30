package main

import (
	_ "embed"
	"fmt"
	"github.com/urfave/cli"
	"go-gql/cmd"
	"go-gql/config"
	"go-gql/migrate"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.InitConfig()

	app := cli.NewApp()
	app.Name = "go-gql"
	app.Usage = "use graphql & go build a simple web application"
	app.Commands = []cli.Command{
		{
			Name:  "api",
			Usage: "start internal web server",
			Action: func(c *cli.Context) error {
				r := cmd.SetupRouter()
				err := r.Run(fmt.Sprintf(":%s", config.Config.App.ServerPort))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "migrate",
			Usage: "start api web server",
			Action: func(c *cli.Context) error {
				migrate.Migrate(config.GetDB())
				return nil
			},
		},
	}
	sigComplete := make(chan struct{})
	go func() {
		defer close(sigComplete)
		err := app.Run(os.Args)
		if err != nil {
			log.Fatal("app run failed, err: ", err.Error())
		}
	}()

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-sigTerm:
		log.Println("receive stop signal")
	case <-sigComplete:
	}
}
