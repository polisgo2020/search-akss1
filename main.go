package main

import (
	"log"
	"os"
	"searchera/indexio"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Searchera",
		Usage: "Make reverse index from directory with text files and search with it",
	}

	idxFlag := &cli.StringFlag{
		Name:     "index",
		Usage:    "Path to index file",
		Required: true,
	}

	app.Commands = []*cli.Command{
		{
			Name:    "make",
			Aliases: []string{"m"},
			Usage:   "Make reverse index from directory",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "dir",
					Usage:    "Path to directory with text files",
					Required: true,
				},
				idxFlag,
			},
			Action: ProcessMakeIdx,
		},
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "Search in reverse index",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "str",
					Usage:    "String for search",
					Required: true,
				},
				idxFlag,
			},
			Action: ProcessSearch,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func ProcessMakeIdx(c *cli.Context) error {
	dirPath := c.String("dir")
	idxPath := c.String("index")

	err := indexio.MakeAndWriteIndex(dirPath, idxPath)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func ProcessSearch(c *cli.Context) error {
	str := c.String("str")
	idxPath := c.String("index")

	err := indexio.ReadAndSearchIndex(idxPath, str)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
