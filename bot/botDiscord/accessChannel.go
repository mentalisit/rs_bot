package botDiscord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"rs_bot/bot/botDiscord/databaseMysqlDs"
	"strings"
)

var ch = []string{}

func accesChat(m *discordgo.MessageCreate) {
	res := strings.HasPrefix(m.Content, ".")
	if res == true && m.Content == ".add" {
		go Delete5s(m.ChannelID, m.ID)
		accessAddChannel(m.ChannelID)
	} else if res == true && m.Content == ".del" {
		go Delete5s(m.ChannelID, m.ID)
		accessDelChannel(m.ChannelID)
	}
}

//init
func readChannel() { // чтение с бд и добавление в масив
	db, er := databaseMysqlDs.DbConnection()
	if er != nil {
		log.Println(er)
	}
	results, err := db.Query("SELECT channel FROM channel")
	if err != nil {
		log.Println(err)
	}
	var channel string
	for results.Next() {
		err = results.Scan(&channel)
		ch = append(ch, channel)
	}
}

func accessDelChannel(chatid string) { //удаление с бд и масива для блокировки
	if !checkChannel(chatid) {
		mes := SendChannel(chatid, "ваш канал и так не подключен к логике бота ")
		go Delete1m(chatid, mes)
	} else {
		db, er := databaseMysqlDs.DbConnection()
		if er != nil {
			log.Println(er)
		}
		_, err := db.Exec("delete from channel where channel = ? ", chatid)
		if err != nil {
			log.Println(err)
		}
		db.Close()
		for i, chat := range ch {
			if chatid == chat {
				ch = append(ch[:i], ch[i+1:]...)
			}
		}
		mes := SendChannel(chatid, "вы отключили мои возможности")
		go Delete1m(chatid, mes)
	}
}

func accessAddChannel(chatid string) { // внесение в дб и добавление в масив
	if checkChannel(chatid) {
		mes := SendChannel(chatid, "Я уже могу работать на вашем канале\n"+
			"повторная активация не требуется.\nнапиши Справка")
		go Delete1m(chatid, mes)
	} else {
		db, er := databaseMysqlDs.DbConnection()
		if er != nil {
			log.Println(er)
		}
		insertChannel := `INSERT INTO channel(channel) VALUES (?)`
		statement, err := db.Prepare(insertChannel)
		if err != nil {
			fmt.Println(err)
		}
		_, err = statement.Exec(chatid)
		if err != nil {
			fmt.Println(err.Error())
		}
		db.Close()
		ch = append(ch, chatid)
		mes := SendChannel(chatid, "Спасибо за активацию.\nпиши Справка")
		go Delete1m(chatid, mes)
	}
}

func checkChannel(chatid string) bool {
	channelGood := false
	for _, chat := range ch {
		if chatid == chat {
			channelGood = true
			break
		}
	}
	return channelGood
}
