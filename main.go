package main

import (
	"fmt"
	"log"
	"os"

	"github.com/giansalex/pluralsight-transcript/repository"
	"github.com/giansalex/pluralsight-transcript/service"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/urfave/cli"
)

const apiUrl = "https://app.pluralsight.com/mobile-api/v2"
const newDbFormat = "pluralsight.%s.db"

func main() {

	app := cli.NewApp()
	app.Name = "Pluralsight Transcript"
	app.Usage = "Translate video to new language"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "es",
			Usage: "language for transcript",
		},
	}

	app.Action = func(c *cli.Context) error {

		if c.NArg() == 0 {
			fmt.Println("Require path Sqlite")
			return nil
		}

		path := c.Args().Get(0)
		lang := c.String("lang")

		newPath := fmt.Sprintf(newDbFormat, lang)
		_, err := service.CopyFile(path, newPath)
		if err != nil {
			fmt.Println("Couldn't copy new database", err.Error())
			return nil
		}

		fmt.Println("Create ", newPath)

		db, err := gorm.Open("sqlite3", newPath)
		if err != nil {
			fmt.Println("Failed to connect database", err.Error())
			return nil
		}
		defer db.Close()

		users := repository.NewUserRepository(db)
		courses := repository.NewCourseRepository(db)
		transcripts := repository.NewTranscriptRepository(db)

		user := users.GetUser()
		api := service.NewApi(apiUrl, user.Jwt.String)
		transcript := service.NewTranscriptService(courses, transcripts, api)
		err = transcript.Transcript(lang)

		if err != nil {
			fmt.Println("Couldn't save transcript", err.Error())
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
