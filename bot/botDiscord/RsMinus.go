package botDiscord

import (
	"database/sql"
)

func RsMinus(db *sql.DB, lvlkz string, m *inMessage) {
	countName := countName(db, lvlkz, m.name, m.chatid)
	if countName == 0 {
		mes := SendChannel(m.chatid, m.nameMention+" ты не в очереди.")
		go Delete5s(m.chatid, mes)
	} else if countName > 0 {
		//чтение айди очереди
		mesid := readMesIDname(db, m.name, lvlkz, m.chatid)
		//удаление с базы данных
		deleteSrorkz(db, m.name, lvlkz, m.chatid)
		mes := SendChannel(m.chatid, m.name+" покинул очередь ")
		go Delete5s(m.chatid, mes)
		//проверяем, есть ли кто в очереди
		countQueue := countQueue(db, lvlkz, m.chatid)
		numberkz := readNumberkz(db, lvlkz, m.chatid)
		// совподения количество  условие
		if countQueue == 0 {
			mes := SendChannel(m.chatid, "Очередь КЗ была удалена .")
			go Delete5s(m.chatid, mes)
			DSBot.ChannelMessageDelete(m.chatid, mesid)
		} else if countQueue == 1 {
			counts1r(db, m.chatid, m.guildid, lvlkz, numberkz)
		} else if countQueue == 2 {
			counts2r(db, m.chatid, m.guildid, lvlkz, numberkz)
		}
	}
}
