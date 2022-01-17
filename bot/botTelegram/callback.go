package botTelegram

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"regexp"
	"time"
)

func callback(cb *tgbotapi.CallbackQuery, db *sql.DB) {
	var lvlkz, kzb, timekz string
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := Bot.Request(callback); err != nil {
		panic(err)
	}
	text := cb.Data
	name := cb.From.UserName
	chatid := cb.Message.Chat.ID
	tm := time.Unix(int64(cb.Message.Date), 0).UTC()
	mdata := (tm.Format("2006-01-02"))
	mtime := (tm.Format("15:04"))
	fmt.Println(name, text, chatid, mtime, mdata)

	re2 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+]|[-])$`) // две переменные
	arr2 := (re2.FindAllStringSubmatch(text, -1))

	if len(arr2) > 0 {
		lvlkz = arr2[0][1]
		kzb = arr2[0][2]
		timekz = "30"
	}
	if kzb == "+" {
		RsPlus(db, name, lvlkz, timekz, chatid)
	} else if kzb == "-" {
		RsMinus(db, name, lvlkz, chatid)
	} else if text == "+" {
		Plus(db, name, chatid)
	} else if text == "-" {
		Minus(db, name, chatid)
	}
}
