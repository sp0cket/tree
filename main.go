package main

import (
	"flag"
	"fmt"
	"os"
	"tree/dir"
)

var flagAll = flag.Bool("a", false, "List all files")
var flagDepth = flag.Int("depth", 0, "Max depth")
var flagPath = flag.String("p", ".", "Working path, default is current path")

func main() {
	flag.Parse()
	dir.FlagAll = *flagAll
	dir.FlagDepth = *flagDepth
	fmt.Println(*flagPath)
	err := dir.Visit(*flagPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
