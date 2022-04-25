package old

/*
import (
	"log"
	"rs_bot/bot/NewBot"
)

func RsMinusOld(in NewBot.inMessage, lvlkz string) {
	if in.tip == "ds" && !in.option.callback {
		go NewBot.dsDelMessage(in.config.DsChannel, in.Ds.mesid)
	} else if in.tip == "tg" && !in.option.callback {
		go NewBot.tgDelMessage(in.config.TgChannel, in.Tg.mesid)
	}
	CountNames := NewBot.countName(in, lvlkz) //проверяем есть ли игрок в очереди
	if CountNames == 0 {
		if in.tip == "ds" {
			go NewBot.dsSendChannelDel5s(in.config.DsChannel, in.nameMention+" ты не в очереди")
		} else if in.tip == "tg" {
			go NewBot.tgSendChannelDel5s(in.config.TgChannel, in.nameMention+" ты не в очереди")
		}
	} else if CountNames > 0 {
		//чтение айди очечреди
		u := NewBot.readAll(in)
		//удаление с БД
		deleteSrorkz(in, lvlkz)
		//проверяем очередь
		countQueue := NewBot.countQueue(in, lvlkz)
		//numkzL := numberQueueLvl(in, lvlkz) + 1
		if in.config.DsChannel != "" {
			go NewBot.dsSendChannelDel5s(in.config.DsChannel, in.name+" покинул очередь")
			if countQueue == 0 {
				go NewBot.dsSendChannelDel5s(in.config.DsChannel, "Очередь КЗ была удалена .")
				NewBot.dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
			}
		}
		if in.config.TgChannel != 0 {
			go NewBot.tgSendChannelDel5s(in.config.TgChannel, in.name+" покинул очередь")
			if countQueue == 0 {
				go NewBot.tgSendChannelDel5s(in.config.TgChannel, "Очередь КЗ была удалена .")
				NewBot.tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
			}
		}
		if in.config.WaChannel != "" {
			//тут логика ватса
		}
		if countQueue > 0 {
			Queue(in, lvlkz)
		}
	}
}

func deleteSrorkz(in NewBot.inMessage, lvlkz string) {
	_, err := NewBot.db.Exec("delete from sborkz where name = ? AND lvlkz = ? AND corpname = ? AND active = 0", in.name, lvlkz, in.config.CorpName)
	if err != nil {
		log.Println(err)
	}
}

*/
