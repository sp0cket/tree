package model

import (
	"container/list"
	"fmt"
	"os"
	"tree/cmd"
)

// File tree symbol.
const (
	symbolNormalNode = "├─ "
	symbolEndNode    = "└─ "
	symbolConnNode   = "│  "
	symbolIdent      = "\t"
)

type FileNode struct {
	FileInfo     os.FileInfo
	Depth        int
	ConnectStack *list.List
	IsLastNode   bool
}

func (node FileNode) PrintNode() {
	str := ""
	connectIdx := node.ConnectStack.Front()
	for i := 0; i < node.Depth; i++ {
		if connectIdx != nil && connectIdx.Value == i {
			str += symbolConnNode
			connectIdx = connectIdx.Next()
		} else {
			str += symbolIdent
		}
	}
	if node.IsLastNode {
		str += symbolEndNode
	} else {
		str += symbolNormalNode
	}
	name := node.FileInfo.Name()
	if node.FileInfo.IsDir() {
		name += "/"
		if cmd.IsTerminal() {
			name = "\033[35m" + name + "\033[0m"
		}
	}
	str += name
	if _, err := fmt.Fprintln(cmd.Output, str); err != nil {
		cmd.PrintError(err)
		os.Exit(-1)
	}
}
