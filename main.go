package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nasanasaQuQ/goProject/src/app/models"
	"github.com/nasanasaQuQ/goProject/src/pkg/base"
	"github.com/nasanasaQuQ/goProject/src/pkg/global"
	"github.com/nasanasaQuQ/goProject/src/pkg/logging"
	"github.com/nasanasaQuQ/goProject/src/routers"
	"log"
	"net/http"
	"time"
)

func init() {
	global.VP = base.Viper()
	global.LOG = base.SetupLogger()
	models.Setup()
	logging.Setup()
}

func main() {

	gin.SetMode(global.CONFIG.Server.RunMode)

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", global.CONFIG.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	global.LOG.Info("[info] start http server listening %s", endPoint)
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
