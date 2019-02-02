package vkbotgo

// GENERATED FILE: DO NOT EDIT

import (
	"fmt"

	"github.com/google/go-querystring/query"
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

// MessagesGetConversationsParams params for messages.getConversations method
//
// https://vk.com/dev/messages.getConversations
type MessagesGetConversationsParams struct {

	// Group ID (for group messages with group access token)
	GroupID int64 `url:"group_id,omitempty"`
	// Offset needed to return a specific subset of conversations.
	Offset int64 `url:"offset,omitempty"`
	// Number of conversations to return.
	Count int64 `url:"count,omitempty"`
	// Filter to apply: &#39;all&#39; — all conversations, &#39;unread&#39; — conversations with unread messages, &#39;important&#39; — conversations, marked as important (only for community messages), &#39;unanswered&#39; — conversations, marked as unanswered (only for community messages)
	Filter string `url:"filter,omitempty"`
	// &#39;1&#39; — return extra information about users and communities
	Extended bool `url:"extended,omitempty"`
	// ID of the message from what to return dialogs.
	StartMessageID int64 `url:"start_message_id,omitempty"`
	// Profile and communities fields to return.
	Fields []string `url:"fields,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesGetConversationsParams) Validate() error {
	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	if param.Count < 0.000000 {
		return fmt.Errorf("parameter Count(=%v) must be greater or equal than 0.000000", param.Count)
	}

	if param.Count > 200.000000 {
		return fmt.Errorf("parameter Count(=%v) must be less or equal than 200.000000", param.Count)
	}

	enum := map[string]struct{}{"all": {}, "unread": {}, "important": {}, "unanswered": {}}
	if _, ok := enum[param.Filter]; !ok {
		return fmt.Errorf("parameter Filter(=%v) expected in [all unread important unanswered]", param.Filter)
	}

	if param.StartMessageID < 0.000000 {
		return fmt.Errorf("parameter StartMessageID(=%v) must be greater or equal than 0.000000", param.StartMessageID)
	}

	return nil
}

// MessagesGetConversations calls VK API method messages.getConversations. Returns a list of the current user&#39;s conversations.
//
// https://vk.com/dev/messages.getConversations
func (vkBot *VkBot) MessagesGetConversations(params *MessagesGetConversationsParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.getConversations", vals)
}

// MessagesGetConversationsByIdParams params for messages.getConversationsById method
//
// https://vk.com/dev/messages.getConversationsById
type MessagesGetConversationsByIdParams struct {

	// Destination IDs. &#34;For user: &#39;User ID&#39;, e.g. &#39;12345&#39;. For chat: &#39;2000000000&#39; &#43; &#39;chat_id&#39;, e.g. &#39;2000000001&#39;. For community: &#39;- community ID&#39;, e.g. &#39;-12345&#39;. &#34;
	PeerIds []int64 `url:"peer_ids"`
	// Return extended properties
	Extended bool `url:"extended,omitempty"`
	// Profile and communities fields to return.
	Fields []string `url:"fields,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesGetConversationsByIdParams) Validate() error {
	if len(param.PeerIds) > 100 {
		return fmt.Errorf("length(=%d) of parameter PeerIds must be less or equal than 100", len(param.PeerIds))
	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesGetConversationsById calls VK API method messages.getConversationsById. Returns conversations by their IDs
//
// https://vk.com/dev/messages.getConversationsById
func (vkBot *VkBot) MessagesGetConversationsById(params *MessagesGetConversationsByIdParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.getConversationsById", vals)
}

// MessagesGetByIdParams params for messages.getById method
//
// https://vk.com/dev/messages.getById
type MessagesGetByIdParams struct {

	// Message IDs.
	MessageIds []int64 `url:"message_ids"`
	// Number of characters after which to truncate a previewed message. To preview the full message, specify &#39;0&#39;. &#34;NOTE: Messages are not truncated by default. Messages are truncated by words.&#34;
	PreviewLength int64 `url:"preview_length,omitempty"`
	// Information whether the response should be extended
	Extended bool `url:"extended,omitempty"`
	// Profile fields to return.
	Fields []string `url:"fields,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesGetByIdParams) Validate() error {
	if len(param.MessageIds) > 100 {
		return fmt.Errorf("length(=%d) of parameter MessageIds must be less or equal than 100", len(param.MessageIds))
	}

	for i, item := range param.MessageIds {
		if item < 0.000000 {
			return fmt.Errorf("item[#%d](=%v) of parameter MessageIds must be greater or equal than 0.000000", i, item)
		}

	}

	if param.PreviewLength < 0.000000 {
		return fmt.Errorf("parameter PreviewLength(=%v) must be greater or equal than 0.000000", param.PreviewLength)
	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesGetById calls VK API method messages.getById. Returns messages by their IDs.
//
// https://vk.com/dev/messages.getById
func (vkBot *VkBot) MessagesGetById(params *MessagesGetByIdParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.getById", vals)
}

// MessagesGetByConversationMessageIdParams params for messages.getByConversationMessageId method
//
// https://vk.com/dev/messages.getByConversationMessageId
type MessagesGetByConversationMessageIdParams struct {

	// Destination ID. &#34;For user: &#39;User ID&#39;, e.g. &#39;12345&#39;. For chat: &#39;2000000000&#39; &#43; &#39;chat_id&#39;, e.g. &#39;2000000001&#39;. For community: &#39;- community ID&#39;, e.g. &#39;-12345&#39;. &#34;
	PeerID int64 `url:"peer_id,omitempty"`
	// Conversation message IDs.
	ConversationMessageIds []int64 `url:"conversation_message_ids"`
	// Information whether the response should be extended
	Extended bool `url:"extended,omitempty"`
	// Profile fields to return.
	Fields []string `url:"fields,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesGetByConversationMessageIdParams) Validate() error {
	if len(param.ConversationMessageIds) > 100 {
		return fmt.Errorf("length(=%d) of parameter ConversationMessageIds must be less or equal than 100", len(param.ConversationMessageIds))
	}

	for i, item := range param.ConversationMessageIds {
		if item < 0.000000 {
			return fmt.Errorf("item[#%d](=%v) of parameter ConversationMessageIds must be greater or equal than 0.000000", i, item)
		}

	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesGetByConversationMessageId calls VK API method messages.getByConversationMessageId. Returns messages by their IDs within the conversation.
//
// https://vk.com/dev/messages.getByConversationMessageId
func (vkBot *VkBot) MessagesGetByConversationMessageId(params *MessagesGetByConversationMessageIdParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.getByConversationMessageId", vals)
}

// MessagesSearchParams params for messages.search method
//
// https://vk.com/dev/messages.search
type MessagesSearchParams struct {

	// Search query string.
	Q string `url:"q,omitempty"`
	// Destination ID. &#34;For user: &#39;User ID&#39;, e.g. &#39;12345&#39;. For chat: &#39;2000000000&#39; &#43; &#39;chat_id&#39;, e.g. &#39;2000000001&#39;. For community: &#39;- community ID&#39;, e.g. &#39;-12345&#39;. &#34;
	PeerID int64 `url:"peer_id,omitempty"`
	// Date to search message before in Unixtime.
	Date int64 `url:"date,omitempty"`
	// Number of characters after which to truncate a previewed message. To preview the full message, specify &#39;0&#39;. &#34;NOTE: Messages are not truncated by default. Messages are truncated by words.&#34;
	PreviewLength int64 `url:"preview_length,omitempty"`
	// Offset needed to return a specific subset of messages.
	Offset int64 `url:"offset,omitempty"`
	// Number of messages to return.
	Count int64 `url:"count,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesSearchParams) Validate() error {
	if param.Date < 0.000000 {
		return fmt.Errorf("parameter Date(=%v) must be greater or equal than 0.000000", param.Date)
	}

	if param.PreviewLength < 0.000000 {
		return fmt.Errorf("parameter PreviewLength(=%v) must be greater or equal than 0.000000", param.PreviewLength)
	}

	if param.Offset < 0.000000 {
		return fmt.Errorf("parameter Offset(=%v) must be greater or equal than 0.000000", param.Offset)
	}

	if param.Count < 0.000000 {
		return fmt.Errorf("parameter Count(=%v) must be greater or equal than 0.000000", param.Count)
	}

	if param.Count > 100.000000 {
		return fmt.Errorf("parameter Count(=%v) must be less or equal than 100.000000", param.Count)
	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesSearch calls VK API method messages.search. Returns a list of the current user&#39;s private messages that match search criteria.
//
// https://vk.com/dev/messages.search
func (vkBot *VkBot) MessagesSearch(params *MessagesSearchParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.search", vals)
}

// MessagesGetHistoryParams params for messages.getHistory method
//
// https://vk.com/dev/messages.getHistory
type MessagesGetHistoryParams struct {

	// Offset needed to return a specific subset of messages.
	Offset int64 `url:"offset,omitempty"`
	// Number of messages to return.
	Count int64 `url:"count,omitempty"`
	// ID of the user whose message history you want to return.
	UserID int64 `url:"user_id,omitempty"`
	//
	PeerID int64 `url:"peer_id,omitempty"`
	// Starting message ID from which to return history.
	StartMessageID int64 `url:"start_message_id,omitempty"`
	// Information whether the response should be extended
	Extended bool `url:"extended,omitempty"`
	// Profile fields to return.
	Fields []string `url:"fields,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int64 `url:"group_id,omitempty"`
	// Sort order: &#39;1&#39; — return messages in chronological order. &#39;0&#39; — return messages in reverse chronological order.
	Rev int64 `url:"rev,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesGetHistoryParams) Validate() error {
	if param.Count < 0.000000 {
		return fmt.Errorf("parameter Count(=%v) must be greater or equal than 0.000000", param.Count)
	}

	if param.Count > 200.000000 {
		return fmt.Errorf("parameter Count(=%v) must be less or equal than 200.000000", param.Count)
	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	enum := map[int64]struct{}{1: {}, 0: {}}
	if _, ok := enum[param.Rev]; !ok {
		return fmt.Errorf("parameter Rev(=%v) expected in [1 0]", param.Rev)
	}

	return nil
}

// MessagesGetHistory calls VK API method messages.getHistory. Returns message history for the specified user or group chat.
//
// https://vk.com/dev/messages.getHistory
func (vkBot *VkBot) MessagesGetHistory(params *MessagesGetHistoryParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.getHistory", vals)
}

// MessagesGetHistoryAttachmentsParams params for messages.getHistoryAttachments method
//
// https://vk.com/dev/messages.getHistoryAttachments
type MessagesGetHistoryAttachmentsParams struct {

	// Peer ID. &#34;, For group chat: &#39;2000000000 &#43; chat ID&#39; , , For community: &#39;-community ID&#39;&#34;
	PeerID int64 `url:"peer_id"`
	// Type of media files to return: *&#39;photo&#39;,, *&#39;video&#39;,, *&#39;audio&#39;,, *&#39;doc&#39;,, *&#39;link&#39;.,*&#39;market&#39;.,*&#39;wall&#39;.,*&#39;share&#39;
	MediaType string `url:"media_type,omitempty"`
	// Message ID to start return results from.
	StartFrom string `url:"start_from,omitempty"`
	// Number of objects to return.
	Count int64 `url:"count,omitempty"`
	// &#39;1&#39; — to return photo sizes in a
	PhotoSizes bool `url:"photo_sizes,omitempty"`
	// Additional profile [vk.com/dev/fields|fields] to return.
	Fields []string `url:"fields,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesGetHistoryAttachmentsParams) Validate() error {
	enum := map[string]struct{}{"photo": {}, "video": {}, "doc": {}, "audio": {}, "link": {}, "market": {}, "wall": {}, "share": {}}
	if _, ok := enum[param.MediaType]; !ok {
		return fmt.Errorf("parameter MediaType(=%v) expected in [photo video doc audio link market wall share]", param.MediaType)
	}

	if param.Count < 0.000000 {
		return fmt.Errorf("parameter Count(=%v) must be greater or equal than 0.000000", param.Count)
	}

	if param.Count > 200.000000 {
		return fmt.Errorf("parameter Count(=%v) must be less or equal than 200.000000", param.Count)
	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesGetHistoryAttachments calls VK API method messages.getHistoryAttachments. Returns media files from the dialog or group chat.
//
// https://vk.com/dev/messages.getHistoryAttachments
func (vkBot *VkBot) MessagesGetHistoryAttachments(params *MessagesGetHistoryAttachmentsParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.getHistoryAttachments", vals)
}

// MessagesSendParams params for messages.send method
//
// https://vk.com/dev/messages.send
type MessagesSendParams struct {

	// User ID (by default — current user).
	UserID int64 `url:"user_id,omitempty"`
	// Unique identifier to avoid resending the message.
	RandomID int64 `url:"random_id,omitempty"`
	// Destination ID. &#34;For user: &#39;User ID&#39;, e.g. &#39;12345&#39;. For chat: &#39;2000000000&#39; &#43; &#39;chat_id&#39;, e.g. &#39;2000000001&#39;. For community: &#39;- community ID&#39;, e.g. &#39;-12345&#39;. &#34;
	PeerID int64 `url:"peer_id,omitempty"`
	// User&#39;s short address (for example, &#39;illarionov&#39;).
	Domain string `url:"domain,omitempty"`
	// ID of conversation the message will relate to.
	ChatID int64 `url:"chat_id,omitempty"`
	// IDs of message recipients (if new conversation shall be started).
	UserIds []int64 `url:"user_ids,omitempty"`
	// (Required if &#39;attachments&#39; is not set.) Text of the message.
	Message string `url:"message,omitempty"`
	// Geographical latitude of a check-in, in degrees (from -90 to 90).
	Lat float64 `url:"lat,omitempty"`
	// Geographical longitude of a check-in, in degrees (from -180 to 180).
	Long float64 `url:"long,omitempty"`
	// (Required if &#39;message&#39; is not set.) List of objects attached to the message, separated by commas, in the following format: &#34;&lt;owner_id&gt;_&lt;media_id&gt;&#34;, &#39;&#39; — Type of media attachment: &#39;photo&#39; — photo, &#39;video&#39; — video, &#39;audio&#39; — audio, &#39;doc&#39; — document, &#39;wall&#39; — wall post, &#39;&lt;owner_id&gt;&#39; — ID of the media attachment owner. &#39;&lt;media_id&gt;&#39; — media attachment ID. Example: &#34;photo100172_166443618&#34;
	Attachment []string `url:"attachment,omitempty"`
	// ID of forwarded messages, separated with a comma. Listed messages of the sender will be shown in the message body at the recipient&#39;s. Example: &#34;123,431,544&#34;
	ForwardMessages string `url:"forward_messages,omitempty"`
	// Sticker id.
	StickerID int64 `url:"sticker_id,omitempty"`
	// &#39;1&#39; if the message is a notification (for community messages).
	Notification bool `url:"notification,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesSendParams) Validate() error {
	if param.ChatID < 0.000000 {
		return fmt.Errorf("parameter ChatID(=%v) must be greater or equal than 0.000000", param.ChatID)
	}

	if param.Lat < -90.000000 {
		return fmt.Errorf("parameter Lat(=%v) must be greater or equal than -90.000000", param.Lat)
	}

	if param.Lat > 90.000000 {
		return fmt.Errorf("parameter Lat(=%v) must be less or equal than 90.000000", param.Lat)
	}

	if param.Long < -180.000000 {
		return fmt.Errorf("parameter Long(=%v) must be greater or equal than -180.000000", param.Long)
	}

	if param.Long > 180.000000 {
		return fmt.Errorf("parameter Long(=%v) must be less or equal than 180.000000", param.Long)
	}

	if param.StickerID < 0.000000 {
		return fmt.Errorf("parameter StickerID(=%v) must be greater or equal than 0.000000", param.StickerID)
	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesSend calls VK API method messages.send. Sends a message.
//
// https://vk.com/dev/messages.send
func (vkBot *VkBot) MessagesSend(params *MessagesSendParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.send", vals)
}

// MessagesEditParams params for messages.edit method
//
// https://vk.com/dev/messages.edit
type MessagesEditParams struct {

	// Destination ID. &#34;For user: &#39;User ID&#39;, e.g. &#39;12345&#39;. For chat: &#39;2000000000&#39; &#43; &#39;chat_id&#39;, e.g. &#39;2000000001&#39;. For community: &#39;- community ID&#39;, e.g. &#39;-12345&#39;. &#34;
	PeerID int64 `url:"peer_id"`
	// (Required if &#39;attachments&#39; is not set.) Text of the message.
	Message string `url:"message,omitempty"`
	// Geographical latitude of a check-in, in degrees (from -90 to 90).
	Lat float64 `url:"lat,omitempty"`
	// Geographical longitude of a check-in, in degrees (from -180 to 180).
	Long float64 `url:"long,omitempty"`
	// (Required if &#39;message&#39; is not set.) List of objects attached to the message, separated by commas, in the following format: &#34;&lt;owner_id&gt;_&lt;media_id&gt;&#34;, &#39;&#39; — Type of media attachment: &#39;photo&#39; — photo, &#39;video&#39; — video, &#39;audio&#39; — audio, &#39;doc&#39; — document, &#39;wall&#39; — wall post, &#39;&lt;owner_id&gt;&#39; — ID of the media attachment owner. &#39;&lt;media_id&gt;&#39; — media attachment ID. Example: &#34;photo100172_166443618&#34;
	Attachment []string `url:"attachment,omitempty"`
	// &#39;1&#39; — to keep forwarded, messages.
	KeepForwardMessages bool `url:"keep_forward_messages,omitempty"`
	// &#39;1&#39; — to keep attached snippets.
	KeepSnippets bool `url:"keep_snippets,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesEditParams) Validate() error {
	if param.Lat < -90.000000 {
		return fmt.Errorf("parameter Lat(=%v) must be greater or equal than -90.000000", param.Lat)
	}

	if param.Lat > 90.000000 {
		return fmt.Errorf("parameter Lat(=%v) must be less or equal than 90.000000", param.Lat)
	}

	if param.Long < -180.000000 {
		return fmt.Errorf("parameter Long(=%v) must be greater or equal than -180.000000", param.Long)
	}

	if param.Long > 180.000000 {
		return fmt.Errorf("parameter Long(=%v) must be less or equal than 180.000000", param.Long)
	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesEdit calls VK API method messages.edit. Edits the message.
//
// https://vk.com/dev/messages.edit
func (vkBot *VkBot) MessagesEdit(params *MessagesEditParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.edit", vals)
}

// MessagesDeleteParams params for messages.delete method
//
// https://vk.com/dev/messages.delete
type MessagesDeleteParams struct {

	// Message IDs.
	MessageIds []int64 `url:"message_ids,omitempty"`
	// &#39;1&#39; — to mark message as spam.
	Spam bool `url:"spam,omitempty"`
	// &#39;1&#39; — delete message for for all.
	DeleteForAll bool `url:"delete_for_all,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesDeleteParams) Validate() error {
	for i, item := range param.MessageIds {
		if item < 0.000000 {
			return fmt.Errorf("item[#%d](=%v) of parameter MessageIds must be greater or equal than 0.000000", i, item)
		}

	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesDelete calls VK API method messages.delete. Deletes one or more messages.
//
// https://vk.com/dev/messages.delete
func (vkBot *VkBot) MessagesDelete(params *MessagesDeleteParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.delete", vals)
}

// MessagesDeleteConversationParams params for messages.deleteConversation method
//
// https://vk.com/dev/messages.deleteConversation
type MessagesDeleteConversationParams struct {

	// User ID. To clear a chat history use &#39;chat_id&#39;
	UserID string `url:"user_id,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID int64 `url:"group_id,omitempty"`
	// Destination ID. &#34;For user: &#39;User ID&#39;, e.g. &#39;12345&#39;. For chat: &#39;2000000000&#39; &#43; &#39;chat_id&#39;, e.g. &#39;2000000001&#39;. For community: &#39;- community ID&#39;, e.g. &#39;-12345&#39;. &#34;
	PeerID int64 `url:"peer_id,omitempty"`
	// Offset needed to delete a specific subset of conversations.
	Offset int64 `url:"offset,omitempty"`
	// Number of conversations to delete. &#34;NOTE: If the number of messages exceeds the maximum, the method shall be called several times.&#34;
	Count int64 `url:"count,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesDeleteConversationParams) Validate() error {
	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	if param.Offset < 0.000000 {
		return fmt.Errorf("parameter Offset(=%v) must be greater or equal than 0.000000", param.Offset)
	}

	if param.Count < 0.000000 {
		return fmt.Errorf("parameter Count(=%v) must be greater or equal than 0.000000", param.Count)
	}

	if param.Count > 10000.000000 {
		return fmt.Errorf("parameter Count(=%v) must be less or equal than 10000.000000", param.Count)
	}

	return nil
}

// MessagesDeleteConversation calls VK API method messages.deleteConversation. Deletes all private messages in a conversation.
//
// https://vk.com/dev/messages.deleteConversation
func (vkBot *VkBot) MessagesDeleteConversation(params *MessagesDeleteConversationParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.deleteConversation", vals)
}

// MessagesRestoreParams params for messages.restore method
//
// https://vk.com/dev/messages.restore
type MessagesRestoreParams struct {

	// ID of a previously-deleted message to restore.
	MessageID int64 `url:"message_id"`
	// Group ID (for group messages with user access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesRestoreParams) Validate() error {
	if param.MessageID < 0.000000 {
		return fmt.Errorf("parameter MessageID(=%v) must be greater or equal than 0.000000", param.MessageID)
	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesRestore calls VK API method messages.restore. Restores a deleted message.
//
// https://vk.com/dev/messages.restore
func (vkBot *VkBot) MessagesRestore(params *MessagesRestoreParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.restore", vals)
}

// MessagesMarkAsReadParams params for messages.markAsRead method
//
// https://vk.com/dev/messages.markAsRead
type MessagesMarkAsReadParams struct {

	// IDs of messages to mark as read.
	MessageIds []int64 `url:"message_ids,omitempty"`
	// Destination ID. &#34;For user: &#39;User ID&#39;, e.g. &#39;12345&#39;. For chat: &#39;2000000000&#39; &#43; &#39;chat_id&#39;, e.g. &#39;2000000001&#39;. For community: &#39;- community ID&#39;, e.g. &#39;-12345&#39;. &#34;
	PeerID int64 `url:"peer_id,omitempty"`
	// Message ID to start from.
	StartMessageID int64 `url:"start_message_id,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesMarkAsReadParams) Validate() error {
	for i, item := range param.MessageIds {
		if item < 0.000000 {
			return fmt.Errorf("item[#%d](=%v) of parameter MessageIds must be greater or equal than 0.000000", i, item)
		}

	}

	if param.StartMessageID < 0.000000 {
		return fmt.Errorf("parameter StartMessageID(=%v) must be greater or equal than 0.000000", param.StartMessageID)
	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesMarkAsRead calls VK API method messages.markAsRead. Marks messages as read.
//
// https://vk.com/dev/messages.markAsRead
func (vkBot *VkBot) MessagesMarkAsRead(params *MessagesMarkAsReadParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.markAsRead", vals)
}

// MessagesMarkAsImportantParams params for messages.markAsImportant method
//
// https://vk.com/dev/messages.markAsImportant
type MessagesMarkAsImportantParams struct {

	// IDs of messages to mark as important.
	MessageIds []int64 `url:"message_ids,omitempty"`
	// &#39;1&#39; — to add a star (mark as important), &#39;0&#39; — to remove the star
	Important bool `url:"important,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesMarkAsImportantParams) Validate() error {
	for i, item := range param.MessageIds {
		if item < 0.000000 {
			return fmt.Errorf("item[#%d](=%v) of parameter MessageIds must be greater or equal than 0.000000", i, item)
		}

	}

	return nil
}

// MessagesMarkAsImportant calls VK API method messages.markAsImportant. Marks and unmarks messages as important (starred).
//
// https://vk.com/dev/messages.markAsImportant
func (vkBot *VkBot) MessagesMarkAsImportant(params *MessagesMarkAsImportantParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.markAsImportant", vals)
}

// MessagesMarkAsImportantConversationParams params for messages.markAsImportantConversation method
//
// https://vk.com/dev/messages.markAsImportantConversation
type MessagesMarkAsImportantConversationParams struct {

	// Group ID (for group messages with group access token)
	GroupID int64 `url:"group_id,omitempty"`
	// ID of conversation to mark as important.
	PeerID int64 `url:"peer_id"`
	// &#39;1&#39; — to add a star (mark as important), &#39;0&#39; — to remove the star
	Important bool `url:"important,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesMarkAsImportantConversationParams) Validate() error {
	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesMarkAsImportantConversation calls VK API method messages.markAsImportantConversation. Marks and unmarks conversations as important.
//
// https://vk.com/dev/messages.markAsImportantConversation
func (vkBot *VkBot) MessagesMarkAsImportantConversation(params *MessagesMarkAsImportantConversationParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.markAsImportantConversation", vals)
}

// MessagesMarkAsAnsweredConversationParams params for messages.markAsAnsweredConversation method
//
// https://vk.com/dev/messages.markAsAnsweredConversation
type MessagesMarkAsAnsweredConversationParams struct {

	// Group ID (for group messages with group access token)
	GroupID int64 `url:"group_id,omitempty"`
	// ID of conversation to mark as important.
	PeerID int64 `url:"peer_id"`
	// &#39;1&#39; — to mark as answered, &#39;0&#39; — to remove the mark
	Answered bool `url:"answered,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesMarkAsAnsweredConversationParams) Validate() error {
	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesMarkAsAnsweredConversation calls VK API method messages.markAsAnsweredConversation. Marks and unmarks conversations as unanswered.
//
// https://vk.com/dev/messages.markAsAnsweredConversation
func (vkBot *VkBot) MessagesMarkAsAnsweredConversation(params *MessagesMarkAsAnsweredConversationParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.markAsAnsweredConversation", vals)
}

// MessagesGetLongPollServerParams params for messages.getLongPollServer method
//
// https://vk.com/dev/messages.getLongPollServer
type MessagesGetLongPollServerParams struct {

	// Long poll version
	LpVersion int64 `url:"lp_version,omitempty"`
	// &#39;1&#39; — to return the &#39;pts&#39; field, needed for the [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
	NeedPts bool `url:"need_pts,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesGetLongPollServerParams) Validate() error {
	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesGetLongPollServer calls VK API method messages.getLongPollServer. Returns data required for connection to a Long Poll server.
//
// https://vk.com/dev/messages.getLongPollServer
func (vkBot *VkBot) MessagesGetLongPollServer(params *MessagesGetLongPollServerParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.getLongPollServer", vals)
}

// MessagesGetLongPollHistoryParams params for messages.getLongPollHistory method
//
// https://vk.com/dev/messages.getLongPollHistory
type MessagesGetLongPollHistoryParams struct {

	// Last value of the &#39;ts&#39; parameter returned from the Long Poll server or by using [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
	Ts int64 `url:"ts,omitempty"`
	// Lsat value of &#39;pts&#39; parameter returned from the Long Poll server or by using [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
	Pts int64 `url:"pts,omitempty"`
	// Number of characters after which to truncate a previewed message. To preview the full message, specify &#39;0&#39;. &#34;NOTE: Messages are not truncated by default. Messages are truncated by words.&#34;
	PreviewLength int64 `url:"preview_length,omitempty"`
	// &#39;1&#39; — to return history with online users only.
	Onlines bool `url:"onlines,omitempty"`
	// Additional profile [vk.com/dev/fields|fields] to return.
	Fields []string `url:"fields,omitempty"`
	// Maximum number of events to return.
	EventsLimit int64 `url:"events_limit,omitempty"`
	// Maximum number of messages to return.
	MsgsLimit int64 `url:"msgs_limit,omitempty"`
	// Maximum ID of the message among existing ones in the local copy. Both messages received with API methods (for example, , ), and data received from a Long Poll server (events with code 4) are taken into account.
	MaxMsgID int64 `url:"max_msg_id,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesGetLongPollHistoryParams) Validate() error {
	if param.Ts < 0.000000 {
		return fmt.Errorf("parameter Ts(=%v) must be greater or equal than 0.000000", param.Ts)
	}

	if param.Pts < 0.000000 {
		return fmt.Errorf("parameter Pts(=%v) must be greater or equal than 0.000000", param.Pts)
	}

	if param.PreviewLength < 0.000000 {
		return fmt.Errorf("parameter PreviewLength(=%v) must be greater or equal than 0.000000", param.PreviewLength)
	}

	if param.EventsLimit < 1000.000000 {
		return fmt.Errorf("parameter EventsLimit(=%v) must be greater or equal than 1000.000000", param.EventsLimit)
	}

	if param.MsgsLimit < 200.000000 {
		return fmt.Errorf("parameter MsgsLimit(=%v) must be greater or equal than 200.000000", param.MsgsLimit)
	}

	if param.MaxMsgID < 0.000000 {
		return fmt.Errorf("parameter MaxMsgID(=%v) must be greater or equal than 0.000000", param.MaxMsgID)
	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesGetLongPollHistory calls VK API method messages.getLongPollHistory. Returns updates in user&#39;s private messages.
//
// https://vk.com/dev/messages.getLongPollHistory
func (vkBot *VkBot) MessagesGetLongPollHistory(params *MessagesGetLongPollHistoryParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.getLongPollHistory", vals)
}

// MessagesCreateChatParams params for messages.createChat method
//
// https://vk.com/dev/messages.createChat
type MessagesCreateChatParams struct {

	// IDs of the users to be added to the chat.
	UserIds []int64 `url:"user_ids"`
	// Chat title.
	Title string `url:"title,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesCreateChatParams) Validate() error {
	for i, item := range param.UserIds {
		if item < 0.000000 {
			return fmt.Errorf("item[#%d](=%v) of parameter UserIds must be greater or equal than 0.000000", i, item)
		}

	}

	return nil
}

// MessagesCreateChat calls VK API method messages.createChat. Creates a chat with several participants.
//
// https://vk.com/dev/messages.createChat
func (vkBot *VkBot) MessagesCreateChat(params *MessagesCreateChatParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.createChat", vals)
}

// MessagesEditChatParams params for messages.editChat method
//
// https://vk.com/dev/messages.editChat
type MessagesEditChatParams struct {

	// Chat ID.
	ChatID int64 `url:"chat_id"`
	// New title of the chat.
	Title string `url:"title"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesEditChatParams) Validate() error {
	return nil
}

// MessagesEditChat calls VK API method messages.editChat. Edits the title of a chat.
//
// https://vk.com/dev/messages.editChat
func (vkBot *VkBot) MessagesEditChat(params *MessagesEditChatParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.editChat", vals)
}

// MessagesGetConversationMembersParams params for messages.getConversationMembers method
//
// https://vk.com/dev/messages.getConversationMembers
type MessagesGetConversationMembersParams struct {

	// Group ID (for group messages with group access token)
	GroupID int64 `url:"group_id,omitempty"`
	// Peer ID.
	PeerID int64 `url:"peer_id,omitempty"`
	// Profile fields to return.
	Fields []string `url:"fields,omitempty"`
	// Case for declension of user name and surname: &#39;nom&#39; — nominative (default), &#39;gen&#39; — genitive, &#39;dat&#39; — dative, &#39;acc&#39; — accusative, &#39;ins&#39; — instrumental, &#39;abl&#39; — prepositional
	NameCase string `url:"name_case,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesGetConversationMembersParams) Validate() error {
	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	enum := map[string]struct{}{"nom": {}, "gen": {}, "dat": {}, "acc": {}, "ins": {}, "abl": {}}
	if _, ok := enum[param.NameCase]; !ok {
		return fmt.Errorf("parameter NameCase(=%v) expected in [nom gen dat acc ins abl]", param.NameCase)
	}

	return nil
}

// MessagesGetConversationMembers calls VK API method messages.getConversationMembers. Returns a list of IDs of users participating in a chat.
//
// https://vk.com/dev/messages.getConversationMembers
func (vkBot *VkBot) MessagesGetConversationMembers(params *MessagesGetConversationMembersParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.getConversationMembers", vals)
}

// MessagesSetActivityParams params for messages.setActivity method
//
// https://vk.com/dev/messages.setActivity
type MessagesSetActivityParams struct {

	// User ID.
	UserID string `url:"user_id,omitempty"`
	// &#39;typing&#39; — user has started to type.
	Type string `url:"type,omitempty"`
	// Destination ID. &#34;For user: &#39;User ID&#39;, e.g. &#39;12345&#39;. For chat: &#39;2000000000&#39; &#43; &#39;chat_id&#39;, e.g. &#39;2000000001&#39;. For community: &#39;- community ID&#39;, e.g. &#39;-12345&#39;. &#34;
	PeerID int64 `url:"peer_id,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesSetActivityParams) Validate() error {
	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesSetActivity calls VK API method messages.setActivity. Changes the status of a user as typing in a conversation.
//
// https://vk.com/dev/messages.setActivity
func (vkBot *VkBot) MessagesSetActivity(params *MessagesSetActivityParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.setActivity", vals)
}

// MessagesSearchConversationsParams params for messages.searchConversations method
//
// https://vk.com/dev/messages.searchConversations
type MessagesSearchConversationsParams struct {

	// Search query string.
	Q string `url:"q,omitempty"`
	// Maximum number of results.
	Count int64 `url:"count,omitempty"`
	// &#39;1&#39; — return extra information about users and communities
	Extended bool `url:"extended,omitempty"`
	// Profile fields to return.
	Fields []string `url:"fields,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID int64 `url:"group_id,omitempty"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesSearchConversationsParams) Validate() error {
	if param.Count < 0.000000 {
		return fmt.Errorf("parameter Count(=%v) must be greater or equal than 0.000000", param.Count)
	}

	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesSearchConversations calls VK API method messages.searchConversations. Returns a list of the current user&#39;s conversations that match search criteria.
//
// https://vk.com/dev/messages.searchConversations
func (vkBot *VkBot) MessagesSearchConversations(params *MessagesSearchConversationsParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.searchConversations", vals)
}

// MessagesAddChatUserParams params for messages.addChatUser method
//
// https://vk.com/dev/messages.addChatUser
type MessagesAddChatUserParams struct {

	// Chat ID.
	ChatID int64 `url:"chat_id"`
	// ID of the user to be added to the chat.
	UserID int64 `url:"user_id"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesAddChatUserParams) Validate() error {
	if param.ChatID < 0.000000 {
		return fmt.Errorf("parameter ChatID(=%v) must be greater or equal than 0.000000", param.ChatID)
	}

	if param.UserID < 0.000000 {
		return fmt.Errorf("parameter UserID(=%v) must be greater or equal than 0.000000", param.UserID)
	}

	return nil
}

// MessagesAddChatUser calls VK API method messages.addChatUser. Adds a new user to a chat.
//
// https://vk.com/dev/messages.addChatUser
func (vkBot *VkBot) MessagesAddChatUser(params *MessagesAddChatUserParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.addChatUser", vals)
}

// MessagesRemoveChatUserParams params for messages.removeChatUser method
//
// https://vk.com/dev/messages.removeChatUser
type MessagesRemoveChatUserParams struct {

	// Chat ID.
	ChatID int64 `url:"chat_id"`
	// ID of the user to be removed from the chat.
	UserID string `url:"user_id"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesRemoveChatUserParams) Validate() error {
	return nil
}

// MessagesRemoveChatUser calls VK API method messages.removeChatUser. Allows the current user to leave a chat or, if the current user started the chat, allows the user to remove another user from the chat.
//
// https://vk.com/dev/messages.removeChatUser
func (vkBot *VkBot) MessagesRemoveChatUser(params *MessagesRemoveChatUserParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.removeChatUser", vals)
}

// MessagesGetLastActivityParams params for messages.getLastActivity method
//
// https://vk.com/dev/messages.getLastActivity
type MessagesGetLastActivityParams struct {

	// User ID.
	UserID int64 `url:"user_id"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesGetLastActivityParams) Validate() error {
	return nil
}

// MessagesGetLastActivity calls VK API method messages.getLastActivity. Returns a user&#39;s current status and date of last activity.
//
// https://vk.com/dev/messages.getLastActivity
func (vkBot *VkBot) MessagesGetLastActivity(params *MessagesGetLastActivityParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.getLastActivity", vals)
}

// MessagesSetChatPhotoParams params for messages.setChatPhoto method
//
// https://vk.com/dev/messages.setChatPhoto
type MessagesSetChatPhotoParams struct {

	// Upload URL from the &#39;response&#39; field returned by the [vk.com/dev/photos.getChatUploadServer|photos.getChatUploadServer] method upon successfully uploading an image.
	File string `url:"file"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesSetChatPhotoParams) Validate() error {
	return nil
}

// MessagesSetChatPhoto calls VK API method messages.setChatPhoto. Sets a previously-uploaded picture as the cover picture of a chat.
//
// https://vk.com/dev/messages.setChatPhoto
func (vkBot *VkBot) MessagesSetChatPhoto(params *MessagesSetChatPhotoParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.setChatPhoto", vals)
}

// MessagesDeleteChatPhotoParams params for messages.deleteChatPhoto method
//
// https://vk.com/dev/messages.deleteChatPhoto
type MessagesDeleteChatPhotoParams struct {

	// Chat ID.
	ChatID int64 `url:"chat_id"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesDeleteChatPhotoParams) Validate() error {
	if param.ChatID < 0.000000 {
		return fmt.Errorf("parameter ChatID(=%v) must be greater or equal than 0.000000", param.ChatID)
	}

	return nil
}

// MessagesDeleteChatPhoto calls VK API method messages.deleteChatPhoto. Deletes a chat&#39;s cover picture.
//
// https://vk.com/dev/messages.deleteChatPhoto
func (vkBot *VkBot) MessagesDeleteChatPhoto(params *MessagesDeleteChatPhotoParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.deleteChatPhoto", vals)
}

// MessagesDenyMessagesFromGroupParams params for messages.denyMessagesFromGroup method
//
// https://vk.com/dev/messages.denyMessagesFromGroup
type MessagesDenyMessagesFromGroupParams struct {

	// Group ID.
	GroupID int64 `url:"group_id"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesDenyMessagesFromGroupParams) Validate() error {
	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesDenyMessagesFromGroup calls VK API method messages.denyMessagesFromGroup. Denies sending message from community to the current user.
//
// https://vk.com/dev/messages.denyMessagesFromGroup
func (vkBot *VkBot) MessagesDenyMessagesFromGroup(params *MessagesDenyMessagesFromGroupParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.denyMessagesFromGroup", vals)
}

// MessagesAllowMessagesFromGroupParams params for messages.allowMessagesFromGroup method
//
// https://vk.com/dev/messages.allowMessagesFromGroup
type MessagesAllowMessagesFromGroupParams struct {

	// Group ID.
	GroupID int64 `url:"group_id"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesAllowMessagesFromGroupParams) Validate() error {
	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	return nil
}

// MessagesAllowMessagesFromGroup calls VK API method messages.allowMessagesFromGroup. Allows sending messages from community to the current user.
//
// https://vk.com/dev/messages.allowMessagesFromGroup
func (vkBot *VkBot) MessagesAllowMessagesFromGroup(params *MessagesAllowMessagesFromGroupParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.allowMessagesFromGroup", vals)
}

// MessagesIsMessagesFromGroupAllowedParams params for messages.isMessagesFromGroupAllowed method
//
// https://vk.com/dev/messages.isMessagesFromGroupAllowed
type MessagesIsMessagesFromGroupAllowedParams struct {

	// Group ID.
	GroupID int64 `url:"group_id"`
	// User ID.
	UserID int64 `url:"user_id"`
}

// Validate is called before sending a request to VK API to validate parameters.
func (param *MessagesIsMessagesFromGroupAllowedParams) Validate() error {
	if param.GroupID < 0.000000 {
		return fmt.Errorf("parameter GroupID(=%v) must be greater or equal than 0.000000", param.GroupID)
	}

	if param.UserID < 0.000000 {
		return fmt.Errorf("parameter UserID(=%v) must be greater or equal than 0.000000", param.UserID)
	}

	return nil
}

// MessagesIsMessagesFromGroupAllowed calls VK API method messages.isMessagesFromGroupAllowed. Returns information whether sending messages from the community to current user is allowed.
//
// https://vk.com/dev/messages.isMessagesFromGroupAllowed
func (vkBot *VkBot) MessagesIsMessagesFromGroupAllowed(params *MessagesIsMessagesFromGroupAllowedParams) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request("messages.isMessagesFromGroupAllowed", vals)
}
