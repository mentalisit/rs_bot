package NewBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

const (
	emOK      = "✅"
	emCancel  = "❎"
	emRsStart = "🚀"
	emPl30    = "⌛"
	emPlus    = "➕"
	emMinus   = "➖"
)

var mesContentNil string

func dsEditComplex(dsmesid, dschatid string) {
	a := &discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      dsmesid,
		Channel: dschatid,
	}
	_, err := DSBot.ChannelMessageEditComplex(a)
	if err != nil {
		fmt.Println(err)
	}
}

func MessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	message, err := DSBot.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		log.Println(err)
	}
	if message.Author.ID == s.State.User.ID {
		readReactionQueue(r, message)
	} else {
		log.Println("message", message)
		log.Println("err", err)
	}
}

func readReactionQueue(r *discordgo.MessageReactionAdd, message *discordgo.Message) {
	user, err := DSBot.User(r.UserID)
	if err != nil {
		fmt.Println(err)
	}
	if user.ID != message.Author.ID {
		ok, config := checkChannelConfigDS(r.ChannelID)
		if ok {
			member, e := DSBot.GuildMember(config.Config.Guildid, user.ID)
			if e != nil {
				fmt.Println("ошибка в функдиск стр57", e)
			}
			name := user.Username
			if member.Nick != "" {
				name = member.Nick
			}

			in := inMessage{
				mtext:       "",
				tip:         "ds",
				name:        name,
				nameMention: user.Mention(),
				Ds: Ds{
					mesid:   r.MessageID,
					nameid:  user.ID,
					guildid: message.GuildID,
				},
				Tg: Tg{
					mesid:  0,
					nameid: 0,
				},
				config: config,
				option: Option{
					callback: true,
					edit:     true,
					update:   false,
				},
			}
			reactionUserRemove(r)
			if r.Emoji.Name == emPlus {
				if in.Plus() {
					dsDeleteMesage5s(in.config.DsChannel, in.Ds.mesid)
				}
			} else if r.Emoji.Name == emMinus {
				if in.Minus() {
					dsDelMessage(in.config.DsChannel, in.Ds.mesid)
				}
			} else if r.Emoji.Name == emOK || r.Emoji.Name == emCancel || r.Emoji.Name == emRsStart || r.Emoji.Name == emPl30 {
				in.lvlkz, err = readMesID(r.MessageID)
				if err == nil && in.lvlkz != "" {
					if r.Emoji.Name == emOK {
						in.timekz = "30"
						in.RsPlus()
					} else if r.Emoji.Name == emCancel {
						in.RsMinus()
					} else if r.Emoji.Name == emRsStart {
						in.RsStart()
					} else if r.Emoji.Name == emPl30 {
						in.Pl30()
					}
				}
			}
		}
	}
}

func reactionUserRemove(r *discordgo.MessageReactionAdd) {
	err := DSBot.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.Name, r.UserID)
	if err != nil {
		fmt.Println("Ошибка удаления эмоджи", err)
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	logicMixDiscord(m)

}

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

func cleanChat(m *discordgo.MessageCreate) {
	res := strings.HasPrefix(m.Content, ".")
	if res == false { //если нет префикса  то удалить через 3 минуты
		go dsDeleteMesageMinuts(m.ChannelID, m.ID, 3)
	}
	if len(m.Attachments) > 0 { //если что-то   то удалить через 3 минуты
		for _, attach := range m.Attachments {
			go dsDeleteMesageMinuts(m.ChannelID, attach.ID, 3)
		}
	}
}

func dsDeleteMesage5s(chatid, mesid string) {
	time.Sleep(5 * time.Second)
	DSBot.ChannelMessageDelete(chatid, mesid)
}
func roleToIdPing(rolePing, guildid string) string {
	//var pingId string          //создаю переменную
	rolPing := "кз" + rolePing // добавляю буквы
	g, err := DSBot.Guild(guildid)
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
func dsSendChannelDel5s(chatid, text string) {
	message, err := DSBot.ChannelMessageSend(chatid, text)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(5 * time.Second)
	DSBot.ChannelMessageDelete(chatid, message.ID)
}
func dsDelMessage(chatid, mesid string) {
	DSBot.ChannelMessageDelete(chatid, mesid)
}

func dsSendChannelDel1m(chatid, text string) {
	message, err := DSBot.ChannelMessageSend(chatid, text)
	if err != nil {
		fmt.Println(err)
	}
	dsDeleteMesageMinuts(chatid, message.ID, 1)
}

func dsSendChannel(chatid, text string) string {
	message, err := DSBot.ChannelMessageSend(chatid, text)
	if err != nil {
		fmt.Println(err)
	}
	return message.ID
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

/*
var (
	// See https://discordapp.com/developers/docs/reference#message-formatting.
	channelMentionRE = regexp.MustCompile("<#[0-9]+>")
	userMentionRE    = regexp.MustCompile("@[^@\n]{1,32}")
	emoteRE          = regexp.MustCompile(`<a?(:\w+:)\d+>`)
)

func enumerateUsernames1(s string) []string {
	onlySpace := true
	for _, r := range s {
		if !unicode.IsSpace(r) {
			onlySpace = false
			break
		}
	}
	if onlySpace {
		return nil
	}

	var username, endSpace string
	var usernames []string
	skippingSpace := true
	for _, r := range s {
		if unicode.IsSpace(r) {
			if !skippingSpace {
				usernames = append(usernames, username)
				skippingSpace = true
			}
			endSpace += string(r)
			username += string(r)
		} else {
			endSpace = ""
			username += string(r)
			skippingSpace = false
		}
	}
	if endSpace == "" {
		usernames = append(usernames, username)
	}
	return usernames
}


func replaceUserMentions1(text string) string {
	replaceUserMentionFunc := func(match string) string {
		var (
			err      error
			member   *discordgo.Member
			username string
		)
fmt.Println(349,match[1:])
		usernames := enumerateUsernames1(match[1:])
		for _, username = range usernames {
			member, err = getGuildMemberByNick1(username)
			if err == nil {
				break
			}
		}
		if member == nil {
			return match
		}
		return strings.Replace(match, "@"+username, member.User.Mention(), 1)
	}
	return userMentionRE.ReplaceAllStringFunc(text, replaceUserMentionFunc)
}
var nickMemberMap = make(map[string]*discordgo.Member)
func getGuildMemberByNick1(nick string) (*discordgo.Member, error) {
	if member, ok := nickMemberMap[nick]; ok {
		return member, nil
	}
	return nil, errors.New("Couldn't find guild member with nick " + nick) // This will most likely get ignored by the caller
}





/*
func (b *Bdiscord) replaceChannelMentions(text string) string {
	replaceChannelMentionFunc := func(match string) string {
		channelID := match[2 : len(match)-1]
		channelName := b.getChannelName(channelID)

		// If we don't have the channel refresh our list.
		if channelName == "" {
			var err error
			b.channels, err = b.c.GuildChannels(b.guildID)
			if err != nil {
				return "#unknownchannel"
			}
			channelName = b.getChannelName(channelID)
		}
		return "#" + channelName
	}
	return channelMentionRE.ReplaceAllStringFunc(text, replaceChannelMentionFunc)
}
*/
//var re = regexp.MustCompile(`(?m)<@!?\d+>`)
//var nre = regexp.MustCompile(`\d+`)
/*
func  replaceUserMentions(text,guildid string) string {
	replaceUserMentionFunc := func(match string) string {
		var (username string)

		member:=repp(text,guildid)


		if member == nil {
			return match
		}
		username=member.User.Username
		return strings.Replace(text,member.User.Mention(),username,  1)
	}
	return userMentionRE.ReplaceAllStringFunc(text, replaceUserMentionFunc)
}

func repp(text,guildid string)*discordgo.Member{

	var member *discordgo.Member
	re.ReplaceAllStringFunc(text, func(m string)string  {
		id := nre.FindString(m)
		member, err = DSBot.State.Member(guildid, id)
		if err != nil {
			member, err = DSBot.GuildMember(guildid, id)
			if err != nil {
				return "unknown#0000" // значение когда не можем получить пользователя
			}
		}
		return ""
	})
	return member
}
/*
func replaceEmotes(text string) string {
	return emoteRE.ReplaceAllString(text, "$1")
}

func (b *Bdiscord) replaceAction(text string) (string, bool) {
	length := len(text)
	if length > 1 && text[0] == '_' && text[length-1] == '_' {
		return text[1 : length-1], true
	}
	return text, false
}

// splitURL splits a webhookURL and returns the ID and token.
func (b *Bdiscord) splitURL(url string) (string, string, bool) {
	const (
		expectedWebhookSplitCount = 7
		webhookIdxID              = 5
		webhookIdxToken           = 6
	)
	webhookURLSplit := strings.Split(url, "/")
	if len(webhookURLSplit) != expectedWebhookSplitCount {
		return "", "", false
	}
	return webhookURLSplit[webhookIdxID], webhookURLSplit[webhookIdxToken], true
}
/*
func (b *Bdiscord) getChannelName(id string) string {
	b.channelsMutex.RLock()
	defer b.channelsMutex.RUnlock()

	for _, c := range b.channelInfoMap {
		if c.Name == "ID:"+id {
			// if we have ID: specified in our gateway configuration return this
			return c.Name
		}
	}

	for _, channel := range b.channels {
		if channel.ID == id {
			return b.getCategoryChannelName(channel.Name, channel.ParentID)
		}
	}
	return ""
}



func enumerateUsernames(s string) []string {
	onlySpace := true
	for _, r := range s {
		if !unicode.IsSpace(r) {
			onlySpace = false
			break
		}
	}
	if onlySpace {
		return nil
	}

	var username, endSpace string
	var usernames []string
	skippingSpace := true
	for _, r := range s {
		if unicode.IsSpace(r) {
			if !skippingSpace {
				usernames = append(usernames, username)
				skippingSpace = true
			}
			endSpace += string(r)
			username += string(r)
		} else {
			endSpace = ""
			username += string(r)
			skippingSpace = false
		}
	}
	if endSpace == "" {
		usernames = append(usernames, username)
	}
	fmt.Println(448,usernames,endSpace)
	return usernames
}

func  getGuildMemberByNick(nick,guildid string) (*discordgo.Member, error) {
	//userMemberMap := map[string]*discordgo.Member
	fmt.Println(453,nick)
	//fmt.Println(DSBot.GuildMember(guildid,"582882137842122773"))
	 member,ok:=DSBot.GuildMember(guildid,nick)
	 fmt.Println(member,err)
	 if ok!=nil{
		return member,nil
	}
	fmt.Println(469,member.User.Username)



	//nickMemberMap := map[string]*discordgo.Member{}
	//if member, ok := nickMemberMap[nick]; ok {
	//	return member, nil
	//}
	return nil, errors.New("Couldn't find guild member with nick " + nick) // This will most likely get ignored by the caller
}
func replaceID(text,guild string )*discordgo.Member{
	var re = regexp.MustCompile(`(?m)<@!?\d+>`)
	var nre = regexp.MustCompile(`\d+`)
	var member *discordgo.Member

	re.ReplaceAllStringFunc(text, func(m string) string{
		id := nre.FindString(m)
		member, err = DSBot.State.Member(guild, id)
		if err != nil {
			member, err = DSBot.GuildMember(guild, id)
			if err != nil {
				return "unknown#0000" // значение когда не можем получить пользователя
			}
		}
		fmt.Println(id,member.Nick)
		return "@"+member.User.Username+member.User.Discriminator
	})
	return member
}

*/

func dsChatName(chatid, guildid string) string {
	g, err := DSBot.Guild(guildid)
	if err != nil {
		fmt.Println(err)
	}
	return g.Name
}
