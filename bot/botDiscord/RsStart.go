package botDiscord

import (
	"database/sql"
	"fmt"
)

//rs start
func RsStart(db *sql.DB, lvlkz string, name string, chatid string, guildid string) {
	countName := countName(db, lvlkz, name, chatid)
	if countName == 0 {
		DSBot.ChannelMessageSend(chatid, "Принудительный старт доступен участникам очереди.")
	} else if countName == 1 {
		numberkz := readNumberkz(db, lvlkz, chatid)
		fmt.Println(numberkz)
		count := countQueue(db, lvlkz, chatid)

		if count == 1 {
			countsS1(db, chatid, guildid, lvlkz)
			fmt.Println("gi")
			//MsqlUpdate(lvlkz, chatid)
			//MsqlNumberUpdate(lvlkz, numberkz, chatid)
		} else if count == 2 {
			countsS2(db, chatid, guildid, lvlkz)
			//MsqlUpdate(lvlkz, chatid)
			//MsqlNumberUpdate(lvlkz, numberkz, chatid)

		} else if count == 3 {
			countsS3(db, chatid, guildid, lvlkz)
			//MsqlUpdate(lvlkz, chatid)
			//MsqlNumberUpdate(lvlkz, numberkz, chatid)

		}
	}
}
