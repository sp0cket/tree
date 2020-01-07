package cmd

import (
	"container/list"
	"os"
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

// Generate tree node string.
func (node *FileNode) String() string {
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
		name = "\033[35m" + name + "\033[0m"
	}
	str += name
	return str
}
