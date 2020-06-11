package cmd

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"os"
	"tree/model"
)

var (
	output      = os.Stdout
	errorOutput = os.Stderr
)

func SetOutput(f *os.File) {
	output = f
}

func GetOutput() *os.File {
	return output
}

func PrintN(node model.FileNode) {
	if _, err := fmt.Fprintln(output, node.String()); err != nil {
		PrintError(err)
		os.Exit(-1)
	}
}

func PrintError(err error, a ...interface{}) {
	fmt.Fprintln(errorOutput, err, a)
}

func Println(a ...interface{}) {
	if _, err := fmt.Fprintln(output, a); err != nil {
		PrintError(err)
		os.Exit(-1)
	}
}

func IsTerminal() bool {
	return isatty.IsTerminal(output.Fd())
}
