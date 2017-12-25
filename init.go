package main

import (
	"flag"
	// "fmt"

	log "github.com/kirillDanshin/dlog"
	"github.com/olebedev/config"
	"github.com/toby3d/patreon"
)

var (
	cfg *config.Config

	flagWebhook = flag.Bool("webhook", false, "enable webhook mode")
)

func errCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func init() {
	var err error
	log.Ln("Parse yaml config...")
	cfg, err = config.ParseYamlFile("config.yaml")
	errCheck(err)

	log.Ln("Check telegram token...")
	_, err = cfg.String("telegram.token")
	errCheck(err)

	if *flagWebhook {
		log.Ln("Check telegram webhook set...")
		_, err = cfg.String("telegram.webhook.set")
		errCheck(err)

		log.Ln("Check telegram webhook listen...")
		_, err = cfg.String("telegram.webhook.listen")
		errCheck(err)

		log.Ln("Check telegram webhook serve...")
		_, err = cfg.String("telegram.webhook.serve")
		errCheck(err)
	}

	log.Ln("Check patreon client_id...")
	_, err = cfg.String("patreon.client_id")
	errCheck(err)

	log.Ln("Check patreon client_secret...")
	_, err = cfg.String("patreon.client_secret")
	errCheck(err)

	patreonClient = patreon.NewClient(
		cfg.UString("patreon.client_id"),
		cfg.UString("patreon.client_secret"),
		"https://toby3d.github.io/PatreonRobot/oauth.html",
		patreon.ScopeUsers, patreon.ScopePledgesToMe, patreon.ScopeMyCampaign,
	)
}
