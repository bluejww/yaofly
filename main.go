package main

import (
	"net/http"
	template2 "html/template"
)

func main() {

	server := &http.Server{
		Addr:"127.0.0.1:8080",
	}

	http.HandleFunc("/",home)


	server.ListenAndServe()

}


func home(writer http.ResponseWriter, request *http.Request) {
	var template *template2.Template
	template = template2.Must(template2.ParseFiles("templates/home.html"))
	template.ExecuteTemplate(writer,"layout",nil)
}


