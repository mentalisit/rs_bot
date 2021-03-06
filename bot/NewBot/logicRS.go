package NewBot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"regexp"
	"strconv"
)

//var rs = make(chan string, 4)

func logicMixDiscord(m *discordgo.MessageCreate) {
	ok, config := checkChannelConfigDS(m.ChannelID) //проверяем чат Ид по конфигу
	accesChatDS(m)                                  //для добавления в конфиг
	if ok {
		var Attachments discordgo.MessageAttachment
		if len(m.Attachments) > 0 {
			for _, attach := range m.Attachments { //вложеные файлы
				Attachments = *attach
			}
		}
		member, e := DSBot.GuildMember(m.GuildID, m.Author.ID) //проверка есть ли изменения имени в этом дискорде
		if e != nil {
			logrus.Println(e)
		}
		name := m.Author.Username
		if member.Nick != "" {
			name = member.Nick
		}

		in := inMessage{
			mtext:       m.Content,
			tip:         "ds",
			name:        name,
			nameMention: m.Author.Mention(),
			Ds: Ds{
				mesid:       m.ID,
				nameid:      m.Message.Author.ID,
				guildid:     m.GuildID,
				Attachments: &Attachments,
			},
			Tg:     Tg{},
			config: config,
		}

		logicRs(in)
		cleanChat(m)
	}
}

func logicMixTelegram(m *tgbotapi.Message) {
	ok, config := checkChannelConfigTG(m.Chat.ID) //проверяем чат Ид по конфигу
	accesChatTg(m)                                //для добавления в конфиг
	if ok {
		in := inMessage{
			mtext:       m.Text,
			tip:         "tg",
			name:        m.From.UserName,
			nameMention: "@" + m.From.UserName,
			Ds:          Ds{},
			Tg: Tg{
				mesid:  m.MessageID,
				nameid: m.From.ID,
			},
			config: config,
			option: Option{
				callback: false,
				edit:     false,
				update:   false,
			},
		}
		logicRs(in)
	}
}

func logicRs(in inMessage) {
	fmt.Println(in.config.CorpName, in.tip, in.name, in.mtext)
	var rss, kzb, subs, subs3, lvlkz, qwery string
	file := false
	if in.tip == "ds" && len(in.Ds.Attachments.ID) > 0 {
		file = true
	}
	if len(in.mtext) > 0 || file {
		str := in.mtext
		re := regexp.MustCompile(`^([4-9]|[1][0-1])([\+]|[-])(\d|\d{2}|\d{3})$`) //три переменные
		arr := (re.FindAllStringSubmatch(str, -1))
		if len(arr) > 0 {
			in.lvlkz = arr[0][1]
			kzb = arr[0][2]
			timekzz, err := strconv.Atoi(arr[0][3])
			if err != nil {
				logrus.Println(err)
			}
			if timekzz > 180 {
				timekzz = 180
			}
			in.timekz = strconv.Itoa(timekzz)
		}
		re2 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+]|[-])$`) // две переменные
		arr2 := (re2.FindAllStringSubmatch(str, -1))
		if len(arr2) > 0 {
			in.lvlkz = arr2[0][1]
			kzb = arr2[0][2]
			in.timekz = "30"
		}

		re3 := regexp.MustCompile(`^([\+]|[-])([4-9]|[1][0-1])$`) // две переменные для добавления или удаления подписок
		arr3 := (re3.FindAllStringSubmatch(str, -1))
		if len(arr3) > 0 {
			lvlkz = arr3[0][2]
			subs = arr3[0][1]
		} else {
			re3 := regexp.MustCompile(`^(Rs|rs)\s(S|s)\s([4-9]|[1][0-1])$`)
			arr3 := (re3.FindAllStringSubmatch(str, -1))
			if len(arr3) > 0 {
				lvlkz = arr3[0][3]
				subs = arr3[0][2]
				if subs == "S" || subs == "s" {
					subs = "+"
				} else if subs == "U" || subs == "u" {
					subs = "-"
				}
				fmt.Println("Тестирование подписок совместимости")
			}
		}

		re4 := regexp.MustCompile(`^(["о"]|["О"]|["o"]|["O"])([4-9]|[1][0-1])$`) // две переменные для чтения  очереди
		arr4 := (re4.FindAllStringSubmatch(str, -1))
		if len(arr4) > 0 {
			qwery = arr4[0][1]
			in.lvlkz = arr4[0][2]
		}

		re5 := regexp.MustCompile(`^([4-9]|[1][0-1])([\+][\+])$`) //rs start
		arr5 := (re5.FindAllStringSubmatch(str, -1))
		if len(arr5) > 0 {
			in.lvlkz = arr5[0][1]
			rss = arr5[0][2]
		} else {
			re5 = regexp.MustCompile(`^(Rs|rs)\s(Start|start)\s([4-9]|[1][0-1])$`) //rs start
			arr5 = (re5.FindAllStringSubmatch(str, -1))
			if len(arr5) > 0 {
				in.lvlkz = arr5[0][3]
				rss = "++"
				fmt.Println("Проверка совместимости принудительного старта ")
			}

			re6 := regexp.MustCompile(`^([\+][\+]|[-][-])([4-9]|[1][0-1])$`) // две переменные
			arr6 := (re6.FindAllStringSubmatch(str, -1))                     // для добавления или удаления подписок 3/4
			if len(arr6) > 0 {
				lvlkz = arr6[0][2]
				subs3 = arr6[0][1]
			} else {
				re6 = regexp.MustCompile(`^(Rs|rs)\s(S|s)\s([4-9]|[1][0-1])(\+)$`)
				arr6 = (re6.FindAllStringSubmatch(str, -1))
				if len(arr6) > 0 {
					lvlkz = arr6[0][3]
					subs3 = arr6[0][2]
					if subs3 == "S" || subs3 == "s" {
						subs3 = "++"
					} else if subs3 == "U" || subs3 == "u" {
						subs3 = "--"
					}
					fmt.Println("проверка совместимости подписок 3 из 4")
				}
			}

			re7 := regexp.MustCompile(`^(["К"]|["к"])\s([0-9]+)\s([0-9]+)$`) // ивент
			arr7 := (re7.FindAllStringSubmatch(str, -1))
			if len(arr7) > 0 {
				points, err := strconv.Atoi(arr7[0][3])
				if err != nil {
					logrus.Println(err)
				}
				numkz, err := strconv.Atoi(arr7[0][2])
				if err != nil {
					logrus.Println(err)
				}
				EventPoints(in, numkz, points)

			}

			re8 := regexp.MustCompile(`^(["T"]|["t"]|["т"]|["Т"])([4-9]|[1][0-1])$`) // запрос топа по уровню
			arr8 := (re8.FindAllStringSubmatch(str, -1))
			if len(arr8) > 0 {
				lvlkz = arr8[0][2]
			}

			var slot, emo string
			reEmodji := regexp.MustCompile("^(Эмоджи)\\s([1-4])\\s(<:\\w+:\\d+>)$") //добавления внутрених эмоджи
			arrEmodji := (reEmodji.FindAllStringSubmatch(str, -1))
			if len(arrEmodji) > 0 {
				slot = arrEmodji[0][2]
				emo = arrEmodji[0][3]
			}
			reEmodji = regexp.MustCompile("^(Эмоджи)\\s([1-4])\\s(\\P{Greek})$") //добавления эмоджи
			arrEmodji = (reEmodji.FindAllStringSubmatch(str, -1))
			if len(arrEmodji) > 0 {
				slot = arrEmodji[0][2]
				emo = arrEmodji[0][3]
			}
			reEmodji = regexp.MustCompile("^(Эмоджи)\\s([1-4])$") //удаление эмоджи с ячейки
			arrEmodji = (reEmodji.FindAllStringSubmatch(str, -1))
			if len(arrEmodji) > 0 {
				slot = arrEmodji[0][2]
				emo = ""
			}

			if kzb == "+" {
				//RsPlus(in, lvlkz, timekz)
				in.RsPlus()

			} else if kzb == "-" {
				//RsMinus(in, lvlkz)
				in.RsMinus()
			} else if len(qwery) > 0 {
				in.Queue()
			} else if len(rss) > 0 {
				in.RsStart()
			} else if subs == "+" {
				go Subscribe(in, lvlkz, 1)
			} else if subs3 == "++" {
				go Subscribe(in, lvlkz, 3)
			} else if subs == "-" {
				go Unsubscribe(in, lvlkz, 1)
			} else if subs3 == "--" {
				go Unsubscribe(in, lvlkz, 3)
			} else if len(arr8) > 0 {
				go TopLevel(in, lvlkz)
			} else if len(slot) > 0 {
				emodjiadd(in, slot, emo)
			} else if ifText(in) {
			} else if str == "1" {
				dsSendChannelDel1m(in.config.DsChannel, "test "+emReadName(in.name))
			} else if in.config.TgChannel != 0 && in.config.DsChannel != "" {
				go bridge(in)
			}
		}
	}
}

func bridge(in inMessage) {
	if in.tip == "ds" {
		//отправляем url файла
		if len(in.Ds.Attachments.ID) > 0 {
			in.mtext = in.mtext + "\n" + in.Ds.Attachments.URL
		}
		//fmt.Println(154,replaceUserMentions1(in.mtext))
		//fmt.Println(155,replaceID(in.mtext,in.Ds.guildid))
		//go dsDeleteMesageMinuts(in.config.DsChannel,in.Ds.mesid,3)
		text := fmt.Sprintf("(DS)%s \n%s", in.name, in.mtext)
		mes := tgSendChannel(in.config.TgChannel, text)
		go tgDeleteMesageMinuts(in.config.TgChannel, mes, 3)

	} else if in.tip == "tg" {
		text := fmt.Sprintf("(TG)%s \n%s", in.name, in.mtext)
		mes := dsSendChannel(in.config.DsChannel, text)
		go dsDeleteMesageMinuts(in.config.DsChannel, mes, 3)
	}
}

func ifText(in inMessage) bool {
	iftext := true
	switch in.mtext {
	case "Ивент старт":
		EventStart(in)
	case "Ивент стоп":
		EventStop(in)
	case "+":
		in.Plus()
	case "-":
		in.Minus()
	case "Справка":
		hhelpName(in)
	case "Справка1":
		if in.tip == "ds" {
			dsDelMessage(in.config.DsChannel, in.Ds.mesid)
			helpChannelUpdate(in.config.DsChannel)
		}
	case "Топ":
		TopAll(in)
	case "Эмоджи":
		emodjis(in)

	default:
		iftext = false
	}
	return iftext
}

func iftipdelete(in inMessage) {
	if in.tip == "ds" {
		dsDelMessage(in.config.DsChannel, in.Ds.mesid)
	} else if in.tip == "tg" {
		tgDelMessage(in.config.TgChannel, in.Tg.mesid)
	}
}

func (in inMessage) iftipdelete() {
	if in.tip == "ds" && !in.option.callback {
		dsDelMessage(in.config.DsChannel, in.Ds.mesid)
	} else if in.tip == "tg" && !in.option.callback {
		tgDelMessage(in.config.TgChannel, in.Tg.mesid)
		if in.nameMention == "@" {
			tgSendChannelDel1m(in.config.TgChannel, nickname)
		}
	}
}

func (in inMessage) ifTipSendMentionText(text string) {
	if in.tip == "ds" {
		go dsSendChannelDel5s(in.config.DsChannel, in.nameMention+text)
	} else if in.tip == "tg" {
		go tgSendChannelDel5s(in.config.TgChannel, in.nameMention+text)
	}
}
