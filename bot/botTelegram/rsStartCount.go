package botTelegram

import (
	"database/sql"
	"fmt"
)

//принудительный старт
func countR1Queue(db *sql.DB, lvlkz string, chatid int64,event int) {
	mesid := readAll(db, lvlkz, chatid);go SendDelMessage5s(chatid, mesid)
	name1:=<-rs;fmt.Printf("%s %d %d",name1,<-rst,<-rsn)
	numberLvl := numberQueueLvl(db, lvlkz, chatid)+1
	//event:=numberQueueChatid(db,chatid)
	var events string
	if event>0{ events=fmt.Sprintf("\nId для внесения очков  %d",event) }else if event==0{events=""}
	updateMesid(db, chatid, lvlkz, mesid)
	message := fmt.Sprintf(
		"Очередь кз%s (%d) была \nзапущена не полной \n\n1. @%s\nВ игру %s",
		lvlkz,numberLvl, name1,events)
	mesid = Send(chatid, message);
	if event==0{go SendDelMessage30m(chatid, mesid)}else if event>0{updateMesid(db,chatid,lvlkz,mesid)}
	numberUpdate(db, lvlkz, chatid)
	updateRsComplite(db, lvlkz, chatid,event)
	numberUpdateChatid(db,chatid)
}
func countR2Queue(db *sql.DB, lvlkz string, chatid int64,event int) {
	mesid := readAll(db, lvlkz, chatid);go SendDelMessage5s(chatid, mesid)
	name1:=<-rs;fmt.Printf("%s %d %d",name1,<-rst,<-rsn)
	name2:=<-rs;fmt.Printf("%s %d %d",name2,<-rst,<-rsn)
	numberLvl := numberQueueLvl(db, lvlkz, chatid)+1
	//event:=numberQueueChatid(db,chatid)
	var events string
	if event>0{ events=fmt.Sprintf("\nId для внесения очков  %d",event) }else if event==0{events=""}
	updateMesid(db, chatid, lvlkz, mesid)
	message := fmt.Sprintf(
		"Очередь кз%s (%d) была \nзапущена не полной \n\n1. @%s\n2. @%s\nВ игру %s",
		lvlkz,numberLvl, name1, name2, events)
	mesid = Send(chatid, message);
	if event==0{go SendDelMessage30m(chatid, mesid)}else if event>0{updateMesid(db,chatid,lvlkz,mesid)}
	numberUpdate(db, lvlkz, chatid)
	updateRsComplite(db, lvlkz, chatid,event)
	numberUpdateChatid(db,chatid)
}
func countR3Queue(db *sql.DB, lvlkz string, chatid int64,event int) {
	mesid := readAll(db, lvlkz, chatid);go SendDelMessage5s(chatid, mesid)
	name1:=<-rs;fmt.Printf("%s %d %d",name1,<-rst,<-rsn)
	name2:=<-rs;fmt.Printf("%s %d %d",name2,<-rst,<-rsn)
	name3:=<-rs;fmt.Printf("%s %d %d",name3,<-rst,<-rsn)
	numberLvl := numberQueueLvl(db, lvlkz, chatid)+1
	//event:=numberQueueChatid(db,chatid)
	var events string
	if event>0{ events=fmt.Sprintf("\nId для внесения очков  %d",event) }else if event==0{events=""}
	updateMesid(db, chatid, lvlkz, mesid)
	message := fmt.Sprintf(
		"Очередь кз%s (%d) была \nзапущена не полной \n\n1. @%s\n2. @%s\n3. @%s\nВ игру %s",
		lvlkz,numberLvl, name1, name2, name3, events)
	mesid = Send(chatid, message)
	if event==0{go SendDelMessage30m(chatid, mesid)}else if event>0{updateMesid(db,chatid,lvlkz,mesid)}
	numberUpdate(db, lvlkz, chatid)
	updateRsComplite(db, lvlkz, chatid,event)
	numberUpdateChatid(db,chatid)
}
