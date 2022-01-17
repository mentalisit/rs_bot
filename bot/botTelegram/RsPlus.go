package botTelegram

import (
	"database/sql"
	"fmt"
)

func RsPlus(db *sql.DB, name, lvlkz, timekz string, chatid int64) {
	CountNames := countName(db, lvlkz, name, chatid)//проверяем есть ли игрок в очереди
	if CountNames == 1 {
		mes := Send(chatid, "@"+name+" ты уже в очереди")
		go SendDelMessage5s(chatid, mes)
	} else {
		countQueue := countQueue(db, lvlkz, chatid)//проверяем, есть ли кто-то в очереди
		numberkzName:=countNumberNameActive1(db,name,chatid,lvlkz)//проверяем количество боёв по уровню кз игрока
		numberQueueLv:=numberQueueLvl(db,lvlkz,chatid)+1 //проверяем какой номер боя определенной красной звезды
		if countQueue == 0 {
			text:=fmt.Sprintf("Очередь кз%s (%d)\n1. @%s - %sмин. - (%d)\n\n%s++ - принудительный старт",
				lvlkz,numberQueueLv, name, timekz,numberkzName, lvlkz)
			mesid := Sends(lvlkz, chatid, text)
			SubscribePing(db, name, lvlkz, chatid)
			//вносим в базу
			insertSborkztgAll(db, mesid, chatid, lvlkz, timekz, name,0)
		} else if countQueue == 1 {
			count1(db, lvlkz, timekz, name, chatid,numberkzName,numberQueueLv)
		} else if countQueue == 2 {
			SubscribePing3(db, name, lvlkz, chatid)
			count2(db, lvlkz, timekz, name, chatid,numberkzName,numberQueueLv)
		} else if countQueue == 3 {
			numberkz := numberQueueLvl(db, lvlkz, chatid)
			count3(db, lvlkz, timekz, name, chatid, numberkz,numberkzName)
		}
	}

}
