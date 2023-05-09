package sendbird

import (
	"errors"
	"fmt"
)

type BotResponse struct {
	Bot
	AI
}

type Bot struct {
	BotUserID                   string            `json:"bot_userid"`
	BotNickname                 string            `json:"bot_nickname"`
	BotProfileURL               string            `json:"bot_profile_url"`
	BotType                     string            `json:"bot_type"`
	BotToken                    string            `json:"bot_token"`
	BotMetadata                 map[string]string `json:"bot_metadata"`
	BotCallbackURL              string            `json:"bot_callback_url"`
	IsPrivacyMode               bool              `json:"is_privacy_mode"`
	EnableMarkAsRead            bool              `json:"enable_mark_as_read"`
	ShowMember                  bool              `json:"show_member"`
	ChannelInvitationPreference int               `json:"channel_invitation_preference"`
	CreatedAt                   int64             `json:"created_at"`
}

type AI struct {
	Backend          string  `json:"backend"`
	SystemMessage    string  `json:"system_message"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             float64 `json:"top_p"`
	PresencePenalty  float64 `json:"presence_penalty"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
}

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

func (c *Client) SendBotMessage(r *SendBotMessageRequest) (BotMessage, error) {
	if r.Message == "" {
		return BotMessage{}, errors.New("send bot message: bot message missing")
	}

	if r.BotID == "" {
		return BotMessage{}, errors.New("send bot message: bot id missing")
	}

	if r.ChannelURL == "" {
		return BotMessage{}, errors.New("send bot message: channel url missing")
	}

	result := BotMessage{}
	if err := c.postAndReturnJSON(c.PrepareUrl(fmt.Sprintf("%s/%s/send", SendbirdURLBots, r.BotID)), r, &result); err != nil {
		return BotMessage{}, err
	}

	return result, nil
}
