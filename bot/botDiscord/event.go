package botDiscord

import (
	"database/sql"
	"fmt"
	"log"
)

func qweryNumevent1(db *sql.DB, chatid string) int {
	var event1 int
	row := db.QueryRow("SELECT numevent FROM rsevent WHERE chatid=? AND activeevent=1 ORDER BY numevent DESC LIMIT 1", chatid)
	err := row.Scan(&event1)
	if err != nil {
	}
	return event1
}
func qweryNumevent0(db *sql.DB, chatid string) int {
	var event0 int
	row := db.QueryRow("SELECT numevent FROM rsevent WHERE chatid=? AND activeevent=0 ORDER BY numevent DESC LIMIT 1", chatid)
	err := row.Scan(&event0)
	if err != nil {
	}
	return event0
}

// старт ивента

func EventStart(db *sql.DB, name string, chatid string) {
	if name == "Mentalisit" {
		//проверяем, есть ли активный ивент
		event1 := qweryNumevent1(db, chatid)
		log.Println("event", event1)
		if event1 > 0 {
			mes := SendChannel(chatid, "Режим ивента уже активирован.")
			go Delete5s(chatid, mes)
		} else {
			event0 := qweryNumevent0(db, chatid)
			if event0 > 0 {
				numberevent := event0 + 1
				insertEvent := `INSERT INTO rsevent (chatid,numevent,activeevent,number) VALUES (?,?,?,?)`
				_, err := db.Exec(insertEvent, chatid, numberevent, 1, 1)
				if err != nil {
					log.Println(err)
				}
			} else {
				insertEvent := `INSERT INTO rsevent (chatid,numevent,activeevent,number) VALUES (?,?,?,?)`
				_, err := db.Exec(insertEvent, chatid, 1, 1, 1)
				if err != nil {
					log.Println(err)
				}
			}
			mes := SendChannel(chatid, "Ивент запущен. После каждого похода на КЗ, "+
				"один из участников КЗ вносит полученные очки в базу командой К (номер катки) (количество набраных очков)")
			Delete1m(chatid, mes)
		}
	} else {
		mes := SendChannel(chatid, "Запуск|ОСтановка Ивента доступен Администратору канала.")
		Delete5s(chatid, mes)
	}
}

func EventStop(db *sql.DB, name string, chatid string) {
	if name == "Mentalisit" {
		event1 := qweryNumevent1(db, chatid)
		if event1 > 0 {
			//update
			_, err := db.Exec("UPDATE rsevent SET activeevent=0 WHERE chatid=? AND numevent=?", chatid, event1)
			if err != nil {
				log.Println(err)
			}
			mes := SendChannel(chatid, "Ивент остановлен.")
			go Delete5s(chatid, mes)
		} else {
			mes := SendChannel(chatid, "Ивент запущен не был. Нечего останавливать.")
			go Delete5s(chatid, mes)
		}
	} else {
		mes := SendChannel(chatid, "Запуск|Остановка Ивента доступен Администратору канала.")
		go Delete5s(chatid, mes)
	}
}

func countEventsPoints(db *sql.DB, chatid string, numberkz int) int {
	var countEventPoints int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE chatid=? AND numberkz=?  AND active=1 AND eventpoints > 0", chatid, numberkz)
	err := row.Scan(&countEventPoints)
	if err != nil {
	}
	return countEventPoints
}

// блок внесения очков за походы на КЗ во время ивента
func EventPoints(db *sql.DB, chatid string, name string, numberkz, points int) {
	// проверяем активен ли ивент
	event1 := qweryNumevent1(db, chatid)
	if event1 > 0 {
		//проверка, был ли участник который добавляет очки, в кз с указанным номером
		var countEventNames int
		row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE chatid=? AND numberkz=?  AND active=1 AND name=?", chatid, numberkz, name)
		err := row.Scan(&countEventNames)
		if err != nil {
		}
		if countEventNames > 0 {
			pointsGood := countEventsPoints(db, chatid, numberkz)
			if pointsGood > 0 {
				mes := SendChannel(chatid, "данные о кз уже внесены ")
				Delete5s(chatid, mes)
			} else {
				// считаем количество участников КЗ опр уровня
				var countEvent int
				row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE chatid=? AND numberkz=?  AND active=1", chatid, numberkz)
				err := row.Scan(&countEvent)
				if err != nil {
				}
				var pointsq int = points / countEvent
				//вносим очки
				_, err = db.Exec(`update sborkz set numberevent=?, eventpoints=? WHERE chatid=? AND numberkz=? AND active=1`, event1, pointsq, chatid, numberkz)
				if err != nil {
					log.Println(err)
				}
				mes := SendChannel(chatid, fmt.Sprintf("%s Очки %d внесены в базу", name, points))
				changeMessageEvent(db, points, countEvent, numberkz, event1, chatid)
				Delete5s(chatid, mes)
			}
		} else {
			mes := SendChannel(chatid, fmt.Sprintf("%s Вы не являетесь участником КЗ под номером %d добавление очков невозможно.", name, numberkz))
			Delete5s(chatid, mes)
		}
	} else {
		mes := SendChannel(chatid, "Ивент не запущен.")
		Delete5s(chatid, mes)
	}
}

func numberUpdateChatid(db *sql.DB, chatid string) {
	_, err := db.Exec(`update rsevent set number=number+1 where chatid = ? AND activeevent = 1`, chatid)
	if err != nil {
		log.Println(err)
	}
}

func numberQueueEvents(db *sql.DB, chatid string) int {

	var number int
	row2 := db.QueryRow("SELECT  number FROM rsevent WHERE activeevent = 1 AND chatid = ? ", chatid)
	err1 := row2.Scan(&number)
	if err1 != nil {
		fmt.Println(err1)
	}
	if number == 0 {
		//fmt.Println("нихрена нет ")
		insertSmt := "INSERT INTO rsevent(chatid,numevent,activeevent,number) VALUES (?, ?, ?, ?)"
		statement, err := db.Prepare(insertSmt)
		if err != nil {
		}
		number = 0
		_, err = statement.Exec(chatid, number)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		return number
	}
	return number
}

func changeMessageEvent(db *sql.DB, points, countEvent, numberkz, numberEvent int, chatid string) {
	results, err := db.Query("SELECT * FROM sborkz WHERE chatid=? AND numberkz=? AND numberevent = ? AND active=1", chatid, numberkz, numberEvent)
	if err != nil {
		fmt.Println(err)
	}
	var t Sborkz

	for results.Next() {
		err = results.Scan(&t.Id, &t.Name, &t.Nameid, &t.Mention, &t.Guildid, &t.Lvlkz, &t.Chatid, &t.Mesid, &t.Time, &t.Date, &t.numberkz, &t.numberevent, &t.eventpoints, &t.Timedown, &t.Active)
		rs <- t.Mention
	}
	cMEvent(t.Mesid, countEvent, points, numberkz, t.Chatid)

}
func cMEvent(mesid string, countEvent, points, numberkz int, chatid string) {
	mes1 := fmt.Sprintf("ивент игра №%d\n", numberkz)
	mesOld := fmt.Sprintf("внесено %d", points)
	if countEvent == 1 {
		text := fmt.Sprintf("%s%s \n%s", mes1, <-rs, mesOld)
		DSBot.ChannelMessageEdit(chatid, mesid, text)
	} else if countEvent == 2 {
		text := fmt.Sprintf("%s%s\n %s\n %s", mes1, <-rs, <-rs, mesOld)
		DSBot.ChannelMessageEdit(chatid, mesid, text)
	} else if countEvent == 3 {
		text := fmt.Sprintf("%s%s\n %s\n %s\n %s", mes1, <-rs, <-rs, <-rs, mesOld)
		DSBot.ChannelMessageEdit(chatid, mesid, text)
	} else if countEvent == 4 {
		text := fmt.Sprintf("%s%s\n %s\n %s %s\n %s", mes1, <-rs, <-rs, <-rs, mesOld)
		DSBot.ChannelMessageEdit(chatid, mesid, text)
	}
}

func event(db *sql.DB, chatid string) (string, int) {
	text := ""
	numE := 0
	//проверяем, есть ли активный ивент
	numberevent := qweryNumevent1(db, chatid)
	if numberevent == 0 { //ивент не активен
		return "", 0
	} else if numberevent > 0 { //активный ивент
		numE = numberQueueEvents(db, chatid) + 1 //номер кз number FROM rsevent
		text = fmt.Sprintf("\nID %d для ивента ", numE)
		numberUpdateChatid(db, chatid)
		return text, numE
	}
	return text, numE
}
