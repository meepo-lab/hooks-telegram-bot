package hooks

import (
	"fmt"
	"regexp"

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
	re := regexp.MustCompile("token:\\S+\\w")
	initMap := fmt.Sprintf("Init %v", m)
	log.Infof(re.ReplaceAllString(initMap, "token:****"))
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
	if _, err := bot.client.SendMessage(bot.message.SuccessMessage(
		bot.repo,
		newVersion,
		config.Changelog)); err != nil {
		return err
	}
	return nil
}

func (bot *TelegramBot) NoRelease(config *hooks.NoReleaseConfig) error {
	log.Infof("reason: " + config.Reason.String())
	log.Infof("message: " + config.Message)
	if _, err := bot.client.SendMessage(bot.message.FailMessage(
		bot.repo,
		config.Reason.String(),
		config.Message)); err != nil {
		return err
	}
	return nil
}
