package model

import "fmt"

type TotalInfo struct {
	DirectoryCount int
	FileCount      int
}

func (info *TotalInfo) String() string {
	return fmt.Sprintf("Total: %d directories, %d files", info.DirectoryCount, info.FileCount)
}

func (info *TotalInfo) Add(otherInfo *TotalInfo) {
	info.FileCount += otherInfo.FileCount
	info.DirectoryCount += otherInfo.DirectoryCount
}
