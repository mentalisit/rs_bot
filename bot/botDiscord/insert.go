package botDiscord

import (
	"database/sql"
	"fmt"
	"time"
)

func insertSborkzAll(db *sql.DB, lvlkz, timekz string, mesid string, name string, nameid string, guildid string, chatid string, mention string) { // внесение в базу данных
	//log.Println("Запись в очередь дискорд ...")
	mdata, mtime := currentTime()
	insertSborkz := `INSERT INTO sborkz(name, nameid,mention,guildid,lvlkz,chatid,mesid,time,date,numberkz,numberevent,eventpoints,timedown,active) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	statement, err := db.Prepare(insertSborkz)
	if err != nil {
		fmt.Println(err)
	}
	_, err = statement.Exec(name, nameid, mention, guildid, lvlkz, chatid, mesid, mtime, mdata, 0, 0, 0, timekz, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func currentTime() (string, string) {
	tm := time.Now()
	mdate := (tm.Format("2006-01-02"))
	mtime := (tm.Format("15:04"))
	return mdate, mtime
}
