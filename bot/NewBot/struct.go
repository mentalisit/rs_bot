package NewBot

import (
	"github.com/bwmarrin/discordgo"
	"sync"
)

type inMessage struct {
	Mutex         sync.Mutex
	mtext         string
	tip           string
	name          string
	nameMention   string
	lvlkz, timekz string
	Ds            Ds
	Tg            Tg
	config        BotConfig
	option        Option
}
type Option struct {
	callback bool
	edit     bool
	update   bool
}
type Ds struct {
	mesid       string
	nameid      string
	guildid     string
	Attachments *discordgo.MessageAttachment
}

type Tg struct {
	mesid  int
	nameid int64
}

type TableConfig struct {
	id             int
	corpname       string
	dschannel      string
	tgchannel      int64
	wachannel      string
	mesiddshelp    string
	mesidtghelp    int
	delmescomplite int
	guildid        string
}
type sborkz struct {
	id          int
	corpname    string
	name        string
	mention     string
	tip         string
	dsmesid     string
	tgmesid     int
	wamesid     string
	time        string
	date        string
	lvlkz       string
	numkzn      int
	numberkz    int
	numberevent int
	eventpoints int
	active      int
	timedown    int
}

type Users struct {
	user1 sborkz
	user2 sborkz
	user3 sborkz
	user4 sborkz
}
type Names struct {
	name1 string
	name2 string
	name3 string
	name4 string
}
type emodjiUser struct {
	id                       int
	name, em1, em2, em3, em4 string
}
