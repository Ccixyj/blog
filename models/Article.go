package models

import "github.com/astaxie/beego/orm"

type Article struct {
	ID       int64       `orm:"column(id)"`
	Title    string      `orm:"index"`
	Content  string      `orm:"type(text)"`
	User     *User       `orm:"rel(fk);null;rel(one);on_delete(set_null)"`
	Category []*Category `orm:"rel(m2m);rel_table(article_category)"`
}

func init() {
	orm.RegisterModel(new(Article))
}
