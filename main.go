package main

import (
	"commons"
	"github.com/gorilla/mux"
	"html/template"
	"item"
	"item/cat"
	"item/param"
	"net/http"
	"user"
)

//显示登录页面
func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/login.html")
	_ = t.Execute(w, nil)
}

func showPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t, _ := template.ParseFiles("view/" + vars["page"] + ".html")
	_ = t.Execute(w, nil)
}

func main() {
	commons.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	commons.Router.HandleFunc("/", welcome)
	commons.Router.HandleFunc("/page/{page}", showPage)

	user.UserHandler()
	item.ItemHandler()
	cat.ItemCatHandler()
	param.ParamHandler()
	http.ListenAndServe(":8090", commons.Router)
}
