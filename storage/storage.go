package storage

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	//"log"
)

func New(storpath, filename string) (*Storage, error) {

	// 检测文件夹是否存在   若不存在  创建文件夹
	if _, err := os.Stat(storpath); err != nil {

		if os.IsNotExist(err) {

			err = os.MkdirAll(storpath, os.ModePerm)

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &Storage{storpath: storpath, name: filename}, nil
}

type Storage struct {
	storpath string
	name     string
}

// 获取文件信息
func (sto *Storage) Get(value interface{}) error {
	var filepath = path.Join(sto.storpath, sto.name)
	return read(filepath, value)
}

// 缓存文件
func (sto *Storage) Store(value interface{}) error {
	var filepath = path.Join(sto.storpath, sto.name)
	return write(filepath+".json", value)
}

// 删除文件
func (sto *Storage) Del() error {
    var filepath = path.Join(sto.storpath, sto.name)
    return os.Remove(filepath)
}

func getFile(storpath string) (*os.File, error) {
	f, err := os.OpenFile(storpath, os.O_RDWR, 0666)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			return os.Create(storpath)
		}
		return nil, err
	}
	return f, nil
}

func read(storpath string, value interface{}) error {
    f, err := os.OpenFile(storpath, os.O_RDWR, 0666)
	defer f.Close()

	if err != nil {
		return err
	}

	return json.NewDecoder(bufio.NewReader(f)).Decode(&value)
}

func write(storpath string, value interface{}) error {
	content, err := json.Marshal(value)

	if err != nil {
		return err
	}
	return ioutil.WriteFile(storpath, content, os.ModePerm)
}
