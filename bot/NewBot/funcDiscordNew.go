package NewBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"time"
)

func (Ds) DeleteMesage5s(chatid, mesid string) {
	time.Sleep(5 * time.Second)
	DSBot.ChannelMessageDelete(chatid, mesid)
}

func (Ds) SendChannelDel5s(chatid, text string) {
	message, err := DSBot.ChannelMessageSend(chatid, text)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(5 * time.Second)
	DSBot.ChannelMessageDelete(chatid, message.ID)
}

func (Ds) roleToIdPing(in inMessage) string {
	//var pingId string          //создаю переменную
	rolPing := "кз" + in.lvlkz // добавляю буквы
	g, err := DSBot.Guild(in.config.Config.Guildid)
	if err != nil {
		logrus.Println(err)
	}
	exist, role := roleExists(g, rolPing)
	if !exist {
		//создаем роль и возврашаем пинг
		newRole, err := DSBot.GuildRoleCreate(in.config.Config.Guildid)
		if err != nil {
			logrus.Println(err, "ошибка создания роли ")

		}
		role, err = DSBot.GuildRoleEdit(in.config.Config.Guildid, newRole.ID, rolPing, newRole.Color, newRole.Hoist, 37080064, true)
		if err != nil {
			logrus.Println(err, "Ошибка редактирования новой роли ")
			err = DSBot.GuildRoleDelete(in.config.Config.Guildid, newRole.ID)
			if err != nil {
				logrus.Println(err, "удаления новой роли")
			}
		}
		return role.Mention()
	} else {
		return role.Mention()
	}

	r, err := DSBot.GuildRoles(in.config.Config.Guildid)
	if err != nil {
		logrus.Println(err)
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

func (Ds) EmbedDS(name1, name2, name3, name4, lvlkz string, numkz int) {
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

func (Ds) EditComplex(dsmesid, dschatid string) {
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

func (Ds) addEnojiRsQueue(chatid, mesid string) {
	DSBot.MessageReactionAdd(chatid, mesid, emOK)
	DSBot.MessageReactionAdd(chatid, mesid, emCancel)
	DSBot.MessageReactionAdd(chatid, mesid, emRsStart)
	DSBot.MessageReactionAdd(chatid, mesid, emPl30)

}
