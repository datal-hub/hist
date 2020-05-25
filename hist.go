package main

import (
	"hist/cli"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hist"
	app.Usage = "App for counting ascii symbols from files"
	app.Commands = []cli.Command{
		cmd.Init,
		cmd.Hist,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
