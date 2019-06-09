package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	apiURL      = "https://app.pluralsight.com/mobile-api/v2"
	newDbFormat = "pluralsight.%s.db"
)

func main() {

	app := cli.NewApp()
	app.Name = "Pluralsight Transcript"
	app.Usage = "Translate video to your preference language"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "es",
			Usage: "language for transcript",
		},
	}

	app.Action = defaultCommand

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
