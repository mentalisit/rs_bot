package NewBot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

type timer struct {
	id       int
	dsmesid  string
	dschatid string
	tgmesid  int
	tgchatid int64
	timed    time.Duration
}

func timerDeleteMessage() {
	_, err := db.Exec(`update timer set timed = timed - 60`)
	if err != nil {
		fmt.Println(err)
	}
	results, err := db.Query("SELECT * FROM timer WHERE timed < 60")
	if err != nil {
		fmt.Println(err)
	}
	for results.Next() {
		var t timer
		err = results.Scan(&t.id, &t.dsmesid, &t.dschatid, &t.tgmesid, &t.tgchatid, &t.timed)

		if t.timed <= 60 {
			go t.timerdeletemesSec()
		}
	}
}

func (t timer) timerdeletemesSec() {
	if t.timed > 0 {
		time.Sleep(t.timed * time.Second)
	}
	if t.dsmesid != "" {
		DSBot.ChannelMessageDelete(t.dschatid, t.dsmesid)
	}
	if t.tgmesid != 0 {
		TgBot.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(t.tgchatid, t.tgmesid)))
	}
	go t.timerDeleteDB()
}

func dsDeleteMesageMinuts(chatid, mesid string, minuts int) {
	minuts = minuts * 60 //seconds
	timerInsert(mesid, chatid, 0, 0, minuts)
}
func tgDeleteMesageMinuts(chatid int64, idSendMessage int, minuts int) {
	minuts = minuts * 60 //seconds
	timerInsert("", "", idSendMessage, chatid, minuts)
}

func timerInsert(dsmesid, dschatid string, tgmesid int, tgchatid int64, timed int) {
	insertTimer := `INSERT INTO timer(dsmesid,dschatid,tgmesid,tgchatid,timed) VALUES (?,?,?,?,?)`
	_, err := db.Exec(insertTimer, dsmesid, dschatid, tgmesid, tgchatid, timed)
	if err != nil {
		log.Println(err)
	}
}

func (t timer) timerDeleteDB() {
	_, err := db.Exec("delete from timer where  id = ? ", t.id)
	if err != nil {
		log.Println(err)
	}
}
