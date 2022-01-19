package botDiscord

import (
	"database/sql"
)

//rs start
func RsStart(db *sql.DB, lvlkz string, name string, chatid string, guildid string) {
	countName := countName(db, lvlkz, name, chatid)
	if countName == 0 {
		mes, _ := DSBot.ChannelMessageSend(chatid, "Принудительный старт доступен участникам очереди.")
		go Delete1m(chatid, mes.ID)
	} else if countName == 1 {
		numberkz := readNumberkz(db, lvlkz, chatid)
		count := countQueue(db, lvlkz, chatid)
		if count == 1 {
			countsS1(db, chatid, guildid, lvlkz, numberkz)
		} else if count == 2 {
			countsS2(db, chatid, guildid, lvlkz, numberkz)
		} else if count == 3 {
			countsS3(db, chatid, guildid, lvlkz, numberkz)
		}
	}
}
