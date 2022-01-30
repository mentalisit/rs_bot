package botDiscord

import (
	"database/sql"
	"fmt"
	"log"
)

func readAll(db *sql.DB, lvlkz string, chatid string) string {
	results, err := db.Query("SELECT name,nameid,mention,mesid,timedown FROM sborkz WHERE lvlkz = ? AND chatid = ? AND active = 0", lvlkz, chatid)
	if err != nil {
		log.Println(err)
	}

	a := []string{}
	for results.Next() {
		var t Sborkz
		err = results.Scan(&t.Name, &t.Nameid, &t.Mention, &t.Mesid, &t.Timedown)
		rs <- fmt.Sprintf("%s", t.Mention)
		rst <- fmt.Sprintf("%d", t.Timedown)
		a = append(a, t.Mesid)
	}
	a = removeDuplicateElementString(a)
	return a[0]
}

func readMesIDname(db *sql.DB, name, lvlkz, chatid string) string {
	mesid := ""
	results, err := db.Query("SELECT mesid FROM sborkz WHERE lvlkz = ? AND chatid = ? AND name = ? AND active = 0", lvlkz, chatid, name)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		var t Sborkz
		err = results.Scan(&t.Mesid)
		mesid = t.Mesid
		return mesid
	}
	return mesid
}

func readMesID(db *sql.DB, mesid string) (string, error) {
	results, err := db.Query("SELECT lvlkz FROM sborkz WHERE mesid = ? AND active = 0", mesid)
	if err != nil {
		log.Println(err)
	}
	a := []string{}
	for results.Next() {
		var t Sborkz
		err = results.Scan(&t.Lvlkz)
		a = append(a, t.Lvlkz)
	}
	a = removeDuplicateElementString(a)
	return a[0], err
}
