// file: holoCommon.go
// purpose: contains all common functions for holo cli project

package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/urfave/cli"
)

func sendArgErr(c *cli.Context, str string) error {
	fmt.Print("Incorrect Usage: ", str, "!\n\n")
	cli.ShowSubcommandHelp(c)
	return errors.New(str)
}

func checkNumArg(c *cli.Context, argNum int) error {
	if c.NArg() != argNum {
		return sendArgErr(c, "incorrect number of arguments")
	}
	return nil
}

func checkFlagVal(c *cli.Context, strFlag string) error {
	if c.String(strFlag) == "" {
		err := "--" + strFlag + " flag is required"
		return sendArgErr(c, err)
	}

	// in case of an example:
	// incorrect command: holo artifacts put FILENAME --release -f
	// correct command: holo artifacts put FILENAME --release RELEASE -f
	// issue: --release flag takes "-f" as its argument
	errFlagArgContainsForceFlag, _ := regexp.MatchString("^(-)+f(orce)*$", c.String(strFlag))
	if errFlagArgContainsForceFlag {
		err := "--" + strFlag + " flag requires an argument"
		return sendArgErr(c, err)
	}

	// in case of an example:
	// incorrect command: holo artifacts put FILENAME WRONG_Extra_Arg --release
	// correct command: holo artifacts put FILENAME --release RELEASE
	// issue: --release flag takes "FILENAME" as its argument
	// => exclude "-f", "-force", "--f", "--force", "-"
	forceFlag, _ := regexp.MatchString("^(-)+f(orce)*$", os.Args[len(os.Args)-1])
	flagPrefix, _ := regexp.MatchString("^(-)+.+$", os.Args[len(os.Args)-1])
	if flagPrefix && !forceFlag {
		err := "--" + strFlag + " flag requires an argument"
		return sendArgErr(c, err)
	}
	return nil
}

func checkOptionalFlagVal(c *cli.Context, strFlag string) error {
	if c.String(strFlag) == "" {
		return nil
	}
	return checkFlagVal(c, strFlag)
}
