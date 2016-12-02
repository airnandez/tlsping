package main

import (
	"os"
)

//  printVersion prints the version information about this application
func printVersion(f *os.File) {
	const template = `
{{.AppName}} {{.AppVersion}}

Compiler:
{{.Tab1}}{{.GoVersion}} ({{.Os}}/{{.Arch}})

Built on:
{{.Tab1}}{{.BuildTime}}

Author:
{{.Tab1}}Fabio Hernandez
{{.Tab1}}IN2P3/CNRS computing center, Lyon (France)

Source code and documentation:
{{.Tab1}}https://github.com/airnandez/{{.AppName}}
`
	render(template, tmplFields, f)
}
