package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vending/app/routes"
	"vending/common/constants"
	"vending/config"
)

func init() {
	//  TODO 默认配置文件路径待修改
	config.NewConfig("cmd/config.yml")
}
func run(mode string) {
	gin.SetMode(mode)
	r := gin.Default()
	err := routes.InitRoute(r)
	if err != nil {

	}
	srv := &http.Server{
		Addr:    ":" + config.Base.Server.Port,
		Handler: r,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server stop ...")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Fatal("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %v \n", err)
	}
	select {
	case <-ctx.Done():
		log.Fatal("timeout of 10 seconds.")
	}
	log.Fatal("Server exiting")
}

// TODO 待移除
func Run() {
	run(constants.DebugMode)
}
