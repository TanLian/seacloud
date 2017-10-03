package models

import (
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
	"seacloud/utils"
)

//下载文件时临时生成的一个下载链接
type TmpDownloadLink struct {
	Id         int    `orm:"pk;auto"`
	Token   string `orm:"unique"`
	Path string
	UserName string
	CreateTime time.Time `orm:"column(create_time);type(datetime);null;auto_now_add" json:"create_time"`
	UsedTimes int
	UsedLimits int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(TmpDownloadLink))
}

/*
注：usedlimits为0时表示token可以使用任意多次
*/
func GenerateTmpDownloadToken(username, p string, usedlimits int) (string, error) {
	token := utils.GenerateToken(username, p)
	link := TmpDownloadLink{
		Token: token,
		Path: p,
		UserName: username,
		UsedTimes: 0,
		UsedLimits: usedlimits}
	o := orm.NewOrm()
	_, err := o.Insert(&link)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return token, nil
}

func GetTmpDownloadObjByToken(token string) (*TmpDownloadLink, error) {
	o := orm.NewOrm()
	obj := TmpDownloadLink{Token: token}

	err := o.Read(&obj, "Token")
	if err != nil {
		return nil, err
	}
	
	return &obj, nil
}

func (this *TmpDownloadLink) TokenIsExpired() bool {
	createTime := this.CreateTime.Unix()
	now := time.Now().Unix()

	return now - createTime > 2 * utils.ONEHOUR * 1000
}

func (this *TmpDownloadLink) TokenUsedTimesAddOne() {
	this.UsedTimes += 1
	o := orm.NewOrm()
	o.Update(this, "UsedTimes")
}