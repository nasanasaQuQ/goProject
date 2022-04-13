package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {

}

func main() {
	log.Println("server is Starting...")

	router := gin.Default()

	router.GET("/api/login", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"Hello": "world",
		})
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Println("server is started, Listen on port 8080")

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	// 通过信号阻塞停止代码 直到手动停止
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	log.Println("Shutting down the server")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shundown: %v\n", err)
	}

}
