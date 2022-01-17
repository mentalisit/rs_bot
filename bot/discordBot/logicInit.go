package discordBot

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func init() {
	createTableSborkz()

}

func createTableSborkz() { // создание таблицы
	db := conDbDs()
	createTable := `CREATE TABLE sborkz (
		"id"	integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name"	TEXT,
		"nameid"	TEXT,
		"guildid"	TEXT,
		"lvlkz"	TEXT,
		"chatid"	TEXT,
		"mesid"	TEXT,
		"timedown"	INTEGER,
		"active"	INTEGER
);`
	statement, err := db.Prepare(createTable)
	if err == nil {
		statement.Exec()
		db.Close()
		log.Println("создание таблицы сбора для дискорда ")
	}
	db.Close()
}
