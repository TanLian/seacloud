package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Favorites struct {
	Id         int    `orm:"pk;auto"`
	UserName string
	Path string
	IsDir bool
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Favorites))
}

func AddFavoritesItem(username, path string, isDir bool) error {
	item := Favorites {
		UserName: username,
		Path: path,
		IsDir: isDir}
	o := orm.NewOrm()
	_, err := o.Insert(&item)
	return err
}

func DeleteFavoritesItem(username, path string) error {
	_, err := orm.NewOrm().QueryTable("favorites").Filter("user_name", username).Filter("path", path).Delete()
	return err
}

func GetAllFavoritesByUsername(username string) ([]*Favorites, error) {
	var items []*Favorites
	num, err := orm.NewOrm().QueryTable("favorites").Filter("user_name", username).All(&items)
	fmt.Println(num)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return items, nil
}

func DoesFavorateItemExist(username, path string) bool {
	return orm.NewOrm().QueryTable("favorites").Filter("user_name", username).Filter("path", path).Exist()
}