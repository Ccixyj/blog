package models

import (
	"time"

	"crypto/md5"
	"fmt"

	"crypto/sha512"
	"encoding/base64"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"strings"
)

type Password string

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func (pwd Password) GenerateSalt() string {
	h := md5.New()
	if _, ok := h.Write([]byte(string(pwd))); ok != nil {
		fmt.Errorf("an error occurred : %v", ok)
	}
	md5key := fmt.Sprintf("%x", h.Sum(nil))
	saltLen := r.Intn(5) + 5 //salt length
	var strs = make([]byte, saltLen+1)
	strLen := len(md5key)
	for i := 0; i <= saltLen; i++ {
		strs[i] = md5key[r.Intn(strLen)]
	}
	return string(strs)
}

func (pwd Password) doCompare(compare Password) bool {
	decoded, error := base64.URLEncoding.DecodeString(string(compare))
	if error != nil {
		fmt.Errorf("decode base64 error %v", error)
	}
	splits := strings.Split(string(decoded), "$")
	if len(splits) != 2 {
		fmt.Errorf("decode base64 size must == 2 : %v", error)
	}
	return splits[1] == string(generateHash(splits[0], pwd))
}

//Password(base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s$%s", salt, pwdStr))))
func generateHash(salt string, oldPwd Password) string {
	sh := sha512.New()
	if _, ok := sh.Write([]byte(salt + string(oldPwd))); ok != nil {
		fmt.Errorf("an error occurred : %v", ok)
	}
	bs := sh.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

type User struct {
	ID      int64     `orm:"column(id)"`
	Name    string    `orm:"size(20);unique;index"`
	Pwd     Password  `json:"-"`
	Profile *Profile  `orm:"rel(fk);null;rel(one);on_delete(set_null)" json:"-"`
	Create  time.Time `orm:"auto_now_add;type(datetime)"`
	Update  time.Time `orm:"auto_now;type(datetime)"`
}

/**用户其他信息
 */
type Profile struct {
	ID     int64 `orm:"column(id)"`
	User   *User `orm:"reverse(one)"`
	Desc   string
	Update time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User), new(Profile))
}

func GenerateAdmin() *User {
	var adamin = new(User)
	adamin.Name = "15152278073"
	pwd := Password("45127996wind")
	salt := pwd.GenerateSalt()
	adamin.Pwd = Password(base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s$%s", salt, generateHash(salt, pwd)))))
	adamin.Profile = &Profile{ID: 1, Desc: "一个吴彦祖"}
	return adamin
}
