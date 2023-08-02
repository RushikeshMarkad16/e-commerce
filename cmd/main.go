package main

import (
	"os"
	"sync"

	"github.com/RushikeshMarkad16/e-commerce/app"
	"github.com/RushikeshMarkad16/e-commerce/config"
	"github.com/RushikeshMarkad16/e-commerce/db"
	"github.com/RushikeshMarkad16/e-commerce/server"
	"github.com/urfave/cli"
)

var wg sync.WaitGroup

func main() {
	config.Load()
	app.Init()
	defer app.Close()

	cliApp := cli.NewApp()
	cliApp.Name = "E-Commerce App"
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start server",
			Action: func(c *cli.Context) error {
				wg.Add(2)
				go server.StartAPIServer(&wg)
				go server.StartgRPCServer(&wg)
				wg.Wait()
				return nil
			},
		},

		{
			Name:  "create_migration",
			Usage: "create migration file",
			Action: func(c *cli.Context) error {
				return db.CreateMigrationFile(c.Args().Get(0))
			},
		},

		{
			Name:  "migrate",
			Usage: "run db migrations",
			Action: func(c *cli.Context) error {
				err := db.RunMigrations()
				return err
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback migrations",
			Action: func(c *cli.Context) error {
				return db.RollbackMigrations(c.Args().Get(0))
			},
		},
	}
	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
