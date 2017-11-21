package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
)

type Profile struct {
	Id         int    `orm:"pk;auto"`
	UserName   string `orm:"unique"`
	Motto	string
	Intro string
	Depart string
	Tele	string
}

func init() {
	orm.RegisterModel(new(Profile))
}

func GetProfileByUsername(username string) (*Profile, error) {
	o := orm.NewOrm()
	profile := Profile{UserName: username}

	err := o.Read(&profile, "UserName")
	if err != nil {
		return nil, err
	}
	
	return &profile, nil
}

func InsertOrUpdateProfile(username, motto, intro, depart, tele string) error {
	o := orm.NewOrm()
	profile, err := GetProfileByUsername(username)
	if profile != nil {
		profile.Motto = motto
		profile.Intro = intro
		profile.Depart = depart
		profile.Tele = tele
		_, err = o.Update(profile)
		return err
	}
	profile = &Profile {
		UserName: username,
		Motto: motto,
		Intro: intro,
		Depart: depart,
		Tele: tele}
	_, err = o.Insert(profile)
	return err
}