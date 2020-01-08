package main

import (
	"flag"
	"os"
	"tree/cmd"
	"tree/dir"
)

var flagAll = flag.Bool("a", false, "List all files")
var flagDepth = flag.Int("depth", 0, "Max depth")
var flagPath = flag.String("p", ".", "Working path, default is current path")
var flagOutput = flag.String("o", "", "Output file, default to console")

func main() {
	flag.Parse()
	dir.FlagAll = *flagAll
	dir.FlagDepth = *flagDepth
	if len(*flagOutput) > 0 {
		if _, err := os.Stat(*flagOutput); err == nil {
			file, err := os.Open(*flagOutput)
			if err != nil {
				cmd.PrintError(err)
				os.Exit(-1)
			}
			cmd.OutputFile(file)
		} else if os.IsNotExist(err) {
			file, err := os.Create(*flagOutput)
			if err != nil {
				cmd.PrintError(err)
				os.Exit(-1)
			}
			cmd.OutputFile(file)
		} else {
			if err != nil {
				cmd.PrintError(err)
				os.Exit(-1)
			}
		}
	}
	cmd.Println(*flagPath)
	err := dir.Visit(*flagPath)
	if err != nil {
		cmd.PrintError(err)
		os.Exit(-1)
	}
}
