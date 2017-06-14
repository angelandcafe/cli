// file: holoInit.go
// purpose: for holo cli - PLAT-314
//          for global settings

package main

import "github.com/urfave/cli"

func setCmdList() []cli.Command {

	cmdList := make([]cli.Command, 0)
	cmdList = append(cmdList, setNotesCmd())
	cmdList = append(cmdList, setRevertCmd())
	cmdList = append(cmdList, setSubmitCmd())
	cmdList = append(cmdList, setWhichCmd())
	cmdList = append(cmdList, setArtifactsCmd())
	cmdList = append(cmdList, setComponentsCmd())
	cmdList = append(cmdList, setReleasesCmd())
	cmdList = append(cmdList, setTrainsCmd())

	return cmdList
}

func setInit() {
	cli.AppHelpTemplate = `{{if .UsageText}}--------------------------------------------
| {{.Name}}{{if .Usage}} - {{.Usage}}{{end}} |
--------------------------------------------
{{end}}USAGE:
  {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} command {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments]{{end}}{{end}}
  {{if .VisibleCommands}}
COMMANDS:{{range .VisibleCategories}}{{if .Name}}
{{.Name}}:{{end}}{{range .VisibleCommands}}
  {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}
       {{end}}
Use "{{.HelpName}} [command] -help" for more information about a command.
       {{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
  {{range $index, $option := .VisibleFlags}}{{if $index}}
  {{end}}{{$option}}{{end}}{{end}}

`
	cli.CommandHelpTemplate = `USAGE:
  {{.HelpName}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments]{{end}}{{if .Category}}

CATEGORY:
  {{.Category}}{{end}}{{if .Description}}

DESCRIPTION:
  {{.Description}}{{end}}{{if .VisibleFlags}}

OPTIONS:
  {{range .VisibleFlags}}{{.}}
  {{end}}{{end}}

`
	cli.SubcommandHelpTemplate = `USAGE:
   {{.HelpName}} command {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments]{{end}}

COMMANDS:{{range .VisibleCategories}}{{if .Name}}
  {{.Name}}:{{end}}{{range .VisibleCommands}}
   {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}
{{end}}
Use "{{.HelpName}} [command] -help" for more information about a command.
{{if .VisibleFlags}}
OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}

`
}
