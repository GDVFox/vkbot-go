package vkbotgo

import (
	"github.com/google/go-querystring/query"
)

const (
	//ChatBase start base for chat IDs
	ChatBase = 2000000000
)

// Message is a struct describing a private message.
type Message struct {
	// ID message ID.
	ID int64 `json:"id"`
	// Date when the message has been sent in Unixtime.
	Date int64 `json:"date"`
	// PeerID peer ID.
	PeerID int64 `json:"peer_id"`
	// Message author's ID.
	FromID int64 `json:"from_id"`
	// Text message text.
	Text string `json:"text"`
	// RandomID ID used for sending messages. It returned only for outgoing messages.
	RandomID int64 `json:"random_id"`
	// Important is it an important message.
	Important bool `json:"important"`
	// Payload service variable message bots.
	Payload string `json:"payload"`
}

// MessagesSendParams params for message.send method
//
// https://vk.com/dev/messages.send
type MessagesSendParams struct {
	// User ID (by default — current user).
	UserID int64 `url:"user_id,omitempty"`
	// Unique identifier to avoid resending the message.
	RandomID int64 `url:"random_id"`
	// Destination ID. \"For user: 'User ID', e.g. '12345'.
	// For chat: '2000000000' + 'chat_id', e.g. '2000000001'.
	// For community: '- community ID', e.g. '-12345'. \"
	PeerID int64 `url:"peer_id"`
	// User's short address (for example, 'gdv_fox').
	Domain string `url:"domain,omitempty"`
	// ID of conversation the message will relate to.
	ChatID int64 `url:"chat_id,omitempty"`
	// IDs of message recipients (if new conversation shall be started).
	UserIDs []int64 `url:"user_ids,omitempty"`
	// (Required if 'attachments' is not set.) Text of the message.
	Message string `url:"message"`
	// Geographical latitude of a check-in, in degrees (from -90 to 90).
	Latitude float64 `url:"lat,omitempty"`
	// Geographical longitude of a check-in, in degrees (from -180 to 180).
	Longitude float64 `url:"long,omitempty"`
	// (Required if 'message' is not set.) List of objects attached to the message, separated by commas,
	// in the following format: \"<owner_id>_<media_id>\", '' — Type of media attachment: 'photo' — photo,
	// 'video' — video, 'audio' — audio, 'doc' — document, 'wall' — wall post, '<owner_id>' — ID of the media attachment owner.
	// '<media_id>' — media attachment ID. Example: \"photo100172_166443618\"
	Attachments []string `url:"attachment,omitempty"`
	// ID of forwarded messages, separated with a comma. Listed messages of the sender will be shown
	// in the message body at the recipient's. Example: \"123,431,544\"
	ForwardMessages string `url:"forward_messages,omitempty"`
	// Sticker id.
	StickerID int64 `url:"sticker_id,omitempty"`
	// '1' if the message is a notification (for community messages).
	Notification bool `url:"notification,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// MessagesSend sends a message.
func (vkBot *VkBot) MessagesSend(param *MessagesSendParams) (*VkResponse, error) {
	vals, err := query.Values(param)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.send", vals)
}
