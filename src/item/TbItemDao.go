package item

import (
	"commons"
	"database/sql"
	"fmt"
)

/*
rows:每页显示的条数
page:当前第几页
*/
func selByPageDao(rows, page int) []TbItem {
	//0开始查询 查询几个
	r, err := commons.Dql("select * from tb_item limit ?,?", rows*(page-1), rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	ts := make([]TbItem, 0)
	for r.Next() {
		var t TbItem
		var s sql.NullString
		//如果直接使用t.Barcode由于数据库中列为Null导致填充错误
		_ = r.Scan(&t.Id, &t.Title, &t.SellPoint, &t.Price, &t.Num, &s, &t.Image, &t.Cid, &t.Status, &t.Created, &t.Updated)
		t.Barcode = s.String
		ts = append(ts, t)
	}
	commons.CloseConn()
	return ts
}

//查询总条数
func selCount() (count int) {
	rows, err := commons.Dql("select count(*) from tb_item")
	if err != nil {
		fmt.Println(err)
		return -1
	}
	rows.Next()
	_ = rows.Scan(&count)
	commons.CloseConn()
	return
}

/**
返回值小于0更新失败
*/
func updStatusByIdsDao(ids []string, status int) int {
	if len(ids) <= 0 {
		return -1
	}
	sql := "update tb_item set status=? where "
	for i := 0; i < len(ids); i++ {
		sql += " id=" + ids[i]
		if i < len(ids)-1 {
			sql += " or "
		}
	}
	count, err := commons.Dml(sql, status)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//新增
func insertItemDao(t TbItem) int {
	count, err := commons.Dml("insert into tb_item values(?,?,?,?,?,?,?,?,?,?,?)", t.Id, t.Title, t.SellPoint, t.Price, t.Num, t.Barcode, t.Image, t.Cid, t.Status, t.Created, t.Updated)
	if err != nil {
		return -1
	}
	return int(count)
}

//delete by id
func delById(id int) int {
	count, err := commons.Dml("delete from tb_item where id=?", id)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}

//根据主键查询内容
func selByIdDao(id int) *TbItem {
	rows, err := commons.Dql("select * from tb_item where id=?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if rows.Next() {
		t := new(TbItem)
		var s sql.NullString
		rows.Scan(&t.Id, &t.Title, &t.SellPoint, &t.Price, &t.Num, &s, &t.Image, &t.Cid, &t.Status, &t.Created, &t.Updated)
		t.Barcode = s.String
		return t
	}
	return nil
}

func updItemByIdWithTx(t TbItem) int {
	return commons.PrepareWithTx("update tb_item set title=?,sell_point=?,price=?,num=?,barcode=?,image=?,cid=?,status=?,updated=? where id=?",
		t.Title, t.SellPoint, t.Price, t.Num, t.Barcode, t.Image, t.Cid, t.Status, t.Updated, t.Id)
}
