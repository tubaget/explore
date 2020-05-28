package routers

import (
	"fmt"
	"net/http"
	/*
		"log"*/

	"github.com/gin-gonic/gin"
	/*	"github.com/astaxie/beego/validation"
		"github.com/Unknwon/com"

		"explore/models"
		"explore/pkg/e"
		"explore/pkg/setting"
		"explore/pkg/util"*/
)

func Index(c *gin.Context) {
	fmt.Println(33)
}

//登录
func Login(c *gin.Context) {
	fmt.Println(33)

	c.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "dsd",
		"data" : nil,
	})

	/*id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface {}
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})*/
}