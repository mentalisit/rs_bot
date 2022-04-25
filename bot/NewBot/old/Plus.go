package old

/*
import (
	"fmt"
	"log"
	"rs_bot/bot/NewBot"
)

func Plus(in NewBot.inMessage) bool {
	if !in.option.callback {
		if in.tip == "ds" {
			NewBot.dsDeleteMesage5s(in.config.DsChannel, in.Ds.mesid)
		} else if in.tip == "tg" {
			NewBot.tgDelMessage10s(in.config.TgChannel, in.Tg.mesid)
		}
	}

	var countNames int
	message := ""
	ins := true
	row := NewBot.db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND active = 0", in.name, in.config.CorpName)
	err := row.Scan(&countNames)
	if err != nil {
		fmt.Println(err)
	}
	if countNames == 0 {
		message = in.nameMention + " ты не в очереди"
		ins = false
	} else if countNames > 0 {
		results, err := NewBot.db.Query("SELECT * FROM sborkz WHERE name = ? AND corpname = ? AND active = 0", in.name, in.config.CorpName)
		if err != nil {
			fmt.Println(err)
		}
		for results.Next() {
			var t NewBot.sborkz
			err = results.Scan(&t.id, &t.corpname, &t.name, &t.mention, &t.tip, &t.dsmesid, &t.tgmesid, &t.wamesid, &t.time, &t.date, &t.lvlkz, &t.numkzn, &t.numberkz, &t.numberevent, &t.eventpoints, &t.active, &t.timedown)
			if t.name == in.name && t.timedown > 3 {
				message = fmt.Sprintf("%s рановато плюсик жмешь, ты в очереди на кз%d будешь еще %dмин", t.mention, t.lvlkz, t.timedown)
			} else if t.name == in.name && t.timedown <= 3 {
				message = t.mention + " время обновлено "
				_, err := NewBot.db.Exec("update sborkz set timedown = timedown + 30 where active = 0 AND name = ? AND corpname = ?", t.name, t.corpname)
				if err != nil {
					log.Println(err)
				}
				in.lvlkz=t.lvlkz
				in.Queue()
			}
		}
	}
	if in.tip == "ds" {
		go NewBot.dsSendChannelDel5s(in.config.DsChannel, message)
	} else if in.tip == "tg" {
		go NewBot.tgSendChannelDel5s(in.config.TgChannel, message)
	}
	return ins
}



*/
