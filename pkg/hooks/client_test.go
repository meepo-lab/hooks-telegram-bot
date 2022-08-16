package hooks

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	tgClient := &TGClient{
		token:   os.Getenv("TOKEN"),
		chat_id: os.Getenv("CHAT_ID"),
	}
	msgSuccess := defaultSuccessTemplate(
		"package",
		EscapeSpecialCharacters("1.0.0"),
		EscapeSpecialCharacters("- Changelog"))
	renderedMsg := RenderedMessage{
		Message: msgSuccess.RawMessage,
		Format:  msgSuccess.Format,
	}

	ok, err := tgClient.SendMessage(renderedMsg)

	require.NoError(t, err)
	require.Equal(t, true, ok)

	msgError := defaultFailTemplate("package", "No change", "error")
	renderedMsgFail := RenderedMessage{
		Message: msgError.RawMessage,
		Format:  msgError.Format,
	}

	ok2, err2 := tgClient.SendMessage(renderedMsgFail)

	require.NoError(t, err2)
	require.Equal(t, true, ok2)
}

func TestEscapeSpecialCharacters(t *testing.T) {
	version := "v1.0.0"
	newVersion := EscapeSpecialCharacters(version)
	require.Equal(t, "v1\\.0\\.0", newVersion)
}
