// file: cmdSubmit.go
// purpose: holo subCMD - submit
//
//      - description: submit a new version of a component
//            - usage: holo submit SUBMISSION SRCROOT [--train TRAIN1,TRAIN2,TRAIN3...] [-f/--force] [--notes NOTES|-]
//          - example: holo submit osd-38 . --notes - <<-EOF
//                     holo submit osd-37.1 . --train ThundercatDeuce
//                     cat notes.txt | holo submit osd-38 . -train train1,2 -notes -
//   - example output: osd-38 fixes a couple crashers:
//
//                     VHAP-89: Some stupid crash
//                     VHAP-90: Some other stupid crash
//                     EOF
//                     This operation will affect the following build trains:
//
//                     - Avalanche: osd-37
//                     - Thundercat: osd-37
//
//                     Really submit osd-38? (y/n) y
//
//                     Tagging git repository...
//                     Pushing git tags...
//                     Archiving HEAD...
//                     Submitting sources...
//                     Done!

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/urfave/cli"
)

func setSubmitCmd() cli.Command {

	return cli.Command{
		Name:      "submit",
		Usage:     "submit a new version of a component",
		ArgsUsage: "SUBMISSION SRCROOT [--train TRAIN1,TRAIN2,TRAIN3...] [-f/--force] [--notes NOTES|-]",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "train",
				Usage: "remove component from the given train(s)",
			},
			cli.StringFlag{
				Name:  "notes",
				Usage: "custom notes for the submission",
			},
			cli.BoolFlag{
				Name:  "force, f",
				Usage: "suppress prompts",
			},
		},
		Action: cmdSubmitAction,
	}
}

func cmdSubmitAction(c *cli.Context) error {
	notes := []string{}

	// check input arguments
	err := submitCheckArg(c, &notes)
	if err != nil {
		return err
	}

	fmt.Println("In the function of submit:")
	fmt.Println("### cmd: holo submit SUBMISSION SRCROOT [--train TRAIN1,TRAIN2,TRAIN3...] [-f/--force] [--notes NOTES|-]###")
	fmt.Println("SUBMISSION: ", c.Args().Get(0))
	fmt.Println("SRCROOT: ", c.Args().Get(1))
	fmt.Println("train list: ", c.String("train"))
	fmt.Println("NOTES: ", notes)
	fmt.Println("force:", c.Bool("force"))
	return nil
}

func submitCheckArg(c *cli.Context, notes *[]string) error {
	var err error

	err = checkNumArg(c, 2)
	if err != nil {
		return err
	}

	err = checkOptionalFlagVal(c, "train")
	if err != nil {
		return err
	}

	err = checkOptionalFlagNotesVal(c, notes)
	if err != nil {
		return err
	}

	return nil
}

func checkOptionalFlagNotesVal(c *cli.Context, notes *[]string) error {
	err := checkOptionalFlagVal(c, "notes")
	if err != nil {
		return err
	}

	// in case of "--notes -", read from os.Stdin
	match, _ := regexp.MatchString("-", os.Args[len(os.Args)-1])
	if match {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "." {
				break
			}
			*notes = append(*notes, line)
		}
	} else {
		*notes = append(*notes, c.String("notes"))
	}
	return nil
}
