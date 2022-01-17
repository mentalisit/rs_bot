package botDiscord

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

//логика очереди
func counts0(db *sql.DB, lvlkz, timekz string, name string, nameid string, guildid string, chatid string) {
	name1 = nameid + "  🕒 " + timekz
	name2 = ""
	name3 = ""
	name4 = ""
	lvlk := roleToIdPing(lvlkz, guildid)
	EmbedDS(name1, name2, name3, name4, lvlk)
	mesCompl, err := DSBot.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
		Content: nameid + " запустил очередь " + lvlk})
	if err != nil {fmt.Println(err)}
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesCompl.ID,
		Channel: chatid,
	})
	insertSborkzAll(db, lvlkz, timekz, mesCompl.ID, name, nameid, guildid, chatid)
}
func counts1(db *sql.DB, lvlkz, timekz string, name string, nameid string, guildid string, chatid string) {
	mesid := readAll(db, lvlkz, chatid)
	fmt.Println(mesid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = nameid + "  🕒  " + timekz
	name3 = ""
	name4 = ""
	lvlk := roleToIdPing(lvlkz, guildid)
	mes := SendChannel(chatid, lvlk+" 2/4 "+nameid+" присоеденился к очереди")
	Delete5s(chatid, mes)
	EmbedDS(name1, name2, name3, name4, lvlk)
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: chatid,
	})
	insertSborkzAll(db, lvlkz, timekz, mesid, name, nameid, guildid, chatid)
}
func counts2(db *sql.DB, lvlkz, timekz string, name string, nameid string, guildid string, chatid string) {
	mesid := readAll(db, lvlkz, chatid)
	fmt.Println(mesid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = <-rs + "  🕒  " + <-rst
	name3 = nameid + "  🕒  " + timekz
	name4 = ""
	lvlk := roleToIdPing(lvlkz, guildid)
	mes := SendChannel(chatid, lvlk+" 3/4 "+nameid+" присоеденился к очереди")
	Delete5s(chatid, mes)
	EmbedDS(name1, name2, name3, name4, lvlk)
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: chatid,
	})
	insertSborkzAll(db, lvlkz, timekz, mesid, name, nameid, guildid, chatid)
}
func counts3(db *sql.DB, lvlkz, timekz string, name string, nameid string, guildid string, chatid string) {
	mesid := readAll(db, lvlkz, chatid)
	fmt.Println(mesid)
	name1 = <-rs; names1 := name1 + "  🕒  " + <-rst
	name2 = <-rs; names2 := name2 + "  🕒  " + <-rst
	name3 = <-rs; names3 := name3 + "  🕒  " + <-rst
	name4 = nameid + "  🕒  " + timekz
	lvlk := roleToIdPing(lvlkz, guildid)
	mes := SendChannel(chatid, " 4/4 "+nameid+" присоеденился к очереди")
	Delete5s(chatid, mes)
	mes = SendChannel(chatid, "очередь заполнена нужно тут кучку всего написать ")
	EmbedDS(names1, names2, names3, name4, lvlk)
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesid,
		Channel: chatid,
	})
	insertSborkzAll(db, lvlkz, timekz, mesid, name, nameid, guildid, chatid)
	updateActive1(db, lvlkz, chatid)
	//update похода +
}
