package dir

import (
	"container/list"
	"io/ioutil"
	"path"
	"strings"
	"tree/cmd"
)

var (
	FlagAll   bool
	FlagDepth int
)

func Visit(visitPath string) error {
	return visit(visitPath, 0, list.New())
}

func visit(visitPath string, depth int, connectStack *list.List) error {
	filesInfo, err := ioutil.ReadDir(visitPath)
	if err == nil {
		fileCount := len(filesInfo) - 1
		for idx, file := range filesInfo {
			if FlagAll == false && strings.HasPrefix(file.Name(), ".") {
				continue
			}
			name := file.Name()
			if file.IsDir() {
				name += "/"

			}
			cmd.PrintN(cmd.FileNode{
				FileInfo:     file,
				Depth:        depth,
				ConnectStack: connectStack,
				IsLastNode:   idx == fileCount,
			})
			if file.IsDir() {
				if FlagDepth > 0 && depth+1 >= FlagDepth {
					continue
				}
				nextConnStack := connectStack
				if idx < fileCount {
					nextConnStack.PushBack(depth)
				}
				if err := visit(path.Join(visitPath, file.Name()), depth+1, nextConnStack); err != nil {
					return err
				}
			}
		}
	}
	return err
}
