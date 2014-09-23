package main

import (
	"fmt"
	"github.com/Masterminds/cookoo"
	"github.com/codegangsta/cli"
	"os"
)

func main() {

	// Setup Cookoo
	reg, router, cxt := cookoo.Cookoo()

	// A test route
	reg.Route("version", "Print the version and exit.").
		Does(showVersion, "_").Using("version").From("cxt:version")

	// Play with codegangsta/cli
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) {
		println("boom! I say!")
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}
	// try: go run cli.go test --version='foo'
	app.Commands = []cli.Command{
		{
			Name:      "test",
			ShortName: "t",
			Usage:     "A test cookoo route",
			Action: func(c *cli.Context) {
				cxt.Put("version", c.String("version"))
				router.HandleRequest("version", cxt, false)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "version",
					Value: "1.0.0-beta",
					Usage: "specified string version",
				},
			},
		},
		// One command level in.
		// try: go run cli.go complete bar
		{
			Name:      "complete",
			ShortName: "c",
			Usage:     "complete a task on the list",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args().First())
			},
		},
		{
			Name:      "template",
			ShortName: "r",
			Usage:     "options for task templates",
			Subcommands: []cli.Command{
				// Two commands in
				// try: go run cli.go r add bar
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) {
						println("new task template: ", c.Args().First())
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) {
						println("removed task template: ", c.Args().First())
					},
					Subcommands: []cli.Command{
						// Three commands
						// try: go run cli.go r remove add bar
						{
							Name:  "add",
							Usage: "add a new template",
							Action: func(c *cli.Context) {
								println("boom: ", c.Args().First())
							},
						},
						{
							Name:  "foo",
							Usage: "add a new template",
							Action: func(c *cli.Context) {
								println("boom: ", c.Args().First())
							},
							Subcommands: []cli.Command{
								// Four commands in
								// try: go run cli.go r remove foo add bar
								{
									Name:  "add",
									Usage: "add a new template",
									Action: func(c *cli.Context) {
										println("stick: ", c.Args().First())
									},
								},
							},
						},
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

func showVersion(c cookoo.Context, p *cookoo.Params) (interface{}, cookoo.Interrupt) {
	version := p.Get("version", "0.1.0").(string)
	fmt.Println(version)
	return version, nil
}
