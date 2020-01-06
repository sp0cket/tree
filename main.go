package main

import (
	"flag"
	"fmt"
	"os"
	"tree/dir"
)

var flagAll = flag.Bool("a", false, "List all files")
var flagDepth = flag.Int("depth", 0, "Max depth")

func main() {
	flag.Parse()
	dir.FlagAll = *flagAll
	dir.FlagDepth = *flagDepth
	currentDir, _ := os.Getwd()
	fmt.Println(currentDir)
	err := dir.Visit(currentDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
