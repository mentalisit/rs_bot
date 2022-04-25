package NewBot

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow/appstate"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"mime"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

type wa struct {
	startedAt int64
}

var historySyncID int32
var startupTime = time.Now().Unix()
var users map[string]events.Contact
var waRuns = false

/*
func eventHandler(evt interface{}) {
	fmt.Println("					",evt)
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println( "				mtext", v.Message.GetConversation())
		fmt.Println( "				mte", v.Info.Sender)
		fmt.Println( "				отправитель ", v.Info.Sender.String())
	case *events.Receipt:
	}
}
*/
func eventHandler(rawEvt interface{}) {
	mtext := ""
	nameid := ""
	mention := ""
	name := ""

	switch evt := rawEvt.(type) {
	case *events.AppStateSyncComplete:
		if len(client.Store.PushName) > 0 && evt.Name == appstate.WAPatchCriticalBlock {
			err := client.SendPresence(types.PresenceAvailable)
			if err != nil {
				log.Warnf("Failed to send available presence: %v", err)
			} else {
				log.Infof("Marked self as available")
			}
		}
	case *events.Connected, *events.PushNameSetting:
		if len(client.Store.PushName) == 0 {
			return
		}
		// Send presence available when connecting and when the pushname is changed.
		// This makes sure that outgoing messages always have the right pushname.
		err := client.SendPresence(types.PresenceAvailable)
		if err != nil {
			log.Warnf("Failed to send available presence: %v", err)
		} else {
			log.Infof("Marked self as available")
		}
	case *events.StreamReplaced:
		os.Exit(0)
	case *events.Message:
		metaParts := []string{fmt.Sprintf("pushname: %s", evt.Info.PushName), fmt.Sprintf("timestamp: %s", evt.Info.Timestamp)}
		if evt.Info.Type != "" {
			metaParts = append(metaParts, fmt.Sprintf("type: %s", evt.Info.Type))
		}
		if evt.Info.Category != "" {
			metaParts = append(metaParts, fmt.Sprintf("category: %s", evt.Info.Category))
		}
		if evt.IsViewOnce {
			metaParts = append(metaParts, "view once")
		}
		if evt.IsViewOnce {
			metaParts = append(metaParts, "ephemeral")
		}

		//log.Infof("Received message %s from %s (%s): %+v", evt.Info.ID, evt.Info.SourceString(), strings.Join(metaParts, ", "), evt.Message)

		if evt.Info.Timestamp.Unix() < w.startedAt {
			waRuns = true
		}
		mtext = evt.Message.GetConversation()
		//fmt.Println("				75",metaParts)
		nameid = evt.Info.Sender.String()
		name = evt.Info.PushName

		//fmt.Println("			text	",*evt.RawMessage.ExtendedTextMessage.Text)
		if evt.Message.ExtendedTextMessage != nil {
			mtext = *evt.Message.ExtendedTextMessage.Text // Если кто то упоминает
			for _, men := range evt.RawMessage.ExtendedTextMessage.ContextInfo.MentionedJid {
				mention = men
			}
		}

		img := evt.Message.GetImageMessage()
		if img != nil {
			data, err := client.Download(img)
			if err != nil {
				log.Errorf("Failed to download image: %v", err)
				return
			}
			exts, _ := mime.ExtensionsByType(img.GetMimetype())
			path := fmt.Sprintf("%s%s", evt.Info.ID, exts[0])
			err = os.WriteFile(path, data, 0600)
			if err != nil {
				log.Errorf("Failed to save image: %v", err)
				return
			}
			log.Infof("Saved image in message to %s", path)

		}
	case *events.Receipt:
		if evt.Type == events.ReceiptTypeRead || evt.Type == events.ReceiptTypeReadSelf {
			//log.Infof("%v was read by %s at %s", evt.MessageIDs, evt.SourceString(), evt.Timestamp)
			//fmt.Println("		99			",evt.MessageIDs)//отчет о доставке те сообщения которые прочитали
			//fmt.Println("		100			",evt.SourceString())//кто и где прочитал
			//fmt.Println("					",evt.MessageIDs)
		} else if evt.Type == events.ReceiptTypeDelivered {
			//log.Infof("%s was delivered to %s at %s", evt.MessageIDs[0], evt.SourceString(), evt.Timestamp)
			//fmt.Println("			104		",evt.MessageIDs[0])//ид сообщения
			//fmt.Println("			105		",evt.SourceString())//
		}

	case *events.Presence:
		if evt.Unavailable {
			if evt.LastSeen.IsZero() {
				//log.Infof("%s is now offline", evt.From)
				fmt.Println("					112", evt.From)
			} else {
				//log.Infof("%s is now offline (last seen: %s)", evt.From, evt.LastSeen)
				fmt.Println("					115", evt.From, evt.LastSeen)
			}
		} else {
			log.Infof("%s is now online", evt.From)
		}
	case *events.HistorySync:
		id := atomic.AddInt32(&historySyncID, 1)
		fileName := fmt.Sprintf("history-%d-%d.json", startupTime, id)
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Errorf("Failed to open file to write history sync: %v", err)
			return
		}
		enc := json.NewEncoder(file)
		enc.SetIndent("", "  ")
		err = enc.Encode(evt.Data)
		if err != nil {
			log.Errorf("Failed to write history sync: %v", err)
			return
		}
		log.Infof("Wrote history sync to %s", fileName)
		_ = file.Close()
	case *events.AppState:
		log.Debugf("App state event: %+v / %+v", evt.Index, evt.SyncActionValue)
	}
	if nameid != "" && mtext != "" || waRuns {
		if name == "" {
			name = "Кто-то"
		}

		numberAndSuffix := strings.SplitN(mtext, "@", 2)
		fmt.Println(155, numberAndSuffix)
		mentions := getSenderNotify(numberAndSuffix[0]) // + "@s.whatsapp.net")
		fmt.Println(159, mentions)
		mtext = strings.Replace(mtext, "@"+numberAndSuffix[0], "@"+mention, 1)

		fmt.Printf(" отправитель:%s,\n id: %s,\n ментионджид:%s, \n Текст сообщения: %s,\n", name, nameid, mention, mtext)
	}
}
func getSenderNotify(senderJid string) string {
	if sender, exists := users[senderJid]; exists {
		fmt.Println(168, sender.Action.String())
		return sender.Action.String()
	}

	return ""
}
