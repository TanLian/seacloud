package models

import (
	"fmt"
	"seacloud/utils"
	"time"

	"os"
	"path/filepath"

	"github.com/astaxie/beego/orm"
)

type DownloadShareLink struct {
	UserName    string    `json:"username"`
	Path        string    `json:"path"`
	Password    string    `json:"password"`
	Expired     time.Time `orm:"column(expired_time);type(datetime);null;" json:"expired_time"`
	CreatorTime time.Time `orm:"column(creator_time);type(datetime);null;auto_now_add" json:"creator_time"`
	ViewCount   int64     `json:"view_count"`
	FileType    string    `json:"file_type"`
	Token       string    `json:"token" orm:"pk;"`
}

func init() {
	orm.RegisterModel(new(DownloadShareLink))
}

func StoreDownloadLink(username, token, path, password string, expired time.Time) error {
	//获取data_dir根目录
	dataDir := utils.GetDataBaseDir()
	p := filepath.Join(dataDir, username, "files", path)
	info, err := os.Stat(p)
	if err != nil {
		return err
	}

	//加密password
	passwordEnc := ""
	if password != "" {
		passwordEnc, err = utils.AesEncrypt(password)
		if err != nil {
			return err
		}
	}

	//获取文件类型
	tp := "f"
	if info.IsDir() {
		tp = "d"
	}

	shareLink := DownloadShareLink{
		UserName:  username,
		Path:      path,
		Password:  passwordEnc,
		Expired:   expired,
		ViewCount: 0,
		FileType:  tp,
		Token:     token}

	o := orm.NewOrm()
	_, err = o.Insert(&shareLink)

	return err
}

type DownloadLinkInfo struct {
	Token    string `json:"token"`
	Expired  string `json:"expired"`
	Password string `json:"password"`
}

func GetDownloadLinkInfo(username, path string) (bool, DownloadLinkInfo, error) {
	var items []*DownloadShareLink
	var info DownloadLinkInfo
	num, err := orm.NewOrm().QueryTable("download_share_link").Filter("user_name", username).Filter("path", path).All(&items)
	if err != nil {
		fmt.Println(err)
		return false, info, err
	}
	if num > 1 {
		orm.NewOrm().QueryTable("download_share_link").Filter("user_name", username).Filter("path", path).Delete()
		return false, info, nil
	}
	if num == 0 {
		return false, info, nil
	}
	password, err := utils.AesDecrypt(items[0].Password)
	if err != nil {
		fmt.Println(err)
		return false, info, err
	}
	tmFormat := items[0].Expired.Format("2006-01-02")
	info = DownloadLinkInfo{
		Token:    items[0].Token,
		Expired:  tmFormat,
		Password: password}
	return true, info, nil
}
