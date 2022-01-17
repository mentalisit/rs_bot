package botTelegram

import (
	"database/sql"
	"fmt"
)

//проверка очереди после удаления или запрос очереди
func count1remove(db *sql.DB, lvlkz string, chatid int64,numberQueueLv int) {
	mesid:=readAll(db, lvlkz, chatid);go SendDelMessage5s(chatid,mesid)
	name1:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	text:=fmt.Sprintf(
		"Очередь кз%s (%d)\n1. @%s\n\n%s++ - принудительный старт",
								lvlkz,numberQueueLv,name1,lvlkz)
	mesid=Sends(lvlkz,chatid,text)
	updateMesid(db,chatid,lvlkz,mesid)
}
func count2remove(db *sql.DB, lvlkz string, chatid int64,numberQueueLv int) {
	mesid := readAll(db, lvlkz, chatid);go SendDelMessage5s(chatid,mesid)
	name1:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	name2:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	text:=fmt.Sprintf(
		"Очередь кз%s (%d)\n1. @%s\n2. @%s\n\n%s++ - принудительный старт",
								lvlkz,numberQueueLv,name1,name2,lvlkz)
	mesid=Sends(lvlkz,chatid,text)
	updateMesid(db,chatid,lvlkz,mesid)
}

