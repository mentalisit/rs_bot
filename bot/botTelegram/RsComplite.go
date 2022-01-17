package botTelegram

import (
	"database/sql"
	"fmt"
	"log"
)

func updateRsComplite(db *sql.DB, lvlkz string, chatid int64,numberkz int) {
	numEvent:=qweryNumevent1(db,chatid)
	fmt.Println("updateRsComplite", numEvent)
	_, err := db.Exec(
		`update sborkztg set active = 1, activedel=0,numberkz=?,numberevent=? where lvlkz = ? AND chatid = ? AND active = 0`,
								numberkz,numEvent, lvlkz, chatid)
	if err != nil {
		log.Println(err)
	}
}
