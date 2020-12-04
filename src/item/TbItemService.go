package item

import (
	"commons"
	"io/ioutil"
	"item/cat"
	"item/desc"
	"item/paramitem"
	"math/rand"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func showItemService(page, rows int) (e *commons.Datagrid) {
	ts := selByPageDao(rows, page)
	if ts != nil {
		itemChildren := make([]TbItemChild, 0)
		for i := 0; i < len(ts); i++ {
			var itemChild TbItemChild
			itemChild.Id = ts[i].Id
			itemChild.Updated = ts[i].Updated
			itemChild.Created = ts[i].Created
			itemChild.Status = ts[i].Status
			itemChild.Barcode = ts[i].Barcode
			itemChild.Price = ts[i].Price
			itemChild.Num = ts[i].Num
			itemChild.SellPoint = ts[i].SellPoint
			itemChild.Title = ts[i].Title
			itemChild.CategoryName = cat.ShowCatByIdService(ts[i].Cid).Name
			itemChildren = append(itemChildren, itemChild)
		}
		e = new(commons.Datagrid)
		e.Rows = itemChildren
		e.Total = selCount()
		return
	}
	return nil
}

func delByIdsService(ids string) (e commons.EgoResult) {
	count := updStatusByIdsDao(strings.Split(ids, ","), 3)
	if count > 0 {
		e.Status = 200
	}
	return
}

func instockService(ids string) (e commons.EgoResult) {
	count := updStatusByIdsDao(strings.Split(ids, ","), 1)
	if count > 0 {
		e.Status = 200
	}
	return
}

func offStockService(ids string) (e commons.EgoResult) {
	count := updStatusByIdsDao(strings.Split(ids, ","), 2)
	if count > 0 {
		e.Status = 200
	}
	return
}

func imageUploadService(f multipart.File, h *multipart.FileHeader) map[string]interface{} {
	m := make(map[string]interface{})
	b, err := ioutil.ReadAll(f)
	if err != nil {
		m["error"] = 1
		m["message"] = "上传失败,服务器错误"
		return m
	}
	//ns+rand+extension name
	rand.Seed(time.Now().UnixNano())
	fileName := "static/images/" + strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(rand.Intn(1000)) + h.Filename[strings.LastIndex(h.Filename, "."):]
	err = ioutil.WriteFile(fileName, b, 0777)
	if err != nil {
		m["error"] = 1
		m["message"] = "上传失败,保存图片时错误"
		return m
	}
	m["error"] = 0
	m["url"] = commons.CurrPath + fileName
	return m
}

//商品新增
func insetService(f url.Values) (e commons.EgoResult) {
	var t TbItem
	cid, _ := strconv.Atoi(f["Cid"][0])
	t.Cid = cid
	t.Title = f["Title"][0]
	t.SellPoint = f["SellPoint"][0]
	price, _ := strconv.Atoi(f["Price"][0])
	t.Price = price
	num, _ := strconv.Atoi(f["Num"][0])
	t.Num = num
	t.Image = f["Image"][0]
	t.Status = 1
	date := time.Now().Format("2006-01-02 15:04:05")
	t.Created = date
	t.Updated = date
	id := commons.GenId()
	t.Id = id
	//商品表新增执行
	count := insertItemDao(t)
	if count > 0 {
		//商品描述新增
		var tbItemDesc desc.TbItemDesc
		tbItemDesc.ItemId = id
		tbItemDesc.Created = date
		tbItemDesc.Updated = date
		tbItemDesc.ItemDesc = f["Desc"][0]
		countDesc := desc.Insert(tbItemDesc)
		if countDesc > 0 {
			paramItem := paramitem.TbItemParamItem{ItemId: id, ParamData: f["paramData"][0]}
			countParamItem := paramitem.InsertParamItem(paramItem)
			if countParamItem > 0 {
				e.Status = 200
			} else {
				//删除商品
				delById(id)
				//删除商品描述中的数据
				desc.DelDescByIdDao(id)
				e.Status = 400
			}

		} else {
			//删除商品中数据, service层引入事务管理  可以引入事务处理 将dao层的事务提到service层处理
			delById(id)
			e.Status = 400
		}
	}
	return
}

//修改页面显示信息
func showItemDescCatService(id int) TbItemDescChild {
	item := selByIdDao(id)
	var c TbItemDescChild
	c.Id = item.Id
	c.Updated = item.Updated
	c.Created = item.Created
	c.Barcode = item.Barcode
	c.Cid = item.Cid
	c.Title = item.Title
	c.SellPoint = item.SellPoint
	c.Price = item.Price
	c.Image = item.Image
	c.Status = item.Status
	c.Num = item.Num
	//商品类目
	c.CategoryName = cat.ShowCatByIdService(c.Cid).Name
	//商品描述
	c.Desc = desc.SelByIdService(c.Id).ItemDesc
	return c
}

func updateService(v url.Values) (e commons.EgoResult) {

	_ = commons.OpenConnWithTx()
	var t TbItem
	id, _ := strconv.Atoi(v["Id"][0])
	t.Id = id
	cid, _ := strconv.Atoi(v["Cid"][0])
	t.Cid = cid
	t.Title = v["Title"][0]
	t.SellPoint = v["SellPoint"][0]
	price, _ := strconv.Atoi(v["Price"][0])
	t.Price = price
	num, _ := strconv.Atoi(v["Num"][0])
	t.Num = num
	t.Image = v["Image"][0]
	status, _ := strconv.Atoi(v["Status"][0])
	t.Status = int8(status)
	date := time.Now().Format("2006-01-02 15:04:05")
	t.Updated = date

	count := updItemByIdWithTx(t)

	if count > 0 {

		var itemDesc desc.TbItemDesc
		itemDesc.ItemId = id
		itemDesc.ItemDesc = v["Desc"][0]
		itemDesc.Updated = date
		count = desc.UpdDescByIdWithTxDao(itemDesc)
		if count > 0 {
			commons.CloseConnWithTx(true)
			e.Status = 200
			return
		}
	}
	commons.CloseConnWithTx(false)
	return

}
