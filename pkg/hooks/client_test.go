package hooks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	tgClient := &TGClient{
		token:   "5531851450:AAG16F3wWtYKapc07z6_uFM7T4sL4iPs1To",
		chat_id: "226619698",
	}
	msgSuccess := defaultSuccessTemplate("package", EscapeSpecialCharacters("1.0.0"), "\\- Changelog")
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
