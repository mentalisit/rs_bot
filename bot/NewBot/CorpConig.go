package NewBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"rs_bot/bot/NewBot/database"
	"strings"
)

var P = New()

type Configs struct {
	DelMesComplite int
	mesidDsHelp    string
	mesidTgHelp    int
	Primer         string
	Guildid        string
}
type BotConfig struct {
	Type      int
	CorpName  string
	DsChannel string
	TgChannel int64
	WaChannel string
	Config    Configs
}

func New() *Proxies {
	var arr Proxies
	return &arr
}

type Proxies []BotConfig

func addCorp(CorpName string, DsChannel string, TgChannel int64, WaChannel string, DelMesComplite int, mesiddshelp string, mesidtghelp int, guildid string) {
	corpConfig := BotConfig{
		Type:      0xff,
		CorpName:  CorpName,
		DsChannel: DsChannel,
		TgChannel: TgChannel,
		WaChannel: WaChannel,
		Config: Configs{
			DelMesComplite: DelMesComplite,
			mesidDsHelp:    mesiddshelp,
			mesidTgHelp:    mesidtghelp,
			Primer:         "",
			Guildid:        guildid,
		},
	}
	*P = append(*P, corpConfig)
}

func checkChannelConfigDS(chatid string) (channelGood bool, config BotConfig) {
	if chatid != "" {
		for _, pp := range *P {
			if chatid == pp.DsChannel {
				channelGood = true
				config = pp
				break
			}
		}
	}
	return channelGood, config
}

func checkChannelConfigTG(chatid int64) (channelGood bool, config BotConfig) {
	if chatid != 0 {
		for _, pp := range *P {
			if chatid == pp.TgChannel {
				channelGood = true
				config = pp
				break
			}
		}
	}
	return channelGood, config
}

func checkCorpNameConfig(corpname string) (channelGood bool, config BotConfig) {
	if corpname != "" { // есть ли корпа
		for _, pp := range *P {
			if corpname == pp.CorpName {
				channelGood = true
				config = pp
				break
			}
		}
	}
	return channelGood, config
}

func readBotConfig() { // чтение с бд и добавление в масив
	db, err := database.DbConnection()
	if err != nil {
		log.Println(err)
	}
	results, err := db.Query("SELECT * FROM config")
	if err != nil {
		log.Println(err)
	}
	var t TableConfig
	for results.Next() {
		err = results.Scan(&t.id, &t.corpname, &t.dschannel, &t.tgchannel, &t.wachannel, &t.mesiddshelp, &t.mesidtghelp, &t.delmescomplite, &t.guildid)
		addCorp(t.corpname, t.dschannel, t.tgchannel, t.wachannel, t.delmescomplite, t.mesiddshelp, t.mesidtghelp, t.guildid)
	}
}

func accesChatDS(m *discordgo.MessageCreate) {
	res := strings.HasPrefix(m.Content, ".")
	if res == true && m.Content == ".add" {
		go dsDeleteMesage5s(m.ChannelID, m.ID)
		accessAddChannelDs(m.ChannelID, m.GuildID)
	} else if res == true && m.Content == ".del" {
		go dsDeleteMesage5s(m.ChannelID, m.ID)
		accessDelChannelDs(m.ChannelID)
	}
}
func accesChatTg(m *tgbotapi.Message) {
	res := strings.HasPrefix(m.Text, ".")
	if res == true && m.Text == ".add" {
		go tgDelMessage10s(m.Chat.ID, m.MessageID)
		accessAddChannelTg(m.Chat.ID)
	} else if res == true && m.Text == ".del" {
		go tgDelMessage10s(m.Chat.ID, m.MessageID)
		accessDelChannelTg(m.Chat.ID)
	}
}
func accessAddChannelDs(chatid, guildid string) { // внесение в дб и добавление в масив
	ok, _ := checkChannelConfigDS(chatid)
	if ok {
		go dsSendChannelDel1m(chatid, "Я уже могу работать на вашем канале\n"+
			"повторная активация не требуется.\nнапиши Справка1")
	} else {
		chatName := dsChatName(chatid, guildid)
		insertConfig := `INSERT INTO config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite,guildid) VALUES (?,?,?,?,?,?,?,?)`
		statement, err := db.Prepare(insertConfig)
		if err != nil {
			fmt.Println(err)
		}
		_, err = statement.Exec(chatName, chatid, 0, "", "", 0, 0, guildid)
		if err != nil {
			fmt.Println(err.Error())
		}
		//db.Close()
		addCorp(chatName, chatid, 0, "", 1, "", 0, guildid)
		go dsSendChannelDel1m(chatid, "Спасибо за активацию.\nпиши Справка1")
	}
}
func accessDelChannelDs(chatid string) { //удаление с бд и масива для блокировки
	ok, _ := checkChannelConfigDS(chatid)
	if !ok {
		go dsSendChannelDel1m(chatid, "ваш канал и так не подключен к логике бота ")
	} else {
		_, err := db.Exec("delete from config where dschannel = ? ", chatid)
		if err != nil {
			log.Println(err)
		}
		*P = *New()
		readBotConfig()
		go dsSendChannelDel1m(chatid, "вы отключили мои возможности")
	}
}

func accessAddChannelTg(chatid int64) { // внесение в дб и добавление в масив
	ok, _ := checkChannelConfigTG(chatid)
	if ok {
		go tgSendChannelDel1m(chatid, "Я уже могу работать на вашем канале\n"+
			"повторная активация не требуется.\nнапиши Справка1")
	} else {
		chatName := tgChatName(chatid)
		insertConfig := `INSERT INTO config (corpname,dschannel,tgchannel,wachannel,mesiddshelp,mesidtghelp,delmescomplite) VALUES (?,?,?,?,?,?,?)`
		statement, err := db.Prepare(insertConfig)
		if err != nil {
			fmt.Println(err)
		}
		_, err = statement.Exec(chatName, "", chatid, "", "", 0, 0)
		if err != nil {
			fmt.Println(err.Error())
		}
		addCorp(chatName, "", chatid, "", 1, "", 0, "")
		go tgSendChannelDel1m(chatid, "Спасибо за активацию.\nпиши Справка1")
	}
}
func accessDelChannelTg(chatid int64) { //удаление с бд и масива для блокировки
	ok, _ := checkChannelConfigTG(chatid)
	if !ok {
		go tgSendChannelDel1m(chatid, "ваш канал и так не подключен к логике бота ")
	} else {
		_, err := db.Exec("delete from config where tgchannel = ? ", chatid)
		if err != nil {
			log.Println(err)
		}
		*P = *New()
		readBotConfig()
		go tgSendChannelDel1m(chatid, "вы отключили мои возможности")
	}
}
