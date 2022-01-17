package botDiscord

import (
	"database/sql"
	"log"
)

func countName(db *sql.DB, lvlkz string, name string, chatid string) int {

	var countNames int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE name = ? AND lvlkz = ? AND chatid = ? AND active = 0", name, lvlkz, chatid)
	err := row.Scan(&countNames)
	if err != nil {
		log.Println("imenDS", countNames, err)
	}
	return countNames
}

func countQueue(db *sql.DB, lvlkz string, chatid string) int {

	var countQueue int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkz WHERE lvlkz = ? AND chatid = ? AND active = 0", lvlkz, chatid)
	err = row.Scan(&countQueue)
	if err != nil {
		log.Println("количество в очереди ошибка", countQueue,err)
	}
	return countQueue
}
