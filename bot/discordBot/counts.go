package discordBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

//const insertSmtg string = "INSERT INTO sborkzds (name, mesid, chatid, time, date, lvlkz, numberkz, eventpoints, active, timedown, activedel) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
//const prstr string= "%s - принудительный старт"

/////////////////////////////////////////

func counts1r(chatid, guildid, lvlkz string) {
	mesid := readAll(lvlkz, chatid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = ""
	name3 = ""
	name4 = ""
	lvlk := roleToId(lvlkz, guildid)
	EmbedDS(name1, name2, name3, name4, lvlk)
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: chatid,
	})
}

func counts2r(chatid, guildid, lvlkz string) {
	mesid := readAll(lvlkz, chatid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = <-rs + "  🕒  " + <-rst
	name3 = ""
	name4 = ""
	lvlk := roleToId(lvlkz, guildid)
	EmbedDS(name1, name2, name3, name4, lvlk)
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: chatid,
	})
}

func hhelp(name, chatid string) {
	DSBot.ChannelMessageSend(chatid, fmt.Sprintf("Справка для  %s \n"+
		"Встать в очередь: [4-11]+  или\n"+
		" [4-11]+[указать время ожидания в минутах]\n"+
		" **9+** - встать в очередь на КЗ 9ур.\n"+
		" **9+60** - встать на КЗ 9ур, время ожидания не более 60 минут.\n"+
		"Покинуть очередь: [4-11] -\n"+
		" **9-**- выйти из очереди на КЗ 9ур.\n"+
		"Посмотреть список активных очередей: **о**\n"+
		" **о9** - вывод очередь для вашей Кз\n"+
		"Включить уведомления: + [5-11]\n"+ ////нужно слеать пинг
		"**+9**  - подписаться на уведомления о сборе на КЗ 9ур.\n"+
		"Отключить уведомления: - [5-11]\n", name))
}

/*		"**-9**  - отключить уведомления о сборе на КЗ 9ур." +
		"Добавить/удалить пользователя в ЧС: rs bl @user" +
		"	<b>rs bl @Nick</b> - добавить/удалить пользователя Nick в ЧС." +
		"		Посмотреть свой ЧС: <b>mybl</b>" +
*/
