package hooks

import (
	"strings"
)

type MessageFmt string

const (
	MARKDOWN MessageFmt = "MarkdownV2"
	HTLM     MessageFmt = "HTML"
)

func ParseMessageFmt(s string) MessageFmt {
	switch strings.ToLower(s) {
	case "markdown":
		return MARKDOWN
	case "html":
		return HTLM
	}
	return MARKDOWN
}

type MessageConfig struct {
	RawMessage string
	Format     MessageFmt
	CustomData map[string]string
}

type Message struct {
	RawMessage string
	Format     MessageFmt
	CustomData map[string]string
}

type RenderedMessage struct {
	Message string
	Format  MessageFmt
}

// func (msg *Message) renderMessage(success bool) RenderedMessage {
// return RenderedMessage{
// Message: msg.RawMessage,
// Format:  msg.Format,
// }
// }

func (msg *Message) SuccessMessage(packageName, newVersion, changelogs string) RenderedMessage {
	if len(msg.RawMessage) == 0 {
		defaultMsg := defaultSuccessTemplate(
			EscapeSpecialCharacters(packageName),
			EscapeSpecialCharacters(newVersion),
			EscapeSpecialCharacters(changelogs))

		msg.RawMessage = defaultMsg.RawMessage
		msg.Format = defaultMsg.Format
	}

	return RenderedMessage{
		Message: msg.RawMessage,
		Format:  msg.Format,
	}
}

func (msg *Message) FailMessage(packageName, reason, errMsg string) RenderedMessage {
	if len(msg.RawMessage) == 0 {
		defaultMsg := defaultFailTemplate(
			EscapeSpecialCharacters(packageName),
			EscapeSpecialCharacters(reason),
			EscapeSpecialCharacters(errMsg))

		msg.RawMessage = defaultMsg.RawMessage
		msg.Format = defaultMsg.Format
	}

	return RenderedMessage{
		Message: msg.RawMessage,
		Format:  msg.Format,
	}
}
