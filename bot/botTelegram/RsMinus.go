package botTelegram

import (
	"database/sql"
)

func RsMinus(db *sql.DB, name, lvlkz string, chatid int64) {
	//проверяем, есть ли в очереди игрок
	CountNames := countName(db, lvlkz, name, chatid)
	if CountNames == 0 {
		mes := Send(chatid, " @"+name+" ты не в очереди.")
		go SendDelMessage5s(chatid, mes)
	} else { //проверяем, есть ли кто в очереди
		if CountNames > 0 {
			//удаление с базы данных
			mesid := readMesid(db, lvlkz, chatid)
			mes := Send(chatid, "@"+name+" покинул очередь ")
			go SendDelMessage5s(chatid, mes)
			deleteSborkz(db, name, lvlkz, chatid)
			//проверяем, есть ли кто в очереди
			countQueue := countQueue(db, lvlkz, chatid)
			numberQueueLv:=numberQueueLvl(db,lvlkz,chatid)

			if countQueue == 0 {
				mes0 := Send(chatid, "Очередь КЗ была удалена .")
				go SendDelMessage5s(chatid, mes0)
				go SendDelMessage5s(chatid, mesid)
			} else if countQueue == 1 {
				count1remove(db, lvlkz, chatid,numberQueueLv)
			} else if countQueue == 2 {
				count2remove(db, lvlkz, chatid,numberQueueLv)
			}

		}
	}
}
