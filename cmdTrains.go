// file: cmdTrains.go
// purpose: holo subCMD - trains
//
//      - description: manage trains
//            - usage: holo trains subcommand
//      - subcommands: ls, show, create
//
//        => subCMD 1: ls
//             - desc: list all known build trains
//            - usage: holo trains ls
//          - example: holo trains ls
//   - example output: Avalanche
//                     Thundercat
//                     ThundercatDeuce
//
//        => subCMD 2: show
//             - desc: describe a given build train
//            - usage: holo trains show TRAIN
//          - example: holo trains show Thundercat
//   - example output: Train: Thundercat
//                     Created: 2015-03-09 09:00AM PST
//                     Latest release: Thundercat1A177
//                     Component membership:
//                     - bash
//                     - jetpack
//                     - osd
//                     - zlib
//                     - ...
//
//        => subCMD 3: create
//             - desc: create a new build train, or clone an existing build train
//            - usage: holo trains create TRAIN [--base TRAIN]
//          - example: holo trains create Avalanche
//                     holo trains create ThundercatDeuce --base Thundercat
//   - example output: Success!
//

package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func setTrainsCmd() cli.Command {

	return cli.Command{
		Name:  "trains",
		Usage: "manage a train",
		Subcommands: []cli.Command{
			{
				Name:      "ls",
				ArgsUsage: " ",
				Usage:     "list all known build trains",
				Action:    lsTrainsAction,
			},
			{
				Name:      "show",
				ArgsUsage: "TRAIN",
				Usage:     "describe a given build train",
				Action:    showTrainsAction,
			},
			{
				Name:      "create",
				ArgsUsage: "TRAIN [--base TRAIN]",
				Usage:     "create a new build train, or clone an existing build train",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "base",
						Usage: "base on existing build train",
					},
				},
				Action: createTrainsAction,
			},
		},
	}
}

// Subcommand: ls
func lsTrainsAction(c *cli.Context) error {
	// check arguments
	err := checkNumArg(c, 0)
	if err != nil {
		return err
	}
	fmt.Println("In the function of trains ls:")
	fmt.Println("### cmd: holo trains ls ###")

	return nil
}

// Subcommand: show
func showTrainsAction(c *cli.Context) error {
	// check arguments
	err := checkNumArg(c, 1)
	if err != nil {
		return err
	}

	fmt.Println("In the function of trains show:")
	fmt.Println("### cmd: holo trains show TRAIN ###")
	fmt.Println("TRAIN: ", c.Args().Get(0))

	return nil
}

// Subcommand: create
func createTrainsAction(c *cli.Context) error {
	// check arguments
	err := createTrainsCheckArg(c)
	if err != nil {
		return err
	}

	fmt.Println("In the function of trains create:")
	fmt.Println("### cmd: holo trains create TRAIN [--base TRAIN] ###")
	fmt.Println("TRAIN: ", c.Args().Get(0))
	fmt.Println("base-TRAIN: ", c.String("base"))
	return nil
}

// check input arguments:
func createTrainsCheckArg(c *cli.Context) error {
	err := checkNumArg(c, 1)
	if err != nil {
		return err
	}
	err = checkOptionalFlagVal(c, "base")
	if err != nil {
		return err
	}
	return nil
}
