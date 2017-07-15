package conf

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//声明一些全局变量
var (
	dbhostsip  = "127.0.0.1:3306" //IP地址
	dbusername = "root"        //用户名
	dbpassword = "root"    //密码
	dbname     = "user"           //表名
)

func Open() *sql.DB { //连接
	db, err := sql.Open("mysql", dbusername+":"+dbpassword+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
	}
	return db

}
