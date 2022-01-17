package botTelegram

import (
	"database/sql"
	"log"
)

func subscribe(db *sql.DB, name string, nameidid int64, lvlkz string, mesid int, chatid int64) {
	//проверка активной подписки
	var counts int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM subscribe WHERE name = ? AND lvlkz = ? AND chatid = ?", name, lvlkz, chatid)
	errs := row.Scan(&counts)
	if errs != nil {
		log.Println(errs)
	}

	if counts == 1 {
		mes := Send(chatid, "@"+name+" ты уже подписан на кз"+lvlkz+"\n для добавления в очередь напиши "+lvlkz+"+")
		go SendDelMessage5s(chatid, mes)
	} else {
		//добавление в оочередь пинга
		insertSubscribe := `INSERT INTO subscribe (name, nameid, lvlkz, messid, chatid, timestart, timeend) VALUES (?,?,?,?,?,?,?)`
		statement, err := db.Prepare(insertSubscribe)
		_, err = statement.Exec(name, nameidid, lvlkz, mesid, chatid, 0, 0)
		if err != nil {
			log.Println(err.Error())
		}
		mes := Send(chatid, "@"+name+" вы подписались на пинг кз"+lvlkz+"\n для добавления в очередь напиши "+lvlkz+"+")
		go SendDelMessage5s(chatid, mes)
	}
}

// отписаться от пинга
func unsubscribe(db *sql.DB, name string, nameidid int64, lvlkz string, mesid int, chatid int64) {
	//проверка активной подписи
	var counts int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM subscribe WHERE name = ? AND lvlkz = ? AND chatid = ?", name, lvlkz, chatid)
	errs := row.Scan(&counts)
	if errs != nil {
		log.Println(errs.Error())
	}

	if counts == 0 {
		mes := Send(chatid, "@"+name+" ты не подписан на пинг")
		go SendDelMessage5s(chatid, mes)
	} else if counts == 1 {
		//удаление с базы данных
		_, err := db.Exec("delete from subscribe where name = ? AND lvlkz = ? AND chatid = ? ", name, lvlkz, chatid)
		mes := Send(chatid, "@"+name+" отписался от пинга ")
		go SendDelMessage5s(chatid, mes)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
func SubscribePing(db *sql.DB, name string, lvlkz string, chatid int64) {
	var name1, names, men string
	if rows, err := db.Query("SELECT name FROM subscribe WHERE lvlkz = ? AND chatid = ?", lvlkz, chatid); err == nil {
		for rows.Next() {
			rows.Scan(&name1)
			if name == name1 {
				continue
			}
			names = " @" + name1 + " "
			men = names + men
		}
		rows.Close()
	}
	mes := Send(chatid, men)
	go SendDelMessage30m(chatid, mes)
}

//         3/4
func subscribe3(db *sql.DB, name string, nameidid int64, lvlkz string, mesid int, chatid int64) {
	//проверка активной подписки
	var counts int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM subscribe3 WHERE name = ? AND lvlkz = ? AND chatid = ?", name, lvlkz, chatid)
	errs := row.Scan(&counts)
	if errs != nil {
		log.Println(errs)
	}

	if counts == 1 {
		mes := Send(chatid, "@"+name+" ты уже подписан на кз"+lvlkz+" 3/4\n для добавления в очередь напиши "+lvlkz+"+")
		go SendDelMessage5s(chatid, mes)
	} else {
		//добавление в оочередь пинга
		insertSubscribe := `INSERT INTO subscribe3 (name, nameid, lvlkz, messid, chatid, timestart, timeend) VALUES (?,?,?,?,?,?,?)`
		statement, err := db.Prepare(insertSubscribe)
		_, err = statement.Exec(name, nameidid, lvlkz, mesid, chatid, 0, 0)
		if err != nil {
			log.Println(err.Error())
		}
		mes := Send(chatid, "@"+name+" вы подписались на пинг кз"+lvlkz+" 3/4\n для добавления в очередь напиши "+lvlkz+"+")
		go SendDelMessage5s(chatid, mes)
	}
}

// отписаться от пинга
func unsubscribe3(db *sql.DB, name string, nameidid int64, lvlkz string, mesid int, chatid int64) {
	//проверка активной подписи
	var counts int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM subscribe3 WHERE name = ? AND lvlkz = ? AND chatid = ?", name, lvlkz, chatid)
	errs := row.Scan(&counts)
	if errs != nil {
		log.Println(errs.Error())
	}

	if counts == 0 {
		mes := Send(chatid, "@"+name+" ты не подписан на пинг  3/4")
		go SendDelMessage5s(chatid, mes)
	} else if counts == 1 {
		//удаление с базы данных
		_, err := db.Exec("delete from subscribe3 where name = ? AND lvlkz = ? AND chatid = ? ", name, lvlkz, chatid)
		mes := Send(chatid, "@"+name+" отписался от пинга  3/4")
		go SendDelMessage5s(chatid, mes)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
func SubscribePing3(db *sql.DB, name string, lvlkz string, chatid int64) {
	var name1, names, men string
	if rows, err := db.Query("SELECT name FROM subscribe3 WHERE lvlkz = ? AND chatid = ?", lvlkz, chatid); err == nil {
		for rows.Next() {
			rows.Scan(&name1)
			if name == name1 {
				continue
			}
			names = " @" + name1 + " "
			men = names + men
		}
		rows.Close()
	}
	if len(men) > 0 {
		mes := Send(chatid, "3/4 "+men)
		go SendDelMessage30m(chatid, mes)
	}

}
