package timer

import (
	"log"
	"rs_bot/bot/botTelegram"
	"time"
)

//цикл минутный
func Timer() {
	for {
		time.Sleep(1 * time.Minute)

		log.Println("минута")

		go botTelegram.MinusMin()

	}
}
