package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pro911/golang-demo/internal/app/middleware"
	"github.com/pro911/golang-demo/internal/pkg/biz/response"
	"github.com/pro911/golang-demo/internal/pkg/handler"
)

func RegisterRouter(r *gin.Engine) {
	// cors
	r.Use(middleware.Cors())

	r.Use(middleware.ZapLogger(), middleware.ZapRecovery(true))

	//一个测试监听探活接口
	r.Any("online", func(c *gin.Context) {
		response.Success(c)
		return
	})

	r.Any("task_send", handler.TaskSend)
}
