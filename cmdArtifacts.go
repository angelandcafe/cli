// file: cmdArtifacts.go
// purpose: holo subCMD - artifacts
//
//      - description: manage artifacts
//            - usage: holo artifacts subcommand
//      - subcommands: ls, put, get
//
//        => subCMD 1: ls
//             - desc: list artifacts within a release
//            - usage: holo artifacts ls RELEASE
//          - example: holo artifacts ls Thundercat1A176
//   - example output: Cake1,1_DVT_1.0_Thundercat1A176_RestoreImage.swpkg
//                     Cake1,1_DVT2_1.0_Thundercat1A176_RestoreImage.swpkg
//                     Cake1,1_DVT_1.0_Thundercat1A176_USBImage.img
//                     Cake1,1_DVT2_1.0_Thundercat1A176_USBImage.img
//                     Cake1,1_DVT_1.0_Thundercat1A176_rootfs.img
//                     Cake1,1_DVT2_1.0_Thundercat1A176_rootfs.img
//                     Thundercat1A176_SDKRoot.txz
//
//        => subCMD 2: put
//             - desc: upload artifacts for a release
//            - usage: holo artifacts put FILENAME --release RELEASE [-f/--force]
//          - example: holo artifacts put ./foo.tar.gz --release Thundercat1A176
//   - example output: Uploading...
//                     [============================>                          ] 50% (30MB/s; ETA: 1s)
//
//        => subCMD 3: get
//             - desc: download artifacts for a release
//            - usage: holo artifacts get FILENAME DESTINATION --release RELEASE
//          - example: holo artifacts get foo.tar.gz /tmp/Thundercat1A176_foo.tar.gz --release Thundercat1A176
//   - example output: Downloading...
//                     [============================>                          ] 50% (30MB/s; ETA: 1s)

package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func setArtifactsCmd() cli.Command {

	return cli.Command{
		Name:  "artifacts",
		Usage: "manage artifacts for a release",
		Subcommands: []cli.Command{
			{
				Name:      "ls",
				ArgsUsage: "RELEASE",
				Usage:     "list artifacts within a release",
				Action:    lsArtifactsAction,
			},
			{
				Name:      "put",
				ArgsUsage: "FILENAME --release RELEASE [-f/--force]",
				Usage:     "upload artifacts for a release",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "release",
						Usage: "the name of a release within a build train",
					},
					cli.BoolFlag{
						Name:  "force, f",
						Usage: "suppress prompts",
					},
				},
				Action: putArtifactsAction,
			},
			{
				Name:      "get",
				ArgsUsage: "FILENAME DESTINATION --release RELEASE",
				Usage:     "download artifacts for a release",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "release",
						Usage: "the name of a release within a build train",
					},
				},
				Action: getArtifactsAction,
			},
		},
	}
}

// Subcommand: ls
func lsArtifactsAction(c *cli.Context) error {
	// check arguments
	err := lsArtifactsCheckArg(c)
	if err != nil {
		return err
	}
	fmt.Println("In the function of artifacts ls:")
	fmt.Println("### cmd: holo artifacts ls RELEASE ###")
	fmt.Print("RELEASE: ", c.Args().Get(0), "\n\n")
	return nil
}

// Subcommand: put
func putArtifactsAction(c *cli.Context) error {
	// check arguments
	err := putArtifactsCheckArg(c)
	if err != nil {
		return err
	}
	fmt.Println("In the function of artifacts put:")
	fmt.Println("### cmd: holo artifacts put FILENAME --release RELEASE [-f/--force] ###")
	fmt.Println("FILENAME:", c.Args().Get(0))
	fmt.Println("RELEASE: ", c.String("release"))
	fmt.Println("force:", c.Bool("force"))
	return nil
}

// Subcommand: get
func getArtifactsAction(c *cli.Context) error {
	// check arguments
	err := getArtifactsCheckArg(c)
	if err != nil {
		return err
	}

	fmt.Println("In the function of artifacts get:")
	fmt.Println("### cmd: holo artifacts get FILENAME DESTINATION --release RELEASE ###")
	fmt.Println("FILENAME: ", c.Args().Get(0))
	fmt.Println("DESTINATION: ", c.Args().Get(1))
	fmt.Println("RELEASE: ", c.String("release"))
	return nil
}

// check input arguments:
func lsArtifactsCheckArg(c *cli.Context) error {
	err := checkNumArg(c, 1)
	if err != nil {
		return err
	}
	return nil
}

func putArtifactsCheckArg(c *cli.Context) error {
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

func getArtifactsCheckArg(c *cli.Context) error {
	err := checkNumArg(c, 2)
	if err != nil {
		return err
	}
	err = checkFlagVal(c, "release")
	if err != nil {
		return err
	}
	return nil
}
