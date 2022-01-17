package discordBot

import "log"

func countnq(lvlkz string, name string, chatid string) (int, int) {
	db := conDbDs()
	var countNames int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND lvlkz = ? AND chatid = ? AND active = 0", name, lvlkz, chatid)
	err := row.Scan(&countNames)
	if err != nil {
	}
	log.Println("imenDS", countNames)
	defer db.Close()
	var count int
	row = db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE lvlkz = ? AND chatid = ? AND active = 0", lvlkz, chatid)
	err = row.Scan(&count)
	if err != nil {
		log.Println("количество в очереди ошибка", count)
	}
	log.Println("количество в очереди ", count)
	db.Close()
	return countNames, count

}
func QueueDs(lvlkz, timekz string, mesid string, name string, nameid string, guildid string, chatid string) {
	countName, countQueue := countnq(lvlkz, name, chatid) //:=countName(lvlkz,name,chatid)
	if countName == 1 {
		dMessage, _ := DSBot.ChannelMessageSend(chatid, nameid+" ты уже в очереди")
		go Delete5s(chatid, dMessage.ID)
	} else {
		//countQueue:=countQueue(lvlkz,chatid)
		if countQueue == 0 {
			//отправить что ты первый в очереди
			counts0(lvlkz, timekz, name, nameid, guildid, chatid)
		} else if countQueue == 1 {
			counts1(lvlkz, timekz, name, nameid, guildid, chatid)
		} else if countQueue == 2 {
			counts2(lvlkz, timekz, name, nameid, guildid, chatid)
		} else if countQueue == 3 {
			counts3(lvlkz, timekz, name, nameid, guildid, chatid)
		}
	}

}
