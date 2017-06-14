// file: cmdNotes.go
// purpose: holo subCMD - notes
//
//      - description: show the release notes for a submission
//            - usage: holo notes SUBMISSION
//          - example: holo notes osd-37
//   - example output: Submitted by: steve@bowerswilkins.com
//                     Date: 2017-06-01 12:34PM PST
//
//                     osd-37 adds new features and fixes some bugs.
//
//                     New features:
//
//                     VHAP-27: Better tiling
//                     VHAP-39: Notifications when song changes happen
//
//                     Bugs fixed:
//
//                     VHAP-47: Crash when trying to tile three devices
//                     VHAP-88: Crash when entering Settings

package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func setNotesCmd() cli.Command {

	return cli.Command{
		Name:      "notes",
		UsageText: "holo notes SUBMISSION",
		Usage:     "show the release notes for a submission",
		ArgsUsage: "SUBMISSION",
		Action:    cmdNotesAction,
	}
}

func cmdNotesAction(c *cli.Context) error {
	// check input arguments
	err := checkNumArg(c, 1)
	if err != nil {
		return err
	}

	fmt.Println("In the function of notes:")
	fmt.Println("### cmd: holo notes SUBMISSION ###")
	fmt.Println("SUBMISSION: ", c.Args().Get(0))
	return nil
}
