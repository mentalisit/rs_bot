package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"rs_bot/bot/NewBot"
)

func init() {
	//	go timer.Timer()
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	log.Println("ЗАПУСК БОТА")
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	//go bot2.StartBot()
	//fmt.Println(config.InitEnv())
	NewBot.StartBot()

	<-make(chan struct{})
	return
}
