package botTelegram

import (
	"database/sql"
	"fmt"
)

func Minus(db *sql.DB, name string, chatid int64) {
	var countNames int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkztg WHERE name = ? AND chatid = ? AND active = 0", name, chatid)
	err := row.Scan(&countNames)
	if err != nil {
	}
	if countNames == 0 {
		mes := Send(chatid, "@"+name+" ты не в очереди")
		go SendDelMessage5s(chatid, mes)
	} else if countNames > 0 {
		results, err := db.Query("SELECT * FROM sborkztg WHERE name = ? AND chatid = ? AND active = 0", name, chatid)
		if err != nil {
			fmt.Println(err)
		}
		for results.Next() {
			var tag Sborkz
			err = results.Scan(&tag.Id, &tag.Name, &tag.Mesid, &tag.Chatid, &tag.Time, &tag.Date, &tag.Lvlkz, &tag.Numberkz,&tag.Numberevent, &tag.Eventpoints, &tag.Active, &tag.Timedown, &tag.Activedel)
			if tag.Name == name && tag.Timedown > 3 && tag.Chatid == chatid {
				mes := Send(chatid, "@"+name+" рановато минус жмешь, ты в очереди на кз"+tag.Lvlkz+" будешь еще "+fmt.Sprintf("%dмин", tag.Timedown))
				go SendDelMessage5s(chatid, mes)
			} else if tag.Name == name && tag.Timedown <= 3 && tag.Chatid == chatid {
				RsMinus(db,tag.Name,tag.Lvlkz,tag.Chatid)
			}
		}
	}
}

