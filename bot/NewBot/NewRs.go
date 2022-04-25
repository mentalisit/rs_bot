package NewBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func (in inMessage) RsPlus() {
	in.Mutex.Lock()
	defer in.Mutex.Unlock()
	go in.iftipdelete()
	if in.countName() == 1 { //проверяем есть ли игрок в очереди
		in.ifTipSendMentionText(" ты уже в очереди")
	} else {
		countQueue := in.countQueue()         //проверяем, есть ли кто-то в очереди
		numkzN := in.countNumberNameActive1() //проверяем количество боёв по уровню кз игрока
		numkzL := in.numberQueueLvl() + 1     //проверяем какой номер боя определенной красной звезды

		dsmesid := ""
		tgmesid := 0
		wamesid := ""

		if countQueue == 0 {
			if in.config.DsChannel != "" {
				name1 = fmt.Sprintf("%s  🕒  %s  (%d)", in.name, in.timekz, numkzN)
				name2 = ""
				name3 = ""
				name4 = ""
				lvlk := in.Ds.roleToIdPing(in)
				in.Ds.EmbedDS(name1, name2, name3, name4, lvlk, numkzL)
				mesCompl, err := DSBot.ChannelMessageSendComplex(in.config.DsChannel, &discordgo.MessageSend{
					Content: in.name + " запустил очередь " + lvlk})
				if err != nil {
					fmt.Println(err)
				}
				in.Ds.EditComplex(mesCompl.ID, in.config.DsChannel)
				in.Ds.addEnojiRsQueue(in.config.DsChannel, mesCompl.ID)
				dsmesid = mesCompl.ID
			}
			if in.config.TgChannel != 0 {
				text := fmt.Sprintf("Очередь кз%s (%d)\n1. %s - %sмин. (%d) \n\n%s++ - принудительный старт",
					in.lvlkz, numkzL, in.name, in.timekz, numkzN, in.lvlkz)
				tgmesid = tgSendEmded(in.lvlkz, in.config.TgChannel, text)
				in.SubscribePing(1)
			}
			if in.config.WaChannel != "" {
				//Тут будет логика ватса
			}

		} else if countQueue == 1 {
			u := in.readAll()
			dsmesid = u.user1.dsmesid

			if in.config.DsChannel != "" {
				name1 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user1.name, u.user1.timedown, u.user1.numkzn)
				name2 = fmt.Sprintf("%s  🕒  %s  (%d)", in.name, in.timekz, numkzN)
				name3 = ""
				name4 = ""
				lvlk := in.Ds.roleToIdPing(in)
				in.Ds.EmbedDS(name1, name2, name3, name4, lvlk, numkzL)
				text := lvlk + " 2/4 " + in.name + " присоединился к очереди"
				go in.Ds.SendChannelDel5s(in.config.DsChannel, text)
				in.Ds.EditComplex(u.user1.dsmesid, in.config.DsChannel)
			}
			if in.config.TgChannel != 0 {
				text1 := fmt.Sprintf("Очередь кз%s (%d)\n", in.lvlkz, numkzL)
				name1 = fmt.Sprintf("1. %s - %dмин. (%d) \n", u.user1.name, u.user1.timedown, u.user1.numkzn)
				name2 = fmt.Sprintf("2. %s - %sмин. (%d) \n", in.name, in.timekz, numkzN)
				text2 := fmt.Sprintf("\n%s++ - принудительный старт", in.lvlkz)
				text := fmt.Sprintf("%s %s %s %s", text1, name1, name2, text2)
				go tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
				tgmesid = tgSendEmded(in.lvlkz, in.config.TgChannel, text)
				mesidTgUpdate(tgmesid, in.lvlkz, in.config.CorpName)
			}
			if in.config.WaChannel != "" {
				//Тут будет логика ватса
			}

		} else if countQueue == 2 {
			u := in.readAll()
			dsmesid = u.user1.dsmesid

			if in.config.DsChannel != "" {
				name1 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user1.name, u.user1.timedown, u.user1.numkzn)
				name2 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user2.name, u.user2.timedown, u.user2.numkzn)
				name3 = fmt.Sprintf("%s  🕒  %s  (%d)", in.name, in.timekz, numkzN)
				name4 = ""
				lvlk := roleToIdPing(in.lvlkz, in.config.Config.Guildid)
				lvlk3 := roleToIdPing(in.lvlkz+"+", in.config.Config.Guildid)
				EmbedDS(name1, name2, name3, name4, lvlk, numkzL)
				text := lvlk + " 3/4 " + in.name + " присоединился к очереди " + lvlk3 + " нужен еще один для фулки"
				go dsSendChannelDel5s(in.config.DsChannel, text)
				dsEditComplex(u.user1.dsmesid, in.config.DsChannel)
			}
			if in.config.TgChannel != 0 {
				text1 := fmt.Sprintf("Очередь кз%s (%d)\n", in.lvlkz, numkzL)
				name1 = fmt.Sprintf("1. %s - %dмин. (%d) \n", u.user1.name, u.user1.timedown, u.user1.numkzn)
				name2 = fmt.Sprintf("2. %s - %dмин. (%d) \n", u.user2.name, u.user2.timedown, u.user2.numkzn)
				name3 = fmt.Sprintf("3. %s - %sмин. (%d) \n", in.name, in.timekz, numkzN)
				text2 := fmt.Sprintf("\n%s++ - принудительный старт", in.lvlkz)
				text := fmt.Sprintf("%s %s %s %s %s", text1, name1, name2, name3, text2)
				go tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
				tgmesid = tgSendEmded(in.lvlkz, in.config.TgChannel, text)
				mesidTgUpdate(tgmesid, in.lvlkz, in.config.CorpName)
				SubscribePing(in, in.lvlkz, 3)
			}
			if in.config.WaChannel != "" {
				//Тут будет логика ватса
			}

		} else if countQueue == 3 {
			u := in.readAll()
			textEvent, numkzEvent := event(in)
			numberevent := qweryNumevent1(in) //получаем номер ивета если он активен
			if numberevent > 0 {
				numkzL = numkzEvent
			}

			dsmesid := u.user1.dsmesid

			if in.config.DsChannel != "" {
				if u.user1.tip == "ds" {
					name1 = u.user1.mention
				} else {
					name1 = u.user1.name
				}
				if u.user2.tip == "ds" {
					name2 = u.user2.mention
				} else {
					name2 = u.user2.name
				}
				if u.user3.tip == "ds" {
					name3 = u.user3.mention
				} else {
					name3 = u.user3.name
				}
				if in.tip == "ds" {
					name4 = in.nameMention
				} else {
					name4 = in.name
				}
				go dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
				go dsSendChannelDel5s(in.config.DsChannel, " 4/4 "+in.name+" присоединился к очереди")
				text := fmt.Sprintf("4/4 Очередь КЗ%s сформирована\n %s %s\n %s %s \nВ ИГРУ %s", in.lvlkz, name1, name2, name3, name4, textEvent)
				dsmesid = dsSendChannel(in.config.DsChannel, text)
				mesidDsUpdate(dsmesid, in.lvlkz, in.config.DsChannel)
			}
			if in.config.TgChannel != 0 {
				if u.user1.tip == "tg" {
					name1 = u.user1.mention
				} else {
					name1 = u.user1.name
				}
				if u.user2.tip == "tg" {
					name2 = u.user2.mention
				} else {
					name2 = u.user2.name
				}
				if u.user3.tip == "tg" {
					name3 = u.user3.mention
				} else {
					name3 = u.user3.name
				}
				if in.tip == "tg" {
					name4 = in.nameMention
				} else {
					name4 = in.name
				}
				go tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
				go tgSendChannelDel5s(in.config.TgChannel, in.name+" закрыл очередь кз"+in.lvlkz)
				text := fmt.Sprintf("Очередь КЗ%s сформирована\n%s %s\n%s %s\n В ИГРУ \n%s",
					in.lvlkz, name1, name2, name3, name4, textEvent)
				tgmesid = tgSendChannel(in.config.TgChannel, text)
				mesidTgUpdate(tgmesid, in.lvlkz, in.config.CorpName)
			}
			if in.config.WaChannel != "" {
				//Тут будет логика ватса
			}

			updateComplite(in.lvlkz, dsmesid, tgmesid, wamesid, numkzL, numberevent, in.config.CorpName)
		}

		numevent := 0 //qweryNumevent1(in)
		mdate, mtime := currentTime()
		insertSborkztg1 := `INSERT INTO sborkz(corpname,name,mention,tip,dsmesid,tgmesid,wamesid,time,date,lvlkz,
                   numkzn,numberkz,numberevent,eventpoints,active,timedown) 
				VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(insertSborkztg1, in.config.CorpName, in.name, in.nameMention, in.tip, dsmesid, tgmesid,
			wamesid, mtime, mdate, in.lvlkz, numkzN, 0, numevent, 0, 0, in.timekz)
		if err != nil {
			log.Println(err)
		}

	}
}

func (in inMessage) Pl30() {
	countName := countName(in, in.lvlkz)
	text := ""
	if countName == 0 {
		text = in.nameMention + " ты не в очереди "
	} else if countName > 0 {
		var timedown int
		results, err := db.Query("SELECT timedown FROM sborkz WHERE lvlkz = ? AND corpname = ? AND active = 0 AND name = ?",
			in.lvlkz, in.config.CorpName, in.name)
		if err != nil {
			fmt.Println(err)
		}
		for results.Next() {
			err = results.Scan(&timedown)
			if timedown >= 150 {
				text = fmt.Sprintf("%s максимальное время в очереди ограничено на 180 минут\n твое время %d мин.  ", in.nameMention, timedown)
			} else {
				text = in.nameMention + " время обновлено +30"
				_, err := db.Exec(`update sborkz set timedown = timedown+30 where lvlkz = ? AND corpname = ? AND name = ?`,
					in.lvlkz, in.config.CorpName, in.name)
				if err != nil {
					log.Println(err)
				}
				in.option.callback = true
				in.option.edit = true
				in.Queue()
			}
		}
	}
	if in.tip == "ds" {
		go dsSendChannelDel5s(in.config.DsChannel, text)
	} else if in.tip == "tg" {
		go tgSendChannelDel5s(in.config.TgChannel, text)
	}
}

func (in inMessage) RsMinus() {
	in.Mutex.Lock()
	if in.tip == "ds" && !in.option.callback {
		go dsDelMessage(in.config.DsChannel, in.Ds.mesid)
	} else if in.tip == "tg" && !in.option.callback {
		go tgDelMessage(in.config.TgChannel, in.Tg.mesid)
	}
	CountNames := countName(in, in.lvlkz) //проверяем есть ли игрок в очереди
	if CountNames == 0 {
		if in.tip == "ds" {
			go dsSendChannelDel5s(in.config.DsChannel, in.nameMention+" ты не в очереди")
		} else if in.tip == "tg" {
			go tgSendChannelDel5s(in.config.TgChannel, in.nameMention+" ты не в очереди")
		}
	} else if CountNames > 0 {
		//чтение айди очечреди
		u := in.readAll()
		//удаление с БД
		_, err := db.Exec("delete from sborkz where name = ? AND lvlkz = ? AND corpname = ? AND active = 0",
			in.name, in.lvlkz, in.config.CorpName)
		if err != nil {
			log.Println(err)
		}
		//проверяем очередь
		countQueue := countQueue(in, in.lvlkz)
		//numkzL := numberQueueLvl(in, lvlkz) + 1
		if in.config.DsChannel != "" {
			go dsSendChannelDel5s(in.config.DsChannel, in.name+" покинул очередь")
			if countQueue == 0 {
				go dsSendChannelDel5s(in.config.DsChannel, "Очередь КЗ была удалена .")
				dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
			}
		}
		if in.config.TgChannel != 0 {
			go tgSendChannelDel5s(in.config.TgChannel, in.name+" покинул очередь")
			if countQueue == 0 {
				go tgSendChannelDel5s(in.config.TgChannel, "Очередь КЗ была удалена .")
				tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
			}
		}
		if in.config.WaChannel != "" {
			//тут логика ватса
		}
		in.Mutex.Unlock()
		if countQueue > 0 {
			in.Queue()
		}
	}
}

func (in inMessage) Queue() {
	in.Mutex.Lock()
	defer in.Mutex.Unlock()
	if in.tip == "ds" && !in.option.callback {
		go dsDelMessage(in.config.DsChannel, in.Ds.mesid)
	} else if in.tip == "tg" && !in.option.callback {
		go tgDelMessage(in.config.TgChannel, in.Tg.mesid)
	}
	count := countQueue(in, in.lvlkz)
	numberLvl := in.numberQueueLvl() + 1
	// совподения количество  условие
	if count == 0 {
		text := "Очередь КЗ " + in.lvlkz + " пуста "
		if in.tip == "ds" {
			go dsSendChannelDel5s(in.config.DsChannel, text)
		} else if in.tip == "tg" {
			go tgSendChannelDel5s(in.config.TgChannel, text)
		}
	} else if count == 1 {
		u := in.readAll()
		if in.config.DsChannel != "" {
			name1 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user1.name, u.user1.timedown, u.user1.numkzn)
			name2 = ""
			name3 = ""
			name4 = ""
			lvlk := roleToIdPing(in.lvlkz, in.config.Config.Guildid)
			EmbedDS(name1, name2, name3, name4, lvlk, numberLvl)
			if in.option.edit {
				dsEditComplex(u.user1.dsmesid, in.config.DsChannel)
			} else if !in.option.edit {
				DSBot.ChannelMessageDelete(in.config.DsChannel, u.user1.dsmesid)
				mesCompl, err := DSBot.ChannelMessageSendComplex(in.config.DsChannel, &discordgo.MessageSend{
					Content: mesContentNil,
					Embed:   Embeds})
				if err != nil {
					fmt.Println(err)
				}
				addEnojiRsQueue(in.config.DsChannel, mesCompl.ID)
				mesidDsUpdate(mesCompl.ID, in.lvlkz, in.config.CorpName)
			}
		}
		if in.config.TgChannel != 0 {
			text1 := fmt.Sprintf("Очередь кз%s (%d)\n", in.lvlkz, numberLvl)
			name1 = fmt.Sprintf("1. %s - %dмин. (%d) \n", u.user1.name, u.user1.timedown, u.user1.numkzn)
			text2 := fmt.Sprintf("\n%s++ - принудительный старт", in.lvlkz)
			text := fmt.Sprintf("%s %s %s", text1, name1, text2)
			if in.option.edit {
				tgEditMessageText(in.config.TgChannel, u.user1.tgmesid, text, in.lvlkz)
			} else if !in.option.edit {
				mesidTg := tgSendEmded(in.lvlkz, in.config.TgChannel, text)
				mesidTgUpdate(mesidTg, in.lvlkz, in.config.CorpName)
				go tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
			}
		}
		if in.config.WaChannel != "" {

		}
	} else if count == 2 {
		u := in.readAll()

		if in.config.DsChannel != "" {
			name1 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user1.name, u.user1.timedown, u.user1.numkzn)
			name2 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user2.name, u.user2.timedown, u.user2.numkzn)
			name3 = ""
			name4 = ""
			lvlk := roleToIdPing(in.lvlkz, in.config.Config.Guildid)
			EmbedDS(name1, name2, name3, name4, lvlk, numberLvl)
			if in.option.edit {
				dsEditComplex(u.user1.dsmesid, in.config.DsChannel)
			} else if !in.option.edit {
				dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
				mesCompl, err := DSBot.ChannelMessageSendComplex(in.config.DsChannel, &discordgo.MessageSend{
					Content: mesContentNil,
					Embed:   Embeds})
				if err != nil {
					fmt.Println(err)
				}
				addEnojiRsQueue(in.config.DsChannel, mesCompl.ID)
				mesidDsUpdate(mesCompl.ID, in.lvlkz, in.config.CorpName)
			}
		}
		if in.config.TgChannel != 0 {
			text1 := fmt.Sprintf("Очередь кз%s (%d)\n", in.lvlkz, numberLvl)
			name1 = fmt.Sprintf("1. %s - %dмин. (%d) \n", u.user1.name, u.user1.timedown, u.user1.numkzn)
			name2 = fmt.Sprintf("2. %s - %dмин. (%d) \n", u.user2.name, u.user2.timedown, u.user2.numkzn)
			text2 := fmt.Sprintf("\n%s++ - принудительный старт", in.lvlkz)
			text := fmt.Sprintf("%s %s %s %s", text1, name1, name2, text2)
			if in.option.edit {
				tgEditMessageText(in.config.TgChannel, u.user1.tgmesid, text, in.lvlkz)
			} else if !in.option.edit {
				mesidTg := tgSendEmded(in.lvlkz, in.config.TgChannel, text)
				mesidTgUpdate(mesidTg, in.lvlkz, in.config.CorpName)
				go tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
			}
		}
		if in.config.WaChannel != "" {

		}
	} else if count == 3 {
		u := in.readAll()

		if in.config.DsChannel != "" {
			name1 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user1.name, u.user1.timedown, u.user1.numkzn)
			name2 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user2.name, u.user2.timedown, u.user2.numkzn)
			name3 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user3.name, u.user3.timedown, u.user3.numkzn)
			name4 = ""
			lvlk := roleToIdPing(in.lvlkz, in.config.Config.Guildid)
			EmbedDS(name1, name2, name3, name4, lvlk, numberLvl)
			if in.option.edit {
				dsEditComplex(u.user1.dsmesid, in.config.DsChannel)
			} else if !in.option.edit {
				dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
				mesCompl, err := DSBot.ChannelMessageSendComplex(in.config.DsChannel, &discordgo.MessageSend{
					Content: mesContentNil,
					Embed:   Embeds})
				if err != nil {
					fmt.Println(err)
				}
				addEnojiRsQueue(in.config.DsChannel, mesCompl.ID)
				mesidDsUpdate(mesCompl.ID, in.lvlkz, in.config.CorpName)
			}
		}
		if in.config.TgChannel != 0 {
			text1 := fmt.Sprintf("Очередь кз%s (%d)\n", in.lvlkz, numberLvl)
			name1 = fmt.Sprintf("1. %s - %dмин. (%d) \n", u.user1.name, u.user1.timedown, u.user1.numkzn)
			name2 = fmt.Sprintf("2. %s - %dмин. (%d) \n", u.user2.name, u.user2.timedown, u.user2.numkzn)
			name3 = fmt.Sprintf("3. %s - %dмин. (%d) \n", u.user3.name, u.user3.timedown, u.user3.numkzn)
			text2 := fmt.Sprintf("\n%s++ - принудительный старт", in.lvlkz)
			text := fmt.Sprintf("%s %s %s %s %s", text1, name1, name2, name3, text2)
			if in.option.edit {
				tgEditMessageText(in.config.TgChannel, u.user1.tgmesid, text, in.lvlkz)
			} else if !in.option.edit {
				mesidTg := tgSendEmded(in.lvlkz, in.config.TgChannel, text)
				mesidTgUpdate(mesidTg, in.lvlkz, in.config.CorpName)
				go tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
			}
		}
		if in.config.WaChannel != "" {

		}
	}
}

func (in inMessage) RsStart() {
	in.Mutex.Lock()
	defer in.Mutex.Unlock()
	if !in.option.callback {
		if in.tip == "ds" {
			dsDeleteMesage5s(in.config.DsChannel, in.Ds.mesid)
		} else if in.tip == "tg" {
			tgDelMessage10s(in.config.TgChannel, in.Tg.mesid)
		}
	}
	countName := countName(in, in.lvlkz)
	if countName == 0 {
		if in.tip == "ds" {
			dsSendChannelDel5s(in.config.DsChannel, "Принудительный старт доступен участникам очереди.")
		} else if in.tip == "tg" {
			tgSendChannelDel5s(in.config.TgChannel, "Принудительный старт доступен участникам очереди.")
		}
	} else if countName == 1 {
		numberkz := in.numberQueueLvl()
		count := countQueue(in, in.lvlkz)
		if count == 1 {
			u := in.readAll()

			textEvent, numkzEvent := event(in)
			numberevent := qweryNumevent1(in)
			if numberevent > 0 {
				numberkz = numkzEvent
			}

			dsmesid := ""
			tgmesid := 0
			wamesid := ""
			if in.config.DsChannel != "" {
				if u.user1.tip == "ds" {
					name1 = u.user1.mention
				} else {
					name1 = u.user1.name
				}
				text := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n\n1. %s\nВ игру %s",
					in.lvlkz, numberkz, name1, textEvent)
				dsmesid = dsSendChannel(in.config.DsChannel, text)
				go dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
				mesidDsUpdate(dsmesid, in.lvlkz, in.config.CorpName)
			}
			if in.config.TgChannel != 0 {
				if u.user1.tip == "tg" {
					name1 = u.user1.mention
				} else {
					name1 = u.user1.name
				}
				go tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
				text := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n\n1. %s\nВ игру %s",
					in.lvlkz, numberkz, name1, textEvent)
				tgmesid := tgSendChannel(in.config.TgChannel, text)
				mesidTgUpdate(tgmesid, in.lvlkz, in.config.CorpName)
			}
			updateComplite(in.lvlkz, dsmesid, tgmesid, wamesid, numberkz, numberevent, in.config.CorpName)
		} else if count == 2 {
			u := in.readAll()

			textEvent, numkzEvent := event(in)
			numberevent := qweryNumevent1(in)
			if numberevent > 0 {
				numberkz = numkzEvent
			}
			dsmesid := ""
			tgmesid := 0
			wamesid := ""

			if in.config.DsChannel != "" { //discord
				if u.user1.tip == "ds" {
					name1 = u.user1.mention
				} else {
					name1 = u.user1.name
				}
				if u.user2.tip == "ds" {
					name2 = u.user2.mention
				} else {
					name2 = u.user2.name
				}
				text1 := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n", in.lvlkz, numberkz)
				text2 := fmt.Sprintf("\n%s %s\nВ игру %s", name1, name2, textEvent)
				text := text1 + text2
				dsmesid = dsSendChannel(in.config.DsChannel, text)
				go dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
				mesidDsUpdate(dsmesid, in.lvlkz, in.config.CorpName)
			}
			if in.config.TgChannel != 0 { //telegram
				if u.user1.tip == "tg" {
					name1 = u.user1.mention
				} else {
					name1 = u.user1.name
				}
				if u.user2.tip == "tg" {
					name2 = u.user2.mention
				} else {
					name2 = u.user2.name
				}
				go tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
				text1 := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n", in.lvlkz, numberkz)
				text2 := fmt.Sprintf("\n%s %s\nВ игру %s", name1, name2, textEvent)
				tgmesid = tgSendChannel(in.config.TgChannel, text1+text2)
				mesidTgUpdate(tgmesid, in.lvlkz, in.config.CorpName)
			}
			updateComplite(in.lvlkz, dsmesid, tgmesid, wamesid, numberkz, numberevent, in.config.CorpName)

		} else if count == 3 {
			u := in.readAll()

			textEvent, numkzEvent := event(in)
			numberevent := qweryNumevent1(in)
			if numberevent > 0 {
				numberkz = numkzEvent
			}
			dsmesid := ""
			tgmesid := 0
			wamesid := ""

			if in.config.DsChannel != "" { //discord
				if u.user1.tip == "ds" {
					name1 = u.user1.mention
				} else {
					name1 = u.user1.name
				}
				if u.user2.tip == "ds" {
					name2 = u.user2.mention
				} else {
					name2 = u.user2.name
				}
				if u.user3.tip == "ds" {
					name3 = u.user3.mention
				} else {
					name3 = u.user3.name
				}
				text := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n\n%s %s %s\nВ игру %s",
					in.lvlkz, numberkz, name1, name2, name3, textEvent)
				dsmesid = dsSendChannel(in.config.DsChannel, text)
				go dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
				mesidDsUpdate(dsmesid, in.lvlkz, in.config.CorpName)
			}
			if in.config.TgChannel != 0 { //telegram
				if u.user1.tip == "tg" {
					name1 = u.user1.mention
				} else {
					name1 = u.user1.name
				}
				if u.user2.tip == "tg" {
					name2 = u.user2.mention
				} else {
					name2 = u.user2.name
				}
				if u.user3.tip == "tg" {
					name3 = u.user3.mention
				} else {
					name3 = u.user3.name
				}
				go tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
				text := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n\n%s %s %s\nВ игру %s",
					in.lvlkz, numberkz, name1, name2, name3, textEvent)
				tgmesid = tgSendChannel(in.config.TgChannel, text)
				mesidTgUpdate(tgmesid, in.lvlkz, in.config.CorpName)
			}
			updateComplite(in.lvlkz, dsmesid, tgmesid, wamesid, numberkz, numberevent, in.config.CorpName)
		}
	}
}

func (in inMessage) Plus() bool {

	if !in.option.callback {
		if in.tip == "ds" {
			go dsDeleteMesage5s(in.config.DsChannel, in.Ds.mesid)
		} else if in.tip == "tg" {
			go tgDelMessage10s(in.config.TgChannel, in.Tg.mesid)
		}
	}

	var countNames int
	message := ""
	ins := true
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND active = 0", in.name, in.config.CorpName)
	err := row.Scan(&countNames)
	if err != nil {
		fmt.Println(err)
	}

	if countNames == 0 {
		message = in.nameMention + " ты не в очереди"
		ins = false
	} else if countNames > 0 {
		results, err := db.Query("SELECT * FROM sborkz WHERE name = ? AND corpname = ? AND active = 0", in.name, in.config.CorpName)
		if err != nil {
			fmt.Println(err)
		}
		for results.Next() {
			var t sborkz
			err = results.Scan(&t.id, &t.corpname, &t.name, &t.mention, &t.tip, &t.dsmesid, &t.tgmesid, &t.wamesid, &t.time, &t.date, &t.lvlkz, &t.numkzn, &t.numberkz, &t.numberevent, &t.eventpoints, &t.active, &t.timedown)
			if t.name == in.name && t.timedown > 3 {
				message = fmt.Sprintf("%s рановато плюсик жмешь, ты в очереди на кз%s будешь еще %dмин",
					t.mention, t.lvlkz, t.timedown)
			} else if t.name == in.name && t.timedown <= 3 {
				message = t.mention + " время обновлено "
				_, err := db.Exec("update sborkz set timedown = timedown + 30 where active = 0 AND name = ? AND corpname = ?", t.name, t.corpname)
				if err != nil {
					log.Println(err)
				}
				in.lvlkz = t.lvlkz
				go in.Queue()
			}
		}
	}
	if in.tip == "ds" {
		go dsSendChannelDel5s(in.config.DsChannel, message)
	} else if in.tip == "tg" {
		go tgSendChannelDel5s(in.config.TgChannel, message)
	}

	return ins
}

func (in inMessage) Minus() bool {

	if !in.option.callback {
		if in.tip == "ds" {
			dsDeleteMesage5s(in.config.DsChannel, in.Ds.mesid)
		} else if in.tip == "tg" {
			tgDelMessage10s(in.config.TgChannel, in.Tg.mesid)
		}
	}

	inm := true
	message := ""
	var countNames int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND active = 0", in.name, in.config.CorpName)
	err := row.Scan(&countNames)
	if err != nil {
		fmt.Println(err)
	}

	if countNames == 0 {
		message = in.nameMention + " ты не в очереди"
		inm = false
	} else if countNames > 0 {
		results, err := db.Query("SELECT * FROM sborkz WHERE name = ? AND corpname = ? AND active = 0", in.name, in.config.CorpName)
		if err != nil {
			fmt.Println(err)
		}
		for results.Next() {
			var t sborkz
			err = results.Scan(&t.id, &t.corpname, &t.name, &t.mention, &t.tip, &t.dsmesid, &t.tgmesid, &t.wamesid, &t.time, &t.date, &t.lvlkz, &t.numkzn, &t.numberkz, &t.numberevent, &t.eventpoints, &t.active, &t.timedown)
			if t.name == in.name && t.timedown > 3 {
				message = fmt.Sprintf("%s рановато минус жмешь, ты в очереди на кз%s будешь еще %dмин", t.mention, t.lvlkz, t.timedown)
			} else if t.name == in.name && t.timedown <= 3 {

				in.lvlkz = t.lvlkz
				in.RsMinus()
			}
		}
	}
	if len(message) > 0 {
		if in.tip == "ds" {
			go dsSendChannelDel5s(in.config.DsChannel, message)
		} else if in.tip == "tg" {
			go tgSendChannelDel5s(in.config.TgChannel, message)
		}
	}

	return inm
}

//проверка есть ли игрок в очереди
func (in inMessage) countName() int {
	var countNames int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND lvlkz = ? AND corpname = ? AND active = 0",
		in.name, in.lvlkz, in.config.CorpName)
	err := row.Scan(&countNames)
	if err != nil {
		logrus.Println(err)
	}
	return countNames
}

//проверка сколько игровок в очереди
func (in inMessage) countQueue() int {
	var count int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE lvlkz = ? AND corpname = ? AND active = 0",
		in.lvlkz, in.config.CorpName)
	err := row.Scan(&count)
	if err != nil {
		fmt.Println(err)
	}
	return count
}

// выковыриваем из базы значение количества походов на кз
func (in inMessage) countNumberNameActive1() int {
	var countNumberNameActive1 int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE lvlkz = ? AND corpname = ? AND name = ? AND active = 1",
		in.lvlkz, in.config.CorpName, in.name)
	err := row.Scan(&countNumberNameActive1)
	if err != nil {
		fmt.Println(err)
	}
	return countNumberNameActive1
}

func (in inMessage) numberQueueLvl() int {
	var number int
	row2 := db.QueryRow("SELECT  number FROM numkz WHERE lvlkz = ? AND corpname = ?",
		in.lvlkz, in.config.CorpName)
	err2 := row2.Scan(&number)
	if err2 != nil {
		fmt.Println(err2)
	}
	if number == 0 {
		insertSmt := "INSERT INTO numkz(lvlkz, number,corpname) VALUES (?,?,?)"
		statement, err := db.Prepare(insertSmt)
		if err != nil {
			fmt.Println(err)
		}
		number = 1
		_, err = statement.Exec(in.lvlkz, number, in.config.CorpName)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		return number
	}
	return number
}

func (in inMessage) readAll() (users Users) {
	u := Users{
		user1: sborkz{},
		user2: sborkz{},
		user3: sborkz{},
		user4: sborkz{},
	}
	user := 1
	results, err := db.Query("SELECT * FROM sborkz WHERE lvlkz = ? AND corpname = ? AND active = 0", in.lvlkz, in.config.CorpName)
	if err != nil {
		fmt.Println(err)
	}
	for results.Next() {
		var t sborkz
		err = results.Scan(&t.id, &t.corpname, &t.name, &t.mention, &t.tip, &t.dsmesid, &t.tgmesid, &t.wamesid, &t.time, &t.date, &t.lvlkz, &t.numkzn, &t.numberkz, &t.numberevent, &t.eventpoints, &t.active, &t.timedown)
		if user == 1 {
			u.user1 = t
		} else if user == 2 {
			u.user2 = t
		} else if user == 3 {
			u.user3 = t
		} else if user == 4 {
			u.user4 = t
		}
		user = user + 1
	}
	return u
}

func currentTime() (string, string) {
	tm := time.Now()
	mdate := (tm.Format("2006-01-02"))
	mtime := (tm.Format("15:04"))

	return mdate, mtime
}
