package botDiscord

import (
	"database/sql"
	"fmt"
)

func insertSborkzAll(db *sql.DB, lvlkz, timekz string, mesid string, name string, nameid string, guildid string, chatid string) { // внесение в базу данных
	//log.Println("Запись в очередь дискорд ...")
	insertSborkz := `INSERT INTO sborkz(name, nameid,guildid,lvlkz,chatid,mesid,timedown,active) VALUES (?, ?, ?, ?, ?, ?,?,?)`
	statement, err := db.Prepare(insertSborkz)
	if err != nil {		fmt.Println(err)}
	_, err = statement.Exec(name, nameid, guildid, lvlkz, chatid, mesid, timekz, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
}
