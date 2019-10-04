package util

import (
	"io/ioutil"
	"os"
)

func WriteFile(path string, content string) error {
	fileObj, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fileObj.Close()
	contents := []byte(content)
	if _, err := fileObj.Write(contents); err != nil {
		return err
	}
	return nil
}

func WriteFileAppend(path string, content string) error {
	fileObj, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer fileObj.Close()
	contents := []byte(content)
	if _, err := fileObj.Write(contents); err != nil {
		return err
	}
	return nil
}

//ReadFile 读取文件
func ReadFile(path string) (string, error) {
	fi, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd), nil
}

//RemoveFile 删除指定文件夹或文件
func RemoveFile(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}

//MakeDirAll 创建指定路径文件
func MakeDirAll(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}
	return nil
}

//ReadDir 查看指定目录下所有文件
func ReadDir(path string) ([]string, error) {
	if dirList, err := ioutil.ReadDir(path); err != nil {
		return nil, err
	} else {
		fileName := make([]string, len(dirList))
		for i, v := range dirList {
			if v.IsDir() {
				fileName[i] = "dir:" + v.Name()
			} else {
				fileName[i] = "file:" + v.Name()
			}
		}
		return fileName, nil
	}
}
