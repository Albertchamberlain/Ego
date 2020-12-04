package item

import (
	"commons"
	"encoding/json"
	"net/http"
	"strconv"
)

func ItemHandler() {
	commons.Router.HandleFunc("/showItem", showItemController)
	commons.Router.HandleFunc("/item/delete", delByIdsController)
	commons.Router.HandleFunc("/item/instock", instockController)
	commons.Router.HandleFunc("/item/offstock", offstockController)
	commons.Router.HandleFunc("/item/imageupload", imagesUploadController)
	commons.Router.HandleFunc("/item/add", insertControllew) //商品新增
	commons.Router.HandleFunc("/item/showItemById", showItemDescCatController)
	commons.Router.HandleFunc("/item/update", updateController) //商品修改页面信息显示
}

//modify
func updateController(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	er := updateService(r.Form)
	b, _ := json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	_, _ = w.Write(b)
}

//显示商品信息
func showItemController(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("rows"))
	datagrid := showItemService(page, rows)
	b, _ := json.Marshal(datagrid)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	_, _ = w.Write(b)
}

func delByIdsController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	er := delByIdsService(ids)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	_, _ = w.Write(b)
}

func instockController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	er := instockService(ids)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	_, _ = w.Write(b)
}

func offstockController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	er := offStockService(ids)
	b, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	_, _ = w.Write(b)
}

func imagesUploadController(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("imgFile")
	if err != nil {
		m := make(map[string]interface{})
		m["error"] = 1
		m["message"] = "接收图片失败"
		b, _ := json.Marshal(m)
		w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
		_, _ = w.Write(b)
		return
	}
	m := imageUploadService(file, fileHeader)
	b, _ := json.Marshal(m)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	_, _ = w.Write(b)
}

//add
func insertControllew(w http.ResponseWriter, r *http.Request) {
	//需要先进行解析
	_ = r.ParseForm()
	er := insetService(r.Form)
	b, _ := json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	_, _ = w.Write(b)
}

func showItemDescCatController(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	c := showItemDescCatService(id)
	b, _ := json.Marshal(c)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}
