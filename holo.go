// file: holo.go
// purpose: for holo cli - PLAT-314

package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func init() {
	setInit()
}

func main() {
	app := cli.NewApp()

	app.Name = "holo"
	app.Usage = "a tool for managing CosmOS builds"
	app.Version = "1.0"
	app.UsageText = "holo command [arguments]"

	app.Commands = setCmdList()
	sort.Sort(cli.CommandsByName(app.Commands))

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Printf("Incorrect Usage: No matching command '%s'\n\n", command)
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	app.Run(os.Args)

}
