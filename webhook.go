package viber

import (
	"encoding/json"
	"fmt"
)

// WebhookReq request
type WebhookReq struct {
	URL        string   `json:"url"`
	EventTypes []string `json:"event_types"`
}

//WebhookResp response
type WebhookResp struct {
	Status        int      `json:"status"`
	StatusMessage string   `json:"status_message"`
	EventTypes    []string `json:"event_types,omitempty"`
}

// WebhookVerify response
type WebhookVerify struct {
	Event        string `json:"event"`
	Timestamp    uint64 `json:"timestamp"`
	MessageToken uint64 `json:"message_token"`
}

// SetWebhook for Viber callbacks
// if eventTypes is nil, all callbacks will be set to webhook
// if eventTypes is empty []string mandatory callbacks will be set
// Mandatory callbacks: "message", "subscribed", "unsubscribed"
// All possible callbacks: "message", "subscribed",  "unsubscribed", "delivered", "seen", "failed", "conversation_started"
func (v *Viber) SetWebhook(url string, eventTypes []string) (WebhookResp, error) {
	var resp WebhookResp

	req := WebhookReq{
		URL:        url,
		EventTypes: eventTypes,
	}
	r, err := v.PostData(fmt.Sprintf("%s/set_webhook", ViberAPI), req)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(r, &resp)
	return resp, err
}

// RemoveWebhook for Viber
// Once you set a webhook to your Public Account/ bot your 1-on-1 conversation button will appear and users will be able to access it.
// At the moment there is no option to disable the 1-on-1 conversation from the Public Account
// bot settings, so to disable this option you’ll need to remove the webhook you set for the account.
// Removing the webhook is done by Posting a set_webhook request with an empty webhook string.
func (v *Viber) RemoveWebhook() (resp WebhookResp, err error) {
	var req = WebhookReq{
		URL: "",
	}
	r, err := v.PostData(fmt.Sprintf("%s/set_webhook", ViberAPI), req)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(r, &resp)
	return resp, err
}
