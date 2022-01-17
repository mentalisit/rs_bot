package botDiscord

import "fmt"

func roleAdd(nameid, guildid, roleKzInt string) {
	roleid := roleKzToId(roleKzInt, guildid)

	DSBot.GuildMemberRoleAdd(guildid, nameid, roleid)
}
func roleKzToId(rolePing, guildid string) string {
	var roleId string          //создаю переменную
	rolPing := "кз" + rolePing // добавляю буквы
	r, err := DSBot.GuildRoles(guildid)
	if err != nil {
		fmt.Println(err)
	}
	l := len(r) // количество ролей на сервере
	i := 0
	for i < l { //ищу роли в цикле
		if r[i].Name == rolPing {
			roleId = r[i].ID
			return roleId // возвращаю ид роли
		} else {
			i = i + 1 // продолжаю перебор
		}
	}
	return "(роль не найдена)" // если не нашол нужной роли
}
