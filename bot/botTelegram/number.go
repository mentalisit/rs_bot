package botTelegram

import (
	"database/sql"
	"fmt"
	"log"
)

func numberQueueLvl(db *sql.DB, lvlkz string, chatid int64) int {
	var number int
	row2 := db.QueryRow("SELECT  number FROM numkz WHERE lvlkz = ? AND chatid = ? ", lvlkz, chatid)
	err1 := row2.Scan(&number)
	if err1 != nil {
		fmt.Println(err1)
	}
	if number == 0 {
		fmt.Println("нихрена нет ")
		insertSmt := "INSERT INTO numkz(lvlkz, number,chatid) VALUES (?, ?, ?)"
		statement, err := db.Prepare(insertSmt)
		if err != nil {
		}
		number = 0
		_, err = statement.Exec(lvlkz, number, chatid)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		return number
	}
	return number
}

func numberUpdate(db *sql.DB, lvlkz string, chatid int64) {
	_, err := db.Exec(`update numkz set number=number+1 where lvlkz = ? AND chatid = ?`, lvlkz, chatid)
	if err != nil {
		log.Println(err)
	}
}

func numberUpdateChatid(db *sql.DB, chatid int64) {
	_, err := db.Exec(`update rsevent set number=number+1 where chatid = ? AND activeevent = 1`, chatid)
	if err != nil {
		log.Println(err)
	}


}


func numberQueueChatid(db *sql.DB,chatid int64) int {
	var number int
	//var numEventString string
	activeEvent := qweryNumevent1(db, chatid)
	if activeEvent == 0 {
		fmt.Println("ивент не запущен ")
		//return ""
	} else if activeEvent > 0 {
		//получаем номер катки
		row2 := db.QueryRow("SELECT  number FROM rsevent WHERE chatid = ? AND activeevent = 1", chatid)
		err1 := row2.Scan(&number)
		if err1 != nil {
			fmt.Println(err1)}
		fmt.Println(number,"строка 61 должно писать ид катки ивента",number)

		if number <1 {
			fmt.Println("нихрена нет ")
			insertSmt := "INSERT INTO rsevent(chatid, numevent,activeevent,number,) VALUES (?, ?, ?, ?)"
			statement, err := db.Prepare(insertSmt)
			if err != nil {}
			number = 0
			_, err = statement.Exec(chatid,activeEvent, 1,0 )
			if err != nil {
				fmt.Println(err)}
			//numEventString=fmt.Sprintf("\n ивент Id %d",number)
			return number
		} else {
			//number=number+1
			//numEventString=fmt.Sprintf("\n ивент Id %d",number)
			return number
		}
	}
	return number
}