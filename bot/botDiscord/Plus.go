package botDiscord

import (
	"database/sql"
	"fmt"
	"log"
)

func Plus(db *sql.DB, name string, chatid string) bool {
	var countNames int
	in := true
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND chatid = ? AND active = 0", name, chatid)
	err := row.Scan(&countNames)
	if err != nil {
	}
	if countNames == 0 {
		mes := SendChannel(chatid, name+" ты не в очереди")
		go Delete5s(chatid, mes)
		in = false
	} else if countNames > 0 {
		fmt.Println(countNames)
		results, err := db.Query("SELECT * FROM sborkz WHERE name = ? AND chatid = ? AND active = 0", name, chatid)
		if err != nil {
			fmt.Println(err)
		}
		for results.Next() {
			var t Sborkz
			err = results.Scan(&t.Id, &t.Name, &t.Nameid, &t.Mention, &t.Guildid, &t.Lvlkz, &t.Chatid, &t.Mesid, &t.Time, &t.Date, &t.numberkz, &t.numberevent, &t.eventpoints, &t.Timedown, &t.Active)
			if t.Name == name && t.Timedown > 3 && t.Chatid == chatid {
				mes := SendChannel(chatid, name+" рановато плюсик жмешь, ты в очереди на кз"+t.Lvlkz+" будешь еще "+fmt.Sprintf("%dмин", t.Timedown))
				go Delete5s(chatid, mes)
			} else if t.Name == name && t.Timedown <= 3 && t.Chatid == chatid {
				mes := SendChannel(chatid, name+" время обновлено ")
				go Delete5s(chatid, mes)
				_, err := db.Exec("update sborkz set timedown = timedown + 30 where active = 0 AND name = ? AND chatid = ?", t.Name, t.Chatid)
				if err != nil {
					log.Println(err)
				}
				Queue(db, t.Lvlkz, t.Chatid, t.Guildid, true)
			}
		}
	}
	return in
}
