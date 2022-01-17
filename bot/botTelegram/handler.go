package botTelegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func removeDuplicateElementInt(languages []int) []int {
	result := make([]int, 0, len(languages))
	temp := map[int]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
func myChatMember(member *tgbotapi.ChatMemberUpdated) {
	if member.NewChatMember.Status == "member" {
		Send(member.Chat.ID, fmt.Sprintf("@%s мне нужны права админа для коректной работы", member.From.UserName))
	} else if member.NewChatMember.Status == "administrator" {
		Send(member.Chat.ID, fmt.Sprintf("@%s спасибо ... я готов к работе ", member.From.UserName))
	}
}
