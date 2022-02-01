package botDiscord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"rs_bot/bot/botDiscord/databaseMysqlDs"
	"strings"
)

var p = New()

type Configs struct {
	DelMesComplite bool
	Primer         string
}
type Channel struct {
	Type      int
	CorpName  string
	DsChannel string
	TgChannel int64

	Config *Configs
}

func New() *Proxies {
	var arr Proxies
	return &arr
}

type Proxies []Channel

func addCorp(CorpName string, DsChannel string, TgChannel int64, DelMesComplite bool) {
	corpConfig := Channel{
		Type:      0xff,
		CorpName:  CorpName,
		DsChannel: DsChannel,
		TgChannel: TgChannel,
		Config: &Configs{
			DelMesComplite: DelMesComplite,
			Primer:         "",
		},
	}
	*p = append(*p, corpConfig)
}

func checkChannelConfig(chatid string) (channelGood bool, config Channel) {
	//channelGood := false
	//var config Channel
	for _, pp := range *p {
		if chatid == pp.DsChannel {
			channelGood = true
			config = pp
			break
		}
	}
	return channelGood, config
}

func readChannelConfig() { // чтение с бд и добавление в масив
	db, er := databaseMysqlDs.DbConnection()
	if er != nil {
		log.Println(er)
	}
	results, err := db.Query("SELECT channel FROM channel")
	if err != nil {
		log.Println(err)
	}
	var channel string
	for results.Next() {
		err = results.Scan(&channel)
		addCorp("", channel, 0, true)
	}
	db.Close()
}

func accesChat(m *discordgo.MessageCreate) {
	res := strings.HasPrefix(m.Content, ".")
	if res == true && m.Content == ".add" {
		go Delete5s(m.ChannelID, m.ID)
		accessAddChannel(m.ChannelID)
	} else if res == true && m.Content == ".del" {
		go Delete5s(m.ChannelID, m.ID)
		accessDelChannel(m.ChannelID)
	}
}
func accessAddChannel(chatid string) { // внесение в дб и добавление в масив
	ok, _ := checkChannelConfig(chatid)
	if ok {
		mes := SendChannel(chatid, "Я уже могу работать на вашем канале\n"+
			"повторная активация не требуется.\nнапиши Справка")
		go Delete1m(chatid, mes)
	} else {
		db, er := databaseMysqlDs.DbConnection()
		if er != nil {
			log.Println(er)
		}
		insertChannel := `INSERT INTO channel(channel) VALUES (?)`
		statement, err := db.Prepare(insertChannel)
		if err != nil {
			fmt.Println(err)
		}
		_, err = statement.Exec(chatid)
		if err != nil {
			fmt.Println(err.Error())
		}
		db.Close()
		addCorp("", chatid, 0, true)
		mes := SendChannel(chatid, "Спасибо за активацию.\nпиши Справка")
		go Delete1m(chatid, mes)
	}
}
func accessDelChannel(chatid string) { //удаление с бд и масива для блокировки
	ok, _ := checkChannelConfig(chatid)
	if !ok {
		mes := SendChannel(chatid, "ваш канал и так не подключен к логике бота ")
		go Delete1m(chatid, mes)
	} else {
		db, er := databaseMysqlDs.DbConnection()
		if er != nil {
			log.Println(er)
		}
		_, err := db.Exec("delete from channel where channel = ? ", chatid)
		if err != nil {
			log.Println(err)
		}
		db.Close()
		*p = *New()
		readChannelConfig()

		mes := SendChannel(chatid, "вы отключили мои возможности")
		go Delete1m(chatid, mes)
	}
}
