package sendbird

import (
	"errors"
	"fmt"
)

type SendMessageRequest struct {
	ChannelURL  string `json:"-"`
	ChannelType string `json:"-"`
	UserID      string `json:"user_id"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
	CustomType  string `json:"custom_type,omitempty"`
	Data        string `json:"data,omitempty"`
	SendPush    bool   `json:"send_push,omitempty"`
	IsSilent    bool   `json:"is_silent,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
}

func (c *Client) SendTextMessage(r *SendMessageRequest) (TextMessage, error) {
	if r.ChannelURL == "" {
		return TextMessage{}, errors.New("send message: channel url missing")
	}

	if r.ChannelType == "" {
		return TextMessage{}, errors.New("send message: channel type missing")
	}

	if r.Message == "" {
		return TextMessage{}, errors.New("send message: message missing")
	}

	if r.UserID == "" {
		return TextMessage{}, errors.New("send message: user id missing")
	}

	if r.MessageType == "" {
		return TextMessage{}, errors.New("send message: message type missing")
	}

	result := TextMessage{}
	if err := c.postAndReturnJSON(c.PrepareUrl(fmt.Sprintf("/%s/%s/messages", r.ChannelType, r.ChannelURL)), r, &result); err != nil {
		return TextMessage{}, err
	}

	return result, nil
}
