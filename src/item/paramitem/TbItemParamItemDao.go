package paramitem

import (
	"commons"
	"fmt"
)

//  新增商品参数与更新时间为系统当前时间，主键自增
func InsertParamItem(p TbItemParamItem) int {
	count, err := commons.Dml("insert into tb_item_param_item values (default,?,?,now(),now())", p.ItemId, p.ParamData)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

func selByItemIdDao(id int) *TbItemParamItem {
	r, err := commons.Dql("select * from tb_item_param_item where item_id=?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if r.Next() {
		item := new(TbItemParamItem)
		r.Scan(&item.Id, &item.ItemId, &item.ParamData, &item.Created, &item.Updated)
		return item
	}
	return nil
}
