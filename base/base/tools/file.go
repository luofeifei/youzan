package tools

import (
	"io/ioutil"
	"os"
	"path"
)

/*
写文件
*/
func WriteFile(filepath string, data []byte) error {
	return ioutil.WriteFile(filepath, data, os.ModePerm)
}

/**
读文件
*/
func ReadFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

/**
删除文件
*/
func Remove(path string) (error) {
	return os.Remove(path)
}

/**
删除文件或者目录
*/
func RemoveAll(path string) (error) {
	return os.RemoveAll(path)
}

/**
检查目录是否存在
*/
func PathExists(path string) (isExist bool) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

/**
创建目录
*/
func Mkdir(dir string) (err error) {
	return os.Mkdir(dir, os.ModePerm)
}

/**
确保目录存在，如果没有，则创建它
*/
func EnsureDir(dir string) (err error) {
	parent := path.Dir(dir)
	if _, err = os.Stat(parent); os.IsNotExist(err) {
		if err = EnsureDir(parent); err != nil {
			return
		}
	}
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
	}
	return
}

/**
确保文件存在，如果不存在，则创建它
*/
func EnsureFile(filepath string) (err error) {
	var (
		file *os.File
	)
	if err = EnsureDir(path.Dir(filepath)); err != nil {
		return err
	}
	if _, err = os.Stat(filepath); os.IsNotExist(err) {
		file, err = os.Create(filepath)
		defer func() {
			file.Close()
		}()
	}
	return
}

//几乎与fs.WriteFile相同，不同之处在于如果目录不存在，则会创建该目录。
func OuputFile(filepath string, data []byte) error {
	if err := EnsureDir(path.Dir(filepath)); err != nil {
		return err
	}
	return WriteFile(filepath, data)
}