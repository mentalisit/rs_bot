package botDiscord

import (
	"database/sql"
	"fmt"
)

func Minus(db *sql.DB, name string, chatid string) bool {
	in := true
	var countNames int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND chatid = ? AND active = 0", name, chatid)
	err := row.Scan(&countNames)
	if err != nil {
	}
	if countNames == 0 {
		mes := SendChannel(chatid, name+" ты не в очереди")
		go Delete5s(chatid, mes)
		in = false
	} else if countNames > 0 {
		results, err := db.Query("SELECT * FROM sborkz WHERE name = ? AND chatid = ? AND active = 0", name, chatid)
		if err != nil {
			fmt.Println(err)
		}
		for results.Next() {
			var t Sborkz
			err = results.Scan(&t.Id, &t.Name, &t.Nameid, &t.Mention, &t.Guildid, &t.Lvlkz, &t.Chatid, &t.Mesid, &t.Time, &t.Date, &t.numberkz, &t.numberevent, &t.eventpoints, &t.Timedown, &t.Active)
			if t.Name == name && t.Timedown > 3 && t.Chatid == chatid {
				mes := SendChannel(chatid, name+" рановато минус жмешь, ты в очереди на кз"+t.Lvlkz+" будешь еще "+fmt.Sprintf("%dмин", t.Timedown))
				go Delete5s(chatid, mes)
			} else if t.Name == name && t.Timedown <= 3 && t.Chatid == chatid {
				m := inMessage{
					mtext:       "",
					name:        t.Name,
					nameMention: t.Mention,
					nameid:      t.Nameid,
					mesid:       t.Mesid,
					guildid:     t.Guildid,
					chatid:      t.Chatid,
				}
				RsMinus(db, t.Lvlkz, &m)
			}
		}
	}
	return in
}
