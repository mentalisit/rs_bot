package botDiscord

import (
	"fmt"
	"time"
)

func mainTime() {
	tm := time.Now()
	mdata := (tm.Format("2006-01-02"))
	mtime := (tm.Format("15:04"))
	fmt.Println(mdata)
	fmt.Println(mtime)
}
func reaact() {

}
