package routes

import (
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	//Default 默认使用了两个中间件
	//gin.Logger() 日志文件 和 gin.Recovery() 故障处理中间件
	//1. 创建路由
	r := gin.Default()
	//2. 初始路由组
	//绑定路由规则，执行的函数
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			//c.JSON(状态,返回值)
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(utils.HttpPort)

}
