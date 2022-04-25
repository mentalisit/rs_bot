package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"rs_bot/bot/NewBot"
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
	//go bot2.StartBot()
	NewBot.StartBot()

	<-make(chan struct{})
	return
}
