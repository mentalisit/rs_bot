package botDiscord

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"rs_bot/bot/botDiscord/databaseMysqlDs"
)

func oneMinutsTimer(db *sql.DB) {
	var count int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE active = 0")
	err := row.Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
	}
	if count > 0 {
		a := []string{}
		aa := []string{}
		results, err := db.Query("SELECT name,nameid,mesid,timedown FROM sborkz WHERE active = 0")
		if err != nil {
			fmt.Println(err)
		}
		var tagDS Sborkzds
		for results.Next() {
			err = results.Scan(&tagDS.Name, &tagDS.Nameid, &tagDS.Mesid, &tagDS.Timedown)
			a = append(a, tagDS.Mesid)
		}
		a = removeDuplicateElementString(a)
		for _, v := range a {
			skip := false
			for _, u := range aa {
				if v == u {
					skip = true
					break
				}
			}
			if !skip {
				messageupdate(db, v)
			}
		}
	}
}
func readALL(db *sql.DB, mesid string) Sborkzds {
	results, err := db.Query("SELECT * FROM sborkz WHERE  mesid = ? AND active = 0", mesid)
	if err != nil {
		log.Println(err)
	}
	var t Sborkzds
	for results.Next() {
		err = results.Scan(&t.Id, &t.Name, &t.Nameid, &t.Mention, &t.Guildid, &t.Lvlkz, &t.Chatid, &t.Mesid, &t.Timedown, &t.Active)
		//rs <- fmt.Sprintf("%s", t.Mention)
		//rst <- fmt.Sprintf("%d", t.Timedown)
	}
	return t
}

func messageupdate(db *sql.DB, mesid string) {
	var count45 int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE mesid = ? AND active = 0", mesid)
	err := row.Scan(&count45)
	if err != nil {
		fmt.Println(err.Error())
	}
	if count45 == 1 {
		t := readALL(db, mesid)
		counts1q(db, t)
	} else if count45 == 2 {
		t := readALL(db, mesid)
		counts2q(db, t)
	} else if count45 == 3 {
		t := readALL(db, mesid)
		counts3q(db, t)

	}
}

func counts1q(db *sql.DB, t Sborkzds) {
	readAll(db, t.Lvlkz, t.Chatid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = ""
	name3 = ""
	name4 = ""
	lvlk := roleToIdPing(t.Lvlkz, t.Guildid)
	numkz := readNumberkz(db, t.Lvlkz, t.Chatid)
	EmbedDS(name1, name2, name3, name4, lvlk, numkz)
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      t.Mesid,
		Channel: t.Chatid,
	})
}

func counts2q(db *sql.DB, t Sborkzds) {
	readAll(db, t.Lvlkz, t.Chatid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = <-rs + "  🕒  " + <-rst
	name3 = ""
	name4 = ""
	lvlk := roleToIdPing(t.Lvlkz, t.Guildid)
	numkz := readNumberkz(db, t.Lvlkz, t.Chatid)
	EmbedDS(name1, name2, name3, name4, lvlk, numkz)
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      t.Mesid,
		Channel: t.Chatid,
	})
}

func counts3q(db *sql.DB, t Sborkzds) {
	readAll(db, t.Lvlkz, t.Chatid)
	name1 = <-rs + "  🕒  " + <-rst
	name2 = <-rs + "  🕒  " + <-rst
	name3 = <-rs + "  🕒  " + <-rst
	name4 = ""
	lvlk := roleToIdPing(t.Lvlkz, t.Guildid)
	numkz := readNumberkz(db, t.Lvlkz, t.Chatid)
	EmbedDS(name1, name2, name3, name4, lvlk, numkz)
	DSBot.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Content: &mesContentNil,
		Embed:   Embeds,
		ID:      t.Mesid,
		Channel: t.Chatid,
	})
}
func MinusMin() {
	db, _ := databaseMysqlDs.DbConnection()
	_, err := db.Exec(`update sborkz set timedown = timedown - 1 where active = 0`)
	if err != nil {
		msqlTimeo(db)
		fmt.Println("errrrr")
	}
	msqlTimeo(db)
}
func msqlTimeo(db *sql.DB) {
	results, err := db.Query("SELECT * FROM sborkz WHERE active = 0")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var t Sborkzds
		err = results.Scan(&t.Id, &t.Name, &t.Nameid, &t.Mention, &t.Guildid, &t.Lvlkz, &t.Chatid, &t.Mesid, &t.Timedown, &t.Active)
		inmes := inMessage{
			mtext:       "",
			name:        t.Name,
			nameMention: t.Mention,
			nameid:      t.Nameid,
			mesid:       t.Mesid,
			guildid:     t.Guildid,
			chatid:      t.Chatid,
		}
		if t.Timedown == 3 {
			mes3s := SendChannel(t.Chatid, t.Mention+" время почти вышло  ...\n если ты еще тут пиши +")
			go Delete3m(t.Chatid, mes3s)
		} else if t.Timedown == 0 {
			RsMinus(db, t.Lvlkz, &inmes)
		} else if t.Timedown <= -1 {
			RsMinus(db, t.Lvlkz, &inmes)
		}
	}
	oneMinutsTimer(db)
}
