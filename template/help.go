/*
   @Time : 2019-05-17 09:12
   @Author : frozenchen
   @File : help
   @Software: studio
*/
package template

var Help_template = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .Flags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
   {{if .Version}}{{if not .HideVersion}}
VERSION:
   {{.Version}}
   {{end}}{{end}}{{if len .Authors}}
AUTHOR(S):
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:{{range .Categories}}{{if .Name}}
  {{.Name}}{{ ":" }}{{end}}{{range .Commands}}
    {{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}{{end}}
{{end}}{{end}}{{if .Flags}}

COPYRIGHT:
   {{.Copyright}}
   {{end}}`
