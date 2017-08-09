package models

type Result struct {
	Code   int32
	Msg    string
	Result interface{}
}

func NewSuccess(msg string, res interface{}) *Result {
	return &Result{
		Code:   200,
		Msg:    msg,
		Result: res,
	}
}

func NewError(code int32, msg string, res interface{}) *Result {
	return &Result{
		Code:   code,
		Msg:    msg,
		Result: res,
	}
}

type PageModel struct {
	Data interface{}
	Page *Page
}

type Page struct {
	CurrentPage int64 //当前页
	RecordSum   int64 //总数
	PageSize    int   //每页数量
	TotalPage   int64 //总页数
}

func init() {

}
