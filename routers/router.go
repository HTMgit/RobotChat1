package routers

import (
	"robot_chat/routers/api"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine) {
	router.POST("/robotchat", api.RobotChat)                 //发送聊天信息
	router.POST("/changerobotstatus", api.ChangeRobotStatus) //修改对应id机器人开关状态
}
