package desc

import (
	"commons"
	"fmt"
)

func insertDescDao(t TbItemDesc) int {
	count, err := commons.Dml("insert into tb_item_desc values(?,?,?,?)", t.ItemId, t.ItemDesc, t.Created, t.Updated)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//根据主键查询
func selByIdDao(id int) *TbItemDesc {
	r, err := commons.Dql("select * from tb_item_desc where item_id=?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if r.Next() {
		t := new(TbItemDesc)
		r.Scan(&t.ItemId, &t.ItemDesc, &t.Created, &t.Updated)
		return t
	}
	return nil
}

//将修改商品描述作为事务
func UpdDescByIdWithTxDao(t TbItemDesc) int {
	return commons.PrepareWithTx("update tb_item_desc set item_desc=?,updated=? where item_id=?", t.ItemDesc, t.Updated, t.ItemId)
}

//删除商品描述

func DelDescByIdDao(id int) int {
	count, err := commons.Dml("delete from tb_item_desc where item_id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	return int(count)

}
