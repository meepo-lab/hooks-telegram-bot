package hooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/apex/log"
)

type TGClient struct {
	token   string
	chat_id string
}

type TGMessageResponse struct {
	Ok     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
}

func (client *TGClient) getUrl() string {
	return fmt.Sprintf("https://api.telegram.org/client%s", client.token)
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
	if err := json.Unmarshal(body, &response); err != nil {
		return false, err
	}
	log.Infof("Response: %v", jsonBody)

	return jsonBody.Ok, nil
}
