package botDiscord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"regexp"
	"rs_bot/bot/botDiscord/databaseMysqlDs"
)

var rs = make(chan string, 4)
var rst = make(chan string, 4)

type inMessage struct {
	mtext       string
	nameMention string
	nameid      string
	mesid       string
	name        string
	guildid     string
	chatid      string
}

func logicRS(s *discordgo.Session, m *discordgo.MessageCreate) {
	db, er := databaseMysqlDs.DbConnection()
	if er != nil {
		log.Println(er)
	}
	inmes := inMessage{
		mtext:       m.Content,
		nameMention: m.Author.Mention(),
		nameid:      m.Message.Author.ID,
		mesid:       m.ID,
		name:        m.Message.Author.Username,
		guildid:     m.GuildID,
		chatid:      m.ChannelID,
	}
	//mtext :=  //текст сообщения
	//nameMention:=m.Author.Mention()
	//nameid := m.Message.Author.ID
	//mesid := m.ID             // ид сообщения 911747673093197844
	//name := m.Message.Author.Username//m.Author.Username // имя Mentalisit
	//guildid := m.GuildID      // id 700238199070523412
	//chatid := m.ChannelID     //chat id 909527364730490890

	if m.Author.ID == s.State.User.ID {
		return
	}
	var kzb, subs, qwery, rss string

	if len(m.Content) > 0 {
		//fmt.Println(name, nameid)
		str := inmes.mtext                                                 //mtext
		re := regexp.MustCompile(`^([4-9]|[1][0-1])([\+]|[-])(\d|\d{2})$`) //три переменные
		arr := (re.FindAllStringSubmatch(str, -1))
		if len(arr) > 0 {
			lvlkz = arr[0][1]
			kzb = arr[0][2]
			timekz = arr[0][3]
		}
		re2 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+]|[-])$`) // две переменные
		arr2 := (re2.FindAllStringSubmatch(str, -1))
		if len(arr2) > 0 {
			lvlkz = arr2[0][1]
			kzb = arr2[0][2]
			timekz = "30"
		}
		if kzb == "+" {
			DSBot.ChannelMessageDelete(inmes.chatid, inmes.mesid)
			RsPlus(db, lvlkz, timekz, &inmes) // mesid, name, nameid, guildid, chatid)
		} else if kzb == "-" {
			DSBot.ChannelMessageDelete(inmes.chatid, inmes.mesid)
			RsMinus(db, lvlkz, &inmes)
		}

		re3 := regexp.MustCompile(`^([\+]|[-])([4-9]|[1][0-1])$`) // две переменные для добавления или удаления подписок
		arr3 := (re3.FindAllStringSubmatch(str, -1))
		if len(arr3) > 0 {
			lvlkz = "кз" + arr3[0][2]
			subs = arr3[0][1]
		}
		if subs == "+" {
			Subscribe(m.GuildID, lvlkz, m.Message.Author.ID, m.ChannelID)
			go Delete1m(m.ChannelID, m.ID)

		} else if subs == "-" {
			Unsubscribe(m.GuildID, lvlkz, m.Message.Author.ID, m.ChannelID)
			go Delete1m(m.ChannelID, m.ID)
		}

		re4 := regexp.MustCompile(`^(["о"]|["О"])([4-9]|[1][0-1])$`) // две переменные для чтения  очереди
		arr4 := (re4.FindAllStringSubmatch(str, -1))
		if len(arr4) > 0 {
			qwery = arr4[0][1]
			lvlkz = arr4[0][2]
		}
		if len(qwery) > 0 {
			//DSBot.ChannelMessageDelete(chatid, mesid)
			//MsqlRsQ(lvlkz,m.ID)

		}

		re5 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+][\+])$`) //rs start
		arr5 := (re5.FindAllStringSubmatch(str, -1))
		if len(arr5) > 0 {
			lvlkz = arr5[0][1]
			rss = arr5[0][2]
		}
		if len(rss) > 0 {
			DSBot.ChannelMessageDelete(m.ChannelID, m.ID)
			RsStart(db, lvlkz, m.Message.Author.Username, m.ID, m.GuildID)
		}

		//if mtext == "Справка" {
		//	DSBot.ChannelMessageDelete(chatid, mesid)
		//hhelp(name, chatid)
		//		} else if mtext == "1" {
		//DSBot.ChannelMessageDelete(chatid, mesid)
		//			go Delete1m(chatid,mesid)
		//			mainTime()

		//		} else if mtext == "2" {
		//DSBot.ChannelMessageDelete(chatid, mesid)

		//		}
		//go Delete5s(chatid,mesid)
	}

	GuildState, err := DSBot.State.Guild(m.GuildID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("serverName: ", GuildState.Name, "Test DS String: "+m.Content)
}

/*
func creater(guildid string, levelkz string) {
	role,err:=DSBot.GuildRoleCreate(guildid)
	if err !=nil {fmt.Println(err)}
	DSBot.GuildRoleEdit(guildid,role.ID,"кз"+lvlkz,)
}
*/
