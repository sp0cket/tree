package cmd

import (
	"container/list"
	"os"
)

const (
	symbolNormalNode = "├ "
	symbolEndNode    = "└ "
	symbolConnNode   = "│ "
	symbolIdent      = "\t"
)

type FileNode struct {
	FileInfo     os.FileInfo
	Depth        int
	ConnectStack *list.List
	IsLastNode   bool
}

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
	str += node.FileInfo.Name()
	return str
}