package cmd

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"os"
)

var (
	Output      = os.Stdout
	errorOutput = os.Stderr
)

func PrintError(err error, a ...interface{}) {
	fmt.Fprintln(errorOutput, err, a)
}

func Println(a ...interface{}) {
	if _, err := fmt.Fprintln(Output, a); err != nil {
		PrintError(err)
		os.Exit(-1)
	}
}

func IsTerminal() bool {
	return isatty.IsTerminal(Output.Fd())
}
