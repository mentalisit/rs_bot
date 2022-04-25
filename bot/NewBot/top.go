package NewBot

import (
	"fmt"
	"log"
)

func TopAll(in inMessage) {
	iftipdelete(in)
	numberevent := qweryNumevent1(in)
	if numberevent == 0 {
		topall(in)
	} else {
		topEvent(in)
	}
}
func TopLevel(in inMessage, lvlkz string) {
	iftipdelete(in)
	numberevent := qweryNumevent1(in)
	if numberevent == 0 {
		topLevel(in, lvlkz)
	} else {
		topEventLevel(in, lvlkz)
	}
}

func topEvent(in inMessage) { //TopEvent
	mesage := "\xF0\x9F\x93\x96 ТОП Участников Ивента:\n"
	number := 1
	numberevent := qweryNumevent1(in)
	var good bool = false
	results, err := db.Query("SELECT name FROM sborkz WHERE corpname=? AND numberevent = ? AND active=1 GROUP BY name ASC LIMIT 40", in.config.CorpName, numberevent)
	if err != nil {
		log.Println(err)
	}
	ifInSend5s(in, "Сканирую базу данных")
	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND active = 1 AND numberevent = ?", name, in.config.CorpName, numberevent)
			err := row.Scan(&countNames)
			if err != nil {
				log.Println("imen", countNames, err)
			}
			var points int
			row4 := db.QueryRow("SELECT  SUM(eventpoints) FROM sborkz WHERE name = ? AND corpname = ? AND active = 1 AND numberevent = ?", name, in.config.CorpName, numberevent)
			err4 := row4.Scan(&points)
			if err4 != nil {
				log.Println("points", points, err)
			}

			insertTempTopEvent := `INSERT INTO temptopevent(name,numkz,points) VALUES (?,?,?)`
			statement, err := db.Prepare(insertTempTopEvent)
			if err != nil {
				fmt.Println(err)
			}
			_, err = statement.Exec(name, countNames, points)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	ifInSend5s(in, "Формирую список ")
	if good {
		resultsq, err := db.Query("SELECT * FROM temptopevent ORDER BY points DESC")
		if err != nil {
			log.Println(err)
		}
		var name string
		var numkz, id, points int
		var message2 string
		for resultsq.Next() {
			err = resultsq.Scan(&id, &name, &numkz, &points)
			message2 = message2 + fmt.Sprintf("%d. %s - %d (%d)\n", number, name, numkz, points)
			number = number + 1
		}
		mesage = mesage + message2
		if in.tip == "ds" {
			dsSendChannelDel1m(in.config.DsChannel, mesage)
		} else if in.tip == "tg" {
			tgSendChannelDel1m(in.config.TgChannel, mesage)
		}

		_, err = db.Exec("DELETE FROM temptopevent")
		if err != nil {
			fmt.Println(err)
		}

	} else if !good {
		ifInSend5s(in, " История не найдена ")
	}
}
func topEventLevel(in inMessage, lvlkz string) { //TopEventLevel
	mesage := "\xF0\x9F\x93\x96 ТОП Участников ивента кз:" + lvlkz + "\n"
	number := 1
	numberevent := qweryNumevent1(in)
	var good bool = false
	results, err := db.Query("SELECT name FROM sborkz WHERE corpname=? AND numberevent = ? AND active=1  AND lvlkz = ? GROUP BY name ASC LIMIT 38", in.config.CorpName, numberevent, lvlkz)
	if err != nil {
		log.Println(err)
	}
	ifInSend5s(in, "Сканирую базу данных")
	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND active = 1 AND numberevent = ? AND lvlkz = ?", name, in.config.CorpName, numberevent, lvlkz)
			err := row.Scan(&countNames)
			if err != nil {
				log.Println("imen", countNames, err)
			}
			var points int
			row4 := db.QueryRow("SELECT  SUM(eventpoints) FROM sborkz WHERE name = ? AND corpname = ? AND active = 1 AND numberevent = ? AND lvlkz = ?", name, in.config.CorpName, numberevent, lvlkz)
			err4 := row4.Scan(&points)
			if err4 != nil {
				log.Println("points", points, err)
			}

			insertTempTopEvent := `INSERT INTO temptopevent(name,numkz,points) VALUES (?,?,?)`
			statement, err := db.Prepare(insertTempTopEvent)
			if err != nil {
				fmt.Println(err)
			}
			_, err = statement.Exec(name, countNames, points)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	ifInSend5s(in, "Формирую список ")
	if good {
		resultsq, err := db.Query("SELECT * FROM temptopevent ORDER BY points DESC")
		if err != nil {
			log.Println(err)
		}
		var name string
		var numkz, id, points int
		var message2 string
		for resultsq.Next() {
			err = resultsq.Scan(&id, &name, &numkz, &points)
			fmt.Println(number, name, numkz)
			message2 = message2 + fmt.Sprintf("%d. %s - %d (%d)\n", number, name, numkz, points)
			number = number + 1
		}
		mesage = mesage + message2
		if in.tip == "ds" {
			dsSendChannelDel1m(in.config.DsChannel, mesage)
		} else if in.tip == "tg" {
			tgSendChannelDel1m(in.config.TgChannel, mesage)
		}

		_, err = db.Exec("DELETE FROM temptopevent")
		if err != nil {
			fmt.Println(err)
		}

	} else if !good {
		ifInSend5s(in, " История не найдена ")
	}
}

func topEventDay(in inMessage) {

}

func topall(in inMessage) { //Top
	mesage := "\xF0\x9F\x93\x96 ТОП Участников:\n"
	number := 1
	var good bool = false
	results, err := db.Query("SELECT name FROM sborkz WHERE corpname=? AND active=1 GROUP BY name ASC LIMIT 40", in.config.CorpName)
	if err != nil {
		log.Println(err)
	}
	ifInSend5s(in, "Сканирую базу данных")
	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND active = 1", name, in.config.CorpName)
			err := row.Scan(&countNames)
			if err != nil {
				log.Println("imen", countNames, err)
			}

			insertTempTopEvent := `INSERT INTO temptopevent(name,numkz,points) VALUES (?,?,?)`
			statement, err := db.Prepare(insertTempTopEvent)
			if err != nil {
				fmt.Println(err)
			}
			_, err = statement.Exec(name, countNames, 0)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	ifInSend5s(in, "Формирую список ")
	if good {
		resultsq, err := db.Query("SELECT * FROM temptopevent ORDER BY numkz DESC")
		if err != nil {
			log.Println(err)
		}
		var name string
		var numkz, id, points int
		var message2 string
		for resultsq.Next() {
			err = resultsq.Scan(&id, &name, &numkz, &points)
			message2 = message2 + fmt.Sprintf("%d. %s - %d \n", number, name, numkz)
			number = number + 1
		}
		mesage = mesage + message2
		if in.tip == "ds" {
			dsSendChannelDel1m(in.config.DsChannel, mesage)
		} else if in.tip == "tg" {
			tgSendChannelDel1m(in.config.TgChannel, mesage)
		}

		_, err = db.Exec("DELETE FROM temptopevent")
		if err != nil {
			fmt.Println(err)
		}

	} else if !good {
		ifInSend5s(in, " История не найдена ")
	}
}

func topLevel(in inMessage, lvlkz string) { //TopLevel
	mesage := "\xF0\x9F\x93\x96 ТОП Участников кз:" + lvlkz + "\n"
	number := 1
	var good bool = false
	results, err := db.Query("SELECT name FROM sborkz WHERE corpname=? AND active=1  AND lvlkz = ? GROUP BY name ASC LIMIT 40", in.config.CorpName, lvlkz)
	if err != nil {
		log.Println(219, err)
	}
	ifInSend5s(in, "Сканирую базу данных")
	var name string
	for results.Next() {
		err = results.Scan(&name)
		if len(name) > 0 {
			good = true
			var countNames int
			row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND active = 1 AND lvlkz = ?", name, in.config.CorpName, lvlkz)
			err := row.Scan(&countNames)
			if err != nil {
				log.Println("imen", countNames, err)
			}
			insertTempTopEvent := `INSERT INTO temptopevent(name,numkz,points) VALUES (?,?,?)`
			statement, err := db.Prepare(insertTempTopEvent)
			if err != nil {
				fmt.Println(err)
			}
			_, err = statement.Exec(name, countNames, 0)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	ifInSend5s(in, "Формирую список ")
	if good {
		resultsq, err := db.Query("SELECT * FROM temptopevent ORDER BY numkz DESC")
		if err != nil {
			log.Println(err)
		}
		var name string
		var numkz, id, points int
		var message2 string
		for resultsq.Next() {
			err = resultsq.Scan(&id, &name, &numkz, &points)
			message2 = message2 + fmt.Sprintf("%d. %s - %d \n", number, name, numkz)
			number = number + 1
		}
		mesage = mesage + message2
		if in.tip == "ds" {
			dsSendChannelDel1m(in.config.DsChannel, mesage)
		} else if in.tip == "tg" {
			tgSendChannelDel1m(in.config.TgChannel, mesage)
		}

		_, err = db.Exec("DELETE FROM temptopevent")
		if err != nil {
			fmt.Println(err)
		}

	} else if !good {
		ifInSend5s(in, " История не найдена ")
	}
}

func topDay(in inMessage) {

}

/*

func top7day()  {
	results2, err := db.Query("SELECT * FROM sborkz WHERE name=? AND chatid=? AND active=1",name, chatid )
	if err != nil {		log.Println(err)	}
	var t Sborkz
	for results2.Next() {
		err = results2.Scan(&t.Id, &t.Name, &t.Nameid, &t.Mention, &t.Guildid, &t.Lvlkz, &t.Chatid, &t.Mesid, &t.Time, &t.Date, &t.numberkz, &t.numberevent, &t.eventpoints, &t.Timedown, &t.Active)

	}
func top24h()  {


}

*/

func ifInSend5s(in inMessage, text string) {
	if in.tip == "ds" {
		go dsDeleteMesage5s(in.config.DsChannel, text)
	} else if in.tip == "tg" {
		go tgSendChannelDel5s(in.config.TgChannel, text)
	}
}
