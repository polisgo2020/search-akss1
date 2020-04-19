package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"os"
	"searchera/config"
	"searchera/indexio"
	"searchera/server"
)

var cfg config.Config

func main() {
	var err error

	cfg, err = config.Load()
	if err != nil {
		log.Fatal().Err(err)
	}

	lvl, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal().Err(err)
	}

	zerolog.SetGlobalLevel(lvl)

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

	if err = app.Run(os.Args); err != nil {
		log.Fatal().Err(err)
	}
}

func ProcessServer(c *cli.Context) error {
	addr := cfg.Listen
	idxPath := c.String("index")

	idx, err := indexio.ReadIndex(idxPath)
	if err != nil {
		log.Fatal().Err(err)
	}

	log.Info().Msgf("Index '%s' read successful", idxPath)

	if err := server.Run(addr, idx, idxPath); err != nil {
		log.Fatal().Err(err)
	}

	return nil
}

func ProcessMakeIdx(c *cli.Context) error {
	dirPath := c.String("dir")
	idxPath := c.String("index")

	err := indexio.MakeAndWriteIndex(dirPath, idxPath)
	if err != nil {
		log.Fatal().Err(err)
	}

	return nil
}

func ProcessSearch(c *cli.Context) error {
	str := c.String("str")
	idxPath := c.String("index")

	err := indexio.ReadAndSearchIndex(idxPath, str)
	if err != nil {
		log.Fatal().Err(err)
	}

	return nil
}
