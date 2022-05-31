package ds

import (
	"github.com/bwmarrin/discordgo"
	"os"
)

type Init interface {
	InitDiscord() (*discordgo.Session, error)
}

func InitDiscord() (*discordgo.Session, error) {
	s, err := discordgo.New("Bot " + os.Getenv("TOKEND"))

	return s, err
}
