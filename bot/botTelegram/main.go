package botTelegram

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"rs_bot/bot/botTelegram/databaseMysql"
	"time"
)

var rs = make(chan string, 4)
var rst = make(chan int, 4)
var rsn = make(chan int, 4)
var Bot, BotErr = tgbotapi.NewBotAPI(os.Getenv("TokenT"))

func StartTgBot() {
	Bot, BotErr = tgbotapi.NewBotAPI(os.Getenv("TokenT"))
	if BotErr != nil {
		logrus.Panic(BotErr)
	}
	Bot.Debug = false
	log.Printf("Бот TELEGRAM загружен  %s", Bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	db, er := databaseMysql.DbConnection()
	if er != nil {
		logrus.Println(er)
	}

	updatesChannelTg(u, db)

}
func updatesChannelTg(u tgbotapi.UpdateConfig, db *sql.DB) {
	updates := Bot.GetUpdatesChan(u)
	for update := range updates {
		if update.CallbackQuery != nil {
			fmt.Println("калбекк")
			callback(update.CallbackQuery, db)

		} else if update.Message != nil {
			if update.Message.Chat.IsPrivate() {
				Send(update.Message.Chat.ID, "сорян это в разработке ")
			} else {
				logicRs(update.Message, db)
			}
		} else if update.MyChatMember != nil {
			myChatMember(update.MyChatMember)

		} else {
			fmt.Println(1, update)
			fmt.Println(2, update.ChatJoinRequest)
			fmt.Println(3, update.ChannelPost)
			fmt.Println(4, update.ChatMember)
			fmt.Println(5, update.ChosenInlineResult)
			fmt.Println(6, update.EditedChannelPost)
			fmt.Println(7, update.InlineQuery)
			fmt.Println(8, update.MyChatMember)
			fmt.Println(9, update.Poll)
			fmt.Println(10, update.PollAnswer)
			fmt.Println(11, update.EditedMessage)
			fmt.Println(12, update.PreCheckoutQuery)
			fmt.Println(13, update.ShippingQuery)

		}

	}
}
func Sends(lvlkz string, chatid int64, text string) int {
	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+", lvlkz+"+"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"-", lvlkz+"-"),
		),
	)
	msg := tgbotapi.NewMessage(chatid, text)
	msg.ReplyMarkup = keyboardQueue
	message, err := Bot.Send(msg)
	if err != nil {
		panic(err)
	}
	return message.MessageID

}
func SendP(chatid int64, text string) int {
	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("+", "+"),
			tgbotapi.NewInlineKeyboardButtonData("-", "-"),
		),
	)
	msg := tgbotapi.NewMessage(chatid, text)
	msg.ReplyMarkup = keyboardQueue
	message, err := Bot.Send(msg)
	if err != nil {
		panic(err)
	}
	return message.MessageID

}

// отправка сообщения в телегу
func Send(chatid int64, text string) int {
	tMessage, _ := Bot.Send(tgbotapi.NewMessage(chatid, text))
	return tMessage.MessageID
}
func SendDelMessage5s(chatid int64, idSendMessage int) {
	time.Sleep(10 * time.Second)
	Bot.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatid, idSendMessage)))
}
func SendDelMessage1m(chatid int64, idSendMessage int) {
	time.Sleep(1 * time.Minute)
	Bot.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatid, idSendMessage)))
}
func SendDelMessage5m(chatid int64, idSendMessage int) {
	time.Sleep(5 * time.Minute)
	Bot.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatid, idSendMessage)))
}
func SendDelMessage30m(chatid int64, idSendMessage int) {
	time.Sleep(30 * time.Minute)
	Bot.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatid, idSendMessage)))
}
func SendDelMessage3m(chatid int64, idSendMessage int) {
	time.Sleep(3 * time.Minute)
	Bot.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatid, idSendMessage)))

}
func EditMessageText(chatid int64, editMesId int, textEdit string, lvlkz string) {
	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+", lvlkz+"+"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"-", lvlkz+"-"),
		),
	)
	tgbotapi.NewEditMessageText(chatid, editMesId, textEdit)
	Bot.Send(&tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:          chatid,
			ChannelUsername: "",
			MessageID:       editMesId,
			InlineMessageID: "",
			ReplyMarkup:     &keyboardQueue,
		},
		Text:                  textEdit,
		ParseMode:             "",
		DisableWebPagePreview: false,
	})
}
