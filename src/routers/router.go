package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nasanasaQuQ/goProject/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	{
		// write routes
	}

	{
		// write routes
	}

	return r

}
