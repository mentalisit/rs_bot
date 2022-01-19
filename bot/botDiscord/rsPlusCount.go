package botDiscord

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var mesContentNil string

//логика очереди
func counts0(db *sql.DB, lvlkz, timekz string, m *inMessage, numkz int) {
	name1 = m.nameMention + "  🕒 " + timekz
	name2 = ""
	name3 = ""
	name4 = ""
	lvlk := roleToIdPing(lvlkz, m.guildid)
	EmbedDS(name1, name2, name3, name4, lvlk, numkz)
	mesCompl, err := DSBot.ChannelMessageSendComplex(m.chatid, &discordgo.MessageSend{
		Content: m.nameMention + " запустил очередь " + lvlk})
	if err != nil {
		fmt.Println(err)
	}
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesCompl.ID,
		Channel: m.chatid,
	})
	insertSborkzAll(db, lvlkz, timekz, mesCompl.ID, m.name, m.nameid, m.guildid, m.chatid, m.nameMention)
}
func counts1(db *sql.DB, lvlkz, timekz string, m *inMessage, numkz int) {
	mesid := readAll(db, lvlkz, m.chatid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = m.nameMention + "  🕒  " + timekz
	name3 = ""
	name4 = ""
	lvlk := roleToIdPing(lvlkz, m.guildid)
	mes := SendChannel(m.chatid, lvlk+" 2/4 "+m.nameMention+" присоеденился к очереди")
	go Delete5s(m.chatid, mes)
	EmbedDS(name1, name2, name3, name4, lvlk, numkz)
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: m.chatid,
	})
	insertSborkzAll(db, lvlkz, timekz, mesid, m.name, m.nameid, m.guildid, m.chatid, m.nameMention)
}
func counts2(db *sql.DB, lvlkz, timekz string, m *inMessage, numkz int) {
	mesid := readAll(db, lvlkz, m.chatid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = <-rs + "  🕒  " + <-rst
	name3 = m.nameMention + "  🕒  " + timekz
	name4 = ""
	lvlk := roleToIdPing(lvlkz, m.guildid)
	mes := SendChannel(m.chatid, lvlk+" 3/4 "+m.nameMention+" присоеденился к очереди")
	go Delete5s(m.chatid, mes)
	EmbedDS(name1, name2, name3, name4, lvlk, numkz)
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: m.chatid,
	})
	insertSborkzAll(db, lvlkz, timekz, mesid, m.name, m.nameid, m.guildid, m.chatid, m.nameMention)
}
func counts3(db *sql.DB, lvlkz, timekz string, m *inMessage, numkz int) {
	mesid := readAll(db, lvlkz, m.chatid)
	name1 = <-rs
	names1 := name1 + "  🕒  " + <-rst
	name2 = <-rs
	names2 := name2 + "  🕒  " + <-rst
	name3 = <-rs
	names3 := name3 + "  🕒  " + <-rst
	name4 = m.nameMention + "  🕒  " + timekz
	lvlk := roleToIdPing(lvlkz, m.guildid)
	mes := SendChannel(m.chatid, " 4/4 "+m.nameMention+" присоеденился к очереди")
	go Delete5s(m.chatid, mes)
	mes = SendChannel(m.chatid, "очередь заполнена нужно тут кучку всего написать ")
	EmbedDS(names1, names2, names3, name4, lvlk, numkz)
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: m.chatid,
	})
	insertSborkzAll(db, lvlkz, timekz, mesid, m.name, m.nameid, m.guildid, m.chatid, m.nameMention)
	updateActive1(db, lvlkz, m.chatid)
	//update похода +
}
