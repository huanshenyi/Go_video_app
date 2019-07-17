package dbops

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
var (
	dbConn *sql.DB
	err error
)

func init()  {
	dbConn,err = sql.Open("mysql","root:root!@#tcp(52.198.99.16:3306)/govideo?charset=utf8")
	if err != nil{
		panic(err.Error())
	}
}
