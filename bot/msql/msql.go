package msql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type Msql struct {

}

func MysqlConn(msqlDataBase string) *sql.DB {
	str:=fmt.Sprintf("root:root@/%s",msqlDataBase)
	db,err :=sql.Open("mysql",str)

	if err != nil{
		fmt.Println(err)
	}
	return db
}

