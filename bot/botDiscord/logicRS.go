package botDiscord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"regexp"
	"rs_bot/bot/botDiscord/databaseMysqlDs"
	"strconv"
)

var rs = make(chan string, 4)
var rst = make(chan string, 4)

type inMessage struct {
	mtext       string
	name        string
	nameMention string
	nameid      string
	mesid       string
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
	mtext := m.Content //текст сообщения
	nameMention := m.Author.Mention()
	//nameid := m.Message.Author.ID
	mesid := m.ID                     // ид сообщения 911747673093197844
	name := m.Message.Author.Username //m.Author.Username // имя Mentalisit
	//guildid := m.GuildID      // id 700238199070523412
	chatid := m.ChannelID //chat id 909527364730490890

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
			go Delete5s(chatid, mesid)
			Queue(db, lvlkz, m.ChannelID, m.GuildID, false)

		}

		re5 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+][\+])$`) //rs start
		arr5 := (re5.FindAllStringSubmatch(str, -1))
		if len(arr5) > 0 {
			lvlkz = arr5[0][1]
			rss = arr5[0][2]
		}
		if len(rss) > 0 {
			DSBot.ChannelMessageDelete(m.ChannelID, m.ID)
			RsStart(db, lvlkz, inmes.name, inmes.chatid, inmes.guildid)
		}
		re7 := regexp.MustCompile(`^(["К"])\s([0-9]+)\s([0-9]+)$`) // ивент
		arr7 := (re7.FindAllStringSubmatch(str, -1))
		if len(arr7) > 0 {
			points, err := strconv.Atoi(arr7[0][3])
			if err != nil {
			}
			numkz, err := strconv.Atoi(arr7[0][2])
			if err != nil {
			}
			EventPoints(db, inmes.chatid, inmes.name, inmes.nameid, numkz, points)
			Delete5s(inmes.chatid, inmes.mesid)
		}
		if mtext == "Ивент старт" {
			EventStart(db, name, inmes.nameid, chatid)
			Delete5s(chatid, mesid)
		} else if mtext == "Ивент стоп" {
			EventStop(db, name, inmes.nameid, chatid)
			Delete5s(chatid, mesid)
		} else if mtext == "Справка" {
			DSBot.ChannelMessageDelete(chatid, mesid)
			Delete3m(chatid, hhelp(nameMention, chatid))
		} else if mtext == "Справка1" {
			DSBot.ChannelMessageDelete(chatid, mesid)
			hhelp1(chatid)
		} else if mtext == "+" { //продление очереди на 30мин
			go Delete5s(chatid, mesid)
			Plus(db, name, chatid)
		} else if mtext == "-" {
			go Delete5s(chatid, mesid)
			Minus(db, name, chatid)
		} else if mtext == "Топ" {
			go Delete5s(chatid, mesid)
			topAll(db, chatid)
		} //else if mtext == "1" {
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

func readReactionQueue(r *discordgo.MessageReactionAdd, message *discordgo.Message) {
	db, er := databaseMysqlDs.DbConnection()
	if er != nil {
		log.Println(er)
	}

	user, err := DSBot.User(r.UserID)
	if err != nil {
		fmt.Println(err)
	}
	if user.ID != message.Author.ID {
		inm := inMessage{
			mtext:       "",
			name:        user.Username,
			nameMention: user.Mention(),
			nameid:      user.ID,
			mesid:       r.MessageID,
			guildid:     r.GuildID,
			chatid:      r.ChannelID,
		}
		if r.Emoji.Name == emPlus {
			if Plus(db, inm.name, inm.chatid) {
				DSBot.ChannelMessageDelete(inm.chatid, inm.mesid)
			}
		} else if r.Emoji.Name == emMinus {
			if Minus(db, inm.name, inm.chatid) {
				DSBot.ChannelMessageDelete(inm.chatid, inm.mesid)
			}
		} else if r.Emoji.Name == emOK || r.Emoji.Name == emCancel || r.Emoji.Name == emRsStart || r.Emoji.Name == emPl30 {
			lvlkz, err = readMesID(db, r.MessageID)
			if r.Emoji.Name == emOK {
				RsPlus(db, lvlkz, "30", &inm)
			} else if r.Emoji.Name == emCancel {
				RsMinus(db, lvlkz, &inm)
			} else if r.Emoji.Name == emRsStart {
				RsStart(db, lvlkz, inm.name, inm.chatid, inm.guildid)
			} else if r.Emoji.Name == emPl30 {
				Pl30(db, lvlkz, &inm)
			}
		}
	}
}

func hhelp(mentionName, chatid string) string {
	mes, _ := DSBot.ChannelMessageSend(chatid, fmt.Sprintf("Справка для  %s \n"+
		"ВНИМАНИЕ БОТ УДАЛЯЕТ СООБЩЕНИЯ \n ОТ ПОЛЬЗОВАТЕЛЕЙ ЧЕРЕЗ 3 МИНУТЫ \n\n"+
		"Встать в очередь: [4-11]+  или\n"+
		" [4-11]+[указать время ожидания в минутах]\n"+
		" 9+  встать в очередь на КЗ 9ур.\n"+
		" 9+60  встать на КЗ 9ур, время ожидания не более 60 минут.\n"+
		"Покинуть очередь: [4-11] -\n"+
		" 9- выйти из очереди КЗ 9ур.\n"+
		"Посмотреть список активных очередей: о[4-11]\n"+
		" о9  вывод очередь для вашей Кз\n"+
		"Получить роль кз: + [5-11]\n"+
		" +9   получить роль КЗ 9ур.\n"+
		" -9   снять роль ", mentionName))
	return mes.ID
}
func hhelp1(chatid string) {
	DSBot.ChannelMessageSend(chatid, fmt.Sprintf("Справка \n"+
		"ВНИМАНИЕ БОТ УДАЛЯЕТ СООБЩЕНИЯ \n ОТ ПОЛЬЗОВАТЕЛЕЙ ЧЕРЕЗ 3 МИНУТЫ \n\n"+
		"Встать в очередь: [4-11]+  или\n"+
		" [4-11]+[указать время ожидания в минутах]\n"+
		" 9+  встать в очередь на КЗ 9ур.\n"+
		" 9+60  встать на КЗ 9ур, время ожидания не более 60 минут.\n"+
		"Покинуть очередь: [4-11] -\n"+
		" 9- выйти из очереди КЗ 9ур.\n"+
		"Посмотреть список активных очередей: о[4-11]\n"+
		" о9  вывод очередь для вашей Кз\n"+
		"Получить роль кз: + [5-11]\n"+
		" +9   получить роль КЗ 9ур.\n"+
		" -9   снять роль "))
}
