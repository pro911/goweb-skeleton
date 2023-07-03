package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pro911/golang-demo/internal/pkg/biz/errno"
	"github.com/pro911/golang-demo/internal/pkg/biz/response"
	"github.com/pro911/golang-demo/internal/pkg/logic/task"
)

func TaskSend(ctx *gin.Context) {
	data, err := task.TaskSendData(ctx)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrServer, "")
		return
	}
	response.SuccessWithData(ctx, data)
	return
}
