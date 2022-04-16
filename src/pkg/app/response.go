package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nasanasaQuQ/goProject/src/pkg/constant"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode int, errCode interface{}, data interface{}) {
	switch errCode.(type) {
	case int:
		intCode := errCode.(int)
		g.C.JSON(httpCode, Response{
			Code: intCode,
			Msg:  constant.GetMsg(intCode),
			Data: data,
		})
	case string:
		strCode := errCode.(string)
		g.C.JSON(httpCode, Response{
			Code: 9999,
			Msg:  strCode,
			Data: data,
		})
	}
	return
}
