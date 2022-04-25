package old

/*
import (
	"fmt"
	"rs_bot/bot/NewBot"
)

//rs start
func RsStart(in NewBot.inMessage, lvlkz string) {
	if !in.option.callback {
		if in.tip == "ds" {
			NewBot.dsDeleteMesage5s(in.config.DsChannel, in.Ds.mesid)
		} else if in.tip == "tg" {
			NewBot.tgDelMessage10s(in.config.TgChannel, in.Tg.mesid)
		}
	}
	countName := NewBot.countName(in, lvlkz)
	if countName == 0 {
		if in.tip == "ds" {
			NewBot.dsSendChannelDel5s(in.config.DsChannel, "Принудительный старт доступен участникам очереди.")
		} else if in.tip == "tg" {
			NewBot.tgSendChannelDel5s(in.config.TgChannel, "Принудительный старт доступен участникам очереди.")
		}
	} else if countName == 1 {
		numberkz := NewBot.numberQueueLvl(in, lvlkz)
		count := NewBot.countQueue(in, lvlkz)
		if count == 1 {
			countsS1(in, lvlkz, numberkz)
		} else if count == 2 {
			countsS2(in, lvlkz, numberkz)
		} else if count == 3 {
			countsS3(in, lvlkz, numberkz)
		}
	}
}

func countsS1(in NewBot.inMessage, lvlkz string, numkz int) {
	u := NewBot.readAll(in)

	textEvent, numkzEvent := NewBot.event(in)
	numberevent := NewBot.qweryNumevent1(in)
	if numberevent > 0 {
		numkz = numkzEvent
	}

	dsmesid := ""
	tgmesid := 0
	wamesid := ""
	if in.config.DsChannel != "" {
		if u.user1.tip == "ds" {
			NewBot.name1 = u.user1.mention
		} else {
			NewBot.name1 = u.user1.name
		}
		text := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n\n1. %s\nВ игру %s",
			lvlkz, numkz, NewBot.name1, textEvent)
		dsmesid = NewBot.dsSendChannel(in.config.DsChannel, text)
		go NewBot.dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
		NewBot.mesidDsUpdate(dsmesid, lvlkz, in.config.CorpName)
	}
	if in.config.TgChannel != 0 {
		if u.user1.tip == "tg" {
			NewBot.name1 = u.user1.mention
		} else {
			NewBot.name1 = u.user1.name
		}
		go NewBot.tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
		text := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n\n1. %s\nВ игру %s",
			lvlkz, numkz, NewBot.name1, textEvent)
		tgmesid := NewBot.tgSendChannel(in.config.TgChannel, text)
		NewBot.mesidTgUpdate(tgmesid, lvlkz, in.config.CorpName)
	}
	NewBot.updateComplite(lvlkz, dsmesid, tgmesid, wamesid, numkz, numberevent, in.config.CorpName)
}

func countsS2(in NewBot.inMessage, lvlkz string, numkz int) {
	u := NewBot.readAll(in)

	textEvent, numkzEvent := NewBot.event(in)
	numberevent := NewBot.qweryNumevent1(in)
	if numberevent > 0 {
		numkz = numkzEvent
	}
	dsmesid := ""
	tgmesid := 0
	wamesid := ""

	if in.config.DsChannel != "" { //discord
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
		text1 := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n", lvlkz, numkz)
		text2 := fmt.Sprintf("\n%s %s\nВ игру %s", NewBot.name1, NewBot.name2, textEvent)
		text := text1 + text2
		dsmesid = NewBot.dsSendChannel(in.config.DsChannel, text)
		go NewBot.dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
		NewBot.mesidDsUpdate(dsmesid, lvlkz, in.config.CorpName)
	}
	if in.config.TgChannel != 0 { //telegram
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
		go NewBot.tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
		text1 := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n", lvlkz, numkz)
		text2 := fmt.Sprintf("\n%s %s\nВ игру %s", NewBot.name1, NewBot.name2, textEvent)
		tgmesid = NewBot.tgSendChannel(in.config.TgChannel, text1+text2)
		NewBot.mesidTgUpdate(tgmesid, lvlkz, in.config.CorpName)
	}
	NewBot.updateComplite(lvlkz, dsmesid, tgmesid, wamesid, numkz, numberevent, in.config.CorpName)

}

func countsS3(in NewBot.inMessage, lvlkz string, numkz int) {
	u := NewBot.readAll(in)

	textEvent, numkzEvent := NewBot.event(in)
	numberevent := NewBot.qweryNumevent1(in)
	if numberevent > 0 {
		numkz = numkzEvent
	}
	dsmesid := ""
	tgmesid := 0
	wamesid := ""

	if in.config.DsChannel != "" { //discord
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
		text := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n\n%s %s %s\nВ игру %s",
			lvlkz, numkz, NewBot.name1, NewBot.name2, NewBot.name3, textEvent)
		dsmesid = NewBot.dsSendChannel(in.config.DsChannel, text)
		go NewBot.dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
		NewBot.mesidDsUpdate(dsmesid, lvlkz, in.config.CorpName)
	}
	if in.config.TgChannel != 0 { //telegram
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
		go NewBot.tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
		text := fmt.Sprintf("Очередь кз%s (%d) была \nзапущена не полной \n\n%s %s %s\nВ игру %s",
			lvlkz, numkz, NewBot.name1, NewBot.name2, NewBot.name3, textEvent)
		tgmesid = NewBot.tgSendChannel(in.config.TgChannel, text)
		NewBot.mesidTgUpdate(tgmesid, lvlkz, in.config.CorpName)
	}
	NewBot.updateComplite(lvlkz, dsmesid, tgmesid, wamesid, numkz, numberevent, in.config.CorpName)
}
*/
