package botTelegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func myChatMember(member *tgbotapi.ChatMemberUpdated){
	if member.NewChatMember.IsMember {
		Send(member.Chat.ID,fmt.Sprintf("@%s мне нужны права админа для коректной работы",member.From.UserName))
	}else if member.NewChatMember.IsAdministrator(){
		Send(member.Chat.ID,fmt.Sprintf("@%s спасибо ... я готов к работе ",member.From.UserName))
	}
}
func updatesComand(c *tgbotapi.Message){
	if c.Command()=="help"{
		hhelp(c.From.UserName,c.Chat.ID)
	}
}