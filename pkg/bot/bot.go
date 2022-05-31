package bot

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"rs_bot/pkg/client/ds"
)

type Bot struct {
	tg *tgbotapi.BotAPI
	ds ds.Init
	db *sql.DB
}

func InitBOT(tg *tgbotapi.BotAPI, ds *discordgo.Session, db *sql.DB) *Bot {
	return &Bot{
		tg: tg,
		ds: ds,
		db: db,
	}

}
func (b *Bot) Start() {
	b.ds.InitDiscord()
}
