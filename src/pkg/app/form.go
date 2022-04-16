package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/nasanasaQuQ/goProject/src/pkg/constant"
	"github.com/nasanasaQuQ/goProject/src/pkg/logging"
	"net/http"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)

	if err != nil {
		logging.Error(err)
		return http.StatusBadRequest, constant.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)

	if err != nil {
		logging.Error(err)
		return http.StatusInternalServerError, constant.ERROR
	}

	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, constant.INVALID_PARAMS
	}

	return http.StatusOK, constant.SUCCESS
}
