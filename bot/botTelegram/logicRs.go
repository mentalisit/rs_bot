package botTelegram

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"time"
)

func conDBsqLite() *sql.DB { //sqLite
	db, err := sql.Open("sqlite3", "./db/tg.db")
	if err != nil {
		logrus.Println(err)
	}
	db.SetConnMaxLifetime(1 * time.Second)
	return db
}

func logicRs(Message *tgbotapi.Message, db *sql.DB) {
	mtext := Message.Text //текст сообщения
	//tm := time.Unix(int64(Message.Date), 0).UTC()
	//mdata := (tm.Format("2006-01-02"))
	//mtime := (tm.Format("15:04"))
	mesid := Message.MessageID    // ид сообщения
	name := Message.From.UserName // имя
	nameidid := Message.From.ID   //chat id
	chatid := Message.Chat.ID     //chat id
	var kzb, subs, subs3, qwery, rss, lvlkz, timekz string

	if Message.From.UserName == "" {
		mes:=Send(Message.Chat.ID, "Для того что бы БОТ мог Вас индентифицировать, создайте уникальный НикНей в настройках. Вы можете использовать a-z, 0-9 и символы подчеркивания. Минимальная длина - 5 символов.")
		SendDelMessage3m(chatid,mes)
	} else if Message.From.UserName != "" {
		if len(Message.Text) > 0 {
			msg := tgbotapi.NewMessage(Message.Chat.ID, "")

			str := mtext //str := "9+55" //полученая строка из телеграм
			fmt.Println("Test string", name, " ", str)
			re := regexp.MustCompile(`^([4-9]|[1][0-1])([\+]|[-])(\d|\d{2})$`) //три переменные
			arr := (re.FindAllStringSubmatch(str, -1))
			if len(arr) > 0 {
				lvlkz = arr[0][1]
				kzb = arr[0][2]
				timekz = arr[0][3]
			}

			re2 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+]|[-])$`) // две переменные
			arr2 := (re2.FindAllStringSubmatch(str, -1))
			if len(arr2) > 0 {
				lvlkz = arr2[0][1]
				kzb = arr2[0][2]
				timekz = "30"
			}

			re3 := regexp.MustCompile(`^([\+]|[-])([4-9]|[1][0-1])$`) // две переменные для добавления или удаления подписок
			arr3 := (re3.FindAllStringSubmatch(str, -1))
			if len(arr3) > 0 {
				lvlkz = arr3[0][2]
				subs = arr3[0][1]
			}

			re4 := regexp.MustCompile(`^(["о"]|["О"])([4-9]|[1][0-1])$`) // две переменные для чтения  очереди
			arr4 := (re4.FindAllStringSubmatch(str, -1))
			if len(arr4) > 0 {
				qwery = arr4[0][1]
				lvlkz = arr4[0][2]
			}
			re5 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+][\+])$`) //rs start
			arr5 := (re5.FindAllStringSubmatch(str, -1))
			if len(arr5) > 0 {
				lvlkz = arr5[0][1]
				rss = arr5[0][2]
			}
			re6 := regexp.MustCompile(`^([\+][\+]|[-][-])([4-9]|[1][0-1])$`) // две переменные для добавления или удаления подписок 3/4
			arr6 := (re6.FindAllStringSubmatch(str, -1))
			if len(arr6) > 0 {
				lvlkz = arr6[0][2]
				subs3 = arr6[0][1]
			}
//("/^(rs|Rs)\s(p|P)\s([0-9]+)\s([0-9]+)$/i"
			re7 := regexp.MustCompile(`^(["К"])\s([0-9]+)\s([0-9]+)$`) // ивент
			arr7 := (re7.FindAllStringSubmatch(str, -1))
			if len(arr7) > 0 {
				points,err := strconv.Atoi(arr7[0][3]);if err!=nil{}
				numkz,err := strconv.Atoi(arr7[0][2]);if err!=nil{}
				fmt.Println(points)
				fmt.Println(numkz)
				EventPoints(db,chatid,name,numkz,points)
				SendDelMessage5s(chatid,mesid)
			}
			if kzb == "+" { //запуск очереди кз
				RsPlus(db, name, lvlkz, timekz, chatid);go SendDelMessage5s(chatid, mesid)
			}else if kzb == "-" { //выход с очереди кз
				RsMinus(db, name, lvlkz, chatid);go SendDelMessage5s(chatid, mesid)
			}else if subs == "+" { // подписаться на пинг
				subscribe(db, name, nameidid, lvlkz, mesid, chatid);go SendDelMessage5s(chatid, mesid)
			}else if subs == "-" { //отписаться от пинга
				unsubscribe(db, name, nameidid, lvlkz, mesid, chatid);go SendDelMessage5s(chatid, mesid)
			}else if subs3 == "++" { //подписаться на пинг 3/4
				subscribe3(db, name, nameidid, lvlkz, mesid, chatid);go SendDelMessage5s(chatid, mesid)
			}else if subs3 == "--" { //отписаться от пинга 3/4
				unsubscribe3(db, name, nameidid, lvlkz, mesid, chatid);go SendDelMessage5s(chatid, mesid)
			}else if len(qwery) > 0 { //проверка кто в очереди
				Queue(db, lvlkz, chatid);go SendDelMessage5s(chatid, mesid)
			}else if len(rss) > 0 { //принудительный старт
				RsStart(db, name, lvlkz, chatid);go SendDelMessage5s(chatid, mesid)
			}else if mtext == "Справка"{go hhelp(name, chatid);go SendDelMessage5s(chatid, mesid)
			}else if mtext == "help" {go hhelp(name, chatid);go SendDelMessage5s(chatid, mesid)
			}else if Message.IsCommand(){updatesComand(Message)
			}else if mtext == "+" { //продление очереди на 30мин
				go SendDelMessage5s(chatid, mesid);Plus(db, name, chatid)
			}else if mtext == "-" {go SendDelMessage5s(chatid, mesid);	Minus(db,name,chatid)
			}else if mtext=="Ивент старт"{EventStart(db,name,chatid);SendDelMessage5s(chatid,mesid)
			}else if mtext=="Ивент стоп"{EventStop(db,name,chatid);SendDelMessage5s(chatid,mesid)
			}else if mtext=="a"{
				nu:=numberQueueChatid(db,chatid)
				fmt.Println(nu)
			}else if mtext=="A"{
				numberUpdateChatid(db,chatid)
			}

			Bot.Send(msg)
		}
	}

}
