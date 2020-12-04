package commons

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//数据库操作的三个对象
var (
	db   *sql.DB
	stmt *sql.Stmt
	rows *sql.Rows
	tx   *sql.Tx
)

func OpenConnWithTx() (err error) {
	db, err = sql.Open("mysql", "root:****@tcp(****:3306)/GO")
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	//开启事务
	tx, err = db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	return nil
}
func PrepareWithTx(sql string, args ...interface{}) int {
	result, err := tx.Exec(sql, args...)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}
func CloseConnWithTx(result bool) {
	if result {
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	if rows != nil {
		_ = rows.Close()
	}
	if stmt != nil {
		_ = stmt.Close()
	}
	if db != nil {
		_ = db.Close()
	}
}
func openConn() (err error) {
	db, err = sql.Open("mysql", "root:***@tcp(****:3306)/GO")
	if err != nil {
		fmt.Println("连接失败", err) //发布时删除
		return
	}
	return nil
}

//关闭连接,首字母大写,需要跨Package访问的
func CloseConn() {
	if rows != nil {
		_ = rows.Close()
	}
	if stmt != nil {
		_ = stmt.Close()
	}
	if db != nil {
		_ = db.Close()
	}
}

//新增,删除,修改
func Dml(sql string, args ...interface{}) (int64, error) {
	err := openConn()
	if err != nil {
		fmt.Println("执行DML时出现错误,打开连接失败")
		return 0, err
	}
	stmt, err = db.Prepare(sql)
	if err != nil {
		fmt.Println("执行DML时出现错误,预处理出现错误")
		return 0, err
	}
	//此处要有...表示切片,如果没有表示数组,会报错
	result, err := stmt.Exec(args...)

	if err != nil {
		fmt.Println("执行DML时出现错误,执行错误")
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println("执行DML时出现错误,获取受影响行数错误")
		return 0, err
	}
	CloseConn() //关闭连接
	return count, err
}

//DQL
func Dql(sql string, args ...interface{}) (*sql.Rows, error) {
	err := openConn()
	if err != nil {
		fmt.Println("执行DQL出现错误,打开连接失败")
		return nil, err
	}
	stmt, err = db.Prepare(sql)
	if err != nil {
		fmt.Println("执行DQL出现错误,预处理实现")
		return nil, err
	}

	rows, err = stmt.Query(args...)
	if err != nil {
		fmt.Println("执行DQL出现错误,执行错误")
		return nil, err
	}
	return rows, nil
}
