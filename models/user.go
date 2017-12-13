package models

import (
	"crypto/rand"
	"fmt"
	"time"

	"golang.org/x/crypto/scrypt"

	"io"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	jwt "github.com/dgrijalva/jwt-go"
)

const pwHashBytes = 64

var (
	key []byte = []byte("brucetan@qq.com")
)

type User struct {
	Id         int    `orm:"pk;auto"`
	UserName   string `orm:"unique"`
	Password   string
	Profile    string    //简介
	IsAdmin    bool
	LastLogin  time.Time `orm:"column(last_login);type(datetime);null" json:"last_login"`
	Telephone  string
	Avatar     string
	CreateTime time.Time `orm:"column(create_time);type(datetime);null;auto_now_add" json:"create_time"`
	Salt       string `json:"salt,omitempty"`
	Source string
}

func init() {
	// 需要在init中注册定义的model
	//orm.RegisterModelWithPrefix("user_", new(User))
	orm.RegisterModel(new(User))
}

func generateSalt() (salt string, err error) {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", buf), nil
}

// 产生json web token
func GenToken() string {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    "brucetan",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		logs.Error(err)
		return ""
	}
	return ss
}

func generatePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h), nil
}

// 校验token是否有效
func CheckToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return false
	}
	return true
}

func (u *User) CheckUserPass(pass string) (bool, error) {
	hash, err := generatePassHash(pass, u.Salt)
	if err != nil {
		return false, err
	}

	return u.Password == hash, nil
}

func (u *User) ChangePassWord(newPass string) error {
	hash, err := generatePassHash(newPass, u.Salt)
	if err != nil {
		return err
	}
	u.Password = hash
	_, err = orm.NewOrm().Update(u, "Password")
	return err
}

func InsertUser(username, password string, isAdmin bool, profile, telephone, avatar, source string) error {
	salt, err := generateSalt()
	if err != nil {
		return err
	}
	hash, err := generatePassHash(password, salt)
	if err != nil {
		return err
	}
	
	user := User{
		UserName:     username,
		Password: hash,
		IsAdmin: isAdmin,
		Profile: profile,
		Telephone: telephone,
		Avatar: avatar,
		Salt:     salt,
		Source: source}

	o := orm.NewOrm()
	_, err = o.Insert(&user)
	return err
}

func DeleteUserByName(username string) error {
	o := orm.NewOrm()
	
	_, err := o.Delete(&User{UserName:username})
	
	return err
}

func GetUserByName(username string) (*User, error) {
	o := orm.NewOrm()
	user := User{UserName: username}

	err := o.Read(&user, "UserName")
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

func GetAllUserNames() ([]string, error) {
	userlist := make([]string, 0)
	o := orm.NewOrm()
	_, err := o.Raw("SELECT user_name FROM user").QueryRows(&userlist)
	if err != nil {
			return userlist, err
	}
	return userlist, nil
}