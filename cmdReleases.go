// file: cmdReleases.go
// purpose: holo subCMD - releases
//
//      - description: manage releases
//            - usage: holo releases subcommand
//      - subcommands: ls, show, cut, clone
//
//        => subCMD 1: ls
//             - desc: list releases within a build train
//            - usage: holo releases ls TRAIN
//          - example: holo releases ls --train Thundercat
//   - example output: Thundercat1A176
//                     Thundercat1A175
//                     Thundercat1A174
//                     Thundercat1A173
//                     Thundercat1A172
//                     Thundercat1A171
//                     Thundercat1A170
//                     ...
//
//        => subCMD 2: show
//             - desc: describe a given release
//            - usage: holo releases show RELEASE
//          - example: holo releases show Thundercat1A176
//   - example output: Release: Thundercat1A176
//                     Created: 2017-06-02 11:42AM PST
//                     Base: Thundercat1A175
//                     Submissions:
//                     - bash-3
//                     - jetpack-98
//                     - osd-38
//                     - zlib-1
//                     - ...
//
//        => subCMD 3: cut
//             - desc: cut a new release by flushing pending submissions,
//                     diffed against the release specified by --base
//            - usage: holo releases cut RELEASE [--base RELEASE] [-f/--force]
//          - example: holo releases cut Thundercat1A177 --base Thundercat1A176
//   - example output: This operation will bring in the following submissions:
//                     - jetpack-99
//                     - osd-39
//                     Really cut release Thundercat1A177? (y/n)
//
//        => subCMD 4: clone
//             - desc: create a new release based on an existing release
//                     (typically only useful after cloning a train)
//            - usage: holo releases clone RELEASE [--base RELEASE]
//          - example: holo releases clone ThundercatDeuce1A177 --base Thundercat1A177
//   - example output: Success!
//

package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func setReleasesCmd() cli.Command {

	return cli.Command{
		Name:  "releases",
		Usage: "manage a release",
		Subcommands: []cli.Command{
			{
				Name:      "ls",
				ArgsUsage: "TRAIN",
				Usage:     "list releases within a build train",
				Action:    lsReleasesAction,
			},
			{
				Name:      "show",
				ArgsUsage: "RELEASE",
				Usage:     "describe a given release",
				Action:    showReleasesAction,
			},
			{
				Name:      "cut",
				ArgsUsage: "RELEASE [--base RELEASE] [-f/--force]",
				Usage:     "cut a new release by flushing pending submissions",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "base",
						Usage: "base release",
					},
					cli.BoolFlag{
						Name:  "force, f",
						Usage: "suppress prompts",
					},
				},
				Action: cutReleasesAction,
			},
			{
				Name:      "clone",
				ArgsUsage: "RELEASE [--base RELEASE]",
				Usage:     "create a new release based on an existing release",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "base",
						Usage: "base release",
					},
				},
				Action: cloneReleasesAction,
			},
		},
	}
}

// Subcommand: ls
func lsReleasesAction(c *cli.Context) error {
	// check arguments
	err := checkNumArg(c, 1)
	if err != nil {
		return err
	}

	fmt.Println("In the function of release ls:")
	fmt.Println("### cmd: holo releases ls TRAIN ###")
	fmt.Println("TRAIN: ", c.Args().Get(0))
	return nil
}

// Subcommand: show
func showReleasesAction(c *cli.Context) error {
	// check arguments
	err := checkNumArg(c, 1)
	if err != nil {
		return err
	}

	fmt.Println("In the function of release show:")
	fmt.Println("### cmd: holo releases show RELEASE ###")
	fmt.Println("RELEASE: ", c.Args().Get(0))
	return nil
}

// Subcommand: cut
func cutReleasesAction(c *cli.Context) error {
	// check arguments
	err := cutReleasesCheckArg(c)
	if err != nil {
		return err
	}

	fmt.Println("In the function of release cut:")
	fmt.Println("### cmd: holo releases cut RELEASE [--base RELEASE] [-f/--force] ###")
	fmt.Println("RELEASE: ", c.Args().Get(0))
	fmt.Println("base-RELEASE: ", c.String("base"))
	fmt.Println("force:", c.Bool("force"))
	return nil
}

// Subcommand: clone
func cloneReleasesAction(c *cli.Context) error {
	// check arguments
	err := cloneReleasesCheckArg(c)
	if err != nil {
		return err
	}

	fmt.Println("In the function of release clone:")
	fmt.Println("### cmd: holo releases clone RELEASE [--base RELEASE] ###")
	fmt.Println("RELEASE: ", c.Args().Get(0))
	fmt.Println("base-RELEASE: ", c.String("base"))
	return nil
}

// check input arguments:
func cutReleasesCheckArg(c *cli.Context) error {
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

func cloneReleasesCheckArg(c *cli.Context) error {
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
