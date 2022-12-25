package common

type Resp struct {
	BaseApi
}
type Info struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// Ok 操作成功
func (r *Resp) Ok(data any) {
	r.s(data, "操作成功", 200)
}

// Error 错误请求
func (r *Resp) Error(msg string) {
	r.s(nil, msg, 500)
}

func (r *Resp) s(data any, msg string, code int) {
	r.Context.JSONP(200, Info{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
