package sendbird_test

import (
	"fmt"
	"testing"

	. "github.com/woodstock-tokyo/sendbird"
)

var testMessagesClient = NewTestClient()

func TestMessagesActions(t *testing.T) {
	testSendTextMessage(t, "group_channels", testChannelURL, testUserID, fmt.Sprintf("test from %s", testUserID))
}

func TestGetAMessageActions(t *testing.T) {
	textMessage := testSendTextMessage(t, "group_channels", testChannelURL, testUserID, fmt.Sprintf("test from %s", testUserID))
	testGetAMessage(t, "group_channels", testChannelURL, textMessage.MessageID, true)
}

func testSendTextMessage(t *testing.T, channelType string, channelURL string, userID string, message string) TextMessage {
	r := &SendTextMessageRequest{
		UserID:      userID,
		Message:     message,
		MessageType: "MESG",
		IsSilent:    false,
	}

	result, err := testMessagesClient.SendTextMessage(channelType, channelURL, r)
	if err != nil {
		t.Errorf("Fail in SendTextMessage(): %+v", err)
	}

	t.Logf("SendTextMessage() Result: %+v", result)

	return result
}

func testGetAMessage(t *testing.T, channelType string, channelURL string, messageID int64, includeReactions bool) {
	r := &GetAMessageRequest{
		IncludeReactions: true,
	}

	result, err := testMessagesClient.GetAMessage(channelType, channelURL, messageID, r)
	if err != nil {
		t.Errorf("Fail in GetAMessage(): %+v", err)
	}

	t.Logf("GetAMessage() Result: %+v", result)
}
