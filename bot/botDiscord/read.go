package botDiscord

import (
	"database/sql"
	"fmt"
	"log"
)

func readAll(db *sql.DB, lvlkz string, chatid string) string {
	MessageComplexID := "empty"
	results, err := db.Query("SELECT name,nameid,mesid,timedown FROM sborkz WHERE lvlkz = ? AND chatid = ? AND active = 0", lvlkz, chatid)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		var t Sborkzds
		err = results.Scan(&t.Name, &t.Nameid, &t.Mesid, &t.Timedown)
		rs <- fmt.Sprintf("%s", t.Name)
		rst <- fmt.Sprintf("%d", t.Timedown)
		MessageComplexID = t.Mesid
		return MessageComplexID
	}
	return MessageComplexID
}

func readMesIDname(db *sql.DB, name, lvlkz, chatid string) string {
	mesid := ""
	results, err := db.Query("SELECT mesid FROM sborkz WHERE lvlkz = ? AND chatid = ? AND name = ? AND active = 0", lvlkz, chatid, name)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		//fmt.Println("подключаемся к дб дс месид")
		var t Sborkzds
		err = results.Scan(&t.Mesid)
		mesid = t.Mesid
		//fmt.Println(t.Mesid, "vvvvvvvvvvvv")
		return mesid
	}
	//fmt.Println(mesid, "ffffffff")
	return mesid
}
