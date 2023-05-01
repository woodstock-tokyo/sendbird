package sendbird

import (
	"errors"
	"fmt"
)

type SendBotMessageRequest struct {
	BotID        string   `json:"bot_userid"`
	Message      string   `json:"message"`
	ChannelURL   string   `json:"channel_url"`
	CustomerType string   `json:"custom_type,omitempty"`
	Data         string   `json:"data,omitempty"`
	SendPush     bool     `json:"send_push,omitempty"`
	Mentioned    []string `json:"mentioned,omitempty"`
	MarkAsRead   bool     `json:"mark_as_read,omitempty"`
	DedUpID      string   `json:"dedup_id,omitempty"`
	CreateAt     int64    `json:"created_at,omitempty"`
}

func (c *Client) SendBotMessage(r *SendBotMessageRequest) (TextMessage, error) {
	if r.Message == "" {
		return TextMessage{}, errors.New("send bot message: bot message missing")
	}

	if r.BotID == "" {
		return TextMessage{}, errors.New("send bot message: bot id missing")
	}

	if r.ChannelURL == "" {
		return TextMessage{}, errors.New("send bot message: channel url missing")
	}

	result := TextMessage{}
	if err := c.postAndReturnJSON(c.PrepareUrl(fmt.Sprintf("%s/%s/send", SendbirdURLBots, r.BotID)), r, &result); err != nil {
		return TextMessage{}, err
	}

	return result, nil
}
