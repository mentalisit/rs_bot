package discordBot

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

type logicDB interface {
	ConDbDs()
	readMesIDname(string, string, string) string
}
type DatabaseLogic struct {
	db *sql.DB
	mu *sync.Mutex
}

func conDbDs() *sql.DB {
	db, err := sql.Open("sqlite3", "./config/ds.db")
	if err != nil {
		log.Println(err)
	}
	db.SetConnMaxLifetime(1 * time.Second)
	return db
}

func (r DatabaseLogic) readMes(name, lvlkz, chatid string) string {
	r.mu.Lock()
	defer r.mu.Unlock()
	mesid := ""
	results, err := r.db.Query("SELECT mesid FROM sborkz WHERE lvlkz = ? AND chatid = ? AND name = ? AND active = 0", lvlkz, chatid, name)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		var t Sborkzds
		err = results.Scan(&t.Mesid)
		mesid = t.Mesid
		r.db.Close()
		return mesid
	}
	r.db.Close()
	return mesid
}

func readMesIDname(name, lvlkz, chatid string) string {
	db := conDbDs()
	mesid := ""
	results, err := db.Query("SELECT mesid FROM sborkz WHERE lvlkz = ? AND chatid = ? AND name = ? AND active = 0", lvlkz, chatid, name)
	if err != nil {
		log.Println(err)
	}
	for results.Next() {
		fmt.Println("подключаемся к дб дс месид")
		var t Sborkzds
		err = results.Scan(&t.Mesid)
		mesid = t.Mesid
		fmt.Println(t.Mesid, "vvvvvvvvvvvv")
		db.Close()
		return mesid
	}
	db.Close()
	fmt.Println(mesid, "ffffffff")
	return mesid
}
func deleteSrorkzTime(name, lvlkz, chatid string) {
	time.Sleep(2 * time.Second)
	deleteSrorkz(name, lvlkz, chatid)
}
func deleteSrorkz(name, lvlkz, chatid string) {
	db := conDbDs()
	fmt.Println(28)
	_, err := db.Exec("delete from sborkz where name = ? AND lvlkz = ? AND chatid = ? AND active = 0", name, lvlkz, chatid)
	if err != nil {
		fmt.Println("31 err")
		log.Println(err)
		fmt.Println(err.Error())
		db.Close()
		//deleteSrorkz(name,lvlkz, chatid)
	}
	fmt.Println("41")
	db.Close()
}
func removeQueue(name, nameid, lvlkz string, chatid, guildId string) {
	countName, countQueue := countnq(lvlkz, name, chatid) //:=countName(lvlkz,name,chatid)
	if countName == 0 {
		mes := SendChannel(chatid, nameid+" ты не в очереди.")
		go Delete5s(chatid, mes)
	} else if countName > 0 {
		//чтение айди очереди
		mesid := readMesIDname(name, lvlkz, chatid)
		fmt.Println(mesid, "qqqqqqqqqq")
		//удаление с базы данных
		deleteSrorkz(name, lvlkz, chatid)
		mes := SendChannel(chatid, nameid+" покинул очередь ")
		go Delete5s(chatid, mes)
		//проверяем, есть ли кто в очереди
		//countQueue:=countQueue(lvlkz,chatid)
		// совподения количество  условие
		if countQueue == 0 {
			mes := SendChannel(chatid, "Очередь КЗ была удалена .")
			go Delete5s(chatid, mes)
			go Delete5s(chatid, mesid)
		} else if countQueue == 1 {
			counts1r(chatid, guildId, lvlkz)
		} else if countQueue == 2 {
			counts2r(chatid, guildId, lvlkz)
		}
	}
}
