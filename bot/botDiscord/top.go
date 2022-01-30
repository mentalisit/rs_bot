package botDiscord

import (
	"database/sql"
	"fmt"
	"log"
	"rs_bot/bot/botDiscord/databaseMysqlDs"
)

func topAll(db *sql.DB, chatid string) {
	databaseMysqlDs.CreateTableTempTop(db)
	mesage := "\xF0\x9F\x93\x96 ТОП Участников:\n"
	number := 1
	var good bool = false
	results, err := db.Query("SELECT name FROM sborkz WHERE chatid=? AND active=1 GROUP BY name ASC LIMIT 30", chatid)
	if err != nil {
		log.Println(err)
	}
	var t Sborkz
	for results.Next() {
		err = results.Scan(&t.Name)
		if len(t.Name) > 0 {
			good = true
			name := t.Name
			var countNames int
			row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND chatid = ? AND active = 1", name, chatid)
			err := row.Scan(&countNames)
			if err != nil {
				log.Println("imen", countNames, err)
			}
			insertTempTop := `INSERT INTO temptop(name,numkz) VALUES (?,?)`
			statement, err := db.Prepare(insertTempTop)
			if err != nil {
				fmt.Println(err)
			}
			_, err = statement.Exec(name, countNames)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	if good {
		resultsq, err := db.Query("SELECT * FROM temptop ORDER BY numkz DESC")
		if err != nil {
			log.Println(err)
		}
		var name string
		var numkz, id int
		var message2 string
		for resultsq.Next() {
			err = resultsq.Scan(&id, &name, &numkz)
			fmt.Println(number, name, numkz)
			message2 = message2 + fmt.Sprintf("%d. %s - %d \n", number, name, numkz)
			number = number + 1

		}
		mes := SendChannel(chatid, mesage+message2)
		go Delete1m(chatid, mes)
		_, err = db.Exec("DELETE FROM temptop")
		if err != nil {
			fmt.Println(err)
		}

	} else if !good {
		mes := SendChannel(chatid, " История не найдена ")
		Delete1m(chatid, mes)
	}
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
