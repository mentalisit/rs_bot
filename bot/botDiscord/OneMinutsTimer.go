package botDiscord

import (
	"database/sql"
	"fmt"
)

func OneMinutsTimer(db *sql.DB) {
	var count int
	row := db.QueryRow("SELECT  COUNT(*) as count FROM sborkzds WHERE active = 0")
	err := row.Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
	}
	if count > 0 {
		a := []string{}
		aa := []string{}
		results, err := db.Query("SELECT name,nameid,mesid,timedown FROM sborkzds WHERE active = 0")
		if err != nil {
			fmt.Println(err)
		}
		var tagDS Sborkzds
		for results.Next() {
			err = results.Scan(&tagDS.Name, &tagDS.Nameid, &tagDS.Mesid, &tagDS.Timedown)
			a = append(a, tagDS.Mesid)
		}
		//fmt.Println(a)
		for _, v := range a {
			skip := false
			for _, u := range aa {
				if v == u {
					skip = true
					break
				}
			}
			if !skip {
				//fmt.Println("chatid found36 ",v)
				//messageupdate(v)
				/*	var count40 int
									row := database.QueryRow("SELECT  COUNT(*) as count FROM sborkzds WHERE mesid = ? AND active = 0",v)
									err := row.Scan(&count40)
									if err !=nil{fmt.Println(err.Error())}
									if count40==1{
										results, err := database.Query("SELECT guildid,lvlkz,chatid FROM sborkzds WHERE mesid = ? AND active = 0",v);
										if err != nil{ fmt.Println(err)	}
										var tagDS Sborkzds
										for results.Next() {
											err = results.Scan(&tagDS.Guildids, &tagDS.Lvlkzs,&tagDS.Chatids)
											counts1q(tagDS.Chatids,tagDS.Guildids,tagDS.Lvlkzs)}

									}else if count40==2{
										results, err := database.Query("SELECT guildid,lvlkz,chatid FROM sborkzds WHERE mesid = ? AND active = 0",v);
										if err != nil{ fmt.Println(err)	}
										var tagDS Sborkzds
										for results.Next() {
											err = results.Scan(&tagDS.Guildids, &tagDS.Lvlkzs,&tagDS.Chatids)
											counts2q(tagDS.Chatids,tagDS.Guildids,tagDS.Lvlkzs)}
									}else if count40==3{
										results, err := database.Query("SELECT guildid,lvlkz,chatid FROM sborkzds WHERE mesid = ? AND active = 0",v);
										if err != nil{ fmt.Println(err)	}
										var tagDS Sborkzds
										for results.Next() {
											err = results.Scan(&tagDS.Guildids, &tagDS.Lvlkzs,&tagDS.Chatids)
											counts3q(tagDS.Chatids,tagDS.Guildids,tagDS.Lvlkzs) }
									}
								}
							defer database.Close()
							}}

					}

					func messageupdate(v string) {
						var count40 int
						database := msql.MysqlConn("sborkzds")
						row := database.QueryRow("SELECT  COUNT(*) as count FROM sborkzds WHERE mesid = ? AND active = 0", v)
						err := row.Scan(&count40)
						if err != nil {
							fmt.Println(err.Error())
						}
						if count40 == 1 {
							results, err := database.Query("SELECT guildid,lvlkz,chatid FROM sborkzds WHERE mesid = ? AND active = 0", v);
							if err != nil {
								fmt.Println(err)
							}
							var tagDS Sborkzds
							for results.Next() {
								err = results.Scan(&tagDS.Guildids, &tagDS.Lvlkzs, &tagDS.Chatids)
								counts1q(tagDS.Chatids, tagDS.Guildids, tagDS.Lvlkzs)
							}

						} else if count40 == 2 {
							results, err := database.Query("SELECT guildid,lvlkz,chatid FROM sborkzds WHERE mesid = ? AND active = 0", v);
							if err != nil {
								fmt.Println(err)
							}
							var tagDS Sborkzds
							for results.Next() {
								err = results.Scan(&tagDS.Guildids, &tagDS.Lvlkzs, &tagDS.Chatids)
								counts2q(tagDS.Chatids, tagDS.Guildids, tagDS.Lvlkzs)
							}
						} else if count40 == 3 {
							results, err := database.Query("SELECT guildid,lvlkz,chatid FROM sborkzds WHERE mesid = ? AND active = 0", v);
							if err != nil {
								fmt.Println(err)
							}
							var tagDS Sborkzds
							for results.Next() {
								err = results.Scan(&tagDS.Guildids, &tagDS.Lvlkzs, &tagDS.Chatids)
								counts3q(tagDS.Chatids, tagDS.Guildids, tagDS.Lvlkzs)
							}
						}
						database.Close()
					}
				*/
			}
		}
	}
}
