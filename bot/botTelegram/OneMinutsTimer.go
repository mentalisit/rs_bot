package botTelegram

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"rs_bot/bot/botTelegram/databaseMysql"
)

func oneMinutsTimer(db *sql.DB) {
	var count int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkztg WHERE active = 0")
	err := row.Scan(&count)
	if err != nil {log.Println(err.Error())}
	if count > 0 {
		a := []int{}
		aa := []int{}
		results, err := db.Query("SELECT name,mesid,timedown FROM sborkztg WHERE active = 0")
		if err != nil {
			log.Println(err)
		}
		var tag Sborkz
		for results.Next() {
			err = results.Scan(&tag.Name, &tag.Mesid, &tag.Timedown)
			a = append(a, tag.Mesid)
			a = removeDuplicateElementInt(a)
		}
		for _, v := range a {
			skip := false
			for _, u := range aa {
				if v == u {
					skip = true
					break
				}
			}
			if !skip {
				messageupdate(db, v)
			}
		}
	}
}

func messageupdate(db *sql.DB, v int) {
	var count40 int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkztg WHERE mesid = ? AND active = 0", v)
	err := row.Scan(&count40)
	if err != nil {log.Println(err.Error())}
	if count40 == 1 {
		results, err := db.Query("SELECT lvlkz,chatid FROM sborkztg WHERE mesid = ? AND active = 0", v)
		if err != nil {log.Println(err)}
		var tag Sborkz
		for results.Next() {err = results.Scan(&tag.Lvlkz, &tag.Chatid)}
		count1update(db, v, tag.Lvlkz, tag.Chatid)

	} else if count40 == 2 {
		results, err := db.Query("SELECT lvlkz,chatid FROM sborkztg WHERE mesid = ? AND active = 0", v)
		if err != nil {log.Println(err)}
		var tag Sborkz
		for results.Next() {err = results.Scan(&tag.Lvlkz, &tag.Chatid)}
		count2update(db, v, tag.Lvlkz, tag.Chatid)

	} else if count40 == 3 {
		results, err := db.Query("SELECT lvlkz,chatid FROM sborkztg WHERE mesid = ? AND active = 0", v)
		if err != nil {	log.Println(err)}
		var tag Sborkz
		for results.Next() {err = results.Scan(&tag.Lvlkz, &tag.Chatid)}
		count3update(db, v, tag.Lvlkz, tag.Chatid)

	}
}

func count1update(db *sql.DB, v int, lvlkz string, chatid int64) {
	readAll(db, lvlkz, chatid)
	numberQueueLv:=numberQueueLvl(db,lvlkz,chatid)+1
	name1 := fmt.Sprintf("%s - %dмин. (%d)", <-rs, <-rst, <-rsn)
	text := fmt.Sprintf(
		"Очередь кз%s (%d)\n1. %s\n\n%s++ - принудительный старт",
		lvlkz,numberQueueLv,name1,lvlkz)
	EditMessageText(chatid, v, text,lvlkz)
}
func count2update(db *sql.DB, v int, lvlkz string, chatid int64) {
	readAll(db, lvlkz, chatid)
	numberQueueLv:=numberQueueLvl(db,lvlkz,chatid)+1
	name1:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	name2:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	text := fmt.Sprintf(
		"Очередь кз%s (%d)\n1. %s\n2. %s\n\n%s++ - принудительный старт",
		lvlkz,numberQueueLv,name1,name2,lvlkz)
	EditMessageText(chatid, v, text, lvlkz)
}
func count3update(db *sql.DB, v int, lvlkz string, chatid int64) {
	readAll(db, lvlkz, chatid)
	numberQueueLv:=numberQueueLvl(db,lvlkz,chatid)+1
	name1:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	name2:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	name3:=fmt.Sprintf("%s - %dмин. (%d)",<-rs,<-rst,<-rsn)
	text := fmt.Sprintf(
		"Очередь кз%s (%d)\n1. %s\n2. %s\n3. %s\n\n%s++ - принудительный старт",
		lvlkz,numberQueueLv,name1,name2,name3,lvlkz)

	EditMessageText(chatid, v, text, lvlkz)
}

//обновление  -1
func MinusMin() {
	db, er := databaseMysql.DbConnection()
	if er != nil {
		logrus.Println(er)
	}
	_, err := db.Exec(`update sborkztg set timedown = timedown - 1 where active = 0`)
	if err != nil {
		msqlTimeo(db)
		fmt.Println("errrrr")
	}
	msqlTimeo(db)
}

//проверка оставшихся минут
func msqlTimeo(db *sql.DB) {
	results, err := db.Query("SELECT * FROM sborkztg WHERE active = 0")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var tag Sborkz
		err = results.Scan(&tag.Id, &tag.Name, &tag.Mesid, &tag.Chatid, &tag.Time, &tag.Date, &tag.Lvlkz, &tag.Numberkz,&tag.Numberevent, &tag.Eventpoints, &tag.Active, &tag.Timedown, &tag.Activedel)
		if tag.Timedown == 3 {
			mes3s := SendP(tag.Chatid, "@"+tag.Name+" время почти вышло  ...\n если ты еще тут пиши +")
			go SendDelMessage3m(tag.Chatid, mes3s)

		} else if tag.Timedown == 0 {
			RsMinus(db, tag.Name, tag.Lvlkz, tag.Chatid)
		} else if tag.Timedown <= -1 {
			RsMinus(db, tag.Name, tag.Lvlkz, tag.Chatid)
		}
	}
	oneMinutsTimer(db)
}
