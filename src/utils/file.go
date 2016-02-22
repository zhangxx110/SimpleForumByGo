package utils

import (
	"errors"
	"os"
)

var tag_file = "file"

func mkDir(path string) bool {
	var Separator string
	Separator = GetSeparator()
	//logger.Println(Separator)
	dir, _ := os.Getwd()                                //当前的目录
	err := os.MkdirAll(dir+Separator+path, os.ModePerm) //在当前目录下生成md目录
	if err != nil {
		//logger.Println(err)
		if os.IsExist(err) {
			//logger.Println(tag_file, dir+Separator+path+" has exist")
			return true
		} else {
			//logger.Println(tag_file, dir+Separator+path+" not exist")
			return false
		}
	} else {
		//logger.Println("creat" + dir + Separator + path + " success")
		return true
	}

}

/*
全路径：path+fileName
*/
func CreatFile(path string, fileName string) (*os.File, error) {
	if mkDir(path) {
		Separator := GetSeparator()
		f, err := os.OpenFile(path+Separator+fileName, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
		//	logger.Println(tag_file, err)
			return nil, errors.New("creat file error")
		}
		return f, nil
	}
	return nil, errors.New("creat file error")
}
func GetSeparator() string {
	var Separator string
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		Separator = "\\"
	} else {
		Separator = "/"
	}
	return Separator
}
