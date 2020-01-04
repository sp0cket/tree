package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	symbolNormalNode = "├ "
	symbolEndNode    = "└ "
	symbolConnNode   = "│ "
	symbolIdent      = "\t"
)

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
			if file.IsDir() {
				nextConnStack := connectStack
				if idx < fileCount {
					nextConnStack = append(nextConnStack, depth)
					printNode(depth, symbolNormalNode, file.Name(), connectStack)
				} else {
					printNode(depth, symbolEndNode, file.Name(), connectStack)
				}
				if err := visit(path.Join(visitPath, file.Name()), depth+1, nextConnStack); err != nil {
					return err
				}
			} else {
				if idx < fileCount {
					printNode(depth, symbolNormalNode, file.Name(), connectStack)
				} else {
					printNode(depth, symbolEndNode, file.Name(), connectStack)
				}
			}
		}
	}
	return err
}

func main() {
	currentDir, _ := os.Getwd()
	visit(currentDir, 0, make([]int, 0))
}
