package ctx

import (
	"auto/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

func (c *Context) Abort(err error) {
	c.Context.Abort()
	c.JSON(nil, err)
}

type ResponseBody struct {
	Code int32             `json:"code" form:"code"`
	Msg  string            `json:"msg" form:"msg"`
	Data datatypes.JSONMap `json:"data" form:"data"`
}

func (c *Context) JSON(data interface{}, err error) {

	if err != nil {
		data = nil
	}

	er := e.CastError(err)
	c.Context.JSON(http.StatusOK, gin.H{
		"code": er.Code(),
		"msg":  er.Error(),
		"data": data,
	})
}

func (c *Context) Fail(err error) {
	c.JSON(nil, err)
}
