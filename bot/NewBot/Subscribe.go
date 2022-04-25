package NewBot

import (
	"fmt"
	"log"
)

func Subscribe(in inMessage, lvlkz string, tipPing int) {
	if in.tip == "ds" {
		go dsDelMessage(in.config.DsChannel, in.Ds.mesid)
		argRoles := "кз" + lvlkz
		if tipPing == 3 {
			argRoles = "кз" + lvlkz + "+"
		}

		g, err := DSBot.State.Guild(in.Ds.guildid)
		if err != nil {
			fmt.Println(err)
		}

		exist, role := roleExists(g, argRoles)

		if !exist { //если нет роли
			newRole, err := DSBot.GuildRoleCreate(in.Ds.guildid)
			if err != nil {
				fmt.Println(err)
			}
			role, err = DSBot.GuildRoleEdit(in.Ds.guildid, newRole.ID, argRoles, newRole.Color, newRole.Hoist, 37080064, true)
			if err != nil {
				err = DSBot.GuildRoleDelete(in.Ds.guildid, newRole.ID)
				if err != nil {
					fmt.Println(err)
				}
			}
		}

		member, err := DSBot.GuildMember(in.Ds.guildid, in.Ds.nameid)
		if err != nil {
			fmt.Println(err)
		}
		var subscribe int = 0
		if exist {
			for _, _role := range member.Roles {
				if _role == role.ID {
					subscribe = 1
				}
			}
		}

		err = DSBot.GuildMemberRoleAdd(in.Ds.guildid, in.Ds.nameid, role.ID)
		if err != nil {
			fmt.Println(err)
			subscribe = 2
		}
		if subscribe == 0 {
			go dsSendChannelDel1m(in.config.DsChannel, fmt.Sprintf("%s Теперь вы подписаны на %s", member.Mention(), role.Name))
		} else if subscribe == 1 {
			go dsSendChannelDel1m(in.config.DsChannel, fmt.Sprintf("%s Вы уже подписаны на %s", member.Mention(), role.Name))
		} else if subscribe == 2 {
			go dsSendChannelDel1m(in.config.DsChannel, "ошибка: недостаточно прав для выдачи роли "+role.Name)
		}
	} else if in.tip == "tg" {
		go tgDelMessage(in.config.TgChannel, in.Tg.mesid)
		//проверка активной подписки
		var counts int
		row := db.QueryRow("SELECT  COUNT(*) as count FROM subscribe WHERE name = ? AND lvlkz = ? AND chatid = ? AND tip = ?", in.name, lvlkz, in.config.TgChannel, tipPing)
		errs := row.Scan(&counts)
		if errs != nil {
			log.Println(errs)
		}

		if counts == 1 {
			text := fmt.Sprintf("%s ты уже подписан на кз%s %d/4\n для добавления в очередь напиши %s+", in.nameMention, lvlkz, tipPing, lvlkz)
			go tgSendChannelDel5s(in.config.TgChannel, text)
		} else {
			//добавление в оочередь пинга
			insertSubscribe := `INSERT INTO subscribe (name, nameid, lvlkz, tip, chatid, timestart, timeend) VALUES (?,?,?,?,?,?,?)`
			statement, err := db.Prepare(insertSubscribe)
			_, err = statement.Exec(in.name, in.nameMention, lvlkz, tipPing, in.config.TgChannel, 0, 0)
			if err != nil {
				log.Println(err.Error())
			}
			if tipPing == 1 {
				go tgSendChannelDel5s(in.config.TgChannel, in.nameMention+" вы подписались на пинг кз"+lvlkz+" 1/4\n для добавления в очередь напиши "+lvlkz+"+")
			} else if tipPing == 3 {
				go tgSendChannelDel5s(in.config.TgChannel, in.nameMention+" вы подписались на пинг кз"+lvlkz+" 3/4\n для добавления в очередь напиши "+lvlkz+"+")
			}

		}
	}
}

func Unsubscribe(in inMessage, lvlkz string, tipPing int) {
	if in.tip == "ds" {
		go dsDelMessage(in.config.DsChannel, in.Ds.mesid)
		argRoles := "кз" + lvlkz
		if tipPing == 3 {
			argRoles = "кз" + lvlkz + "+"
		}
		var unsubscribe int = 0
		g, err := DSBot.State.Guild(in.Ds.guildid)
		if err != nil {
			fmt.Println(err)
		}

		exist, role := roleExists(g, argRoles)
		if !exist { //если нет роли
			unsubscribe = 1
		}

		member, err := DSBot.GuildMember(in.Ds.guildid, in.Ds.nameid)
		if err != nil {
			fmt.Println(err)
		}
		if exist {
			for _, _role := range member.Roles {
				if _role == role.ID {
					unsubscribe = 2
				}
			}
		}
		if unsubscribe == 2 {
			err = DSBot.GuildMemberRoleRemove(in.Ds.guildid, in.Ds.nameid, role.ID)
			if err != nil {
				fmt.Println(err)
				unsubscribe = 3
			}
		}
		if unsubscribe == 0 {
			go dsSendChannelDel1m(in.config.DsChannel, fmt.Sprintf("%s Вы не подписаны на роль %s", member.Mention(), role.Name))
		} else if unsubscribe == 1 {
			go dsSendChannelDel1m(in.config.DsChannel, fmt.Sprintf("%s Роли %s нет на сервере  ", member.Mention(), argRoles))
		} else if unsubscribe == 2 {
			go dsSendChannelDel1m(in.config.DsChannel, fmt.Sprintf("%s Вы отписались от роли %s", member.Mention(), argRoles))
		} else if unsubscribe == 3 {
			go dsSendChannelDel1m(in.config.DsChannel, "ошибка: недостаточно прав для снятия роли  "+role.Name)
		}
	} else if in.tip == "tg" {
		go tgDelMessage(in.config.TgChannel, in.Tg.mesid)
		//проверка активной подписи
		var counts int
		row := db.QueryRow("SELECT  COUNT(*) as count FROM subscribe WHERE name = ? AND lvlkz = ? AND chatid = ? AND tip = ?", in.name, lvlkz, in.config.TgChannel, tipPing)
		errs := row.Scan(&counts)
		if errs != nil {
			log.Println(errs.Error())
		}

		if counts == 0 {
			text := fmt.Sprintf("%s ты не подписан на пинг кз%s %d/4", in.nameMention, lvlkz, tipPing)
			go tgSendChannelDel5s(in.config.TgChannel, text)
		} else if counts == 1 {
			//удаление с базы данных
			_, err := db.Exec("delete from subscribe where name = ? AND lvlkz = ? AND chatid = ? AND tip = ?", in.name, lvlkz, in.config.TgChannel, tipPing)
			text := fmt.Sprintf("%s отписался от пинга кз%s %d/4", in.nameMention, lvlkz, tipPing)
			go tgSendChannelDel5s(in.config.TgChannel, text)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}

func SubscribePing(in inMessage, lvlkz string, tipPing int) {
	var name1, names, men string
	if rows, err := db.Query("SELECT nameid FROM subscribe WHERE lvlkz = ? AND chatid = ? AND tip = ?", lvlkz, in.config.TgChannel, tipPing); err == nil {
		for rows.Next() {
			rows.Scan(&name1)
			if in.nameMention == name1 {
				continue
			}
			names = name1 + " "
			men = names + men
		}
		rows.Close()
	}
	mes := tgSendChannel(in.config.TgChannel, men)
	go tgDeleteMesageMinuts(in.config.TgChannel, mes, 10)
}
func (in inMessage) SubscribePing(tipPing int) {

	if in.config.TgChannel != 0 {
		var name1, names, men string
		if rows, err := db.Query("SELECT nameid FROM subscribe WHERE lvlkz = ? AND chatid = ? AND tip = ?", in.lvlkz, in.config.TgChannel, tipPing); err == nil {
			for rows.Next() {
				rows.Scan(&name1)
				if in.nameMention == name1 {
					continue
				}
				names = name1 + " "
				men = names + men
			}
			rows.Close()
		}
		mes := tgSendChannel(in.config.TgChannel, men)
		go tgDeleteMesageMinuts(in.config.TgChannel, mes, 10)
	}
}
