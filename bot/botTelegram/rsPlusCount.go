package botTelegram

import (
	"database/sql"
	"fmt"
)

//проверка очереди при добавлении в очередь
func count1(db *sql.DB, lvlkz, timekz string, name string, chatid int64,numberkzName,numberQueueLv int) {
	mesid := readAll(db, lvlkz, chatid);go SendDelMessage5s(chatid, mesid)
	name1:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	newname:=fmt.Sprintf("%s - %sмин. (%d)",name,timekz,numberkzName)
	mesc1 := Send(chatid, "@"+name+" добавился в очередь кз"+lvlkz);go SendDelMessage5s(chatid, mesc1)
	text:=fmt.Sprintf("Очередь КЗ%s (%d) \n1. @%s\n2. @%s\n\n%s++ - принудительный старт",
								lvlkz,numberQueueLv,name1,newname,lvlkz)
	mesid = Sends(lvlkz, chatid, text)
	updateMesid(db, chatid, lvlkz, mesid)
	insertSborkztgAll(db, mesid, chatid, lvlkz, timekz, name,numberkzName)
}
func count2(db *sql.DB, lvlkz, timekz string, name string, chatid int64,numberkzName int,numberQueueLv int) {
	mesid := readAll(db, lvlkz, chatid);go SendDelMessage5s(chatid, mesid)
	name1:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	name2:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	newname:=fmt.Sprintf("%s - %sмин. (%d)",name,timekz,numberkzName)
	mesc2 := Send(chatid, "@"+name+" добавился в очередь кз"+lvlkz);go SendDelMessage5s(chatid, mesc2)
	text:=fmt.Sprintf("Очередь КЗ%s (%d) \n1. @%s\n2. @%s\n3. @%s\n\n%s++ - принудительный старт",
										lvlkz,numberQueueLv,name1,name2,newname,lvlkz)
	mesid = Sends(lvlkz, chatid, text)
	updateMesid(db, chatid, lvlkz, mesid)
	//вносим в базу
	insertSborkztgAll(db, mesid, chatid, lvlkz, timekz, name,numberkzName)
}
func count3(db *sql.DB, lvlkz, timekz string, name string, chatid int64, numberkz int,numberkzName int) {
	mesid := readAll(db, lvlkz, chatid);go SendDelMessage5s(chatid, mesid)
	name1:=<-rs; fmt.Printf("%s %d %d",name1,<-rst,<-rsn)
	name2:=<-rs; fmt.Printf("%s %d %d",name2,<-rst,<-rsn)
	name3:=<-rs; fmt.Printf("%s %d %d",name3,<-rst,<-rsn)
	mesc3 := Send(chatid, "@"+name+" закрыл очередь кз"+lvlkz);go SendDelMessage5s(chatid, mesc3)
	text:=fmt.Sprintf("Очередь КЗ%s сформирована\n1. @%s\n2. @%s\n3. @%s\n4. @%s\n В ИГРУ \n",
										lvlkz, name1, name2, name3, name)
	mesid = Send(chatid,text)
	updateMesid(db, chatid, lvlkz, mesid)
	//вносим в базу
	insertSborkztgAll(db, mesid, chatid, lvlkz, timekz, name,numberkzName)
	updateRsComplite(db, lvlkz, chatid,numberkz) //active1
	numberUpdate(db, lvlkz, chatid)     //number+1
}

