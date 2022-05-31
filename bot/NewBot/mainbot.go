package NewBot

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"os"
	"rs_bot/bot/NewBot/database"
)

var (
	BotId                      string
	TgBot, BotErr              = tgbotapi.NewBotAPI(os.Getenv("TokenT"))
	DSBot                      *discordgo.Session
	db                         *sql.DB
	err                        error
	name1, name2, name3, name4 string
	time1, time2, time3, time4 string
	//wabotStart                 = false
	//client                     *whatsmeow.Client
	//w                          wa
)

func StartBot() {
	DSBot, err = discordgo.New("Bot " + os.Getenv("TokenD"))
	if err != nil {
		logrus.Println(err)
		go StartBot()
	}
	u, err := DSBot.User("@me")
	if err != nil {
		logrus.Println(err.Error())
	}
	BotId = u.ID
	DSBot.AddHandler(messageHandler)
	DSBot.AddHandler(MessageReactionAdd)
	err = DSBot.Open()
	if err != nil {
		logrus.Println(err.Error())
	}
	fmt.Println("Бот DISCORD запущен!!!")

	readBotConfig()
	TgBot, BotErr = tgbotapi.NewBotAPI(os.Getenv("TokenT"))
	if BotErr != nil {
		logrus.Panic(BotErr)
	}
	TgBot.Debug = false
	fmt.Printf("Бот TELEGRAM загружен  %s", TgBot.Self.UserName)
	ut := tgbotapi.NewUpdate(0)
	ut.Timeout = 60
	updatesChannelTg(ut)

	db, err = database.DbConnection()
	if err != nil {
		logrus.Println(err)
	}

	/*
		//whatsapp
		if wabotStart {

			dbLog := waLog.Stdout("Database", "error", true)
			container, err := sqlstore.New("sqlite3", "file:examplestore.db?_foreign_keys=on", dbLog)
			if err != nil {
				panic(err)
			}
			// Если вам нужно несколько сеансов, запомните их JID и используйте вместо них .GetDevice(jid) или .GetAllDevices().
			deviceStore, err := container.GetFirstDevice()
			if err != nil {
				panic(err)
			}
			clientLog := waLog.Stdout("Client", "error", true)
			client = whatsmeow.NewClient(deviceStore, clientLog)
			client.AddEventHandler(eventHandler)

			if client.Store.ID == nil {
				// ID не сохранен, новый логин
				qrChan, _ := client.GetQRChannel(context.Background())
				err = client.Connect()
				if err != nil {
					panic(err)
				}
				for evt := range qrChan {
					if evt.Event == "code" {
						fmt.Println("QR code:", evt.Code)
						qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
					} else {
						fmt.Println("Login event:", evt.Event)
					}
				}
			} else {
				// Already logged in, just connect
				err = client.Connect()
				if err != nil {
					panic(err)
				}

				w.startedAt = int64(time.Now().Unix())
			}

			// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
			c := make(chan os.Signal)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			<-c

			client.Disconnect()
		}

	*/

	//
}
