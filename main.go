package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go-crack/jobs"
	"go-crack/util"
	"log"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	util.DownloadLists()
}

func main() {
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("Version=%s Commit=%s Build Time=%s\n", version, commit, date)
	}

	app := &cli.App{
		Name:    "go-crack",
		Usage:   "Crack archive folders with various strategies along with multi-threading support.",
		Version: version,
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Konotorii",
				Email: "github@konotorii.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "strategy",
				Usage:       "Which cracking strategy to use.",
				Aliases:     []string{"s", "str"},
				Destination: &util.SelectedOptions.Strategy,
				Required:    true,
				DefaultText: "rockyou",
			},
			&cli.BoolFlag{
				Name:        "thread",
				Usage:       "Should process be threaded",
				Aliases:     []string{"t"},
				Destination: &util.SelectedOptions.IsPooled,
				Required:    false,
				Value:       false,
			},
			&cli.BoolFlag{
				Name:        "linear",
				Usage:       "Should list be split into chunks",
				Aliases:     []string{"l"},
				Destination: &util.SelectedOptions.IsLinear,
				Required:    false,
				Value:       false,
			},
			&cli.IntFlag{
				Name:        "tc",
				Aliases:     []string{"thread-count"},
				Usage:       "Total amount of threads to use",
				Destination: &util.SelectedOptions.TotalThreads,
				Required:    false,
				DefaultText: "4",
				Value:       4,
			},
			&cli.StringFlag{
				Name:        "file",
				Usage:       "Where the archive file is located",
				Aliases:     []string{"location", "save-location"},
				Destination: &util.SelectedOptions.FileLocation,
				Required:    true,
				DefaultText: "/.././file.zip",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Printf("%v\n", util.SelectedOptions)

			if util.SelectedOptions.Strategy == "rockyou" {
				jobs.RockYou()
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
