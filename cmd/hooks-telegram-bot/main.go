package main

import (
	"github.com/apex/log"
	hooksTgBot "github.com/meepo-lab/hooks-telegram-bot/pkg/hooks"
	"github.com/ted-vo/semantic-release/v3/pkg/hooks"
	"github.com/ted-vo/semantic-release/v3/pkg/plugin"
)

func main() {
	log.SetHandler(hooksTgBot.NewLogHandler())
	plugin.Serve(&plugin.ServeOpts{
		Hooks: func() hooks.Hooks {
			return &hooksTgBot.TelegramBot{}
		},
	})
}
