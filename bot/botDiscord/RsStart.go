package botDiscord

import "database/sql"

//rs start
func RsStart(db *sql.DB, lvlkz string, name string, chatid string) {
	countName := countName(db, lvlkz, name, chatid)
	if countName == 0 {
		DSBot.ChannelMessageSend(chatid, "Принудительный старт доступен участникам очереди.")
	} else if countName == 1 {
		//numberkz := MsqlNumber(lvlkz, chatid)
		count := countQueue(db, lvlkz, chatid)

		if count == 1 {
			//counts1q(lvlkz,chatid,numberkz)
			//MsqlUpdate(lvlkz, chatid)
			//MsqlNumberUpdate(lvlkz, numberkz, chatid)
		} else if count == 2 {
			//Count2r(lvlkz,chatid,numberkz)
			//MsqlUpdate(lvlkz, chatid)
			//MsqlNumberUpdate(lvlkz, numberkz, chatid)

		} else if count == 3 {
			//Count3r(lvlkz,chatid,numberkz)
			//MsqlUpdate(lvlkz, chatid)
			//MsqlNumberUpdate(lvlkz, numberkz, chatid)

		}
	}
}
