package sendbird

import (
	"errors"
	"net/url"

	"github.com/woodstock-tokyo/sendbird/templates"
)

type baseMessage struct {
	MessageID       int64             `json:"message_id"`
	Type            string            `json:"type"`
	CustomType      string            `json:"custom_type"`
	ChannelURL      string            `json:"channel_url"`
	MentionType     string            `json:"mention_type"`
	MentionedUsers  []User            `json:"mentioned_users"`
	IsRemoved       bool              `json:"is_removed"`
	Message         string            `json:"message"`
	SortedMetaArray []SortedMetaArray `json:"sorted_metaarray"`
	MessageEvents   MessageEvents     `json:"message_events"`
	CreatedAt       int64             `json:"created_at"`
	UpdatedAt       int64             `json:"updated_at"`
}

type TextMessage struct {
	baseMessage

	User                 User              `json:"user"`
	Translations         Translations      `json:"translations"`
	Data                 string            `json:"data"`
	OGTag                OGTag             `json:"og_tag"`
	File                 File              `json:"file"`
	IsAppleCriticalAlert bool              `json:"is_apple_critical_alert"`
	ThreadInfo           ThreadInfo        `json:"thread_info"`
	ParentMessageID      int64             `json:"parent_message_id"`
	ParentMessageInfo    ParentMessageInfo `json:"parent_message_info"`
	IsReplyToChannel     bool              `json:"is_reply_to_channel"`
	Reactions            []Reaction        `json:"reactions"`

	// for bot
	Text string `json:"text"`
}

type FileMessage struct {
	baseMessage

	User              User              `json:"user"`
	File              File              `json:"file"`
	Thumbnails        []string          `json:"thumbnails"`
	RequireAuth       bool              `json:"require_auth"`
	ThreadInfo        ThreadInfo        `json:"thread_info"`
	ParentMessageID   int64             `json:"parent_message_id"`
	ParentMessageInfo ParentMessageInfo `json:"parent_message_info"`
	IsReplyToChannel  bool              `json:"is_reply_to_channel"`
	Reactions         []Reaction        `json:"reactions"`
}

type AdminMessage struct {
	baseMessage

	Data  string `json:"data"`
	OGTag OGTag  `json:"og_tag"`
}

type BotMessage struct {
	Message baseMessage `json:"message"`
}

type Translations struct {
}

type ParentMessageInfo struct {
}

type ThreadInfo struct {
	ReplyCount    int64  `json:"reply_count"`
	MostReplies   []User `json:"most_replies"`
	LastRepliedAt int64  `json:"last_replied_at"`
	UpdatedAt     int64  `json:"updated_at"`
}

type OGTag struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type SortedMetaArray struct {
	Concept  string   `json:"concept"`
	Priority []string `json:"priority"`
}

type MessageEvents struct {
	SendPushNotification string `json:"send_push_notification"`
	UpdateUnreadCount    bool   `json:"update_unread_count"`
	UpdateLastMesage     bool   `json:"update_last_message"`
}

type Reaction struct {
	Key       string   `json:"key"`
	UserIDs   []string `json:"user_ids"`
	UpdatedAt int64    `json:"updated_at"`
}

type File struct {
	URL  string `json:"url"`
	Name string `json:"name"`
	Type string `json:"type"`
	Size int64  `json:"size"`
	Data string `json:"data"`
}

type getMessagesTemplateData struct {
	ChannelType string
	ChannelURL  string
	MessageID   int64
}

type SendTextMessageRequest struct {
	UserID      string `json:"user_id"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
	CustomType  string `json:"custom_type,omitempty"`
	Data        string `json:"data,omitempty"`
	SendPush    bool   `json:"send_push,omitempty"`
	IsSilent    bool   `json:"is_silent,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
}

func (c *Client) SendTextMessage(channelType string, channelURL string, r *SendTextMessageRequest) (TextMessage, error) {
	if r.Message == "" {
		return TextMessage{}, errors.New("send message: message missing")
	}

	if r.UserID == "" {
		return TextMessage{}, errors.New("send message: user id missing")
	}

	if r.MessageType == "" {
		return TextMessage{}, errors.New("send message: message type missing")
	}

	pathString, err := templates.GetMessagesTemplate(getMessagesTemplateData{ChannelType: url.PathEscape(channelType), ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLMessagesWithChannelTypeAndChannelURL)

	if err != nil {
		return TextMessage{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := TextMessage{}
	if err := c.postAndReturnJSON(parsedURL, r, &result); err != nil {
		return TextMessage{}, err
	}

	return result, nil
}

type GetAMessageRequest struct {
	IncludeReactions bool `json:"include_reactions"`
}

type GetAMessageResp struct {
	MessageID            int64             `json:"message_id"`
	Type                 string            `json:"type"`
	CustomType           string            `json:"custom_type"`
	ChannelURL           string            `json:"channel_url"`
	MentionType          string            `json:"mention_type"`
	MentionedUsers       []User            `json:"mentioned_users"`
	IsRemoved            bool              `json:"is_removed"`
	Message              string            `json:"message"`
	SortedMetaArray      []SortedMetaArray `json:"sorted_metaarray"`
	MessageEvents        MessageEvents     `json:"message_events"`
	CreatedAt            int64             `json:"created_at"`
	UpdatedAt            int64             `json:"updated_at"`
	User                 User              `json:"user"`
	Translations         string            `json:"translations"`
	Data                 string            `json:"data"`
	OGTag                OGTag             `json:"og_tag"`
	File                 File              `json:"file"`
	IsAppleCriticalAlert bool              `json:"is_apple_critical_alert"`
	ThreadInfo           ThreadInfo        `json:"thread_info"`
	ParentMessageID      int64             `json:"parent_message_id"`
	ParentMessageInfo    ParentMessageInfo `json:"parent_message_info"`
	IsReplyToChannel     bool              `json:"is_reply_to_channel"`
	Reactions            []Reaction        `json:"reactions"`
	Thumbnails           []string          `json:"thumbnails"`
	RequireAuth          bool              `json:"require_auth"`
}

func (r *GetAMessageRequest) params() url.Values {
	q := make(url.Values)

	if r.IncludeReactions {
		q.Set("include_reactions", "true")
	}

	return q
}

func (c *Client) GetAMessage(channelType string, channelURL string, messageId int64, r *GetAMessageRequest) (GetAMessageResp, error) {
	pathString, err := templates.GetMessagesTemplate(getMessagesTemplateData{ChannelType: url.PathEscape(channelType), ChannelURL: url.PathEscape(channelURL), MessageID: messageId}, templates.SendbirdURLMessagesWithChannelTypeAndChannelURLAndMessageID)

	if err != nil {
		return GetAMessageResp{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	raw := r.params().Encode()
	result := GetAMessageResp{}
	if err := c.getAndReturnJSON(parsedURL, raw, &result); err != nil {
		return GetAMessageResp{}, err
	}

	return result, nil
}
