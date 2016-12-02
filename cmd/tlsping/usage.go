package main

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"text/template"
)

type usageType uint32

const (
	usageShort usageType = iota
	usageLong
)

func printUsage(f *os.File, kind usageType) {
	const template = `
USAGE:
{{.Tab1}}{{.AppName}} [-c count] [-tcponly] [-json] [-ca=<file>]
{{.Tab1}}{{.AppNameFiller}} [-insecure] <server address>
{{.Tab1}}{{.AppName}} -help
{{.Tab1}}{{.AppName}} -version
{{if eq .UsageVersion "short"}}
Use '{{.AppName}} -help' to get detailed information about options and examples
of usage.{{else}}

DESCRIPTION:
{{.Tab1}}{{.AppName}} is a basic tool to measure the time required to establish a
{{.Tab1}}TCP connection and perform a TLS handshake with a remote server.
{{.Tab1}}It reports summary statistics of the measurements obtained over a number
{{.Tab1}}of successful connections.

{{.Tab1}}The address of the remote server, i.e. <server address>, is of the form
{{.Tab1}}'host:port', for instance 'mail.google.com:443'.

OPTIONS:
{{.Tab1}}-c count
{{.Tab2}}Perform count concurrent measurements.
{{.Tab2}}Default: {{.DefaultCount}}

{{.Tab1}}-tcponly
{{.Tab2}}Establish the TCP connection with the remote server but do not perform
{{.Tab2}}the TLS handshake.

{{.Tab1}}-insecure
{{.Tab2}}Don't verify the validity of the server certificate. Only relevant when
{{.Tab2}}TLS handshake is performed (see '-tcponly' option).
{{.Tab2}}This option is intended to be used for measuring times for connecting
{{.Tab2}}to servers which use custom not widely trusted certificates, for
{{.Tab2}}instance, development servers using self-signed certificates.

{{.Tab1}}-ca <file>
{{.Tab2}}PEM-formatted file containing the CA certificates this program trusts.
{{.Tab2}}Default: use this host's CA certificate store.

{{.Tab1}}-json
{{.Tab2}}Format the result in JSON format and print to standard output. Reported
{{.Tab2}}times are understood in seconds.

{{.Tab1}}-help
{{.Tab2}}Prints this help

{{.Tab1}}-version
{{.Tab2}}Show detailed version information about this application

EXAMPLE:
{{.Tab1}}{{.AppName}} -tcponly mail.google.com:443
{{end}}
`
	if kind == usageLong {
		tmplFields["UsageVersion"] = "long"
	}
	tmplFields["DefaultCount"] = fmt.Sprintf("%d", defaultIterations)
	render(template, tmplFields, f)
}

func render(tpl string, fields map[string]string, out io.Writer) {
	minWidth, tabWidth, padding := 4, 4, 0
	tabwriter := tabwriter.NewWriter(out, minWidth, tabWidth, padding, byte(' '), 0)
	templ := template.Must(template.New("").Parse(tpl))
	templ.Execute(tabwriter, fields)
	tabwriter.Flush()
}
