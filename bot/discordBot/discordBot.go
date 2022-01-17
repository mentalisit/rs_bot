package discordBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"regexp"
	"rs_bot/config"
	"time"
)

var (
	BotId       string
	DSBot       *discordgo.Session
	err         error
	MentionUser *discordgo.User
	MentionRole discordgo.Role
	AddReact    *discordgo.MessageReactionAdd
	timekz      string
	lvlkz       string
)

type DiscordBot struct {
	DSBot       *discordgo.Session
	MentionUser *discordgo.User
	MentionRole discordgo.Role
	addReact    *discordgo.MessageReactionAdd
}

var (
	name1, name2, name3, name4 string
	time1, time2, time3, time4 string
)
var Embeds = &discordgo.MessageEmbed{}

func EmbedDS(name1, name2, name3, name4, lvlkz string) {
	Embeds = &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  16711680,
		Description: fmt.Sprintf("Желающие:👇 |  <:rs:918545444425072671> на %s ", lvlkz) +
			fmt.Sprintf(
				"\n1️⃣ %s "+
					"\n2️⃣ %s "+
					"\n3️⃣ %s "+
					"\n4️⃣ %s "+
					"\nпорядковый номер КЗ #: ", name1, name2, name3, name4),

		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "тут бедет что-то ",
				Value:  "Данные обновлены: ",
				Inline: true,
			}},
		Timestamp: time.Now().Format(time.RFC3339), // ТЕКУЩЕЕ ВРЕМЯ ДИСКОРДА
		Title:     "ОЧЕРЕДЬ КЗ  ",
	}
}

func Start() {
	DSBot, err = discordgo.New("Bot " + config.TokenD)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := DSBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	BotId = u.ID
	DSBot.AddHandler(messageHandler)
	err = DSBot.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Println("Бот DISCORD запущен!!!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	mtext := m.Content //текст сообщения
	//tm := time.Local
	//mtime := tm
	//nameid:="<@!"+m.Message.Author.ID+">" //582882137842122773 ping
	nameid := m.Message.Author.Mention()
	mesid := m.ID             // ид сообщения 911747673093197844
	name := m.Author.Username // имя Mentalisit
	guildid := m.GuildID      // id 700238199070523412
	chatid := m.ChannelID     //chat id 909527364730490890

	if m.Author.ID == s.State.User.ID {
		return
	}
	var kzb string
	var subs string
	var qwery string
	var rss string

	if len(m.Content) > 0 {
		str := mtext
		re := regexp.MustCompile(`^([4-9]|[1][0-1])([\+]|[-])(\d|\d{2})$`) //три переменные
		arr := (re.FindAllStringSubmatch(str, -1))
		if len(arr) > 0 {
			lvlkz = arr[0][1]
			kzb = arr[0][2]
			timekz = arr[0][3]
		}

		re2 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+]|[-])$`) // две переменные
		arr2 := (re2.FindAllStringSubmatch(str, -1))
		if len(arr2) > 0 {
			lvlkz = arr2[0][1]
			kzb = arr2[0][2]
			timekz = "30"
		}

		re3 := regexp.MustCompile(`^([\+]|[-])([4-9]|[1][0-1])$`) // две переменные для добавления или удаления подписок
		arr3 := (re3.FindAllStringSubmatch(str, -1))
		if len(arr3) > 0 {
			lvlkz = arr3[0][2]
			subs = arr3[0][1]
		}

		re4 := regexp.MustCompile(`^(["о"]|["О"])([4-9]|[1][0-1])$`) // две переменные для чтения  очереди
		arr4 := (re4.FindAllStringSubmatch(str, -1))
		if len(arr4) > 0 {
			qwery = arr4[0][1]
			lvlkz = arr4[0][2]
		}
		re5 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+][\+])$`) //rs start
		arr5 := (re5.FindAllStringSubmatch(str, -1))
		if len(arr5) > 0 {
			lvlkz = arr5[0][1]
			rss = arr5[0][2]
		}

		if kzb == "+" {
			DSBot.ChannelMessageDelete(chatid, mesid)
			QueueDs(lvlkz, timekz, mesid, name, nameid, guildid, chatid)
		} else if kzb == "-" {
			DSBot.ChannelMessageDelete(chatid, mesid)
			removeQueue(name, nameid, lvlkz, chatid, guildid)

		} else if subs == "+" {
			//Subscribe(name, nameidid, lvlkz, mesid, chatid)

		} else if subs == "-" {
			//Unsubscribe(name,nameidid,lvlkz,mesid,chatid)
		} else if len(qwery) > 0 {
			DSBot.ChannelMessageDelete(chatid, mesid)
			//MsqlRsQ(lvlkz,chatid)

		} else if len(rss) > 0 {
			DSBot.ChannelMessageDelete(chatid, mesid)
			//MsqlRsStart(lvlkz,name,chatid)
		} else if mtext == "Справка" {
			DSBot.ChannelMessageDelete(chatid, mesid)
			hhelp(name, chatid)
		} else if mtext == "1" {
			DSBot.ChannelMessageDelete(chatid, mesid)
			//embedtest(chatid,nameid,guildid)
			//roleAdd()
			deleteSrorkz("Mentalisit", "4", "909527364730490890")

		} else if mtext == "2" {
			DSBot.ChannelMessageDelete(chatid, mesid)
			roleRemove()
		}
		//go Delete5s(chatid,mesid)
	}

	//MsqlPR_DS(lvlkz,timekz,mesid,name,nameid,guildID,chatid)
	GuildState, err := DSBot.State.Guild(guildid)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("serverName: ", GuildState.Name, "Test DS String: "+m.Content)
}
func roleToId(rolePing, guildid string) string {

	//	fmt.Println("resuit.roles ",GuildState.Roles)
	//	fmt.Println("errr",err)
	//GuildState,err:=GoBot.State.Guild(guildid)
	var pingId string          //создаю переменную
	rolPing := "кз" + rolePing // добавляю буквы
	r, err := DSBot.GuildRoles(guildid)
	if err != nil {
		fmt.Println(err)
	}
	//r:=GuildState.Roles
	l := len(r) // количество ролей на сервере
	i := 0
	for i < l { //ищу роли в цикле
		if r[i].Name == rolPing {
			pingId = r[i].ID
			//fmt.Println(pingId)
			return "<@&" + pingId + ">" // возвращаю пинг роли
		} else {
			i = i + 1 // продолжаю перебор
		}
	}
	return "(роль не найдена)" // если не нашол нужной роли
}
func Delete5s(chatid, dMessageId string) {
	time.Sleep(5 * time.Second)
	DSBot.ChannelMessageDelete(chatid, dMessageId)
}
func Delete1m(chatid, dMessageId string) {
	time.Sleep(1 * time.Minute)
	DSBot.ChannelMessageDelete(chatid, dMessageId)
}
func Delete3m(chatid, dMessageId string) {
	time.Sleep(3 * time.Minute)
	DSBot.ChannelMessageDelete(chatid, dMessageId)
}
func SendChannel(chatid, text string) string {
	message, err := DSBot.ChannelMessageSend(chatid, text)
	if err != nil {
		fmt.Println(err)
	}
	return message.ID
}

func embedtest(chatid, nameid, guildid string) {
	lvlkz = "6"
	name1 = nameid + "  🕒 " + timekz
	name2 = ""
	name3 = ""
	name4 = ""
	lvlk := roleToId(lvlkz, guildid)
	//fmt.Println(lvlk)
	EmbedDS(name1, name2, name3, name4, lvlk)
	mes, err := DSBot.ChannelMessageSendComplex(chatid, &discordgo.MessageSend{
		Content: nameid + " запустил очередь " + lvlk})
	if err != nil {
		fmt.Println(err)
	}
	var mesContentNil string
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      mes.ID,
		Channel: chatid,
	})
}

func removeDuplicateElementString(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
