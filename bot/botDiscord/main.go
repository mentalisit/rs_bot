package botDiscord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strings"
	"time"
)

var (
	BotId     string
	DSBot     *discordgo.Session
	err       error
	timekz    string
	lvlkz     string
	emOK      = "✅"
	emCancel  = "❎"
	emRsStart = "🚀"
	emPl30    = "⌛"
	emPlus    = "➕"
	emMinus   = "➖"
)

type DiscordBot struct {
	DSBot       *discordgo.Session
	MentionUser *discordgo.User
	MentionRole discordgo.Role
}

var (
	name1, name2, name3, name4 string
	time1, time2, time3, time4 string
)
var Embeds = &discordgo.MessageEmbed{}

func EmbedDS(name1, name2, name3, name4, lvlkz string, numkz int) {
	Embeds = &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  16711680,
		Description: fmt.Sprintf("Желающие:👇 |  <:rs:918545444425072671> на %s (%d) ", lvlkz, numkz) +
			fmt.Sprintf(
				"\n1️⃣ %s "+
					"\n2️⃣ %s "+
					"\n3️⃣ %s "+
					"\n4️⃣ %s "+
					"\n", name1, name2, name3, name4),

		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   fmt.Sprintf(" %s для добавления в очередь\n%s для выхода из очереди\n%s принудительный старт", emOK, emCancel, emRsStart),
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
	DSBot.AddHandler(MessageReactionAdd)
	err = DSBot.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Println("Бот DISCORD запущен!!!")
	readChannelConfig()
}

func MessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	message, err := DSBot.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		fmt.Println(err)
	}
	if message.Author.ID == s.State.User.ID {
		readReactionQueue(r, message)
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || len(m.Message.Content) < 1 {
		return
	}
	if len(m.Content) > 0 {
		ok, config := checkChannelConfig(m.ChannelID)
		if ok {
			fmt.Println(config)
			logicRS(s, m)
			cleanChat(m)
		}
		accesChat(m)
	}

	if m.ChannelID == "909527364730490890" {
		fmt.Println(m.Content)
		logicRS(s, m)
		cleanChat(m)
	}

}

func cleanChat(m *discordgo.MessageCreate) {
	res := strings.HasPrefix(m.Content, ".")
	if res == false {
		go Delete3m(m.ChannelID, m.ID)
	}
}

func roleToIdPing(rolePing, guildid string) string {
	//var pingId string          //создаю переменную
	rolPing := "кз" + rolePing // добавляю буквы
	g, err := DSBot.State.Guild(guildid)
	if err != nil {
		fmt.Println(err)
	}
	exist, role := roleExists(g, rolPing)
	if !exist {
		//создаем роль и возврашаем пинг
		newRole, err := DSBot.GuildRoleCreate(guildid)
		if err != nil {
			fmt.Println(err)
		}
		role, err = DSBot.GuildRoleEdit(guildid, newRole.ID, rolPing, newRole.Color, newRole.Hoist, 37080064, true)
		if err != nil {
			fmt.Println(err)
			err = DSBot.GuildRoleDelete(guildid, newRole.ID)
			if err != nil {
				fmt.Println(err)
			}
		}
		return role.Mention()
	} else {
		return role.Mention()
	}

	r, err := DSBot.GuildRoles(guildid)
	if err != nil {
		fmt.Println(err)
	}
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
	EmbedDS(name1, name2, name3, name4, lvlk, 0)
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

func removeDuplicateElementString(mesididid []string) []string {
	result := make([]string, 0, len(mesididid))
	temp := map[string]struct{}{}
	for _, item := range mesididid {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func addEnojiRsQueue(chatid, mesid string) {
	DSBot.MessageReactionAdd(chatid, mesid, emOK)
	DSBot.MessageReactionAdd(chatid, mesid, emCancel)
	DSBot.MessageReactionAdd(chatid, mesid, emRsStart)
	DSBot.MessageReactionAdd(chatid, mesid, emPl30)

}

//получаем есть ли роль и саму роль
func roleExists(g *discordgo.Guild, nameRoles string) (bool, *discordgo.Role) {
	nameRoles = strings.ToLower(nameRoles)

	for _, role := range g.Roles {
		if role.Name == "@everyone" {
			continue
		}
		if strings.ToLower(role.Name) == nameRoles {
			return true, role
		}
	}
	return false, nil
}

func checkAdmin(nameid string, chatid string) bool {
	perms, err := DSBot.UserChannelPermissions(nameid, chatid)
	if err != nil {
		fmt.Println(err)
	}
	if perms&discordgo.PermissionAdministrator != 0 {
		fmt.Println("админ")
		return true
	} else {
		fmt.Println("не админ")
		return false
	}
}
