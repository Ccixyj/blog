package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Category struct {
	ID      int64  `orm:"column(id)"`
	Name    string `orm:"size(20);index;unique"`
	Order   int
	Create  time.Time  `orm:"auto_now_add;type(datetime)"`
	Article []*Article `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Category))
}
