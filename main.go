package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"github.com/hhxsv5/gin-x-skeleton/app/http/controllers"
	"github.com/hhxsv5/gin-x/config"
	"github.com/hhxsv5/gin-x/enums/codes"
	"github.com/hhxsv5/gin-x/middleware"
	"github.com/hhxsv5/gin-x/response"
	"github.com/hhxsv5/gin-x/router"

	_ "github.com/hhxsv5/gin-x-skeleton/bootstrap"
)

func main() {
	gin.DisableConsoleColor()
	gin.SetMode(config.AppMode())

	e := gin.New()
	if config.AppModeIs(gin.DebugMode) {
		e.Use(gin.Logger())
	}
	e.Use(middleware.Recovery())
	e.Use(middleware.Cors())

	e.NoRoute(func(c *gin.Context) {
		response.FailJSON(c, codes.NotFoundError, "no matched route")
	})

	// Gin原始的路由绑定方法
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "pong",
			"timestamp": time.Now().Unix(),
			"mode":      gin.Mode(),
		})
	})

	// 静态文件
	//e.Static("/static", "./static")
	// 打包静态文件到二进制
	box := packr.NewBox("./static")
	e.StaticFS("/static", box)

	r := router.NewSlimRouter(e)

	// Slim的路由绑定方法，基于控制器
	r.RegisterGroup("test").RegisterController(controllers.Test{}.NewController())

	s := &http.Server{
		Addr:           config.AppListenAddress(),
		Handler:        http.TimeoutHandler(e, 30*time.Second, "request timeout"),
		ReadTimeout:    31 * time.Second,
		WriteTimeout:   31 * time.Second,
		MaxHeaderBytes: 1 << 19,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln("listen fail", err)
	}
}
