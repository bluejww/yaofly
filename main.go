package main

import (
	"net/http"
	template2 "html/template"
	"fmt"
	"github.com/gin-gonic/gin"
	"context"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/haha",haha)
	http.ListenAndServe(":8080",router)
}

//haha接口
func haha(ctx *gin.Context)  {
	fmt.Println("你好啊，i love you baby")
}

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("in home func 哈哈哈哈")
	request.ParseForm()
	var template *template2.Template
	template = template2.Must(template2.ParseFiles("./templates/home.html"))
	template.ExecuteTemplate(writer,"layout",nil)
}


