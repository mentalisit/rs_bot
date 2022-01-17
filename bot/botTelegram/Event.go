package botTelegram

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)
func qweryNumevent1(db *sql.DB,chatid int64)int{
	var event1 int
	row := db.QueryRow("SELECT numevent FROM rsevent WHERE chatid=? AND activeevent=1 ORDER BY numevent DESC LIMIT 1", chatid)
	err := row.Scan(&event1)
	if err != nil {}
	return event1
}
func qweryNumevent0(db *sql.DB,chatid int64)int{
	var event0 int
	row := db.QueryRow("SELECT numevent FROM rsevent WHERE chatid=? AND activeevent=0 ORDER BY numevent DESC LIMIT 1", chatid)
	err := row.Scan(&event0)
	if err != nil {}
	return event0
}

// старт ивента
//if (preg_match("/^(rs|Rs)\s(event|Event)\s(start|Start|stop|Stop)$/i", $mtext, $m)){
func EventStart(db *sql.DB,name string,chatid int64) {
	if name == "Mentalisit" {
		//проверяем, есть ли активный ивент
		event1 := qweryNumevent1(db, chatid)
		log.Println("event", event1)
		if event1 > 0 {
			mes := Send(chatid, "Режим ивента уже активирован.");
			go SendDelMessage5s(chatid, mes)
		} else {
			event0 := qweryNumevent0(db, chatid)
			if event0 > 0 {
				numberevent := event0 + 1
				insertEvent := `INSERT INTO rsevent (chatid,numevent,activeevent,number) VALUES (?,?,?,?)`
				_, err := db.Exec(insertEvent, chatid, numberevent, 1,1);
				if err != nil {
					log.Println(err)
				}
			} else {
				insertEvent := `INSERT INTO rsevent (chatid,numevent,activeevent,number) VALUES (?,?,?,?)`
				_, err := db.Exec(insertEvent, chatid, 1, 1, 1);
				if err != nil {
					log.Println(err)
				}
			}
			mes := Send(chatid, "Ивент запущен. После каждого похода на КЗ, " +
				"один из участников КЗ вносит полученные очки в базу командой К (номер катки) (количество набраных очков)")
			SendDelMessage5s(chatid, mes)
		}
	} else {
		mes := Send(chatid, "Запуск|ОСтановка Ивента доступен Администратору канала.")
		SendDelMessage5s(chatid, mes)
	}
}
func EventStop(db *sql.DB,name string,chatid int64) {
	if name == "Mentalisit" {
			event1 := qweryNumevent1(db, chatid)
			if event1 > 0 {
				//update
				_, err := db.Exec("UPDATE rsevent SET activeevent=0 WHERE chatid=? AND numevent=?", chatid, event1)
				if err != nil {
					log.Println(err)
				}
				mes := Send(chatid, "Ивент остановлен.")
				go SendDelMessage5s(chatid, mes)
			} else {
				mes := Send(chatid, "Ивент запущен не был. Нечего останавливать.")
				go SendDelMessage5s(chatid, mes)
			}
		} else {
			mes := Send(chatid, "Запуск|Остановка Ивента доступен Администратору канала.")
			go SendDelMessage5s(chatid, mes)
		}
}

func countEventsPoints(db *sql.DB, chatid int64, numberkz int)int{
	var countEventPoints int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkztg WHERE chatid=? AND numberkz=?  AND active=1 AND eventpoints > 0",chatid,numberkz)
	err := row.Scan(&countEventPoints)
	if err != nil {}
	return countEventPoints
}
// блок внесения очков за походы на КЗ во время ивента
func EventPoints(db *sql.DB,chatid int64,name string,numberkz,points int){
	// проверяем активен ли ивент
	event1:=qweryNumevent1(db,chatid)
	if event1>0{
		//проверка, был ли участник который добавляет очки, в кз с указанным номером
		var countEventNames int
		row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkztg WHERE chatid=? AND numberkz=?  AND active=1 AND name=?",chatid,numberkz,name)
		err := row.Scan(&countEventNames)
		if err != nil {}
		if countEventNames>0{
			pointsGood:=countEventsPoints(db,chatid,numberkz)
			if pointsGood>0{
				mes:=Send(chatid,"данные о кз уже внесены ")
				SendDelMessage5s(chatid,mes)
			}else {
			// считаем количество участников КЗ опр уровня
			var countEvent int
			row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkztg WHERE chatid=? AND numberkz=?  AND active=1",chatid,numberkz)
			err := row.Scan(&countEvent)
			if err != nil {}
			var pointsq int =points/countEvent
			//вносим очки
			_, err = db.Exec(`update sborkztg set numberevent=?, eventpoints=? WHERE chatid=? AND numberkz=? AND active=1`, event1,pointsq,chatid,numberkz)
			if err != nil {log.Println(err)	}
			mes:=Send(chatid,fmt.Sprintf("%s Очки %d внесены в базу",name,points))
				changeMessageEvent(db,points,countEvent,numberkz,event1,chatid)
			SendDelMessage5s(chatid,mes)}
		}else {
			mes:=Send(chatid,fmt.Sprintf("%s Вы не являетесь участником КЗ под номером %d добавление очков невозможно.",name,numberkz))
			SendDelMessage5s(chatid,mes)
		}
	}else{
		mes:=Send(chatid,"Ивент не запущен.")
		SendDelMessage5s(chatid,mes)
	}
}

func numberQueueEvents(db *sql.DB, chatid int64) int {

	var number int
	row2 := db.QueryRow("SELECT  number FROM rsevent WHERE activeevent = 1 AND chatid = ? ",  chatid)
	err1 := row2.Scan(&number)
	if err1 != nil {
		fmt.Println(err1)
	}
	if number == 0 {
		fmt.Println("нихрена нет ")
		insertSmt := "INSERT INTO rsevent(chatid,numevent,activeevent,number) VALUES (?, ?, ?, ?)"
		statement, err := db.Prepare(insertSmt)
		if err != nil {}
		number = 0
		_, err = statement.Exec(chatid,number,)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		return number
	}
	return number
}

func changeMessageEvent(db *sql.DB,points,countEvent,numberkz,numberEvent int,chatid int64){
	results, err := db.Query("SELECT * FROM sborkztg WHERE chatid=? AND numberkz=? AND numberevent = ? AND active=1",chatid,numberkz,numberEvent)
	if err != nil {fmt.Println(err)	}
	var tag Sborkz
	for results.Next() {
		err = results.Scan(&tag.Id, &tag.Name, &tag.Mesid, &tag.Chatid, &tag.Time, &tag.Date, &tag.Lvlkz, &tag.Numberkz,&tag.Numberevent, &tag.Eventpoints, &tag.Active, &tag.Timedown, &tag.Activedel)
		rs<-tag.Name
	}
	fmt.Println(tag.Mesid,countEvent,points,tag.Chatid)
	cMEvent(tag.Mesid,countEvent,points,numberkz,tag.Chatid)

}
func cMEvent(mesid,countEvent,points,numberkz int,chatid int64){
	mes1:=fmt.Sprintf("ивент игра №%d\n",numberkz)
	mesOld:=fmt.Sprintf("внесено %d",points)
	if countEvent==1{
		text:=fmt.Sprintf("%s@%s \n%s",mes1,<-rs,mesOld)
		EditMesText(chatid,mesid,text)
	}else if countEvent==2{
		text:=fmt.Sprintf("%s@%s\n @%s\n %s",mes1,<-rs,<-rs,mesOld)
		EditMesText(chatid,mesid,text)
	}else if countEvent==3{
		text:=fmt.Sprintf("%s@%s\n @%s\n @%s\n %s",mes1,<-rs,<-rs,<-rs,mesOld)
		EditMesText(chatid,mesid,text)
	}else if countEvent==4{
		text:=fmt.Sprintf("%s@%s\n @%s\n @%s @%s\n %s",mes1,<-rs,<-rs,<-rs,mesOld)
		EditMesText(chatid,mesid,text)
	}
}

func EditMesText(chatid int64, editMesId int, textEdit string) {
	tgbotapi.NewEditMessageText(chatid, editMesId, textEdit)
	Bot.Send(&tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:          chatid,
			ChannelUsername: "",
			MessageID:       editMesId,
			InlineMessageID: "",
			ReplyMarkup:     nil,
		},
		Text:                  textEdit,
		ParseMode:             "",
		DisableWebPagePreview: false,
	})
}
