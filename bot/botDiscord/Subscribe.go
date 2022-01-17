package botDiscord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)
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

func Subscribe(guildid string, argRoles string, userid,chatid string ){
	g,err:=DSBot.State.Guild(guildid)
	if err !=nil{fmt.Println(25,err)}

	exist,role:=roleExists(g,argRoles)

	if !exist { //если нет роли
		newRole,err:=DSBot.GuildRoleCreate(guildid)
		if err!=nil{fmt.Println(32,err)}
		role,err=DSBot.GuildRoleEdit(guildid,newRole.ID,argRoles,newRole.Color,newRole.Hoist,37080064,true)
			if err!=nil {
				fmt.Println(36,err)
				err = DSBot.GuildRoleDelete(guildid, newRole.ID)
				if err!=nil{fmt.Println(38,err)}
			}
		}

	member,err:=DSBot.GuildMember(guildid,userid)
	if err!=nil{fmt.Println(45,err)}
	var subscribe int=0
	if exist{
		for _,_role:=range member.Roles {
			if _role == role.ID {
				subscribe=1
			}
		}
	}

	err=DSBot.GuildMemberRoleAdd(guildid,userid,role.ID)
	if err!=nil{fmt.Println(57,err)
		subscribe=2
	}
	var mesid string
	if subscribe==0{
		mesid=SendChannel(chatid,fmt.Sprintf("%s Теперь вы подписаны на %s",member.Mention(),role.Name))
	}else if subscribe==1{
		mesid=SendChannel(chatid, fmt.Sprintf("%s Вы уже подписаны на %s",member.Mention(), role.Name))
	}else if subscribe==2{
		mesid=SendChannel(chatid,"ошибка: недостаточно прав для выдачи роли "+role.Name)
	}
	go Delete1m(chatid,mesid)
}

func Unsubscribe(guildid string, argRoles string, userid,chatid string ){
	var unsubscribe int=0

	g,err:=DSBot.State.Guild(guildid)
	if err !=nil{fmt.Println(err)}

	exist,role:=roleExists(g,argRoles)
	if !exist { //если нет роли
		unsubscribe=1
	}

	member,err:=DSBot.GuildMember(guildid,userid)
	if err!=nil{fmt.Println(err)}
	if exist{
		for _,_role:=range member.Roles {
			if _role == role.ID {
				unsubscribe=2
			}
		}
	}
	if unsubscribe==2{
		err=DSBot.GuildMemberRoleRemove(guildid,userid,role.ID)
			if err!=nil{
				fmt.Println(err)
				unsubscribe=3
			}
	}

	var mesid string
	if unsubscribe==0{
		mesid=SendChannel(chatid,fmt.Sprintf("%s Вы не подписаны на роль %s",member.Mention(),role.Name))
	}else if unsubscribe==1{
		mesid=SendChannel(chatid, fmt.Sprintf("%s Роли %s нет на сервере  ",member.Mention(), argRoles))
	}else if unsubscribe==2{
		mesid=SendChannel(chatid,fmt.Sprintf("%s Вы отписались от роли %s",member.Mention(),argRoles))
	}else if unsubscribe==3{
		mesid=SendChannel(chatid,"ошибка: недостаточно прав для снятия роли  "+role.Name)
	}
	go Delete1m(chatid,mesid)
	fmt.Println(104,unsubscribe)
}
