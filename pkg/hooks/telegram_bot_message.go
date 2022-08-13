package hooks

import (
	"strings"
)

type MessageFmt string

const (
	MARKDOWN MessageFmt = "markdown"
	HTLM     MessageFmt = "html"
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
	message := DefaultSuccessTemplate(packageName, newVersion, changelogs)
	return RenderedMessage{
		Message: message.RawMessage,
		Format:  message.Format,
	}
}

func (msg *Message) FailMessage(packageName, reason, errMsg string) RenderedMessage {
	renderedMessage := DefaultFailTemplate(packageName, reason, errMsg)
	return RenderedMessage{
		Message: renderedMessage.RawMessage,
		Format:  renderedMessage.Format,
	}
}
