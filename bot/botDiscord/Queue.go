package botDiscord

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Queue(db *sql.DB, lvlkz string, chatid, guildid string, edit bool) {
	count := countQueue(db, lvlkz, chatid)
	numberLvl := readNumberkz(db, lvlkz, chatid) + 1
	// совподения количество  условие
	if count == 0 {
		mes0 := SendChannel(chatid, "Очередь КЗ "+lvlkz+" пуста ")
		go Delete5s(chatid, mes0)
	} else if count == 1 {
		count1Queue(db, lvlkz, chatid, guildid, numberLvl, edit)
	} else if count == 2 {
		count2Queue(db, lvlkz, chatid, guildid, numberLvl, edit)
	} else if count == 3 {
		count3Queue(db, lvlkz, chatid, guildid, numberLvl, edit)
	}
}

func count1Queue(db *sql.DB, lvlkz string, chatid, guildid string, numberLvlkz int, edit bool) {
	mesid := readAll(db, lvlkz, chatid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = ""
	name3 = ""
	name4 = ""
	lvlk := roleToIdPing(lvlkz, guildid)
	EmbedDS(name1, name2, name3, name4, lvlk, numberLvlkz)
	if edit {
		DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
			Content: &mesContentNil,
			Embed:   Embeds,
			ID:      mesid,
			Channel: chatid,
		})
	} else if !edit {
		DSBot.ChannelMessageDelete(chatid, mesid)
		mesCompl, err := DSBot.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
			Content: mesContentNil,
			Embed:   Embeds})
		if err != nil {
			fmt.Println(err)
		}
		addEnojiRsQueue(chatid, mesCompl.ID)
		updateMesid(db, lvlkz, chatid, mesid, mesCompl.ID)
	}
}
func count2Queue(db *sql.DB, lvlkz string, chatid, guildid string, numberLvlkz int, edit bool) {
	mesid := readAll(db, lvlkz, chatid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = <-rs + "  🕒  " + <-rst
	name3 = ""
	name4 = ""
	lvlk := roleToIdPing(lvlkz, guildid)
	EmbedDS(name1, name2, name3, name4, lvlk, numberLvlkz)
	if edit {
		DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
			Content: &mesContentNil,
			Embed:   Embeds,
			ID:      mesid,
			Channel: chatid,
		})
	} else if !edit {
		DSBot.ChannelMessageDelete(chatid, mesid)
		mesCompl, err := DSBot.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
			Content: mesContentNil,
			Embed:   Embeds})
		if err != nil {
			fmt.Println(err)
		}
		addEnojiRsQueue(chatid, mesCompl.ID)
		updateMesid(db, lvlkz, chatid, mesid, mesCompl.ID)
	}

}
func count3Queue(db *sql.DB, lvlkz string, chatid, guildid string, numberLvlkz int, edit bool) {
	mesid := readAll(db, lvlkz, chatid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = <-rs + "  🕒  " + <-rst
	name3 = <-rs + "  🕒  " + <-rst
	name4 = ""
	lvlk := roleToIdPing(lvlkz, guildid)
	EmbedDS(name1, name2, name3, name4, lvlk, numberLvlkz)
	if edit {
		DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
			Content: &mesContentNil,
			Embed:   Embeds,
			ID:      mesid,
			Channel: chatid,
		})
	} else if !edit {
		DSBot.ChannelMessageDelete(chatid, mesid)
		mesCompl, err := DSBot.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
			Content: mesContentNil,
			Embed:   Embeds})
		if err != nil {
			fmt.Println(err)
		}
		addEnojiRsQueue(chatid, mesCompl.ID)
		updateMesid(db, lvlkz, chatid, mesid, mesCompl.ID)
	}

}
