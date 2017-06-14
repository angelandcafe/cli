// file: cmdRevert.go
// purpose: holo subCMD - revert
//
//      - description: roll a submission back
//            - usage: holo revert SUBMISSION [--train TRAIN1,TRAIN2,TRAIN3...] [-f/--force]
//          - example: holo revert osd-36
//                     holo revert --train ThundercatDeuce osd-36
//   - example output: This operation will affect the following build trains:
//
//                     - Avalanche: osd-37
//                     - Thundercat: osd-37
//
//                     Really revert to osd-36? (y/n)

package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func setRevertCmd() cli.Command {

	return cli.Command{
		Name:      "revert",
		Usage:     "roll a submission back",
		ArgsUsage: "SUBMISSION [--train TRAIN1,TRAIN2,TRAIN3...] [-f/--force]",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "train",
				Usage: "remove component from the given train(s)",
			},
			cli.BoolFlag{
				Name:  "force, f",
				Usage: "suppress prompts",
			},
		},
		Action: cmdRevertAction,
	}
}

func cmdRevertAction(c *cli.Context) error {
	// check input arguments
	err := checkNumArg(c, 1)
	if err != nil {
		return err
	}

	fmt.Println("In the function of revert:")
	fmt.Println("### cmd: holo revert SUBMISSION [--train TRAIN1,TRAIN2,TRAIN3...] [-f/--force] ###")
	fmt.Println("SUBMISSION: ", c.Args().Get(0))
	fmt.Println("train list: ", c.String("train"))
	fmt.Println("force:", c.Bool("force"))
	return nil
}
