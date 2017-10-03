package models

import (
	//"fmt"
	"time"
	"github.com/astaxie/beego/orm"
	"seacloud/utils"
)

//上传文件时生成一个临时的上传链接
type TmpUploadLink struct {
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
	orm.RegisterModel(new(TmpUploadLink))
}

/*
注：usedlimits为0时表示token可以使用任意多次
*/
func GenerateTmpUploadToken(username, p string, usedlimits int) string {
	token := utils.GenerateToken(username, p)
	link := TmpUploadLink{
		Token: token,
		Path: p,
		UserName: username,
		UsedTimes: 0,
		UsedLimits: usedlimits}
	o := orm.NewOrm()
	o.Insert(&link)
	return token
}

func GetTmpUploadObjByToken(token string) (*TmpUploadLink, error) {
	o := orm.NewOrm()
	obj := TmpUploadLink{Token: token}

	err := o.Read(&obj, "Token")
	if err != nil {
		return nil, err
	}
	
	return &obj, nil
}

func (this *TmpUploadLink) TokenIsExpired() bool {
	createTime := this.CreateTime.Unix()
	now := time.Now().Unix()

	return now - createTime > 2 * utils.ONEHOUR * 1000
}

func (this *TmpUploadLink) TokenUsedTimesAddOne() {
	this.UsedTimes += 1
	o := orm.NewOrm()
	o.Update(this, "UsedTimes")
}