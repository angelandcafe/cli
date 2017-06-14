// file: cmdWhich.go
// purpose: holo subCMD - which
//
//      - description: show which version of a component appeared in a release
//            - usage: holo which COMPONENT --release RELEASE
//          - example: holo which osd --release Thundercat1A176
//   - example output: osd-35

package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func setWhichCmd() cli.Command {

	return cli.Command{
		Name:      "which",
		Usage:     "show which version of a component appeared in a release",
		ArgsUsage: "COMPONENT --release RELEASE",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "release",
				Usage: "the name of a release within a build train",
			},
		},
		Action: cmdWhichAction,
	}
}

func cmdWhichAction(c *cli.Context) error {
	// check arguments
	err := whichCheckArg(c)
	if err != nil {
		return err
	}

	fmt.Println("In the function of which:")
	fmt.Println("### cmd: holo which COMPONENT --release RELEASE ###")
	fmt.Println("COMPONENT: ", c.Args().Get(0))
	fmt.Println("release: ", c.String("release"))
	return nil
}

func whichCheckArg(c *cli.Context) error {
	err := checkNumArg(c, 1)
	if err != nil {
		return err
	}

	err = checkFlagVal(c, "release")
	if err != nil {
		return err
	}
	return nil
}
