package botDiscord

import (
	"database/sql"
	"fmt"
)

func RsMinus(db *sql.DB, name, nameid, lvlkz string, chatid, guildId string) {
	countName := countName(db, lvlkz, name, chatid)
	if countName == 0 {
		mes := SendChannel(chatid, nameid+" ты не в очереди.")
		go Delete5s(chatid, mes)
	} else if countName > 0 {
		//чтение айди очереди
		mesid := readMesIDname(db, name, lvlkz, chatid)
		fmt.Println(mesid, "qqqqqqqqqq")
		//удаление с базы данных
		deleteSrorkz(db, name, lvlkz, chatid)
		mes := SendChannel(chatid, nameid+" покинул очередь ")
		go Delete5s(chatid, mes)
		//проверяем, есть ли кто в очереди
		countQueue := countQueue(db, lvlkz, chatid)
		// совподения количество  условие
		if countQueue == 0 {
			mes := SendChannel(chatid, "Очередь КЗ была удалена .")
			go Delete5s(chatid, mes)
			go Delete5s(chatid, mesid)
		} else if countQueue == 1 {
			counts1r(db, chatid, guildId, lvlkz)
		} else if countQueue == 2 {
			counts2r(db, chatid, guildId, lvlkz)
		}
	}
}
