package utils

import (
	"fmt"
	"path/filepath"
	"github.com/widuu/goini"
	"io/ioutil"
)

type File struct {
	Name string
	Size int64
	Type string
	Mtime int64
	MtimeRelative string
}

func GetConfPath() string {
	return "/Users/tanlian/Documents/goprj/src/seacloud/conf/seacloud.ini"
}

/*
给定用户名和path，返回文件列表
*/
func GetFilelistByPath(username, p string) ([]File, error) {
	ret := make([]File, 0)

	//获取data_dir根目录
	confPath := GetConfPath()
	conf := goini.SetConfig(confPath)
	dataDir := conf.GetValue("GENERA", "data_dir")

	//遍历指定目录，返回文件列表
	dir, err := ioutil.ReadDir(filepath.Join(dataDir, username, p))
	if err != nil {
		fmt.Println(err)
		return ret, err
	}

	for _, fi := range dir {
		tp := "file"
		if fi.IsDir() {
			tp = "dir"
		}
		obj := File{
			Name: fi.Name(),
			Size: fi.Size(),
			Type: tp,
			Mtime: fi.ModTime().Unix(),
			MtimeRelative: fi.ModTime().Format("2006-01-02 15:04:05")}
		ret = append(ret, obj)
	}

	return ret, nil
}