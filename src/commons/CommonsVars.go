package commons

import "github.com/gorilla/mux"

var (
	Router              = mux.NewRouter()
	CurrPath            = "http://localhost:80/"           //当前项目url
	HEADER_CONTENT_TYPE = "Content-Type"                   //Content-Type
	JSON_HEADER         = "application/json;charset=utf-8" //json
)
