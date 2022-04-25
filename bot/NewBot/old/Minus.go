package old

/*
import (
	"fmt"
	"rs_bot/bot/NewBot"
)

func Minus(in NewBot.inMessage) bool {
	if !in.option.callback {
		if in.tip == "ds" {
			NewBot.dsDeleteMesage5s(in.config.DsChannel, in.Ds.mesid)
		} else if in.tip == "tg" {
			NewBot.tgDelMessage10s(in.config.TgChannel, in.Tg.mesid)
		}
	}

	inm := true
	message := ""
	var countNames int
	row := NewBot.db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND corpname = ? AND active = 0", in.name, in.config.CorpName)
	err := row.Scan(&countNames)
	if err != nil {
		fmt.Println(err)
	}
	if countNames == 0 {
		message = in.nameMention + " ты не в очереди"
		inm = false
	} else if countNames > 0 {
		results, err := NewBot.db.Query("SELECT * FROM sborkz WHERE name = ? AND corpname = ? AND active = 0", in.name, in.config.CorpName)
		if err != nil {
			fmt.Println(err)
		}
		for results.Next() {
			var t NewBot.sborkz
			err = results.Scan(&t.id, &t.corpname, &t.name, &t.mention, &t.tip, &t.dsmesid, &t.tgmesid, &t.wamesid, &t.time, &t.date, &t.lvlkz, &t.numkzn, &t.numberkz, &t.numberevent, &t.eventpoints, &t.active, &t.timedown)
			if t.name == in.name && t.timedown > 3 {
				message = fmt.Sprintf("%s рановато минус жмешь, ты в очереди на кз%s будешь еще %dмин", t.mention, t.lvlkz, t.timedown)
			} else if t.name == in.name && t.timedown <= 3 {
				in.lvlkz=t.lvlkz
				in.RsMinus()
				//RsMinus(in, t.lvlkz)
			}
		}
	}
	if len(message)>0{
		if in.tip == "ds" {
			go NewBot.dsSendChannelDel5s(in.config.DsChannel, message)
		} else if in.tip == "tg" {
			go NewBot.tgSendChannelDel5s(in.config.TgChannel, message)
		}
	}

	return inm
}


*/
