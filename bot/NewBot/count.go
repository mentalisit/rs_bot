package NewBot

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

//проверка есть ли игрок в очереди
func countName(in inMessage, lvlkz string) int {
	var countNames int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND lvlkz = ? AND corpname = ? AND active = 0", in.name, lvlkz, in.config.CorpName)
	err := row.Scan(&countNames)
	if err != nil {
		logrus.Println(err)
	}
	return countNames
}

//проверка сколько игровок в очереди
func countQueue(in inMessage, lvlkz string) int {
	var count int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE lvlkz = ? AND corpname = ? AND active = 0",
		lvlkz, in.config.CorpName)
	err := row.Scan(&count)
	if err != nil {
		fmt.Println(err)
	}
	return count
}

// выковыриваем из базы значение количества походов на кз
func countNumberNameActive1(in inMessage, lvlkz string) int {
	var countNumberNameActive1 int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE lvlkz = ? AND corpname = ? AND name = ? AND active = 1", lvlkz, in.config.CorpName, in.name)
	err := row.Scan(&countNumberNameActive1)
	if err != nil {
		fmt.Println(err)
	}
	return countNumberNameActive1
}

// выковыриваем из базы значение количества походов на кз
func countNumberNameAct1(tip string, dschatid string, tgchatid int64, name string) int {
	var countNumberNameActive1 int
	if tip == "ds" {
		row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE dschatid = ? AND name = ? AND active = 1", dschatid, name)
		err := row.Scan(&countNumberNameActive1)
		if err != nil {
			fmt.Println(err)
		}
	} else if tip == "tg" {
		row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE tgchatid = ? AND name = ? AND active = 1", tgchatid, name)
		err := row.Scan(&countNumberNameActive1)
		if err != nil {
			fmt.Println(err)
		}
	}
	return countNumberNameActive1
}
