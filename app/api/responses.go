package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func ReturnErrCode(r *ghttp.Request, code int, msg string) {
	r.Response.Status = code
	r.Response.WriteJson(g.Map{
		"Message": msg,
		"code":    code,
	})
}
