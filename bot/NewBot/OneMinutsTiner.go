package NewBot

import (
	"fmt"
)

func MinusMin() {
	_, err := db.Exec(`update sborkz set timedown = timedown - 1 where active = 0`)
	if err != nil {
		fmt.Println(err)
	}
	msqlTimeo()
	timerDeleteMessage()
	autohelp()
}

func msqlTimeo() {
	results, err := db.Query("SELECT * FROM sborkz WHERE active = 0")
	if err != nil {
		fmt.Println(err)
	}
	for results.Next() {
		var t sborkz
		err = results.Scan(&t.id, &t.corpname, &t.name, &t.mention, &t.tip, &t.dsmesid, &t.tgmesid, &t.wamesid, &t.time, &t.date, &t.lvlkz, &t.numkzn, &t.numberkz, &t.numberevent, &t.eventpoints, &t.active, &t.timedown)
		ok, config := checkCorpNameConfig(t.corpname)
		if ok {
			in := inMessage{
				mtext:       "",
				tip:         t.tip,
				name:        t.name,
				nameMention: t.mention,
				lvlkz:       t.lvlkz,
				Ds: Ds{
					mesid:   t.dsmesid,
					nameid:  "",
					guildid: config.Config.Guildid,
				},
				Tg: Tg{
					mesid:  t.tgmesid,
					nameid: 0,
				},
				config: config,
			}

			if t.timedown == 3 {
				text := t.mention + " время почти вышло  ...\n если ты еще тут пиши +"
				if t.tip == "ds" {
					mes3s := dsSendChannel(in.config.DsChannel, text)
					DSBot.MessageReactionAdd(in.config.DsChannel, mes3s, emPlus)
					DSBot.MessageReactionAdd(in.config.DsChannel, mes3s, emMinus)
					go dsDeleteMesageMinuts(in.config.DsChannel, mes3s, 3)
				} else if t.tip == "tg" {
					mes3s := tgSendEmbedTime(in.config.TgChannel, text)
					go tgDeleteMesageMinuts(in.config.TgChannel, mes3s, 3)
				}
			} else if t.timedown == 0 {
				in.RsMinus()
				//RsMinus(in, t.lvlkz)
			} else if t.timedown <= -1 {
				in.RsMinus()
				//RsMinus(in, t.lvlkz)
			}

		}
	}
	oneMinutsTimer()
}

func oneMinutsTimer() {
	var count int //количество активных игроков
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE active = 0")
	err := row.Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
	}
	if count > 0 {
		a := []string{}
		aa := []string{}
		results, err := db.Query("SELECT corpname FROM sborkz WHERE active = 0")
		if err != nil {
			fmt.Println(err)
		}
		var corpname string // ищим корпорации
		for results.Next() {
			err = results.Scan(&corpname)
			a = append(a, corpname)
		}
		a = removeDuplicateElementString(a)

		for _, corp := range a {
			skip := false
			for _, u := range aa {
				if corp == u {
					skip = true
					break
				}
			}
			if !skip {
				messageupdate(corp)
			}
		}
	}
}

func messageupdate(corpname string) {
	ok, config := checkCorpNameConfig(corpname)
	if ok {
		var count109 int
		ds := []string{}
		tg := []int{}
		wa := []string{}
		row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE corpname = ? AND active = 0", corpname)
		err := row.Scan(&count109)
		if err != nil {
			fmt.Println(err.Error())
		}
		if count109 > 0 {
			results, err := db.Query("SELECT * FROM sborkz WHERE corpname = ? AND active = 0", corpname)
			if err != nil {
				fmt.Println(err)
			}
			for results.Next() {
				var t sborkz
				err = results.Scan(&t.id, &t.corpname, &t.name, &t.mention, &t.tip, &t.dsmesid, &t.tgmesid, &t.wamesid, &t.time, &t.date, &t.lvlkz, &t.numkzn, &t.numberkz, &t.numberevent, &t.eventpoints, &t.active, &t.timedown)

				if config.DsChannel != "" {
					ds = append(ds, t.dsmesid)
				}
				if config.TgChannel != 0 {
					tg = append(tg, t.tgmesid)
				}
				if config.WaChannel != "" {
					wa = append(wa, t.wamesid)
				}
			}
		}
		ds = removeDuplicateElementString(ds)
		tg = removeDuplicateElementInt(tg)
		wa = removeDuplicateElementString(wa)

		if config.DsChannel != "" {
			messageupdateDS(ds, config)
		}
		if config.TgChannel != 0 {
			messageupdateTG(tg, config)
		}
		if config.WaChannel != "" {

		}
	}
}

func messageupdateDS(ds []string, config BotConfig) {
	var aa []string
	for _, dsmesid := range ds {
		skip := false
		for _, u := range aa {
			if dsmesid == u {
				skip = true
				break
			}
		}

		if !skip {
			results, err := db.Query("SELECT * FROM sborkz WHERE dsmesid = ? AND active = 0", dsmesid)
			if err != nil {
				fmt.Println(err)
			}
			var t sborkz
			for results.Next() {
				err = results.Scan(&t.id, &t.corpname, &t.name, &t.mention, &t.tip, &t.dsmesid, &t.tgmesid, &t.wamesid, &t.time, &t.date, &t.lvlkz, &t.numkzn, &t.numberkz, &t.numberevent, &t.eventpoints, &t.active, &t.timedown)
			}
			in := inMessage{
				tip:         "ds",
				name:        t.name,
				nameMention: t.mention,
				lvlkz:       t.lvlkz,
				Ds: Ds{
					mesid:   t.dsmesid,
					nameid:  "",
					guildid: config.Config.Guildid,
				},
				config: config,
				option: Option{
					callback: true,
					edit:     true,
					update:   false,
				},
			}
			in.Queue()
		}
	}
}

func messageupdateTG(tg []int, config BotConfig) {
	var aa []int
	for _, tgmesid := range tg {
		skip := false
		for _, u := range aa {
			if tgmesid == u {
				skip = true
				break
			}
		}

		if !skip {
			results, err := db.Query("SELECT * FROM sborkz WHERE tgmesid = ? AND active = 0", tgmesid)
			if err != nil {
				fmt.Println(err)
			}
			var t sborkz
			for results.Next() {
				err = results.Scan(&t.id, &t.corpname, &t.name, &t.mention, &t.tip, &t.dsmesid, &t.tgmesid, &t.wamesid, &t.time, &t.date, &t.lvlkz, &t.numkzn, &t.numberkz, &t.numberevent, &t.eventpoints, &t.active, &t.timedown)
			}
			in := inMessage{
				tip:         "tg",
				name:        t.name,
				nameMention: t.mention,
				lvlkz:       t.lvlkz,
				Tg:          Tg{mesid: t.tgmesid},
				config:      config,
				option: Option{
					callback: true,
					edit:     true,
					update:   false,
				},
			}
			in.Queue()
		}
	}
}
