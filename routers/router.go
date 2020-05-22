package routers

import (
	"github.com/gin-gonic/gin"
	"explore/middleware/log"
	"explore/pkg/setting"
)

func
InitRouter() *gin.Engine {
	r := gin.New()//初始化gin
/*	gin.ForceConsoleColor()//控制台的日志颜色控制
	r.Use(gin.Logger())//将日志输出到控制台*/
	r.Use(gin.Recovery())
	gin.SetMode(setting.Config().Run_mode)
	//r.GET("/auth", j.GetAuth)

	//j文件夹下是给外部提供的
	group_x := r.Group("/x")
	//group_j.Use(log.LoggerToFile(), jwt.JWT())
	group_x.Use()
	{
		//获取标签列表
	}

	return r
}
