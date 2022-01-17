package botTelegram

import "database/sql"

// проверка очереди

func Queue(db *sql.DB, lvlkz string, chatid int64) {
	count := countQueue(db, lvlkz, chatid)
	numberLvl := numberQueueLvl(db, lvlkz, chatid)+1
	// совподения количество  условие
	if count == 0 {
		mes0 := Send(chatid, "Очередь КЗ "+lvlkz+" пуста ")
		go SendDelMessage5s(chatid, mes0)
	} else if count == 1 {
		count1Queue(db, lvlkz, chatid,numberLvl)
	} else if count == 2 {
		count2Queue(db, lvlkz, chatid,numberLvl)
	} else if count == 3 {
		count3Queue(db, lvlkz, chatid,numberLvl)
	}
}
