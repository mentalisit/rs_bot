package botTelegram

import "database/sql"

func RsStart(db *sql.DB, name string, lvlkz string, chatid int64) {
	countName := countName(db, lvlkz, name, chatid)
	event:=numberQueueChatid(db,chatid)
	if countName == 0 {
		mes := Send(chatid, "@"+name+" сначала добавься в очередь ")
		SendDelMessage5s(chatid, mes)
	} else {
		count := countQueue(db, lvlkz, chatid)
		if count == 1 {
			countR1Queue(db, lvlkz, chatid,event)
		} else if count == 2 {
			countR2Queue(db, lvlkz, chatid,event)
		} else if count == 3 {
			countR3Queue(db, lvlkz, chatid,event)
		}

	}
}
