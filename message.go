package main

import (
	"fmt"
	// "strings"

	log "github.com/kirillDanshin/dlog"
	"github.com/toby3d/patreon"
	tg "github.com/toby3d/telegram"
	// "golang.org/x/oauth2"
)

const (
	cmdStart = "start"
	cmdHelp  = "help"
)

var patreonClient *patreon.Client

func message(msg *tg.Message) {
	if !msg.IsCommand() {
		return
	}

	switch msg.Command() {
	case cmdStart:
		if !msg.HasArgument() {
			return
		}

		log.Ln("Token:", "["+msg.CommandArgument()+"]")

		token, err := patreonClient.Exchange(msg.CommandArgument())
		errCheck(err)

		reply := tg.NewMessage(msg.Chat.ID, fmt.Sprintf("%#v", token))
		_, err = bot.SendMessage(reply)
		errCheck(err)
	case cmdHelp:
		reply := tg.NewMessage(msg.Chat.ID, "Authorize first")
		reply.ReplyMarkup = tg.NewInlineKeyboardMarkup(
			tg.NewInlineKeyboardRow(
				tg.NewInlineKeyboardButtonURL(
					"Login via Patreon",
					patreonClient.AuthCodeURL("state"),
				),
			),
		)

		_, err := bot.SendMessage(reply)
		errCheck(err)
	}
}

/*
func command(msg *tg.Message) {
	switch strings.ToLower(msg.Command()) {
	case cmdStart:
		commandStart(msg)
	case cmdHelp:
		commandHelp(msg)
	}
}

func commandStart(msg *tg.Message) {
	if msg.HasArgument() {
		commandStart(msg)
		return
	}
}

func commandStartWithArgument(msg *tg.Message) {
	msg.CommandArgument()

	NewPatreonClient(ctx, token)
}

func commandHelp(msg *tg.Message) {

}
*/
