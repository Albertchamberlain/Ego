package user

import (
	"commons"
	"encoding/json"
	"net/http"
)

func UserHandler() {
	commons.Router.HandleFunc("/login", loginController)

}

func loginController(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	er := LoginService(username, password)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)

}
