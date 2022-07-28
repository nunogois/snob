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

	app := &cli.App{
		Name:    "snob",
		Version: "v0.0.5",
		Usage:   "Simple fetcher for movies and TV shows info",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "key",
				Aliases:     []string{"k", "s"},
				Usage:       "OMDb API key to use",
				Destination: &key,
			},
			&cli.BoolFlag{
				Name:        "json",
				Aliases:     []string{"j"},
				Usage:       "Get result as JSON",
				Destination: &jsonPrint,
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
				printJSON(result)
			} else {
				printPretty(result)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
