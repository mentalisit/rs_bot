package botTelegram

import (
	"database/sql"
	"fmt"
)

func readAll(db *sql.DB, lvlkz string, chatid int64) int {
	var mesID int
	a := []int{}
	aa := []int{}
	results, err := db.Query("SELECT * FROM sborkztg WHERE lvlkz = ? AND chatid = ? AND active = 0", lvlkz, chatid)
	if err != nil {
		fmt.Println(err)
	}
	for results.Next() {

		var tag Sborkz
		err = results.Scan(&tag.Id, &tag.Name, &tag.Mesid, &tag.Chatid, &tag.Time, &tag.Date, &tag.Lvlkz, &tag.Numberkz,&tag.Numberevent, &tag.Eventpoints, &tag.Active, &tag.Timedown, &tag.Activedel)
		//fmt.Println(tag)
		rs <- tag.Name
		rst <- tag.Timedown
		numberkzName:=countNumberNameActive1(db,tag.Name,chatid,lvlkz)//проверяем количество боёв по уровню кз игрока
		rsn <- numberkzName
		if tag.Timedown < 0 {
			deleteSborkz(db, tag.Name, tag.Lvlkz, tag.Chatid)
		}
		a = append(a, tag.Mesid)
	}

	for _, v := range a {
		skip := false
		for _, u := range aa {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			//fmt.Println(v)
			mesID = v
			return mesID
		}
		return mesID
	}
	return mesID
}

func readMesid(db *sql.DB, lvlkz string, chatid int64) int {
	var mesID int
	results := db.QueryRow("SELECT mesid FROM sborkztg WHERE lvlkz = ? AND chatid = ? AND active = 0", lvlkz, chatid)
	results.Scan(&mesID)
	return mesID
}
