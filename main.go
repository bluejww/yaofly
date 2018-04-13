package main

import (
	"net/http"
	template2 "html/template"
	"fmt"
)

func main() {
	server := &http.Server{
		Addr:"106.14.222.26:80",
	}
	http.HandleFunc("/",home)
	server.ListenAndServe()

}


func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("in home func 哈哈哈哈")
	request.ParseForm()
	var template *template2.Template
	template = template2.Must(template2.ParseFiles("./templates/home.html"))
	template.ExecuteTemplate(writer,"layout",nil)
}


