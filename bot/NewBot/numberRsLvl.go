package NewBot

import (
	"fmt"
)

func numberQueueLvl(in inMessage, lvlkz string) int {
	var number int
	row2 := db.QueryRow("SELECT  number FROM numkz WHERE lvlkz = ? AND corpname = ?", lvlkz, in.config.CorpName)
	err2 := row2.Scan(&number)
	if err2 != nil {
		fmt.Println(err2)
	}
	if number == 0 {
		insertSmt := "INSERT INTO numkz(lvlkz, number,corpname) VALUES (?,?,?)"
		statement, err := db.Prepare(insertSmt)
		if err != nil {
			fmt.Println(err)
		}
		number = 1
		_, err = statement.Exec(lvlkz, number, in.config.CorpName)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		return number
	}
	return number
}

func updateNumberkz(in inMessage, lvlkz string) { //обновление номера кз
	_, err := db.Exec(`update numkz set number=number+1 where lvlkz = ? AND corpname = ?`, lvlkz, in.config.CorpName)
	if err != nil {
		fmt.Println(err)
	}
}
