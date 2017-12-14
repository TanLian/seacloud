package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Group struct {
	GroupId         int    `orm:"pk;auto"`
	GroupName   string `orm:"unique"`
	CreatorName string	`orm:"size(20000)"`
	CreatorTime time.Time `orm:"column(modify_time);type(datetime);null;auto_now_add" json:"modify_time"`
}

func init() {
	orm.RegisterModel(new(Group))
}

func CreateGroup(groupName, creatorName string) error {
	group := Group{
		GroupName: groupName,
		CreatorName: creatorName}

	o := orm.NewOrm()
	_, err := o.Insert(&group)
	return err
}