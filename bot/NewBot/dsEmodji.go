package NewBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func emodjis(in inMessage) {
	iftipdelete(in)
	e := emReadUsers(in.name)
	text := "	Для установки эмоджи пиши текст \n" +
		"Эмоджи пробел (номер ячейки1-4) пробел эмоджи \n" +
		"	пример \n" +
		"Эмоджи 1 🚀\n" +
		"	Ваши слоты" +
		"\n1" + e.em1 +
		"\n2" + e.em2 +
		"\n3" + e.em3 +
		"\n4" + e.em4
	dsSendChannelDel1m(in.config.DsChannel, text)
}
func emodjiadd(in inMessage, slot, emo string) {
	iftipdelete(in)
	if in.tip == "ds" {
		t := emReadUsers(in.name)
		if len(t.name) > 0 {
			emUpdateEmodji(in, slot, emo)
		} else {
			insert := `INSERT INTO users(name,em1,em2,em3,em4) VALUES (?,?,?,?,?)`
			_, err := db.Exec(insert, in.name, "", "", "", "")
			if err != nil {
				log.Println(err)
			}
			emUpdateEmodji(in, slot, emo)
		}
	}
}

//вносим емоджи в бд
func emUpdateEmodji(in inMessage, slot, emo string) {
	switch slot {
	case "1":
		_, err := db.Exec(`update users set em1 = ? where name = ?`, emo, in.name)
		if err != nil {
			fmt.Println("oooшибка1", err)
		}
		dsSendChannelDel1m(in.config.DsChannel, fmt.Sprintf("Слот %s обновлен\n%s", slot, emo))
	case "2":
		_, err := db.Exec(`update users set em2 = ? where name = ?`, emo, in.name)
		if err != nil {
			fmt.Println("oooшибка2", err)
		}
		dsSendChannelDel1m(in.config.DsChannel, fmt.Sprintf("Слот %s обновлен\n%s", slot, emo))
	case "3":
		_, err := db.Exec(`update users set em3 = ? where name = ?`, emo, in.name)
		if err != nil {
			fmt.Println("oooшибка3", err)
		}
		dsSendChannelDel1m(in.config.DsChannel, fmt.Sprintf("Слот %s обновлен\n%s", slot, emo))
	case "4":
		_, err := db.Exec(`update users set em4 = ? where name = ?`, emo, in.name)
		if err != nil {
			fmt.Println("oooшибка4", err)
		}
		dsSendChannelDel1m(in.config.DsChannel, fmt.Sprintf("Слот %s обновлен\n%s", slot, emo))
	}
}

// склеиваем имя и эмоджи
func emReadName(name string) string {
	t := emReadUsers(name)
	return fmt.Sprintf("%s %s%s%s%s", name, t.em1, t.em2, t.em3, t.em4)
}

//читаем структуру с бд
func emReadUsers(name string) emodjiUser {
	results, err := db.Query("SELECT * FROM users WHERE name = ?", name)
	if err != nil {
		log.Println(err)
	}
	var t emodjiUser
	for results.Next() {
		err = results.Scan(&t.id, &t.name, &t.em1, &t.em2, &t.em3, &t.em4)
	}
	return t
}

func readReactionEmodji(r *discordgo.MessageReactionAdd) {
	t := emReadUsers(r.Member.User.Username)
	if t.name == r.Member.User.Username {
		emodji := ""
		if r.Emoji.ID != "" {
			emodji = fmt.Sprintf("<:%s:%s>", r.Emoji.Name, r.Emoji.ID)
		} else {
			emodji = r.Emoji.Name
		}
		var em1, em2, em3, em4 string
		em1 = emodji
		fmt.Println("hhh", em1, em2, em3, em4)
	}
}
