package botDiscord

import (
	"database/sql"
	"fmt"
	"log"
)

func updateActive1(db *sql.DB, lvlkz string, chatid string) {
	_, err := db.Exec(`update sborkz set active = 1 where lvlkz = ? AND chatid = ?`, lvlkz, chatid)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("обновлено фулка")

}
