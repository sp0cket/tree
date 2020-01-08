package cmd

import (
	"fmt"
	"os"
)

var (
	outputFile  = os.Stdout
	errorOutput = os.Stderr
)

func OutputFile(f *os.File) {
	outputFile = f
}

func PrintN(node FileNode) {
	if _, err := fmt.Fprintln(outputFile, node.String()); err != nil {
		PrintError(err)
		os.Exit(-1)
	}
}

func PrintError(err error, a ...interface{}) {
	fmt.Fprintln(errorOutput, err, a)
}

func Println(a ...interface{}) {
	if _, err := fmt.Fprintln(outputFile, a); err != nil {
		PrintError(err)
		os.Exit(-1)
	}
}
