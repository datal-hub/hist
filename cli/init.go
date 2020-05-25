package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/urfave/cli"
)

var countFiles int
var countSymbolsPerFile int
var dirPath string

//description run the application from the command line for initialize ascii files
var Init = cli.Command{
	Name:        "init",
	Usage:       "Initialize test ascii files",
	Description: ``,
	Action:      initFiles,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:        "files_count, fc",
			Value:       1000,
			Usage:       "Count of files fot test",
			Destination: &countFiles},
		cli.IntFlag{
			Name:        "symbols_count, sc",
			Value:       100,
			Usage:       "Count of symbols per file",
			Destination: &countSymbolsPerFile},
		cli.StringFlag{
			Name:        "path_to_dir, path",
			Value:       "./examples",
			Usage:       "Path to directory with example files",
			Destination: &dirPath},
	},
}

func makeTestFiles(path string, countFiles int, countSymbolsPerFile int) error {
	rand.Seed(time.Now().Unix())
	for i := 0; i < countFiles; i++ {
		file, err := os.Create(fmt.Sprintf("%s/%d", path, i))
		if err != nil {
			return err
		}
		defer file.Close()
		arrForWrite := make([]byte, 0, countSymbolsPerFile)
		for j := 0; j < countSymbolsPerFile; j++ {
			arrForWrite = append(arrForWrite, byte(rand.Intn(128)))
		}
		if _, err = file.Write(arrForWrite); err != nil {
			return err
		}
	}
	return nil
}

func initFiles(c *cli.Context) error {
	if err := makeTestFiles(dirPath, countFiles, countSymbolsPerFile); err != nil {
		fmt.Printf("error msg: %s", err.Error())
		return err
	}
	return nil
}
