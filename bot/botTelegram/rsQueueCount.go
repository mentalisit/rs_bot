package botTelegram

import (
	"database/sql"
	"fmt"
)

//запрос очереди
func count1Queue(db *sql.DB, lvlkz string, chatid int64,numberLvl int) {
	mesid:=readAll(db, lvlkz, chatid);go SendDelMessage5s(chatid,mesid)
	name1 := fmt.Sprintf("%s - %dмин. (%d)", <-rs, <-rst, <-rsn)
	text := fmt.Sprintf(
		"Очередь кз%s (%d)\n1. %s\n\n%s++ - принудительный старт",
							lvlkz,numberLvl,name1,lvlkz)
	mesid = Sends(lvlkz, chatid, text)
	updateMesid(db, chatid, lvlkz, mesid)
}
func count2Queue(db *sql.DB, lvlkz string, chatid int64,numberLvl int) {
	mesid := readAll(db, lvlkz, chatid);go SendDelMessage5s(chatid,mesid)
	name1:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	name2:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	text := fmt.Sprintf(
		"Очередь кз%s (%d)\n1. %s\n2. %s\n\n%s++ - принудительный старт",
							lvlkz,numberLvl,name1,name2,lvlkz)
	mesid=Sends(lvlkz,chatid,text)
	updateMesid(db,chatid,lvlkz,mesid)
}
func count3Queue(db *sql.DB, lvlkz string, chatid int64,numberLvl int) {
	mesid := readAll(db, lvlkz, chatid);go SendDelMessage5s(chatid,mesid)
	name1:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	name2:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	name3:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	text := fmt.Sprintf(
		"Очередь кз%s (%d)\n1. %s\n2. %s\n3. %s\n\n%s++ - принудительный старт",
							lvlkz,numberLvl,name1,name2,name3,lvlkz)
	mesid=Sends(lvlkz,chatid,text)
	updateMesid(db,chatid,lvlkz,mesid)
}
