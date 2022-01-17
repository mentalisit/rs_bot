package timer

/*
import (
	"fmt"
	"tgbot/bot/discordBot"
	"tgbot/bot/msql"
)

func  timeo()  {
	database:= msql.MysqlConn("sborkzds")
	results, err := database.Query("SELECT * FROM sborkzds WHERE active = 0");
	if err != nil{panic(err.Error())}
	for results.Next(){
		var tag discordBot.Sborkzds
		err = results.Scan(&tag.Id,&tag.Name,&tag.Nameids,&tag.Guildids,&tag.Lvlkzs,&tag.Chatids,&tag.Mesid,&tag.Timedowns,tag.Actives)
		//fmt.Println(tag.Mesid)
		if tag.Timedowns == 3 {
			mes:=discordBot.SendChannel(tag.Chatids,tag.Nameids+" время почти вышло  ...\n если ты еще тут пиши +")
			go discordBot.Delete3m(tag.Chatids,mes)

		}else if tag.Timedowns == 0 {
			discordBot.MsqlRemove(tag.Name,tag.Nameids,tag.Lvlkzs,tag.Chatids,tag.Guildids)

		}else if tag.Timedowns <= -1 {
			discordBot.MsqlRemove(tag.Name,tag.Nameids,tag.Lvlkzs,tag.Chatids,tag.Guildids)
		}else{}
		discordBot.MessageUpdate1m()
		//fmt.Println(tag.Lvlkzs, tag.Chatids, tag.Name, tag.Timedowns,tag.Mesid)
		database.Close()


	}
}


//обновление  -1
func  time1m()  {
	database:= msql.MysqlConn("sborkzds")
	_,err:=database.Exec(`update sborkzds set timedown = timedown - 1 where active = 0` )
	if err != nil {
		timeo()
		fmt.Println(err)
	}
	defer database.Close()
	timeo()

}

*/
