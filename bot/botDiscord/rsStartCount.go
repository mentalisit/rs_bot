package botDiscord

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func countsS1(db *sql.DB, chatid, guildid, lvlkz string, numkz int) {
	mesid := readAll(db, lvlkz, chatid)
	name1 = <-rs
	_ = fmt.Sprintf("имя %s время %s", name1, <-rst)
	name2 = ""
	name3 = ""
	name4 = ""
	lvlk := roleToIdPing(lvlkz, guildid)
	EmbedDS(name1, name2, name3, name4, lvlk, numkz)
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: chatid,
	})
	go Delete3m(chatid, mesid)
	text := fmt.Sprintf("Запущен принудительный старт %s \n %s прошу в игру ", lvlk, name1)
	SendChannel(chatid, text)
	updateActive1(db, lvlkz, chatid)
	updateNumberkz(db, lvlkz, numkz, chatid)
}

func countsS2(db *sql.DB, chatid, guildid, lvlkz string, numkz int) {
	mesid := readAll(db, lvlkz, chatid)
	name1 = <-rs
	_ = fmt.Sprintf("%s %s", name1, <-rst)
	name2 = <-rs
	_ = fmt.Sprintf("%s %s", name2, <-rst)
	name3 = ""
	name4 = ""
	lvlk := roleToIdPing(lvlkz, guildid)
	EmbedDS(name1, name2, name3, name4, lvlk, numkz)
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: chatid,
	})
	go Delete3m(chatid, mesid)
	text := fmt.Sprintf("Запущен принудительный старт %s \n %s %s прошу в игру ", lvlk, name1, name2)
	SendChannel(chatid, text)
	updateActive1(db, lvlkz, chatid)
	updateNumberkz(db, lvlkz, numkz, chatid)
}

func countsS3(db *sql.DB, chatid, guildid, lvlkz string, numkz int) {
	mesid := readAll(db, lvlkz, chatid)
	name1 = <-rs
	_ = fmt.Sprintf("%s %s", name1, <-rst)
	name2 = <-rs
	_ = fmt.Sprintf("%s %s", name2, <-rst)
	name3 = <-rs
	_ = fmt.Sprintf("%s %s", name3, <-rst)
	name4 = ""
	lvlk := roleToIdPing(lvlkz, guildid)
	EmbedDS(name1, name2, name3, name4, lvlk, numkz)
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: chatid,
	})
	go Delete3m(chatid, mesid)
	text := fmt.Sprintf("Запущен принудительный старт %s \n %s %s %s прошу в игру ", lvlk, name1, name2, name3)
	SendChannel(chatid, text)
	updateActive1(db, lvlkz, chatid)
	updateNumberkz(db, lvlkz, numkz, chatid)
}
