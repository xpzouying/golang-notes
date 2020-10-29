package main

import (
	"net/http"

	_ "./docs"
	"github.com/gin-gonic/gin"
	sf "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

// handleHello 响应请求返回
// @Summary 处理测试业务逻辑
// @Description 入口接口
// @accept plain
// @produce plain
// @param name path string true "用户名"
// @Router /hello/{name} [get]
func handleHello(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.String(http.StatusOK, "Hello "+name)
}

// @title zy的swagger测试用例
// @version 0.1
// @description 你好，我是zy。这里是我的swagger demo文档

// @contact.name zy
// @contact.url https://zouying.world
// @contact.email xpzouying@gmail.com

func main() {
	r := gin.New()

	// 增加swagger接口
	r.GET("/swagger/*any", gs.WrapHandler(sf.Handler))

	// 增加HTTP API
	r.GET("/hello/:name", handleHello)

	r.Run(":8080")
}
