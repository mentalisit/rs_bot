package old

/*
import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"rs_bot/bot/NewBot"
)

func Queue(in NewBot.inMessage, lvlkz string) {
	if in.tip == "ds" && !in.option.callback {
		go NewBot.dsDelMessage(in.config.DsChannel, in.Ds.mesid)
	} else if in.tip == "tg" && !in.option.callback {
		go NewBot.tgDelMessage(in.config.TgChannel, in.Tg.mesid)
	}
	count := NewBot.countQueue(in, lvlkz)
	numberLvl := NewBot.numberQueueLvl(in, lvlkz) + 1
	// совподения количество  условие
	if count == 0 {
		text := "Очередь КЗ " + lvlkz + " пуста "
		if in.tip == "ds" {
			go NewBot.dsSendChannelDel5s(in.config.DsChannel, text)
		} else if in.tip == "tg" {
			go NewBot.tgSendChannelDel5s(in.config.TgChannel, text)
		}
	} else if count == 1 {
		count1Queue(in, lvlkz, numberLvl)
	} else if count == 2 {
		count2Queue(in, lvlkz, numberLvl)
	} else if count == 3 {
		count3Queue(in, lvlkz, numberLvl)
	}
}

func count1Queue(in NewBot.inMessage, lvlkz string, numberLvlkz int) {
	u := NewBot.readAll(in)
	if in.config.DsChannel != "" {
		NewBot.name1 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user1.name, u.user1.timedown, u.user1.numkzn)
		NewBot.name2 = ""
		NewBot.name3 = ""
		NewBot.name4 = ""
		lvlk := NewBot.roleToIdPing(lvlkz, in.config.Config.Guildid)
		NewBot.EmbedDS(NewBot.name1, NewBot.name2, NewBot.name3, NewBot.name4, lvlk, numberLvlkz)
		if in.option.edit {
			NewBot.dsEditComplex(u.user1.dsmesid, in.config.DsChannel)
		} else if !in.option.edit {
			NewBot.DSBot.ChannelMessageDelete(in.config.DsChannel, u.user1.dsmesid)
			mesCompl, err := NewBot.DSBot.ChannelMessageSendComplex(in.config.DsChannel, &discordgo.MessageSend{
				Content: NewBot.mesContentNil,
				Embed:   NewBot.Embeds})
			if err != nil {
				fmt.Println(err)
			}
			NewBot.addEnojiRsQueue(in.config.DsChannel, mesCompl.ID)
			NewBot.mesidDsUpdate(mesCompl.ID, lvlkz, in.config.CorpName)
		}
	}
	if in.config.TgChannel != 0 {
		text1 := fmt.Sprintf("Очередь кз%s (%d)\n", lvlkz, numberLvlkz)
		NewBot.name1 = fmt.Sprintf("1. %s - %dмин. (%d) \n", u.user1.name, u.user1.timedown, u.user1.numkzn)
		text2 := fmt.Sprintf("\n%s++ - принудительный старт", lvlkz)
		text := fmt.Sprintf("%s %s %s", text1, NewBot.name1, text2)
		if in.option.edit {
			NewBot.tgEditMessageText(in.config.TgChannel, u.user1.tgmesid, text, lvlkz)
		} else if !in.option.edit {
			mesidTg := NewBot.tgSendEmded(lvlkz, in.config.TgChannel, text)
			NewBot.mesidTgUpdate(mesidTg, lvlkz, in.config.CorpName)
			go NewBot.tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
		}
	}
	if in.config.WaChannel != "" {

	}
}

func count2Queue(in NewBot.inMessage, lvlkz string, numberLvlkz int) {
	u := NewBot.readAll(in)

	if in.config.DsChannel != "" {
		NewBot.name1 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user1.name, u.user1.timedown, u.user1.numkzn)
		NewBot.name2 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user2.name, u.user2.timedown, u.user2.numkzn)
		NewBot.name3 = ""
		NewBot.name4 = ""
		lvlk := NewBot.roleToIdPing(lvlkz, in.config.Config.Guildid)
		NewBot.EmbedDS(NewBot.name1, NewBot.name2, NewBot.name3, NewBot.name4, lvlk, numberLvlkz)
		if in.option.edit {
			NewBot.dsEditComplex(u.user1.dsmesid, in.config.DsChannel)
		} else if !in.option.edit {
			NewBot.dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
			mesCompl, err := NewBot.DSBot.ChannelMessageSendComplex(in.config.DsChannel, &discordgo.MessageSend{
				Content: NewBot.mesContentNil,
				Embed:   NewBot.Embeds})
			if err != nil {
				fmt.Println(err)
			}
			NewBot.addEnojiRsQueue(in.config.DsChannel, mesCompl.ID)
			NewBot.mesidDsUpdate(mesCompl.ID, lvlkz, in.config.CorpName)
		}
	}
	if in.config.TgChannel != 0 {
		text1 := fmt.Sprintf("Очередь кз%s (%d)\n", lvlkz, numberLvlkz)
		NewBot.name1 = fmt.Sprintf("1. %s - %dмин. (%d) \n", u.user1.name, u.user1.timedown, u.user1.numkzn)
		NewBot.name2 = fmt.Sprintf("2. %s - %dмин. (%d) \n", u.user2.name, u.user2.timedown, u.user2.numkzn)
		text2 := fmt.Sprintf("\n%s++ - принудительный старт", lvlkz)
		text := fmt.Sprintf("%s %s %s %s", text1, NewBot.name1, NewBot.name2, text2)
		if in.option.edit {
			NewBot.tgEditMessageText(in.config.TgChannel, u.user1.tgmesid, text, lvlkz)
		} else if !in.option.edit {
			mesidTg := NewBot.tgSendEmded(lvlkz, in.config.TgChannel, text)
			NewBot.mesidTgUpdate(mesidTg, lvlkz, in.config.CorpName)
			go NewBot.tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
		}
	}
	if in.config.WaChannel != "" {

	}
}

func count3Queue(in NewBot.inMessage, lvlkz string, numberLvlkz int) {
	u := NewBot.readAll(in)

	if in.config.DsChannel != "" {
		NewBot.name1 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user1.name, u.user1.timedown, u.user1.numkzn)
		NewBot.name2 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user2.name, u.user2.timedown, u.user2.numkzn)
		NewBot.name3 = fmt.Sprintf("%s  🕒  %d  (%d)", u.user3.name, u.user3.timedown, u.user3.numkzn)
		NewBot.name4 = ""
		lvlk := NewBot.roleToIdPing(lvlkz, in.config.Config.Guildid)
		NewBot.EmbedDS(NewBot.name1, NewBot.name2, NewBot.name3, NewBot.name4, lvlk, numberLvlkz)
		if in.option.edit {
			NewBot.dsEditComplex(u.user1.dsmesid, in.config.DsChannel)
		} else if !in.option.edit {
			NewBot.dsDelMessage(in.config.DsChannel, u.user1.dsmesid)
			mesCompl, err := NewBot.DSBot.ChannelMessageSendComplex(in.config.DsChannel, &discordgo.MessageSend{
				Content: NewBot.mesContentNil,
				Embed:   NewBot.Embeds})
			if err != nil {
				fmt.Println(err)
			}
			NewBot.addEnojiRsQueue(in.config.DsChannel, mesCompl.ID)
			NewBot.mesidDsUpdate(mesCompl.ID, lvlkz, in.config.CorpName)
		}
	}
	if in.config.TgChannel != 0 {
		text1 := fmt.Sprintf("Очередь кз%s (%d)\n", lvlkz, numberLvlkz)
		NewBot.name1 = fmt.Sprintf("1. %s - %dмин. (%d) \n", u.user1.name, u.user1.timedown, u.user1.numkzn)
		NewBot.name2 = fmt.Sprintf("2. %s - %dмин. (%d) \n", u.user2.name, u.user2.timedown, u.user2.numkzn)
		NewBot.name3 = fmt.Sprintf("3. %s - %dмин. (%d) \n", u.user3.name, u.user3.timedown, u.user3.numkzn)
		text2 := fmt.Sprintf("\n%s++ - принудительный старт", lvlkz)
		text := fmt.Sprintf("%s %s %s %s %s", text1, NewBot.name1, NewBot.name2, NewBot.name3, text2)
		if in.option.edit {
			NewBot.tgEditMessageText(in.config.TgChannel, u.user1.tgmesid, text, lvlkz)
		} else if !in.option.edit {
			mesidTg := NewBot.tgSendEmded(lvlkz, in.config.TgChannel, text)
			NewBot.mesidTgUpdate(mesidTg, lvlkz, in.config.CorpName)
			go NewBot.tgDelMessage(in.config.TgChannel, u.user1.tgmesid)
		}
	}
	if in.config.WaChannel != "" {

	}

}


*/
