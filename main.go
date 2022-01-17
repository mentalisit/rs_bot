package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"rs_bot/bot/botDiscord"
	"rs_bot/bot/botTelegram"
	"rs_bot/bot/timer"
)

func init() {
	go timer.Timer()
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	fmt.Println("ЗАПУСК БОТА")
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	go botDiscord.Start()
	go botTelegram.StartTgBot()

	<-make(chan struct{})
	return
}
