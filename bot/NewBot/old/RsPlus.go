package old

/*
import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"rs_bot/bot/NewBot"
)

func RsPlusOld(in NewBot.inMessage, lvlkz, timekz string) {
	if in.tip == "ds" && !in.option.callback {
		go NewBot.dsDelMessage(in.config.DsChannel, in.Ds.mesid)
	} else if in.tip == "tg" && !in.option.callback {
		go NewBot.tgDelMessage(in.config.TgChannel, in.Tg.mesid)
	}
	CountNames := NewBot.countName(in, lvlkz) //проверяем есть ли игрок в очереди
	if CountNames == 1 {
		if in.tip == "ds" {
			go NewBot.dsSendChannelDel5s(in.config.DsChannel, in.nameMention+" ты уже в очереди")
		} else if in.tip == "tg" {
			go NewBot.tgSendChannelDel5s(in.config.TgChannel, in.nameMention+" ты уже в очереди")
		}
	} else {
		countQueue := NewBot.countQueue(in, lvlkz)               //проверяем, есть ли кто-то в очереди
		numberkzName := NewBot.countNumberNameActive1(in, lvlkz) //проверяем количество боёв по уровню кз игрока
		numberQueueLv := NewBot.numberQueueLvl(in, lvlkz) + 1    //проверяем какой номер боя определенной красной звезды
		if countQueue == 0 {
			counts0(in, lvlkz, timekz, numberkzName, numberQueueLv)
		} else if countQueue == 1 {
			counts1(in, lvlkz, timekz, numberkzName, numberQueueLv)
		} else if countQueue == 2 {
			counts2(in, lvlkz, timekz, numberkzName, numberQueueLv)
		} else if countQueue == 3 {
			counts3(in, lvlkz, timekz, numberkzName, numberQueueLv)
		}
	}
}

//логика очереди
func counts0(in NewBot.inMessage, lvlkz, timekz string, numkzN, numkzL int) {
	dsmesid := ""
	tgmesid := 0
	wamesid := ""
	if in.config.DsChannel != "" {
		NewBot.name1 = fmt.Sprintf("%s  🕒  %s  (%d)", in.name, timekz, numkzN)
		NewBot.name2 = ""
		NewBot.name3 = ""
		NewBot.name4 = ""
		lvlk := NewBot.roleToIdPing(lvlkz, in.config.Config.Guildid)
		NewBot.EmbedDS(NewBot.name1, NewBot.name2, NewBot.name3, NewBot.name4, lvlk, numkzL)
		mesCompl, err := NewBot.DSBot.ChannelMessageSendComplex(in.config.DsChannel, &discordgo.MessageSend{
			Content: in.name + " запустил очередь " + lvlk})
		if err != nil {
			fmt.Println(err)
		}
		NewBot.dsEditComplex(mesCompl.ID, in.config.DsChannel)
		NewBot.addEnojiRsQueue(in.config.DsChannel, mesCompl.ID)
		dsmesid = mesCompl.ID
	}
	if in.config.TgChannel != 0 {
		text := fmt.Sprintf("Очередь кз%s (%d)\n1. %s - %sмин. (%d) \n\n%s++ - принудительный старт",
			lvlkz, numkzL, in.name, timekz, numkzN, lvlkz)
		tgmesid = NewBot.tgSendEmded(lvlkz, in.config.TgChannel, text)
		NewBot.SubscribePing(in, lvlkz, 1)
	}
	if in.config.WaChannel != "" {
		//Тут будет логика ватса
	}
	insertSborkzAll(in, lvlkz, timekz, dsmesid, tgmesid, wamesid, numkzN)
}

func counts1(in NewBot.inMessage, lvlkz, timekz string, numkzN, numkzL int) {
	u := NewBot.readAll(in)
	dsmesid := u.user1.dsmesid
	tgmesid := 0
	wamesid := ""
	if in.config.DsChannel != "" {
		NewBot.name1 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user1.name, u.user1.timedown, u.user1.numkzn)
		NewBot.name2 = fmt.Sprintf("%s  🕒  %s  (%d)", in.name, timekz, numkzN)
		NewBot.name3 = ""
		NewBot.name4 = ""
		lvlk := NewBot.roleToIdPing(lvlkz, in.config.Config.Guildid)
		NewBot.EmbedDS(NewBot.name1, NewBot.name2, NewBot.name3, NewBot.name4, lvlk, numkzL)
		text := lvlk + " 2/4 " + in.name + " присоединился к очереди"
		go NewBot.dsSendChannelDel5s(in.config.DsChannel, text)
		NewBot.dsEditComplex(u.user1.dsmesid, in.config.DsChannel)
	}
	if in.config.TgChannel != 0 {
		text1 := fmt.Sprintf("Очередь кз%s (%d)\n", lvlkz, numkzL)
		NewBot.name1 = fmt.Sprintf("1. %s - %dмин. (%d) \n", u.user1.name, u.user1.timedown, u.user1.numkzn)
		NewBot.name2 = fmt.Sprintf("2. %s - %sмин. (%d) \n", in.name, timekz, numkzN)
		text2 := fmt.Sprintf("\n%s++ - принудительный старт", lvlkz)
		text := fmt.Sprintf("%s %s %s %s", text1, NewBot.name1, NewBot.name2, text2)
		go NewBot.tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
		tgmesid = NewBot.tgSendEmded(lvlkz, in.config.TgChannel, text)
		NewBot.mesidTgUpdate(tgmesid, lvlkz, in.config.CorpName)
	}
	if in.config.WaChannel != "" {
		//Тут будет логика ватса
	}
	insertSborkzAll(in, lvlkz, timekz, dsmesid, tgmesid, wamesid, numkzN)
}

func counts2(in NewBot.inMessage, lvlkz, timekz string, numkzN, numkzL int) {
	u := NewBot.readAll(in)
	dsmesid := u.user1.dsmesid
	tgmesid := 0
	wamesid := ""
	if in.config.DsChannel != "" {
		NewBot.name1 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user1.name, u.user1.timedown, u.user1.numkzn)
		NewBot.name2 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user2.name, u.user2.timedown, u.user2.numkzn)
		NewBot.name3 = fmt.Sprintf("%s  🕒  %s  (%d)", in.name, timekz, numkzN)
		NewBot.name4 = ""
		lvlk := NewBot.roleToIdPing(lvlkz, in.config.Config.Guildid)
		NewBot.EmbedDS(NewBot.name1, NewBot.name2, NewBot.name3, NewBot.name4, lvlk, numkzL)
		text := lvlk + " 3/4 " + in.name + " присоединился к очереди"
		go NewBot.dsSendChannelDel5s(in.config.DsChannel, text)
		NewBot.dsEditComplex(u.user1.dsmesid, in.config.DsChannel)
	}
	if in.config.TgChannel != 0 {
		text1 := fmt.Sprintf("Очередь кз%s (%d)\n", lvlkz, numkzL)
		NewBot.name1 = fmt.Sprintf("1. %s - %dмин. (%d) \n", u.user1.name, u.user1.timedown, u.user1.numkzn)
		NewBot.name2 = fmt.Sprintf("2. %s - %dмин. (%d) \n", u.user2.name, u.user2.timedown, u.user2.numkzn)
		NewBot.name3 = fmt.Sprintf("3. %s - %sмин. (%d) \n", in.name, timekz, numkzN)
		text2 := fmt.Sprintf("\n%s++ - принудительный старт", lvlkz)
		text := fmt.Sprintf("%s %s %s %s %s", text1, NewBot.name1, NewBot.name2, NewBot.name3, text2)
		go NewBot.tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
		tgmesid = NewBot.tgSendEmded(lvlkz, in.config.TgChannel, text)
		NewBot.mesidTgUpdate(tgmesid, lvlkz, in.config.CorpName)
		NewBot.SubscribePing(in, lvlkz, 3)
	}
	if in.config.WaChannel != "" {
		//Тут будет логика ватса
	}
	insertSborkzAll(in, lvlkz, timekz, dsmesid, tgmesid, wamesid, numkzN)
}

func counts3(in NewBot.inMessage, lvlkz, timekz string, numkzN, numkzL int) {
	u := NewBot.readAll(in)
	textEvent, numkzEvent := NewBot.event(in)
	numberevent := NewBot.qweryNumevent1(in) //получаем номер ивета если он активен
	if numberevent > 0 {
		numkzL = numkzEvent
	}

	dsmesid := u.user1.dsmesid
	tgmesid := 0
	wamesid := ""
	if in.config.DsChannel != "" {
		if u.user1.tip == "ds" {
			NewBot.name1 = u.user1.mention
		} else {
			NewBot.name1 = u.user1.name
		}
		if u.user2.tip == "ds" {
			NewBot.name2 = u.user2.mention
		} else {
			NewBot.name2 = u.user2.name
		}
		if u.user3.tip == "ds" {
			NewBot.name3 = u.user3.mention
		} else {
			NewBot.name3 = u.user3.name
		}
		if in.tip == "ds" {
			NewBot.name4 = in.nameMention
		} else {
			NewBot.name4 = in.name
		}
		go NewBot.dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
		go NewBot.dsSendChannelDel5s(in.config.DsChannel, " 4/4 "+in.name+" присоединился к очереди")
		text := fmt.Sprintf("4/4 Очередь КЗ%s сформирована\n %s %s\n %s %s \nВ ИГРУ %s", lvlkz, NewBot.name1, NewBot.name2, NewBot.name3, NewBot.name4, textEvent)
		dsmesid = NewBot.dsSendChannel(in.config.DsChannel, text)
		NewBot.mesidDsUpdate(dsmesid, lvlkz, in.config.DsChannel)
	}
	if in.config.TgChannel != 0 {
		if u.user1.tip == "tg" {
			NewBot.name1 = u.user1.mention
		} else {
			NewBot.name1 = u.user1.name
		}
		if u.user2.tip == "tg" {
			NewBot.name2 = u.user2.mention
		} else {
			NewBot.name2 = u.user2.name
		}
		if u.user3.tip == "tg" {
			NewBot.name3 = u.user3.mention
		} else {
			NewBot.name3 = u.user3.name
		}
		if in.tip == "tg" {
			NewBot.name4 = in.nameMention
		} else {
			NewBot.name4 = in.name
		}
		go NewBot.tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
		go NewBot.tgSendChannelDel5s(in.config.TgChannel, in.name+" закрыл очередь кз"+lvlkz)
		text := fmt.Sprintf("Очередь КЗ%s сформирована\n%s %s\n%s %s\n В ИГРУ \n%s",
			lvlkz, NewBot.name1, NewBot.name2, NewBot.name3, NewBot.name4, textEvent)
		tgmesid = NewBot.tgSendChannel(in.config.TgChannel, text)
		NewBot.mesidTgUpdate(tgmesid, lvlkz, in.config.CorpName)
	}
	if in.config.WaChannel != "" {
		//Тут будет логика ватса
	}
	insertSborkzAll(in, lvlkz, timekz, dsmesid, tgmesid, wamesid, numkzN)
	NewBot.updateComplite(lvlkz, dsmesid, tgmesid, wamesid, numkzL, numberevent, in.config.CorpName)
}

func insertSborkzAll(in NewBot.inMessage, lvlkz string, timekz, dsmesid string, tgmesid int, wamesid string, numkzn int) { // внесение в базу данных
	numevent := 0 //qweryNumevent1(in)
	mdate, mtime := currentTime()
	insertSborkztg1 := `INSERT INTO sborkz(corpname,name,mention,tip,dsmesid,tgmesid,wamesid,time,date,lvlkz,
                   numkzn,numberkz,numberevent,eventpoints,active,timedown)
				VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := NewBot.db.Exec(insertSborkztg1, in.config.CorpName, in.name, in.nameMention, in.tip, dsmesid, tgmesid,
		wamesid, mtime, mdate, lvlkz, numkzn, 0, numevent, 0, 0, timekz)
	if err != nil {
		log.Println(err)
	}
}



func Pl30Old(in NewBot.inMessage, lvlkz string) {
	countName := NewBot.countName(in, lvlkz)
	text := ""
	if countName == 0 {
		text = in.nameMention + " ты не в очереди "
	} else if countName > 0 {
		var timedown int
		results, err := NewBot.db.Query("SELECT timedown FROM sborkz WHERE lvlkz = ? AND corpname = ? AND active = 0 AND name = ?",
			lvlkz, in.config.CorpName, in.name)
		if err != nil {
			fmt.Println(err)
		}
		for results.Next() {
			err = results.Scan(&timedown)
			if timedown >= 150 {
				text = fmt.Sprintf("%s максимальное время в очереди ограничено на 180 минут\n твое время %d мин.  ", in.nameMention, timedown)
			} else {
				text = in.nameMention + " время обновлено +30"
				_, err := NewBot.db.Exec(`update sborkz set timedown = timedown+30 where lvlkz = ? AND corpname = ? AND name = ?`,
					lvlkz, in.config.CorpName, in.name)
				if err != nil {
					log.Println(err)
				}
				in.option.callback = true
				in.option.edit = true
				Queue(in, lvlkz)
			}
		}
	}
	if in.tip == "ds" {
		go NewBot.dsSendChannelDel5s(in.config.DsChannel, text)
	} else if in.tip == "tg" {
		go NewBot.tgSendChannelDel5s(in.config.TgChannel, text)
	}
}


*/
