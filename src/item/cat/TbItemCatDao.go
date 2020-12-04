package cat

import (
	"commons"
	"fmt"
)

func selByIdDao(id int) (t *TbItemCat) {
	rows, err := commons.Dql("select * from tb_item_cat where id=?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if rows.Next() {
		t = new(TbItemCat)
		rows.Scan(&t.Id, &t.ParentId, &t.Name, &t.Status, &t.SortOrder, &t.IsParent, &t.Created, &t.Updated)
	}
	commons.CloseConn()
	return
}

/*
根据parent_id查询所有子类目
*/
func selByPidDao(pid int) (c []TbItemCat) {
	rows, err := commons.Dql("select * from tb_item_cat where parent_id=?", pid)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	c = make([]TbItemCat, 0)
	for rows.Next() {
		var t TbItemCat
		rows.Scan(&t.Id, &t.ParentId, &t.Name, &t.Status, &t.SortOrder, &t.IsParent, &t.Created, &t.Updated)
		c = append(c, t)
	}
	commons.CloseConn()
	return
}
