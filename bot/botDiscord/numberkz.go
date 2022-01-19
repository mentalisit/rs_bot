package botDiscord

import (
	"database/sql"
	"fmt"
)

func readNumberkz(db *sql.DB, lvlkz string, chatid string) int {
	var number int
	row2 := db.QueryRow("SELECT  number FROM numkz WHERE lvlkz = ? AND chatid = ? ", lvlkz, chatid)
	err1 := row2.Scan(&number)
	if err1 != nil {
		fmt.Println(err1)
	}
	if number == 0 {
		insertSmt := "INSERT INTO numkz (lvlkz, number,chatid) VALUES (?, ?, ?)"
		_, err := db.Query(insertSmt, lvlkz, 1, chatid)
		if err != nil {
			fmt.Println(err)
		}
		number = 1
	} else {
		return number
	}
	return number
}

func updateNumberkz(db *sql.DB, lvlkz string, number int, chatid string) {
	fmt.Println(number)
	//number=number+1
	_, err := db.Exec(`update numkz set number=number+1 where lvlkz = ? AND chatid = ?`, lvlkz, chatid)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("обновлено номер кз ")

}
