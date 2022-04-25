package NewBot

import (
	"log"
)

func updateComplite(lvlkz string, dsmesid string, tgmesid int, wamesid string, numberkz int, numberevent int, corpname string) {
	_, err := db.Exec(
		`update sborkz set active = 1,dsmesid = ?,tgmesid = ?,wamesid = ?,numberkz = ?,numberevent = ? 
where lvlkz = ? AND corpname = ? AND active = 0`, dsmesid, tgmesid, wamesid,
		numberkz, numberevent, lvlkz, corpname)
	if err != nil {
		log.Println(err)
	}
	if numberevent > 0 {
		_, err := db.Exec(
			`update rsevent set number = number+1  where corpname = ? AND activeevent = 1`, corpname)
		if err != nil {
			log.Println(err)
		}
	}
}

func mesidDsUpdate(mesidds, lvlkz, corpname string) {
	_, err := db.Exec(
		`update sborkz set dsmesid = ? where lvlkz = ? AND corpname = ? `,
		mesidds, lvlkz, corpname)
	if err != nil {
		log.Println(err)
	}
}

func mesidTgUpdate(mesidtg int, lvlkz string, corpname string) {
	_, err := db.Exec(
		`update sborkz set tgmesid = ? where lvlkz = ? AND corpname = ? `,
		mesidtg, lvlkz, corpname)
	if err != nil {
		log.Println(err)
	}
}
