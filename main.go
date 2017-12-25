package main

import (
	log "github.com/kirillDanshin/dlog"
	tg "github.com/toby3d/telegram"
)

var bot *tg.Bot

func main() {
	var err error
	bot, err = tg.NewBot(cfg.UString("telegram.token"))
	errCheck(err)
	log.Ln("Authorized as", bot.Self.Username)

	for update := range bot.NewLongPollingChannel(nil) {
		switch {
		case update.Message != nil:
			if bot.IsMessageFromMe(update.Message) {
				continue
			}

			message(update.Message)
		default:
			continue
		}
	}
}
