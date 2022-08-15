package hooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/apex/log"
)

type TGClient struct {
	token   string
	chat_id string
}

type TGMessageResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
}

func (client *TGClient) getUrl() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", client.token)
}

func (client *TGClient) GetUpdates() {

}

func (client *TGClient) SendMessage(message RenderedMessage) (bool, error) {
	url := fmt.Sprintf("%s/sendMessage", client.getUrl())
	body, _ := json.Marshal(map[string]string{
		"chat_id":    client.chat_id,
		"text":       message.Message,
		"parse_mode": string(message.Format),
	})
	log.Infof("Request body: %s", body)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	log.Info("Message was sent")

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	var jsonBody TGMessageResponse
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		return false, err
	}
	log.Infof("Ok: %v", jsonBody.Ok)
	if !jsonBody.Ok {
		log.Infof("Description: %s, ErrorCode: %v", jsonBody.Description, jsonBody.ErrorCode)
	}
	return jsonBody.Ok, nil
}

func EscapeSpecialCharacters(original string) string {
	chars := []string{"_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
	replacers := make([]string, 0)
	for _, c := range chars {
		replacers = append(replacers, c)
		replacers = append(replacers, fmt.Sprintf("\\%s", c))
	}
	return strings.NewReplacer(replacers...).Replace(original)
}
