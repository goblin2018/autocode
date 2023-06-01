package {{.group}}

import (
	"auto/api"
	"auto/app/logic/{{.package}}"
	"auto/pkg/ctx"
	"auto/pkg/e"
)

func {{.api}}(c *ctx.Context) {
	req := new(api.{{.reqName}})
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}
	l := {{.package}}.New{{.logicFuncName}}Logic(c, sv)
	{{if .hasResp}}resp, {{end}}err := l.{{.logicFuncName}}(req)
	c.JSON({{if .hasResp}}resp{{else}}nil{{end}}, err)
}
