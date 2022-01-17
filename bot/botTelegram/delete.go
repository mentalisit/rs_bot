package botTelegram

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

func deleteSborkz(db *sql.DB, name string, lvlkz string, chatid int64) {
	_, err := db.Exec("delete from sborkztg where name = ? AND lvlkz = ? AND chatid = ? AND active = 0", name, lvlkz, chatid)
	if err != nil {
		logrus.Println(err.Error())
	}
}
