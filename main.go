package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"searchera/indexio"
	"searchera/server"
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
			Name:  "server",
			Usage: "Read index and starts the http server on the specified host and port",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "host",
					DefaultText: "127.0.0.1",
				},
				&cli.StringFlag{
					Name:        "port",
					DefaultText: "8080",
				},
				idxFlag,
			},
			Action: ProcessServer,
		},
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

func ProcessServer(c *cli.Context) error {
	host := c.String("host")
	port := c.String("port")
	idxPath := c.String("index")

	idx, err := indexio.ReadIndex(idxPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Index '%s' read succesfull", idxPath)

	if err := server.Run(host, port, idx, idxPath); err != nil {
		log.Fatal(err)
	}

	return nil
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
