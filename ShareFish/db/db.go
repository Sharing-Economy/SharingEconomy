package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var db *sql.DB //连接池对象
func InitDB() (err error) {
	//数据库
	//用户名:密码@tcp(ip:端口)/数据库的名字
	dsn := "root:123456@tcp(127.0.0.1:3306)/faucet?parseTime=true&loc=Local"
	//连接数据集
	db, err = sql.Open("mysql", dsn) //open不会检验用户名和密码
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("连接中。。。。")
	err = db.Ping() //尝试连接数据库
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("连接数据库成功~")
	//设置数据库连接池的最大连接数
	db.SetMaxIdleConns(10)
	return
}
//新增一条请求记录
func Insert(address string,amount int64) error {
	var sqlStr = "insert into address_huangcuihua (address,amount) values (?,?)"
	_, err := db.Exec(sqlStr, address, amount)
	if err != nil {
		fmt.Println("insertErr:",err)
		return err
	}
	return nil
}
//新增一条请求记录
func Update() error {
	var sqlStr = "update con_num set etherNum=etherNum+1"
	_, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println("updateErr:",err)
		return err
	}
	return nil
}


