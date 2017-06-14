// file: cmdComponents.go
// purpose: holo subCMD - components
//
//      - description: manage components
//            - usage: holo components subcommand
//      - subcommands: ls, show, create, link, unlink
//
//        => subCMD 1: ls
//             - desc: list all known components
//            - usage: holo components ls
//          - example: holo components ls
//   - example output: autoconf
//                     automake
//                     bash
//                     jetpack
//                     osd
//                     zlib
//                     ...
//
//        => subCMD 2: show
//             - desc: describe a given component
//            - usage: holo components show COMPONENT
//          - example: holo components show osd
//   - example output: Component: osd
//                     Created: 2015-03-09 09:00AM PST
//                     Owner: steve@bowerswilkins.com
//                     Repo: git@github.com:bowerswilkins/app.git
//                     Train membership:
//                     - Avalanche
//                     - Thundercat
//                     - ThundercatDeuce
//
//        => subCMD 3: create
//             - desc: create a new component
//            - usage: holo components create COMPONENT --owner EMAIL [--repo URL]
//          - example: holo components create bash --owner cody@bowerswilkins.com --repo git@github.com:bw-oss/bash.git
//   - example output: Success!
//
//        => subCMD 4: link
//             - desc: make a component a member of a train
//            - usage: holo components link COMPONENT --train TRAIN
//          - example: holo components link bash --train Thundercat
//   - example output: Success!
//
//        => subCMD 5: unlink
//             - desc: remove a component from a train
//            - usage: holo components unlink COMPONENT --train TRAIN [-f/--force]
//          - example: holo components unlink bash --train Thundercat
//   - example output: Are you sure? (y/n)
//

package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func setComponentsCmd() cli.Command {

	return cli.Command{
		Name:  "components",
		Usage: "manage a component",
		Subcommands: []cli.Command{
			{
				Name:      "ls",
				ArgsUsage: " ",
				Usage:     "list all known components",
				Action:    lsComponentsAction,
			},
			{
				Name:      "show",
				ArgsUsage: "COMPONENT",
				Usage:     "describe a given component",
				Action:    showComponentsAction,
			},
			{
				Name:      "create",
				ArgsUsage: "COMPONENT --owner EMAIL [--repo URL]",
				Usage:     "create a new component",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "owner",
						Usage: "the email of a owner",
					},
					cli.StringFlag{
						Name:  "repo",
						Usage: "URL for the repo",
					},
				},
				Action: createComponentsAction,
			},
			{
				Name:      "link",
				ArgsUsage: "COMPONENT --train TRAIN",
				Usage:     "make a component a member of a train",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "train",
						Usage: "name of a build train",
					},
				},
				Action: linkComponentsAction,
			},
			{
				Name:      "unlink",
				ArgsUsage: "COMPONENT --train TRAIN [-f/--force]",
				Usage:     "remove a component from a train",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "train",
						Usage: "name of a build train",
					},
					cli.BoolFlag{
						Name:  "force, f",
						Usage: "suppress prompts",
					},
				},
				Action: unlinkComponentsAction,
			},
		},
	}
}

// Subcommand: ls
func lsComponentsAction(c *cli.Context) error {
	// check arguments
	err := checkNumArg(c, 0)
	if err != nil {
		return err
	}
	fmt.Println("In the function of components ls:")
	fmt.Println("### cmd: holo components ls ###")

	return nil
}

// Subcommand: show
func showComponentsAction(c *cli.Context) error {
	// check arguments
	err := checkNumArg(c, 1)
	if err != nil {
		return err
	}

	fmt.Println("In the function of components show:")
	fmt.Println("### cmd: holo components show COMPONENT ###")
	fmt.Println("COMPONENT: ", c.Args().Get(0))

	return nil
}

// Subcommand: create
func createComponentsAction(c *cli.Context) error {
	// check arguments
	err := createComponentsCheckArg(c)
	if err != nil {
		return err
	}

	fmt.Println("In the function of components create:")
	fmt.Println("### cmd: holo components create COMPONENT --owner EMAIL [--repo URL] ###")
	fmt.Println("COMPONENT: ", c.Args().Get(0))
	fmt.Println("owner-EMAIL: ", c.String("owner"))
	fmt.Println("repo-URL: ", c.String("repo"))
	return nil
}

// Subcommand: link
func linkComponentsAction(c *cli.Context) error {
	// check arguments
	err := linkComponentsCheckArg(c)
	if err != nil {
		return err
	}

	fmt.Println("In the function of components link:")
	fmt.Println("### cmd: holo components link COMPONENT --train TRAIN ###")
	fmt.Println("COMPONENT: ", c.Args().Get(0))
	fmt.Println("train: ", c.String("train"))
	return nil
}

// Subcommand: unlink
func unlinkComponentsAction(c *cli.Context) error {
	// check arguments (use the same checkArg function as linkComponentsAction)
	err := linkComponentsCheckArg(c)
	if err != nil {
		return err
	}

	fmt.Println("In the function of components unlink:")
	fmt.Println("### cmd: holo components unlink COMPONENT --train TRAIN [-f/--force] ###")
	fmt.Println("COMPONENT: ", c.Args().Get(0))
	fmt.Println("train: ", c.String("train"))
	fmt.Println("force:", c.Bool("force"))
	return nil
}

// check input arguments:
func createComponentsCheckArg(c *cli.Context) error {
	err := checkNumArg(c, 1)
	if err != nil {
		return err
	}
	err = checkFlagVal(c, "owner")
	if err != nil {
		return err
	}
	err = checkOptionalFlagVal(c, "repo")
	if err != nil {
		return err
	}
	return nil
}

func linkComponentsCheckArg(c *cli.Context) error {
	err := checkNumArg(c, 1)
	if err != nil {
		return err
	}
	err = checkFlagVal(c, "train")
	if err != nil {
		return err
	}
	return nil
}
