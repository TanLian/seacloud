package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Avatar struct {
	Id         int    `orm:"pk;auto"`
	UserName   string `orm:"unique"`
	Data string	`orm:"size(20000)"`
	ModifyTime time.Time `orm:"column(modify_time);type(datetime);null;auto_now_add" json:"modify_time"`
}

func init() {
	orm.RegisterModel(new(Avatar))
}

func GetAvatarDataByUsername(username string) (string, error) {
	o := orm.NewOrm()
	avatar := Avatar{UserName: username}

	err := o.Read(&avatar, "UserName")
	if err != nil {
		return "", err
	}
	
	return avatar.Data, nil
}

func InsertOrUpdateAvatar(username, data string) error{
	o := orm.NewOrm()
	avatar := &Avatar {
		UserName: username,
		Data: data}
	data, err := GetAvatarDataByUsername(username)
	if err != nil {
		_, err = o.Insert(avatar)
		return err
	}
	_, err = o.Update(avatar)
	return err
}