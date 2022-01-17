package discordBot

import (
	"fmt"
	"log"
)

var rs = make(chan string, 4)
var rst = make(chan string, 4)

//var rsMesId = make(chan string, 1)

func insertSborkzAll(lvlkz, timekz string, mesid string, name string, nameid string, guildid string, chatid string) { // внесение в базу данных
	db := conDbDs()
	log.Println("Запись в очередь дискорд ...")
	insertSborkz := `INSERT INTO sborkz(name, nameid,guildid,lvlkz,chatid,mesid,timedown,active) VALUES (?, ?, ?, ?, ?, ?,?,?)`
	statement, err := db.Prepare(insertSborkz)
	if err != nil {
		fmt.Println(err)
		db.Close()
	}
	_, err = statement.Exec(name, nameid, guildid, lvlkz, chatid, mesid, timekz, 0)
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
	}
	db.Close()
}

func readAll(lvlkz string, chatid string) string {
	MessageComplexID := "empty"
	db := conDbDs()
	results, err := db.Query("SELECT name,nameid,mesid,timedown FROM sborkz WHERE lvlkz = ? AND chatid = ? AND active = 0", lvlkz, chatid)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		var t Sborkzds
		err = results.Scan(&t.Name, &t.Nameid, &t.Mesid, &t.Timedown)
		rs <- fmt.Sprintf("%s", t.Nameid)
		rst <- fmt.Sprintf("%d", t.Timedown)
		MessageComplexID = t.Mesid
		return MessageComplexID
	}
	db.Close()
	return MessageComplexID
}

func updateActive(lvlkz string, chatid string) {
	db := conDbDs()
	_, err := db.Exec(`update sborkz set active = 1 where lvlkz = ? AND chatid = ?`, lvlkz, chatid)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("обновлено фулка")
	db.Close()

}
