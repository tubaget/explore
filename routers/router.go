package routers

import (
	"explore/middleware/log"
	"github.com/gin-gonic/gin"
	//"explore/middleware/log"
	"explore/pkg/setting"
)

func
InitRouter() *gin.Engine {
	r := gin.New()//初始化gin
/*	gin.ForceConsoleColor()//控制台的日志颜色控制
	r.Use(gin.Logger())//将日志输出到控制台*/
	r.Use(gin.Recovery())
	gin.SetMode(setting.Config().Run_mode)

	//j文件夹的api是给内部提供的json
	group_admin := r.Group("/")
	//group_j.Use(log.LoggerToFile(), jwt.JWT())
	group_admin.Use(log.LoggerToFile())
	{
		//获取标签列表
		group_admin.GET("/login/login", Login)
	}

	return r
}
