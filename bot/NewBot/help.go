package NewBot

import (
	"fmt"
	"log"
	"time"
)

const hhelpText = "В боте используются только кирилица \n" +
	"Встать в очередь: [4-11]+  или\n" +
	" [4-11]+[указать время ожидания в минутах]\n" +
	"(уровень кз)+(время ожидания)\n" +
	" 9+  встать в очередь на КЗ 9ур.\n" +
	" 9+60  встать на КЗ 9ур, время ожидания не более 60 минут.\n" +
	"Покинуть очередь: [4-11] -\n" +
	" 9- выйти из очереди КЗ 9ур.\n" +
	"Посмотреть список активных очередей: о[4-11]\n" +
	" о9 вывод очередь для вашей Кз\n" +
	"Получить роль кз: + [5-11]\n" +
	" +9 получить роль КЗ 9ур.\n" +
	" -9 снять роль "

func hhelpName(in inMessage) {
	iftipdelete(in)
	if in.tip == "ds" {
		mes, _ := DSBot.ChannelMessageSend(in.config.DsChannel, fmt.Sprintf("Справка для  %s \n"+
			"ВНИМАНИЕ БОТ УДАЛЯЕТ СООБЩЕНИЯ \n ОТ ПОЛЬЗОВАТЕЛЕЙ ЧЕРЕЗ 3 МИНУТЫ \n\n"+hhelpText, in.nameMention))
		dsDeleteMesageMinuts(in.config.DsChannel, mes.ID, 3)
	} else if in.tip == "tg" {
		mes := tgSendChannel(in.config.TgChannel, fmt.Sprintf("Справка для  %s \n"+hhelpText, in.nameMention))
		tgDeleteMesageMinuts(in.config.TgChannel, mes, 3)
	}
}

func hhelpDS1(chatid string) (mesid string, error error) {
	mes, err := DSBot.ChannelMessageSend(chatid, fmt.Sprintf("Справка \n"+
		"ВНИМАНИЕ БОТ УДАЛЯЕТ СООБЩЕНИЯ \n ОТ ПОЛЬЗОВАТЕЛЕЙ ЧЕРЕЗ 3 МИНУТЫ \n\n"+
		hhelpText))
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
	if mtime == "12:00" {
		results, err := db.Query("SELECT dschannel,mesiddshelp FROM config")
		if err != nil {
			log.Println(err)
		}
		var dschannel, mesiddshelp string
		for results.Next() {
			err = results.Scan(&dschannel, &mesiddshelp)
			if mesiddshelp != "" {
				go dsDeleteMesage5s(dschannel, mesiddshelp)
				helpChannelUpdate(dschannel)
			} else {
				helpChannelUpdate(dschannel)
			}
		}
	}
}
func errorchanneldelete(dschatid string) {
	_, err := db.Exec("delete from channel where channel = ? ", dschatid)
	if err != nil {
		log.Println(err)
	}
}

func helpChannelUpdate(dschannel string) {
	newMesidHelp, err := hhelpDS1(dschannel)
	if err == nil {
		_, err := db.Exec(`update config set mesiddshelp = ? where dschannel = ? `, newMesidHelp, dschannel)
		if err != nil {
			log.Println(err)
		}
	} else {
		errorchanneldelete(dschannel)
	}
}
