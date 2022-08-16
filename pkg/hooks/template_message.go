package hooks

import (
	"fmt"
)

func defaultFailTemplate(packageName, reason, errMsg string) *Message {
	var message = fmt.Sprintf("🚧 Something went wrong while trying to publish the new version of *%s*\\!", packageName)
	message += fmt.Sprintf("\nReason: %s", reason)
	if len(errMsg) > 0 {
		message += fmt.Sprintf("\nError Message: %s", errMsg)
	}
	return &Message{
		RawMessage: message,
		Format:     MARKDOWN,
	}
}

func defaultSuccessTemplate(packageName, newVersion, changelogs string) *Message {
	var message = fmt.Sprintf("🎉 *%s* *v%s* has been released\\!", packageName, newVersion)
	if len(changelogs) > 0 {
		message += fmt.Sprintf("\n%s", changelogs)
	}
	return &Message{
		RawMessage: message,
		Format:     MARKDOWN,
	}
}

type MessageTemplate struct {
	Path       string
	CustomData map[string]string
}
