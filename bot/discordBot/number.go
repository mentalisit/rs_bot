package discordBot

/*
import (
	"fmt"
	"tgbot/bot/msql"
)

func MsqlNumber(lvlkz string ,chatid string) int  {
	database:= msql.MysqlConn("sborkzds")
	var number int
	row2 := database.QueryRow("SELECT  number FROM numkz WHERE lvlkz = ? AND chatid = ? ",lvlkz, chatid)
	err1 := row2.Scan(&number)
	if err1!=nil {
		fmt.Println(err1)
	}
	defer database.Close()
	if number ==0 {
		insertSmt:= "INSERT INTO numkz (lvlkz, number,chatid) VALUES (?, ?, ?)"
		_, err := database.Query(insertSmt, lvlkz, 1, chatid)
		if err != nil {	fmt.Println(err) }
		number=1
	}else {
		return number
	}
	return number
}
func  MsqlNumberUpdate(lvlkz string,number int, chatid string)  {
	fmt.Println("vnosim v dazu numer")
	database:= msql.MysqlConn("sborkzds")
	fmt.Println(number)
	//number=number+1
	_,err:=database.Exec(`update numkz set number=number+1 where lvlkz = ? AND chatid = ?`,lvlkz,chatid )
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("обновлено")

}

*/
