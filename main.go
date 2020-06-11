package main

import (
	"flag"
	"fmt"
	"os"
	"tree/cmd"
	"tree/dir"
)

var flagAll = flag.Bool("a", false, "List all files, default false")
var flagDepth = flag.Int("depth", 0, "Max depth, default all")
var flagPath = flag.String("p", ".", "Working path, default is current path")
var flagOutput = flag.String("o", "", "Set output file, default to console")

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
			cmd.Output = file
		} else if os.IsNotExist(err) {
			file, err := os.Create(*flagOutput)
			if err != nil {
				cmd.PrintError(err)
				os.Exit(-1)
			}
			cmd.Output = file
		} else {
			if err != nil {
				cmd.PrintError(err)
				os.Exit(-1)
			}
		}
	}
	cmd.Println(*flagPath)
	info, err := dir.Walk(*flagPath)
	if err != nil {
		cmd.PrintError(err)
		os.Exit(-1)
	}
	fmt.Fprintln(cmd.Output, "\n", info.String())
}
