package botDiscord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
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
	DSBot, err = discordgo.New("Bot " + os.Getenv("TokenD"))
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
	fmt.Println(len(m.Message.Content))
	if m.Author.ID == s.State.User.ID ||m.Author.Bot || len(m.Message.Content)<1 {
		return
	}
	if m.ChannelID == "909527364730490890" {
		logicRS(s, m)
	}

	if m.Content=="A" {
		//Subscribe(m.GuildID, "5", m.Message.Author.ID, m.Message.ChannelID)
	}

}
func roleToIdPing(rolePing, guildid string) string {
	//var pingId string          //создаю переменную
	rolPing := "кз" + rolePing // добавляю буквы
	g,err:=DSBot.State.Guild(guildid)
	if err !=nil{fmt.Println(err)}
	exist,role:=roleExists(g,rolPing)
	if !exist{
		//создаем роль и возврашаем пинг
		newRole,err:=DSBot.GuildRoleCreate(guildid)
		if err!=nil{fmt.Println(err)}
		role,err=DSBot.GuildRoleEdit(guildid,newRole.ID,rolPing,newRole.Color,newRole.Hoist,37080064,true)
		if err!=nil {
			fmt.Println(err)
			err = DSBot.GuildRoleDelete(guildid, newRole.ID)
			if err!=nil{fmt.Println(err)}
		}
		return role.Mention()
	}else {
		return role.Mention()
	}



	r, err := DSBot.GuildRoles(guildid)
	if err != nil {		fmt.Println(err)}
	l := len(r) // количество ролей на сервере
	i := 0
	for i < l { //ищу роли в цикле
		if r[i].Name == rolPing {
			//pingId = r[i].ID
			return r[i].Mention()
			//return "<@&" + pingId + ">" // возвращаю пинг роли
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
	lvlk := roleToIdPing(lvlkz, guildid)
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
