package main

import (
	"log"
	"os"
)

var (
	outfile,_= os.OpenFile("var/log/bot.log",os.O_RDWR|os.O_CREATE|os.O_APPEND,0755)
	LogFile = log.New(outfile,"",0)
)
func ForError(er error) {
	if er !=nil{
		LogFile.Fatalln(er)
	}
}


func CheckErr(err error) {
	if err != nil{
		panic(err)}
}

func cut(text string, limit int) string {
	runes := []rune(text)
	if len(runes) >= limit {
		return string(runes[:limit])
	}
	return text
}

