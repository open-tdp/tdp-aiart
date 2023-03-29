package httpd

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"

	"tdp-aiart/api"
	"tdp-aiart/cmd/args"
)

func Daemon() {

	if args.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	Server(args.Server.Listen, Engine())

}

func Engine() *gin.Engine {

	// 初始化
	engine := gin.New()
	engine.Use(Logger())
	engine.Use(Recovery(true))

	// 接口路由
	api.Router(engine)

	// 静态文件路由
	ui, _ := fs.Sub(args.Efs, "front")
	engine.StaticFS("/ui", http.FS(ui))

	// 默认首页路由
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/ui/")
	})

	return engine

}
