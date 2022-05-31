package timer

import (
	"rs_bot/bot/NewBot"
	"time"
)

//цикл минутный
func Timer() {

	for {

		tm := time.Now().Second()
		if tm == 0 {
			//log.Println("минута")

			NewBot.MinusMin()
		}
		time.Sleep(1 * time.Second)

	}
}
