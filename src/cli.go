package main

import (
	"log"
	"os"
	"sort"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/cli"

	urfavecli "github.com/urfave/cli"
)

const defaultURL = "http://localhost:3000"

func mainCLI() {
	app := urfavecli.NewApp()
	app.Name = "adoption"
	app.Description = "adoption cli for interacting with adoption server"
	app.Version = "0.1.0"
	app.Usage = ""
	app.UsageText = "adoption command [command options] [arguments...]"

	app.Flags = []urfavecli.Flag{
		urfavecli.StringFlag{
			Name:  "URL, u",
			Value: defaultURL,
			Usage: "The URL of the running instance of adoption server",
		},
	}

	app.Commands = []urfavecli.Command{
		{
			Name:    "load",
			Aliases: []string{"l"},
			Usage:   "load csv data into adoption server",
			Action:  loadData,
			Flags: []urfavecli.Flag{
				urfavecli.StringFlag{
					Name:  "file, f",
					Usage: "Load data from `FILE`",
				},
				urfavecli.StringFlag{
					Name:  "kind, k",
					Usage: "Kind of data to load",
				},
			},
		},
	}

	sort.Sort(urfavecli.FlagsByName(app.Flags))
	sort.Sort(urfavecli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func loadData(c *urfavecli.Context) error {
	filename := c.String("file")
	kind := c.String("kind")
	client := &cli.Client{
		URL:      c.GlobalString("URL"),
		Filename: filename,
	}
	switch kind {
	case "tools":
		client.LoadTools()
	case "people":
		client.LoadPeople()
	default:
		glog.Fatalf("kind %s not supported", kind)
	}

	return nil
}
