package botDiscord

import (
	"database/sql"
	"fmt"
	"log"
)

func RsPlus(db *sql.DB, lvlkz, timekz string, m *inMessage) {
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

func Pl30(db *sql.DB, lvlkz string, m *inMessage) {
	countName := countName(db, lvlkz, m.name, m.chatid)
	if countName == 0 {
		dmess, _ := DSBot.ChannelMessageSend(m.chatid, m.nameMention+" ты не в очереди ")
		go Delete5s(m.chatid, dmess.ID)
	} else if countName > 0 {
		dmes, _ := DSBot.ChannelMessageSend(m.chatid, m.nameMention+" время обновлено +30")
		go Delete5s(m.chatid, dmes.ID)
		fmt.Println("mmm ", timekz, lvlkz, m.chatid, m.name)
		_, err := db.Exec(`update sborkz set timedown = timedown+30 where lvlkz = ? AND chatid = ?  AND name = ?`,
			lvlkz, m.chatid, m.name)
		if err != nil {
			log.Println(err)
		}
		Queue(db, lvlkz, m.chatid, m.guildid, true)
	}
}
