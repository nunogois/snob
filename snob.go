package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	loadConfig()
	var key string
	var jsonPrint bool
	var avoidSpoilers bool

	app := &cli.App{
		Name:    "snob",
		Version: "v0.0.6",
		Usage:   "Simple fetcher for movies and TV shows info",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "key",
				Aliases:     []string{"k"},
				Usage:       "OMDb API key to use",
				Destination: &key,
			},
			&cli.BoolFlag{
				Name:        "json",
				Aliases:     []string{"j"},
				Usage:       "Get result as JSON",
				Destination: &jsonPrint,
			},
			&cli.BoolFlag{
				Name:        "spoilers",
				Aliases:     []string{"s"},
				Usage:       "Get result without spoiler-sensitive information",
				Destination: &avoidSpoilers,
			},
		},
		Action: func(ctx *cli.Context) error {
			query := strings.Join(ctx.Args().Slice(), " ")

			if query == "" {
				if key != "" {
					setKey(key)
					fmt.Println("Key set successfully")
					return nil
				} else {
					cli.ShowAppHelpAndExit(ctx, 0)
				}
			}

			result := new(Result)
			if err := search(query, result, key); err != nil {
				return err
			}

			if jsonPrint {
				printJSON(result, avoidSpoilers)
			} else {
				printPretty(result, avoidSpoilers)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
