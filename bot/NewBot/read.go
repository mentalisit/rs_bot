package NewBot

import (
	"fmt"
	"log"
)

func readAll(in inMessage) (users Users) {
	u := Users{
		user1: sborkz{},
		user2: sborkz{},
		user3: sborkz{},
		user4: sborkz{},
	}
	user := 1
	results, err := db.Query("SELECT * FROM sborkz WHERE lvlkz = ? AND corpname = ? AND active = 0", in.lvlkz, in.config.CorpName)
	if err != nil {
		fmt.Println(err)
	}
	for results.Next() {
		var t sborkz
		err = results.Scan(&t.id, &t.corpname, &t.name, &t.mention, &t.tip, &t.dsmesid, &t.tgmesid, &t.wamesid, &t.time, &t.date, &t.lvlkz, &t.numkzn, &t.numberkz, &t.numberevent, &t.eventpoints, &t.active, &t.timedown)
		if user == 1 {
			u.user1 = t
		} else if user == 2 {
			u.user2 = t
		} else if user == 3 {
			u.user3 = t
		} else if user == 4 {
			u.user4 = t
		}
		user = user + 1
	}
	return u
}

func readMesID(mesid string) (string, error) {
	results, err := db.Query("SELECT lvlkz FROM sborkz WHERE dsmesid = ? AND active = 0", mesid)
	if err != nil {
		log.Println(err)
	}
	a := []string{}
	var dsmesid string
	for results.Next() {
		var t sborkz
		err = results.Scan(&t.lvlkz)
		a = append(a, t.lvlkz)
	}
	a = removeDuplicateElementString(a)
	if len(a) > 0 {
		dsmesid = a[0]
		return dsmesid, err
	} else {
		return "", err
	}
	return dsmesid, nil
}

/*
func readMesIDname(in inMessage, lvlkz string) (string, int) {
	mesidDS := ""
	mesidTG := 0
	results, err := db.Query("SELECT dsmesid,tgmesid FROM sborkz WHERE lvlkz = ? AND dschatid = ? AND tgchatid = ? AND name = ? AND active = 0", lvlkz, in.config.DsChannel, in.config.TgChannel, in.name)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		var t sborkz
		err = results.Scan(&t.dsmesid, &t.tgmesid)
		mesidDS = t.dsmesid
		mesidTG = t.tgmesid
	}
	return mesidDS, mesidTG
}
*/
