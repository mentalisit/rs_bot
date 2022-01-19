package botDiscord

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
)

func counts1r(db *sql.DB, chatid, guildid, lvlkz string, numkz int) {
	mesid := readAll(db, lvlkz, chatid)
	name1 = <-rs + "  🕒  " + <-rst
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
}

func counts2r(db *sql.DB, chatid, guildid, lvlkz string, numkz int) {
	mesid := readAll(db, lvlkz, chatid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = <-rs + "  🕒  " + <-rst
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
}
