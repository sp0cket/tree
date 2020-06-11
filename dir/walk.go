package dir

import (
	"container/list"
	"io/ioutil"
	"path"
	"strings"
	"tree/model"
)

var (
	FlagAll   bool
	FlagDepth int
)

func Walk(walkPath string) (*model.TotalInfo, error) {
	return walk(walkPath, 0, list.New())
}

func walk(walkPath string, depth int, connectStack *list.List) (*model.TotalInfo, error) {
	filesInfo, err := ioutil.ReadDir(walkPath)
	totalInfo := &model.TotalInfo{}
	if err == nil {
		fileCount := len(filesInfo) - 1
		for idx, file := range filesInfo {
			if FlagAll == false && strings.HasPrefix(file.Name(), ".") {
				continue
			}
			model.FileNode{
				FileInfo:     file,
				Depth:        depth,
				ConnectStack: connectStack,
				IsLastNode:   idx == fileCount,
			}.PrintNode()
			if file.IsDir() {
				if FlagDepth > 0 && depth+1 >= FlagDepth {
					continue
				}
				nextConnStack := connectStack
				if idx < fileCount {
					nextConnStack.PushBack(depth)
				}
				if info, err := walk(path.Join(walkPath, file.Name()), depth+1, nextConnStack); err != nil {
					return nil, err
				} else {
					totalInfo.Add(info)
				}
				if idx < fileCount {
					nextConnStack.Remove(nextConnStack.Back())
				}
				totalInfo.FileCount += 1
			} else {
				totalInfo.DirectoryCount += 1
			}
		}
	}
	return totalInfo, err
}
