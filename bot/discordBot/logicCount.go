package discordBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func countName(lvlkz string, name string, chatid string) int {
	db := conDbDs()
	var countNames int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND lvlkz = ? AND chatid = ? AND active = 0", name, lvlkz, chatid)
	err := row.Scan(&countNames)
	if err != nil {
	}
	log.Println("imenDS", countNames)
	defer db.Close()
	return countNames
}
func countQueue(lvlkz string, chatid string) int {
	db := conDbDs()
	var count int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE lvlkz = ? AND chatid = ? AND active = 0", lvlkz, chatid)
	err := row.Scan(&count)
	if err != nil {
		log.Println("количество в очереди ошибка", count)
	}
	log.Println("количество в очереди ", count)
	db.Close()
	return count
}

//логика очереди
func counts0(lvlkz, timekz string, name string, nameid string, guildid string, chatid string) {
	name1 = nameid + "  🕒 " + timekz
	name2 = ""
	name3 = ""
	name4 = ""
	lvlk := roleToId(lvlkz, guildid)
	EmbedDS(name1, name2, name3, name4, lvlk)
	mesCompl, err := DSBot.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
		Content: nameid + " запустил очередь " + lvlk})
	if err != nil {
		fmt.Println(err)
	}
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mesCompl.ID,
		Channel: chatid,
	})
	insertSborkzAll(lvlkz, timekz, mesCompl.ID, name, nameid, guildid, chatid)
}
func counts1(lvlkz, timekz string, name string, nameid string, guildid string, chatid string) {
	mesid := readAll(lvlkz, chatid)
	fmt.Println(mesid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = nameid + "  🕒  " + timekz
	name3 = ""
	name4 = ""
	lvlk := roleToId(lvlkz, guildid)
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
	insertSborkzAll(lvlkz, timekz, mesid, name, nameid, guildid, chatid)
}
func counts2(lvlkz, timekz string, name string, nameid string, guildid string, chatid string) {
	mesid := readAll(lvlkz, chatid)
	fmt.Println(mesid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = <-rs + "  🕒  " + <-rst
	name3 = nameid + "  🕒  " + timekz
	name4 = ""
	lvlk := roleToId(lvlkz, guildid)
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
	insertSborkzAll(lvlkz, timekz, mesid, name, nameid, guildid, chatid)
}
func counts3(lvlkz, timekz string, name string, nameid string, guildid string, chatid string) {
	mesid := readAll(lvlkz, chatid)
	fmt.Println(mesid)
	name1 = <-rs
	names1 := name1 + "  🕒  " + <-rst
	name2 = <-rs
	names2 := name2 + "  🕒  " + <-rst
	name3 = <-rs
	names3 := name3 + "  🕒  " + <-rst
	name4 = nameid + "  🕒  " + timekz
	lvlk := roleToId(lvlkz, guildid)
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
	insertSborkzAll(lvlkz, timekz, mesid, name, nameid, guildid, chatid)
	//update очереди + похода +
}

//логика
