package botTelegram

import (
	"database/sql"
	"fmt"
)

//проверка есть ли игрок в очереди
func countName(db *sql.DB, lvlkz string, name string, chatid int64) int {
	var countNames int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkztg WHERE name = ? AND lvlkz = ? AND chatid = ? AND active = 0", name, lvlkz, chatid)
	err := row.Scan(&countNames)
	if err != nil {}
	return countNames
}

//проверка сколько игровок в очереди
func countQueue(db *sql.DB, lvlkz string, chatid int64) int {
	var count int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkztg WHERE lvlkz = ? AND chatid = ? AND active = 0", lvlkz, chatid)
	err := row.Scan(&count)
	if err != nil {}
	return count
}



//обновляем айди сообщения в базе данных
func updateMesid(db *sql.DB, chatid int64, lvlkz string, mesid int) {
	_, err := db.Exec(`update sborkztg set mesid = ? where lvlkz = ? AND chatid = ? AND active = 0`, mesid, lvlkz, chatid)
	if err != nil {
		fmt.Println(err)
	}
}
// выковыриваем из базы значение количества походов на кз
func countNumberNameActive1(db *sql.DB,name string,chatid int64,lvlkz string)int{
	var countNumberNameActive1 int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkztg WHERE lvlkz = ? AND chatid = ? AND name = ? AND active = 1", lvlkz, chatid,name)
	err := row.Scan(&countNumberNameActive1)
	if err != nil {}
	return countNumberNameActive1
}
