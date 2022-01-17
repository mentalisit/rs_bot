package botDiscord

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func countsS1(db *sql.DB, chatid, guildid, lvlkz string) {
	mesid := readAll(db, lvlkz, chatid)
	name1 = <-rs
	fmt.Printf("%s %s", name1, <-rst)
	name2 = ""
	name3 = ""
	name4 = ""
	lvlk := roleToIdPing(lvlkz, guildid)
	EmbedDS(name1, name2, name3, name4, lvlk)
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: chatid,
	})
}

func countsS2(db *sql.DB, chatid, guildid, lvlkz string) {
	mesid := readAll(db, lvlkz, chatid)
	name1 = <-rs
	fmt.Printf("%s %s", name1, <-rst)
	name2 = <-rs
	fmt.Printf("%s %s", name2, <-rst)
	name3 = ""
	name4 = ""
	lvlk := roleToIdPing(lvlkz, guildid)
	EmbedDS(name1, name2, name3, name4, lvlk)
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: chatid,
	})
}

func countsS3(db *sql.DB, chatid, guildid, lvlkz string) {
	mesid := readAll(db, lvlkz, chatid)
	name1 = <-rs
	fmt.Printf("%s %s", name1, <-rst)
	name2 = <-rs
	fmt.Printf("%s %s", name2, <-rst)
	name3 = <-rs
	fmt.Printf("%s %s", name3, <-rst)
	name4 = ""
	lvlk := roleToIdPing(lvlkz, guildid)
	EmbedDS(name1, name2, name3, name4, lvlk)
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: chatid,
	})
}
