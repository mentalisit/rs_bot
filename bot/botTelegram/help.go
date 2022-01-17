package botTelegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func updatesComand(c *tgbotapi.Message) {
	if c.Command() == "help" {
		hhelp(c.From.UserName, c.Chat.ID)
	}
}

func hhelp(name string, chatid int64) {
	mes := Send(chatid, fmt.Sprintf("Справка для  @%s \n"+
		"Встать в очередь: [4-11]+  или\n"+
		" [4-11]+[указать время ожидания в минутах]\n"+
		" 9+  встать в очередь на КЗ 9ур.\n"+
		" 9+60  встать на КЗ 9ур, время ожидания не более 60 минут.\n"+
		"Покинуть очередь: [4-11] -\n"+
		" 9- выйти из очереди на КЗ 9ур.\n"+
		"Посмотреть список активных очередей: **о**\n"+
		" о9  вывод очередь для вашей Кз\n"+
		"Включить уведомления: + [4-11] или ++\n"+
		" +9   подписаться на уведомления о сборе на КЗ 9ур.\n"+
		"Отключить уведомления: - [4-11] или --\n"+
		" -9   отключить уведомления о сборе на КЗ 9ур.", name))
	go SendDelMessage5m(chatid, mes)
}

/*
	"Добавить/удалить пользователя в ЧС: rs bl @user" +
	"	<b>rs bl @Nick</b> - добавить/удалить пользователя Nick в ЧС." +
	"		Посмотреть свой ЧС: <b>mybl</b>" +
*/
