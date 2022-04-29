package NewBot

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

//проверяем всех игроков этой очереди на присутствие в других очередях или корпорациях
func elseChat(u Users, name4 string) {
	if countNameQueue(u.user1.name) > 0 {
		elsetrue(u.user1.name)
	}
	if countNameQueue(u.user2.name) > 0 {
		elsetrue(u.user2.name)
	}
	if countNameQueue(u.user3.name) > 0 {
		elsetrue(u.user3.name)
	}
	if countNameQueue(name4) > 0 {
		elsetrue(name4)
	}
}

//проверяем есть ли игрок в других очередях
func countNameQueue(name string) (countNames int) {
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND active = 0", name)
	err := row.Scan(&countNames)
	if err != nil {
		logrus.Println(err)
	}
	return countNames
}

//удаляем игрока с очереди
func elsetrue(name string) {
	results, err := db.Query("SELECT * FROM sborkz WHERE name = ? AND active = 0", name)
	if err != nil {
		fmt.Println(err)
	}
	for results.Next() {
		var t sborkz
		err = results.Scan(&t.id, &t.corpname, &t.name, &t.mention, &t.tip, &t.dsmesid, &t.tgmesid, &t.wamesid, &t.time, &t.date, &t.lvlkz, &t.numkzn, &t.numberkz, &t.numberevent, &t.eventpoints, &t.active, &t.timedown)

		ok, config := checkCorpNameConfig(t.corpname)
		if ok {
			in := inMessage{
				tip:         t.tip,
				name:        t.name,
				nameMention: t.mention,
				lvlkz:       t.lvlkz,
				timekz:      string(t.timedown),
				Ds: Ds{
					mesid:       t.dsmesid,
					nameid:      "",
					guildid:     "",
					Attachments: nil,
				},
				Tg:     Tg{},
				config: config,
				option: Option{
					callback: false,
					edit:     true,
					update:   false,
				},
			}
			in.RsMinus()
		}

	}
}
