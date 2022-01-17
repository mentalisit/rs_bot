package botTelegram

import (
	"database/sql"
	"log"
	"time"
)

func insertSborkztgAll(db *sql.DB, mesid int, chatid int64, lvlkz string, timekz, name string,numberkz int) { // внесение в базу данных
	numevent:=qweryNumevent1(db,chatid)
	mdate,mtime:=currentTime()
	log.Println("Запись в очередь телеги ...")
	insertSborkztg1 := `INSERT INTO sborkztg(name,mesid,chatid,time,date,lvlkz,numberkz,numberevent,eventpoints,active,timedown,activedel) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := db.Exec(insertSborkztg1, name, mesid, chatid, mtime, mdate, lvlkz, numberkz, numevent, 0, 0, timekz, 0)
	if err != nil {
		log.Println(err)
	}
}
func currentTime()(string,string){
	tm:=time.Now()
	mdate := (tm.Format("2006-01-02"))
	mtime := (tm.Format("15:04"))
	return mdate,mtime
}