package botDiscord

import "database/sql"

func RsPlus(db *sql.DB, lvlkz, timekz string, m *inMessage) { //mesid string, name string, nameid string, guildid string, chatid string) {
	countName := countName(db, lvlkz, m.name, m.chatid)
	if countName == 1 {
		dMessage, _ := DSBot.ChannelMessageSend(m.chatid, m.nameMention+" ты уже в очереди")
		go Delete5s(m.chatid, dMessage.ID)
	} else {
		countQueue := countQueue(db, lvlkz, m.chatid)
		numberkz := readNumberkz(db, lvlkz, m.chatid)
		if countQueue == 0 {
			//отправить что ты первый в очереди
			counts0(db, lvlkz, timekz, m, numberkz)
		} else if countQueue == 1 {
			counts1(db, lvlkz, timekz, m, numberkz)
		} else if countQueue == 2 {
			counts2(db, lvlkz, timekz, m, numberkz)
		} else if countQueue == 3 {
			counts3(db, lvlkz, timekz, m, numberkz)
		}
	}

}
