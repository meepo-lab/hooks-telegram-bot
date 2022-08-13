package hooks

import (
	"github.com/apex/log"
	"github.com/ted-vo/semantic-release/v3/pkg/hooks"
)

var NAME = "Telegram Bot"
var FUVERSION = "dev"

type TelegramBot struct {
	repo    string
	message *Message
	client  *TGClient
}

func (bot *TelegramBot) Init(m map[string]string) error {
	log.Infof("Init %v, m")
	projectName := m["project_name"]
	token := m["token"]
	chat_id := m["chat_id"]
	rawMessage := m["message"]
	format := m["format"]
	// customData := m["CustomData"]

	bot.repo = projectName
	bot.client = &TGClient{
		token:   token,
		chat_id: chat_id,
	}

	bot.message = &Message{
		RawMessage: rawMessage,
		Format:     ParseMessageFmt(format),
		CustomData: map[string]string{},
	}

	return nil
}

func (bot *TelegramBot) Name() string {
	return NAME
}

func (bot *TelegramBot) Version() string {
	return FUVERSION
}

func (bot *TelegramBot) Success(config *hooks.SuccessHookConfig) error {
	oldVersion := config.PrevRelease.Version
	newVersion := config.NewRelease.Version
	log.Infof("old version: " + oldVersion)
	log.Infof("new Version: " + newVersion)
	bot.client.SendMessage(bot.message.SuccessMessage(
		bot.repo,
		newVersion,
		config.Changelog))
	return nil
}

func (bot *TelegramBot) NoRelease(config *hooks.NoReleaseConfig) error {
	log.Infof("reason: " + config.Reason.String())
	log.Infof("message: " + config.Message)
	bot.client.SendMessage(bot.message.FailMessage(
		bot.repo,
		config.Reason.String(),
		config.Message))
	return nil
}
