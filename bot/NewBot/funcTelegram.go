package NewBot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"regexp"
	"time"
)

func updatesChannelTg(u tgbotapi.UpdateConfig) {
	updates := TgBot.GetUpdatesChan(u)
	for update := range updates {
		if update.CallbackQuery != nil {
			callback(update.CallbackQuery)

		} else if update.Message != nil {
			if update.Message.Chat.IsPrivate() {
				tgSendChannel(update.Message.Chat.ID, "сорян это в разработке ")
			} else if update.Message.IsCommand() {
				updatesComand(update.Message)
			} else {
				logicMixTelegram(update.Message)
			}

		} else if update.MyChatMember != nil {
			myChatMember(update.MyChatMember)

		} else if update.EditedMessage != nil {
			logrus.Println("Измененный текст в телеге ", update.EditedMessage.Text)
		} else {
			fmt.Println(1, update)
		}
	}
}

func callback(cb *tgbotapi.CallbackQuery) {
	var lvlkz, kzb, timekz, rss string
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := TgBot.Request(callback); err != nil {
		logrus.Panic(err)
	}
	ok, config := checkChannelConfigTG(cb.Message.Chat.ID)
	if ok {
		re2 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+]|[-])$`) // две переменные
		arr2 := (re2.FindAllStringSubmatch(cb.Data, -1))

		if len(arr2) > 0 {
			lvlkz = arr2[0][1]
			kzb = arr2[0][2]
			timekz = "30"
		}
		re5 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+][\+])$`) //rs start
		arr5 := (re5.FindAllStringSubmatch(cb.Data, -1))
		if len(arr5) > 0 {
			lvlkz = arr5[0][1]
			rss = arr5[0][2]
		}

		re := regexp.MustCompile(`^([4-9]|[1][0-1])([\+]|[-])(\d{2})$`) //три переменные
		arr := (re.FindAllStringSubmatch(cb.Data, -1))
		if len(arr) > 0 {
			lvlkz = arr[0][1]
		}
		in := inMessage{
			mtext:       cb.Data,
			tip:         "tg",
			name:        cb.From.UserName,
			nameMention: "@" + cb.From.UserName,
			Ds: Ds{
				mesid:   "",
				nameid:  "",
				guildid: "",
			},
			Tg: Tg{
				mesid:  cb.Message.MessageID,
				nameid: cb.From.ID,
			},
			config: config,
			option: Option{
				callback: true,
				edit:     true,
				update:   false,
			},
		}
		in.timekz = timekz
		in.lvlkz = lvlkz
		if kzb == "+" {
			in.RsPlus()
		} else if kzb == "-" {
			in.RsMinus()
		} else if cb.Data == "+" {
			in.Plus()
			//if Plus(in) {	tgDelMessage(in.config.TgChannel, in.Tg.mesid)}
		} else if cb.Data == "-" {
			in.Minus()
			//if Minus(in) {tgDelMessage(in.config.TgChannel, in.Tg.mesid)}
		} else if len(rss) > 0 {
			in.RsStart()
		} else if len(arr) > 0 {
			in.Pl30()
		}
	}
}

func myChatMember(member *tgbotapi.ChatMemberUpdated) {
	if member.NewChatMember.Status == "member" {
		tgSendChannel(member.Chat.ID, fmt.Sprintf("@%s мне нужны права админа для коректной работы", member.From.UserName))
	} else if member.NewChatMember.Status == "administrator" {
		tgSendChannel(member.Chat.ID, fmt.Sprintf("@%s спасибо ... я готов к работе \nподтверди активацию .add", member.From.UserName))
	}
}

func tgSendEmded(lvlkz string, chatid int64, text string) int {
	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+", lvlkz+"+"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"-", lvlkz+"-"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"++", lvlkz+"++"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+30", lvlkz+"+30"),
		),
	)
	msg := tgbotapi.NewMessage(chatid, text)
	msg.ReplyMarkup = keyboardQueue
	message, err := TgBot.Send(msg)
	if err != nil {
		logrus.Println(err)
	}
	return message.MessageID

}
func tgSendEmbedTime(chatid int64, text string) int {
	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("+", "+"),
			tgbotapi.NewInlineKeyboardButtonData("-", "-"),
		),
	)
	msg := tgbotapi.NewMessage(chatid, text)
	msg.ReplyMarkup = keyboardQueue
	message, err := TgBot.Send(msg)
	if err != nil {
		logrus.Panic(err)
	}
	return message.MessageID

}

// отправка сообщения в телегу
func tgSendChannel(chatid int64, text string) int {
	tMessage, _ := TgBot.Send(tgbotapi.NewMessage(chatid, text))
	return tMessage.MessageID
}
func tgSendChannelDel1m(chatid int64, text string) {
	tMessage, _ := TgBot.Send(tgbotapi.NewMessage(chatid, text))
	tgDeleteMesageMinuts(chatid, tMessage.MessageID, 1)
}

func tgSendChannelDel5s(chatid int64, text string) {
	tMessage, _ := TgBot.Send(tgbotapi.NewMessage(chatid, text))
	time.Sleep(5 * time.Second)
	TgBot.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatid, tMessage.MessageID)))
}
func tgDelMessage(chatid int64, idSendMessage int) {
	TgBot.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatid, idSendMessage)))
}
func tgDelMessage10s(chatid int64, idSendMessage int) {
	time.Sleep(10 * time.Second)
	TgBot.Request(tgbotapi.DeleteMessageConfig(tgbotapi.NewDeleteMessage(chatid, idSendMessage)))
}

func tgEditMessageText(chatid int64, editMesId int, textEdit string, lvlkz string) {
	var keyboardQueue = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+", lvlkz+"+"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"-", lvlkz+"-"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"++", lvlkz+"++"),
			tgbotapi.NewInlineKeyboardButtonData(lvlkz+"+30", lvlkz+"+30"),
		),
	)
	tgbotapi.NewEditMessageText(chatid, editMesId, textEdit)
	TgBot.Send(&tgbotapi.EditMessageTextConfig{
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

func tgEditText(chatid int64, editMesId int, textEdit string) {
	TgBot.Send(tgbotapi.NewEditMessageText(chatid, editMesId, textEdit))
}

func checkAdminTg(in inMessage) bool {
	admin := false
	admins, err := TgBot.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{struct {
		ChatID             int64
		SuperGroupUsername string
	}{ChatID: in.config.TgChannel, SuperGroupUsername: ""}})
	if err != nil {
		logrus.Println(err)
	}
	for _, ad := range admins {
		if in.name == ad.User.UserName && (ad.IsAdministrator() || ad.IsCreator()) {
			admin = true
			break
		}
	}
	return admin
}

func removeDuplicateElementInt(mesididid []int) []int {
	result := make([]int, 0, len(mesididid))
	temp := map[int]struct{}{}
	for _, item := range mesididid {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func updatesComand(c *tgbotapi.Message) {
	ok, config := checkChannelConfigTG(c.Chat.ID)
	if ok {
		in := inMessage{
			tip:         "tg",
			nameMention: "@" + c.From.UserName,
			Tg:          Tg{mesid: c.MessageID},
			config:      config,
		}
		switch c.Command() {
		case "help":
			help(in)
		case "helpqueue":
			helpQueue(in)
		case "helpnotification":
			helpNotification(in)
		case "helpevent":
			helpEvent(in)
		case "helptop":
			helpTop(in)
		case "helpicon":
			helpIcon(in)
		}
	}
}

func tgChatName(chatid int64) string {
	r, err := TgBot.GetChat(tgbotapi.ChatInfoConfig{struct {
		ChatID             int64
		SuperGroupUsername string
	}{ChatID: chatid}})
	if err != nil {
		logrus.Println(err)
	}
	return r.Title
}
