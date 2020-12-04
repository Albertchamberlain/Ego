package param

import (
	"commons"
	"strconv"
	"strings"
	"time"
)
import a "item/cat"

//显示规格参数
func showParamService(page, rows int) (d commons.Datagrid) {
	t := selByPageDao(page, rows)
	d.Total = selCount()
	cats := make([]TbItemParamCat, 0)
	for i := 0; i < len(t); i++ {
		var cat TbItemParamCat
		cat.Id = t[i].Id
		cat.Updated = t[i].Updated
		cat.Created = t[i].Created
		cat.ParamData = t[i].ParamData
		cat.ItemCatId = t[i].ItemCatId
		cat.CatName = a.ShowCatByIdService(t[i].ItemCatId).Name
		cats = append(cats, cat)
	}
	d.Rows = cats
	return
}

func deleByIdsService(ids string) (e commons.EgoResult) {
	idStr := strings.Split(ids, ",")
	idInt := make([]int, 0)

	for _, n := range idStr {
		id, _ := strconv.Atoi(n)
		idInt = append(idInt, id)
	}
	count := delByIdsDao(idInt)
	if count > 0 {
		e.Status = 200
	}
	return
}

func catidService(catid int) (e commons.EgoResult) {
	p := selByCatIdDao(catid)
	if p == nil {
		e.Status = 200
	}
	return
}

func insertParamService(catid int, paramData string) (e commons.EgoResult) {
	date := time.Now().Format("2006-01-02 15:04:05")
	param := TbItemParam{ItemCatId: catid, ParamData: paramData, Created: date, Updated: date}
	count := insertParamDao(param)
	if count > 0 {
		e.Status = 200

	}
	return
}

//编辑规格参数
func updateParamService(id int, itemCatId int, paramData string) (e commons.EgoResult) {
	param := TbItemParam{Id: id, ItemCatId: itemCatId, ParamData: paramData}
	count := updParamByIdDao(param)
	if count > 0 {
		e.Status = 200
	}
	return
}
