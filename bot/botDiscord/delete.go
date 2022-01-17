package botDiscord

import (
	"database/sql"
	"log"
)

func deleteSrorkz(db *sql.DB, name, lvlkz, chatid string) {
	_, err := db.Exec("delete from sborkz where name = ? AND lvlkz = ? AND chatid = ? AND active = 0", name, lvlkz, chatid)
	if err != nil {
		log.Println(err)
	}
}
