package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	ID      int64     `orm:"column(id)"`
	Name    string    `orm:"size(20)"`
	Pwd     string    `orm:"size(50)"`
	Profile *Profile  `orm:"rel(fk);null;rel(one);on_delete(set_null)"`
	Create  time.Time `orm:"auto_now_add;type(datetime)"`
	Update  time.Time `orm:"auto_now;type(datetime)"`
}

/**用户其他信息
 */
type Profile struct {
	ID     int64     `orm:"column(id)"`
	User   *User     `orm:"reverse(one)"`
	Update time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User), new(Profile))
}
