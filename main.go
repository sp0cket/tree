package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	symbolNormalNode = "├ "
	symbolEndNode    = "└ "
	symbolConnNode   = "│ "
	symbolIdent      = "\t"
)

var flagAll = flag.Bool("a", false, "List all files")
var flagDepth = flag.Int("depth", 0, "Max depth")

func printNode(depth int, symbol, name string, connectStack []int) {
	node := ""
	connStackIdx := -1
	if len(connectStack) > 0 {
		connStackIdx = 0
	}
	for i := 0; i < depth; i++ {
		if connectStack[connStackIdx] == i {
			node += symbolConnNode
			if connStackIdx < len(connectStack)-1 {
				connStackIdx += 1
			}
		} else {
			node += symbolIdent
		}
	}
	node += symbol + name
	fmt.Println(node)
}

func visit(visitPath string, depth int, connectStack []int) error {
	filesInfo, err := ioutil.ReadDir(visitPath)
	if err == nil {
		fileCount := len(filesInfo) - 1
		for idx, file := range filesInfo {
			if *flagAll == false && strings.HasPrefix(file.Name(), ".") {
				continue
			}
			name := file.Name()
			if file.IsDir() {
				name += "/"
			}
			if idx < fileCount {
				printNode(depth, symbolNormalNode, name, connectStack)
			} else {
				printNode(depth, symbolEndNode, name, connectStack)
			}
			if file.IsDir() {
				if *flagDepth > 0 && depth+1 >= *flagDepth {
					continue
				}
				nextConnStack := connectStack
				if idx < fileCount {
					nextConnStack = append(nextConnStack, depth)
				}
				if err := visit(path.Join(visitPath, file.Name()), depth+1, nextConnStack); err != nil {
					return err
				}
			}
		}
	}
	return err
}

func main() {
	flag.Parse()
	currentDir, _ := os.Getwd()
	fmt.Println(currentDir)
	err := visit(currentDir, 0, make([]int, 0))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
