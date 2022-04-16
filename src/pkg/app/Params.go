package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nasanasaQuQ/goProject/src/app/models/dto"
	"github.com/unknwon/com"
)

func GetParams(ctx *gin.Context) dto.BasePage {
	var (
		page   int
		size   int
		blurry string
	)

	page = com.StrTo(ctx.DefaultQuery("page", "1")).MustInt()
	size = com.StrTo(ctx.DefaultQuery("size", "1")).MustInt()
	blurry = ctx.DefaultQuery("blurry", "")

	return dto.BasePage{Page: page, Size: size, Blurry: blurry}
}
