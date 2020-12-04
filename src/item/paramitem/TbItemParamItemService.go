package paramitem

import "commons"

//根据商品id显示商品规格参数
func showItemByIdService(id int) (e commons.EgoResult) {
	item := selByItemIdDao(id)
	if item != nil {
		e.Status = 200
		e.Data = item.ParamData
	}
	return
}
