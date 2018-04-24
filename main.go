package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/api/userLoginIn",haha)
	router.Run(":80")


}

//haha接口
func haha(ctx *gin.Context)  {
	fmt.Println("你好啊，i love you baby")
	userName := ctx.Param("userName")
	password := ctx.Param("password")
	fmt.Println("userName from app %s",userName)
	fmt.Println("password from app %s",password)
	response := gin.H{
		"message" : "我是返回的数据，啊哈哈",
		"resultCode" : 200,
	}
	ctx.JSON(200,response)
}



