package utils

import (
	"fmt"
	"path/filepath"
	"github.com/widuu/goini"
	"io/ioutil"
	"time"
	"strconv"
	"crypto/md5"
	"io"
)

type File struct {
	Name string
	Size int64
	Type string
	Mtime int64
	MtimeRelative string
}

func Translate_seacloud_time(mtime int64) string{
	now := time.Now().Unix()
	fmt.Println(mtime)
	fmt.Println(now)
	if mtime > now {
		return "1秒前"
	}
	switch  {
	case now - mtime > TWOWEEKS:
		return time.Unix(mtime, 0).Format("2006-01-02 15:04:05")
	case now - mtime > ONEDAY:
		return strconv.FormatInt((now-mtime)/ONEDAY, 10) + "天前"
	case now - mtime > ONEHOUR:
		return strconv.FormatInt((now-mtime)/ONEHOUR, 10) + "小时前"
	case now - mtime > ONEMINUTE:
		return strconv.FormatInt((now-mtime)/ONEMINUTE, 10) + "分钟前"
	default:
		return strconv.FormatInt(now-mtime, 10) + "秒前"
	}
}

func GetConfPath() string {
	return "/Users/tanlian/Documents/goprj/src/seacloud/conf/seacloud.ini"
}

func GetDataBaseDir() string {
	confPath := GetConfPath()
	conf := goini.SetConfig(confPath)
	dataDir := conf.GetValue("GENERA", "data_dir")
	return dataDir
}

/*
给定用户名和path，返回文件列表
*/
func GetFilelistByPath(username, p string) ([]File, error) {
	ret := make([]File, 0)

	//获取data_dir根目录
	dataDir := GetDataBaseDir()

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
		mtime := fi.ModTime().Unix()
		seacloud_time := Translate_seacloud_time(mtime)
		fmt.Println(mtime)
		fmt.Println(seacloud_time)
		obj := File{
			Name: fi.Name(),
			Size: fi.Size(),
			Type: tp,
			Mtime: mtime,
			MtimeRelative: seacloud_time}
		ret = append(ret, obj)
	}

	return ret, nil
}

//生成上传链接、下载链接token
func GenerateToken(username, path string) string {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	str := now + username + path
	h := md5.New()
	io.WriteString(h, str)
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

func GetTmpDownloadLink(token string) string {
	return "/api/file/download?token=" + token
}