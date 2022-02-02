package botDiscord

import (
	"database/sql"
	"fmt"
	"log"
	"rs_bot/bot/botDiscord/databaseMysqlDs"
	"time"
)

func hhelp(mentionName, chatid string) string {
	mes, _ := DSBot.ChannelMessageSend(chatid, fmt.Sprintf("Справка для  %s \n"+
		"ВНИМАНИЕ БОТ УДАЛЯЕТ СООБЩЕНИЯ \n ОТ ПОЛЬЗОВАТЕЛЕЙ ЧЕРЕЗ 3 МИНУТЫ \n\n"+
		"Встать в очередь: [4-11]+  или\n"+
		" [4-11]+[указать время ожидания в минутах]\n"+
		"(уровень кз)+(время ожидания)\n"+
		" 9+  встать в очередь на КЗ 9ур.\n"+
		" 9+60  встать на КЗ 9ур, время ожидания не более 60 минут.\n"+
		"Покинуть очередь: [4-11] -\n"+
		" 9- выйти из очереди КЗ 9ур.\n"+
		"Посмотреть список активных очередей: о[4-11]\n"+
		" о9  вывод очередь для вашей Кз\n"+
		"Получить роль кз: + [5-11]\n"+
		" +9   получить роль КЗ 9ур.\n"+
		" -9   снять роль ", mentionName))
	return mes.ID
}
func hhelp1(chatid string) (mesid string, error error) {
	mes, err := DSBot.ChannelMessageSend(chatid, fmt.Sprintf("Справка \n"+
		"ВНИМАНИЕ БОТ УДАЛЯЕТ СООБЩЕНИЯ \n ОТ ПОЛЬЗОВАТЕЛЕЙ ЧЕРЕЗ 3 МИНУТЫ \n\n"+
		"Встать в очередь: [4-11]+  или\n"+
		" [4-11]+[указать время ожидания в минутах]\n"+
		"(уровень кз)+(время ожидания)\n"+
		" 9+  встать в очередь на КЗ 9ур.\n"+
		" 9+60  встать на КЗ 9ур, время ожидания не более 60 минут.\n"+
		"Покинуть очередь: [4-11] -\n"+
		" 9- выйти из очереди КЗ 9ур.\n"+
		"Посмотреть список активных очередей: о[4-11]\n"+
		" о9  вывод очередь для вашей Кз\n"+
		"Получить роль кз: + [5-11]\n"+
		" +9   получить роль КЗ 9ур.\n"+
		" -9   снять роль "))
	if err != nil {
		fmt.Println(err)
		return "", err
	} else if err == nil {
		mesid = mes.ID
		return mesid, nil
	}
	return mesid, error
}
func autohelp() {
	tm := time.Now()
	mtime := (tm.Format("15:04"))
	if mtime != "12:00" {
		db, er := databaseMysqlDs.DbConnection()
		if er != nil {
			log.Println(er)
		}
		results, err := db.Query("SELECT channel,mesidhelp FROM channel")
		if err != nil {
			log.Println(err)
		}
		var channel, mesidhelp string
		for results.Next() {
			err = results.Scan(&channel, &mesidhelp)
			if mesidhelp != "" {
				go Delete5s(channel, mesidhelp)
				helpChannelUpdate(db, channel)
			} else {
				helpChannelUpdate(db, channel)
			}
		}
		db.Close()

	}
}
func errorchanneldelete(db *sql.DB, chatid string) {
	_, err := db.Exec("delete from channel where channel = ? ", chatid)
	if err != nil {
		log.Println(err)
	}
}
func helpChannelUpdate(db *sql.DB, channel string) {
	newMesidHelp, err := hhelp1(channel)
	if err == nil {
		_, err := db.Exec(`update channel set mesidhelp = ? where channel = ? `, newMesidHelp, channel)
		if err == nil {
			log.Println(err)
		}
	} else {
		errorchanneldelete(db, channel)
	}
}
