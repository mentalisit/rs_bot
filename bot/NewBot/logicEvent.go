package NewBot

import (
	"fmt"
	"log"
)

func qweryNumevent1(in inMessage) (event1 int) { //запрос номера ивента
	row := db.QueryRow("SELECT numevent FROM rsevent WHERE corpname=? AND activeevent=1 ORDER BY numevent DESC LIMIT 1", in.config.CorpName)
	err := row.Scan(&event1)
	if err != nil {
		fmt.Println(err)
		event1 = 0
	}
	return event1
}

func qweryNumevent0(in inMessage) (event0 int) { //запрос номера последнего ивента
	row := db.QueryRow("SELECT numevent FROM rsevent WHERE corpname=? AND activeevent=0 ORDER BY numevent DESC LIMIT 1", in.config.CorpName)
	err := row.Scan(&event0)
	if err != nil {
		fmt.Println(err)
	}
	return event0
}

func countEventsPoints(in inMessage, numberkz int) int {
	var countEventPoints int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE corpname=? AND numberkz=?  AND active=1 AND eventpoints > 0", in.config.CorpName, numberkz)
	err := row.Scan(&countEventPoints)
	if err != nil {
		fmt.Println(err)
	}
	return countEventPoints
}

func EventStart(in inMessage) {
	//проверяем, есть ли активный ивент
	event1 := qweryNumevent1(in)
	text := "Ивент запущен. После каждого похода на КЗ, " +
		"один из участников КЗ вносит полученные очки в базу командой К (номер катки) (количество набраных очков)"
	if in.tip == "ds" && (in.name == "Mentalisit" || checkAdmin(in.Ds.nameid, in.config.DsChannel)) {
		if event1 > 0 {
			dsSendChannelDel5s(in.config.DsChannel, "Режим ивента уже активирован.")
		} else {
			insrsevevent(in)
			if in.config.TgChannel != 0 {
				tgSendChannel(in.config.TgChannel, text)
				dsSendChannel(in.config.DsChannel, text)
			} else {
				dsSendChannel(in.config.DsChannel, text)
			}
		}
	} else if in.tip == "tg" && (in.name == "Mentalisit" || checkAdminTg(in)) {
		if event1 > 0 {
			tgSendChannelDel5s(in.config.TgChannel, "Режим ивента уже активирован.")
		} else {
			insrsevevent(in)
			if in.config.DsChannel != "" {
				dsSendChannel(in.config.DsChannel, text)
				tgSendChannel(in.config.TgChannel, text)
			} else {
				tgSendChannel(in.config.TgChannel, text)
			}
		}
	} else {
		text := "Запуск | Оcтановка Ивента доступен Администратору канала."
		if in.tip == "ds" {
			dsSendChannelDel5s(in.config.DsChannel, text)
		} else if in.tip == "tg" {
			tgSendChannelDel5s(in.config.TgChannel, text)
		}
	}
}

func insrsevevent(in inMessage) {
	event0 := qweryNumevent0(in)
	if event0 > 0 {
		numberevent := event0 + 1
		insertEvent := `INSERT INTO rsevent (corpname,numevent,activeevent,number) VALUES (?,?,?,?)`
		_, err := db.Exec(insertEvent, in.config.CorpName, numberevent, 1, 1)
		if err != nil {
			log.Println(err)
		}
	} else {
		insertEvent := `INSERT INTO rsevent (corpname,numevent,activeevent,number) VALUES (?,?,?,?)`
		_, err := db.Exec(insertEvent, in.config.CorpName, 1, 1, 1)
		if err != nil {
			log.Println(err)
		}
	}
}

func EventStop(in inMessage) {
	event1 := qweryNumevent1(in)
	if in.tip == "ds" && (in.name == "Mentalisit" || checkAdmin(in.Ds.nameid, in.config.DsChannel)) {
		if event1 > 0 {
			updateactive0(in, event1)
			go dsSendChannelDel5s(in.config.DsChannel, "Ивент остановлен.")
		} else {
			go dsSendChannelDel5s(in.config.DsChannel, "Ивент и так не активен. Нечего останавливать ")
		}
	} else if in.tip == "tg" && (in.name == "Mentalisit" || checkAdminTg(in)) {
		if event1 > 0 {
			//update
			_, err := db.Exec("UPDATE rsevent SET activeevent=0 WHERE tgchatid=? AND numevent=?", in.config.TgChannel, event1)
			if err != nil {
				log.Println(err)
			}
			go tgSendChannelDel5s(in.config.TgChannel, "Ивент остановлен.")
		} else {
			go tgSendChannelDel5s(in.config.TgChannel, "Ивент и так не активен. Нечего останавливать ")
		}
	} else {
		text := "Запуск|Остановка Ивента доступен Администратору канала."
		if in.tip == "ds" {
			go dsSendChannelDel5s(in.config.DsChannel, text)
		} else if in.tip == "tg" {
			go tgSendChannelDel5s(in.config.TgChannel, text)
		}
	}
}

func updateactive0(in inMessage, event1 int) {
	_, err := db.Exec("UPDATE rsevent SET activeevent=0 WHERE corpname=? AND numevent=?", in.config.CorpName, event1)
	if err != nil {
		log.Println(err)
	}
}

// блок внесения очков за походы на КЗ во время ивента
func EventPoints(in inMessage, numberkz, points int) {
	if in.tip == "ds" {
		dsDelMessage(in.config.DsChannel, in.Ds.mesid)
	} else if in.tip == "tg" {
		tgDelMessage(in.config.TgChannel, in.Tg.mesid)
	}
	// проверяем активен ли ивент
	event1 := qweryNumevent1(in)
	message := ""
	if event1 > 0 {
		var countEventNames int
		row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE corpname = ? AND numberkz=?  AND active=1 AND name=?", in.config.CorpName, numberkz, in.name)
		err := row.Scan(&countEventNames)
		if err != nil {
			fmt.Println(err)
		}
		admin := false
		if in.tip == "ds" {
			admin = checkAdmin(in.Ds.nameid, in.config.DsChannel)
		} else if in.tip == "tg" {
			admin = checkAdminTg(in)
		}
		if countEventNames > 0 || admin {
			pointsGood := countEventsPoints(in, numberkz)
			if pointsGood > 0 && !admin {
				message = "данные о кз уже внесены "
			} else if pointsGood == 0 || admin {
				countEvent := countIupdate(in, numberkz, points, event1)
				message = fmt.Sprintf("%s Очки %d внесены в базу", in.name, points)
				changeMessageEvent(in, points, countEvent, numberkz, event1)
			}
		} else {
			message = fmt.Sprintf("%s Вы не являетесь участником КЗ под номером %d добавление очков невозможно.", in.nameMention, numberkz)
		}

	} else {
		message = "Ивент не запущен."
	}
	if in.tip == "ds" {
		go dsSendChannelDel5s(in.config.DsChannel, message)
	} else if in.tip == "tg" {
		go tgSendChannelDel5s(in.config.TgChannel, message)
	}
}
func countIupdate(in inMessage, numberkz, points, event1 int) int {
	// считаем количество участников КЗ опр уровня
	var countEvent int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE numberevent = ? AND corpname=? AND numberkz=?  AND active=1", event1, in.config.CorpName, numberkz)
	err := row.Scan(&countEvent)
	if err != nil {
		fmt.Println(err)
	}
	var pointsq int = points / countEvent
	//вносим очки
	_, err = db.Exec(`update sborkz set eventpoints=? WHERE numberevent = ? AND corpname =? AND numberkz=? AND active=1`, pointsq, event1, in.config.CorpName, numberkz)
	if err != nil {
		log.Println(err)
	}
	return countEvent
}

func changeMessageEvent(in inMessage, points, countEvent, numberkz, numberEvent int) {
	var name string
	results, err := db.Query("SELECT * FROM sborkz WHERE corpname=? AND numberkz=? AND numberevent = ? AND active=1", in.config.CorpName, numberkz, numberEvent)
	if err != nil {
		fmt.Println(err)
	}
	var t sborkz
	var nd, nt Names

	num := 1
	for results.Next() {
		err = results.Scan(&t.id, &t.corpname, &t.name, &t.mention, &t.tip, &t.dsmesid, &t.tgmesid, &t.wamesid, &t.time, &t.date, &t.lvlkz, &t.numkzn, &t.numberkz, &t.numberevent, &t.eventpoints, &t.active, &t.timedown)
		if t.tip == "ds" {
			name = t.mention
		} else {
			name = t.name
		}
		if num == 1 {
			nd.name1 = name
		} else if num == 2 {
			nd.name2 = name
		} else if num == 3 {
			nd.name3 = name
		} else if num == 4 {
			nd.name4 = name
		}
		if t.tip == "tg" {
			name = t.mention
		} else {
			name = t.name
		}
		if num == 1 {
			nt.name1 = name
		} else if num == 2 {
			nt.name2 = name
		} else if num == 3 {
			nt.name3 = name
		} else if num == 4 {
			nt.name4 = name
		}
		num = num + 1
	}
	cMEvent(in, nd, nt, t, countEvent, points)
	//cMEventDS(t.dsmesid, countEvent, points, numberkz, t.dschatid)
	//cMEventTG(t.tgmesid, countEvent, points, numberkz, t.tgchatid)

}
func cMEvent(in inMessage, nd, nt Names, t sborkz, countEvent, points int) {
	mes1 := fmt.Sprintf("ивент игра №%d\n", t.numberkz)
	mesOld := fmt.Sprintf("внесено %d", points)
	if countEvent == 1 {
		if in.config.DsChannel != "" {
			DSBot.ChannelMessageEdit(in.config.DsChannel, t.dsmesid, fmt.Sprintf("%s %s \n%s", mes1, nd.name1, mesOld))
		}
		if in.config.TgChannel != 0 {
			tgEditText(in.config.TgChannel, t.tgmesid, fmt.Sprintf("%s %s \n%s", mes1, nt.name1, mesOld))
		}
	} else if countEvent == 2 {
		if in.config.DsChannel != "" {
			DSBot.ChannelMessageEdit(in.config.DsChannel, t.dsmesid, fmt.Sprintf("%s %s\n %s\n %s", mes1, nd.name1, nd.name2, mesOld))
		}
		if in.config.TgChannel != 0 {
			text := fmt.Sprintf("%s %s\n %s\n %s", mes1, nt.name1, nt.name2, mesOld)
			tgEditText(in.config.TgChannel, t.tgmesid, text)
		}
	} else if countEvent == 3 {
		if in.config.DsChannel != "" {
			text := fmt.Sprintf("%s %s\n %s\n %s\n %s", mes1, nd.name1, nd.name2, nd.name3, mesOld)
			DSBot.ChannelMessageEdit(in.config.DsChannel, t.dsmesid, text)
		}
		if in.config.TgChannel != 0 {
			text := fmt.Sprintf("%s %s\n %s\n %s\n %s", mes1, nt.name1, nt.name2, nt.name3, mesOld)
			tgEditText(in.config.TgChannel, t.tgmesid, text)
		}
	} else if countEvent == 4 {
		if in.config.DsChannel != "" {
			text := fmt.Sprintf("%s %s\n %s\n %s\n %s\n %s", mes1, nd.name1, nd.name2, nd.name3, nd.name4, mesOld)
			DSBot.ChannelMessageEdit(in.config.DsChannel, t.dsmesid, text)
		}
		if in.config.TgChannel != 0 {
			text := fmt.Sprintf("%s %s\n %s\n %s\n %s\n %s", mes1, nt.name1, nt.name2, nt.name3, nt.name4, mesOld)
			tgEditText(in.config.TgChannel, t.tgmesid, text)
		}
	}
}

func event(in inMessage) (string, int) {
	text := ""
	numE := 0
	//проверяем, есть ли активный ивент
	numberevent := qweryNumevent1(in)
	if numberevent == 0 { //ивент не активен
		return "", 0
	} else if numberevent > 0 { //активный ивент
		numE = numberQueueEvents(in) //номер кз number FROM rsevent
		text = fmt.Sprintf("\nID %d для ивента ", numE)
		//numberUpdateChatid(db, chatid)
		return text, numE
	}
	return text, numE
}

func numberQueueEvents(in inMessage) int {
	var number int
	row := db.QueryRow("SELECT  number FROM rsevent WHERE activeevent = 1 AND corpname = ? ", in.config.CorpName)
	err := row.Scan(&number)
	if err != nil {
		fmt.Println(err)
	}
	return number
}
