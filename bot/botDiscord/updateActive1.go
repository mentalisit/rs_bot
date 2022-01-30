package botDiscord

import (
	"database/sql"
	"log"
)

func updateActive1(db *sql.DB, lvlkz string, chatid string, mesidOld string, mesid string, numberkz int) {
	numberevent := qweryNumevent1(db, chatid)
	_, err := db.Exec(`update sborkz set active = 1, mesid = ?,numberkz = ?,numberevent = ? where lvlkz = ? AND chatid = ? AND mesid =?`,
		mesid, numberkz, numberevent, lvlkz, chatid, mesidOld)
	if err != nil {
		log.Println(err)
	}
}

func updateMesid(db *sql.DB, lvlkz string, chatid string, mesidOld string, mesid string) {
	_, err := db.Exec(`update sborkz set  mesid = ? where lvlkz = ? AND chatid = ? AND mesid =?`,
		mesid, lvlkz, chatid, mesidOld)
	if err != nil {
		log.Println(err)
	}
}
