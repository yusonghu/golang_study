package main

import (
	"days/day16/fileinfo"
	"fmt"
)

func main() {
	//fileinfo.FileInfo()
	//fileinfo.FileRead()
	//fileinfo.FileWrite()
	//fileCopy, err := fileinfo.FileCopy1("iowrite.txt", "iocopy.txt")
	//fmt.Println(fileCopy, err)
	fileCopy, err := fileinfo.FileCopy2("iowrite.txt", "iocopy.txt")
	fmt.Println(fileCopy, err)
}
