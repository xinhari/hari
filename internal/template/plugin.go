package template

var (
	Plugin = `package main
{{if .Plugins}}
import ({{range .Plugins}}
	_ "github.com/xinhari/plugins/{{.}}"{{end}}
){{end}}
`
)
