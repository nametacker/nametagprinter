package main

import (
	"github.com/coderbyheart/nametagprinter"
	"github.com/wsxiaoys/terminal/color"
	"os"
)

func error(msg string) {
	color.Fprintln(os.Stderr, "@{!r}ERROR @{|}"+msg)
}

func main() {
	err := nametagprinter.Serve(nametagprinter.NewConfig())
	if err != nil {
		error(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
