package main

import (
	"fmt"

	"github.com/giansalex/pluralsight-transcript/service"
	"github.com/urfave/cli"
)

func defaultCommand(c *cli.Context) error {

	if c.NArg() == 0 {
		fmt.Println("Require Sqlite Path")
		return nil
	}

	path := c.Args().Get(0)
	lang := c.String("lang")

	newPath := fmt.Sprintf(newDbFormat, lang)
	_, err := service.CopyFile(path, newPath)
	if err != nil {
		fmt.Println("Cannot create new database", err.Error())
		return nil
	}

	fmt.Println("Create ", newPath)

	container := buildContainer(newPath)

	err = container.Invoke(func(transcript service.TranscriptService) {
		err := transcript.Transcript(lang)

		if err != nil {
			fmt.Println("Cannot save transcript", err.Error())
		}
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
