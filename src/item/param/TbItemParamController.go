package param

import (
	"commons"
	"encoding/json"
	"net/http"
	"strconv"
)

func ParamHandler() {
	commons.Router.HandleFunc("/item/param/show", showParamController)
	commons.Router.HandleFunc("/item/param/delete", deleByIdsController)
	commons.Router.HandleFunc("/item/param/iscat", isCatController)
	commons.Router.HandleFunc("/item/param/add", insertParamController)
	commons.Router.HandleFunc("/item/param/edit", updateParamController)
}

func insertParamController(w http.ResponseWriter, r *http.Request) {
	catid, _ := strconv.Atoi(r.FormValue("itemCatId"))
	paramData := r.FormValue("paramData")
	er := insertParamService(catid, paramData)
	b, _ := json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

func isCatController(w http.ResponseWriter, r *http.Request) {
	catid, _ := strconv.Atoi(r.FormValue("catid"))
	er := catidService(catid)
	b, _ := json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

func deleByIdsController(w http.ResponseWriter, r *http.Request) {
	er := deleByIdsService(r.FormValue("ids"))
	marshal, _ := json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	_, _ = w.Write(marshal)

}

func showParamController(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("rows"))

	data := showParamService(page, rows)

	marshal, _ := json.Marshal(data)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	_, _ = w.Write(marshal)

}

func updateParamController(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	itemCatId, _ := strconv.Atoi(r.FormValue("itemCatId"))
	paramData := r.FormValue("paramData")
	er := updateParamService(id, itemCatId, paramData)
	marshal, _ := json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	_, _ = w.Write(marshal)

}
