package vkbotgo

// All types of notifications and their corresponding objects
// that are supported in the Callback API
// and Bots Long Poll API are listed below.
//
// https://vk.com/dev/groups_events

// MessageNewEvent incoming message
type MessageNewEvent = MessagesMessage

// MessageReplyEvent new outgoing message
type MessageReplyEvent = MessagesMessage

// MessageEditEvent to edit the message
type MessageEditEvent = MessagesMessage

// MessageAllowEvent subscribe to messages from the community
type MessageAllowEvent struct {

	// UserID user ID
	UserID int64 `json:"user_id"`
	// Key the parameter passed in the messages method.allowMessagesFromGroup.
	Key string `json:"key"`
}

// MessageDenyEvent new ban messages from the community
type MessageDenyEvent struct {

	// UserID user ID
	UserID int64 `json:"user_id"`
}

// PhotoNewEvent adding a photo
type PhotoNewEvent = PhotosPhoto

// PhotoCommentNewEvent add a comment to a photo
type PhotoCommentNewEvent struct {
	WallWallComment
	// PhotoID photo ID
	PhotoID int64 `json:"photo_id"`
	// PhotoOwnerID ID of the owner of the photo.
	PhotoOwnerID int64 `json:"photo_owner_id"`
}

// PhotoCommentEditEvent editing a photo comment
type PhotoCommentEditEvent = PhotoCommentNewEvent

// PhotoCommentRestoreEvent restore a comment to a photo
type PhotoCommentRestoreEvent = PhotoCommentNewEvent

// PhotoCommentDeleteEvent delete a comment for a photo
type PhotoCommentDeleteEvent struct {

	// OwnerID ID of photo owner
	OwnerID int64 `json:"owner_id"`
	// ID of the review
	ID int64 `json:"id"`
	// UserID ID of the comment author
	UserID int64 `json:"user_id"`
	// DeleterID ID of the user who deleted the comment
	DeleterID int64 `json:"deleter_id"`
	// PhotoID photo ID
	PhotoID int64 `json:"photo_id"`
}

// AudioNewEvent adding audio
type AudioNewEvent = AudioAudio

// VideoNewEvent adding a video
type VideoNewEvent = VideoVideo

// VideoCommentNewEvent add a comment to a video
type VideoCommentNewEvent struct {
	WallWallComment
	// VideoID video ID
	VideoID int64 `json:"video_id"`
	// VideoOwnerID ID of the owner of the video.
	VideoOwnerID int64 `json:"video_owner_id"`
}

// VideoCommentEditEvent editing a video comment
type VideoCommentEditEvent = VideoCommentNewEvent

// VideoCommentRestoreEvent restore a comment to a video
type VideoCommentRestoreEvent = VideoCommentNewEvent

// VideoCommentDeleteEvent delete a comment for a video
type VideoCommentDeleteEvent struct {

	// OwnerID ID of video owner
	OwnerID int64 `json:"owner_id"`
	// ID of the review
	ID int64 `json:"id"`
	// UserID ID of the comment author
	UserID int64 `json:"user_id"`
	// DeleterID ID of the user who deleted the comment
	DeleterID int64 `json:"deleter_id"`
	// VideoID video ID
	VideoID int64 `json:"video_id"`
}

// WallPostNewEvent writing on the wall
type WallPostNewEvent struct {
	WallWallpost
	// PostponedID ID of the pending entry
	PostponedID int64 `json:"postponed_id"`
}

// WallRepostEvent repost entries from the community
type WallRepostEvent = WallPostNewEvent

// WallReplyNewEvent adding a comment on the wall
type WallReplyNewEvent struct {
	WallWallComment
	// PostID post ID
	PostID int64 `json:"post_id"`
	// PostOwnerID ID of the record owner
	PostOwnerID int64 `json:"post_owner_id"`
}

// WallReplyEditEvent editing a comment on the wall
type WallReplyEditEvent = WallReplyNewEvent

// WallReplyRestoreEvent restoring a comment on the wall
type WallReplyRestoreEvent = WallReplyNewEvent

// WallReplyDeleteEvent deleting a comment on the wall
type WallReplyDeleteEvent struct {

	// OwnerID ID of the owner of the wall
	OwnerID int64 `json:"owner_id"`
	// ID of the review
	ID int64 `json:"id"`
	// DeleterID ID of the user who deleted the comment
	DeleterID int64 `json:"deleter_id"`
	// PostID ID of the post to which the comment was left
	PostID int64 `json:"post_id"`
}

// BoardPostNewEvent create a comment in a discussion
type BoardPostNewEvent struct {
	// No comment_board object (https://vk.com/dev/objects/comment_board)
	// TopicID discussion ID
	TopicID int64 `json:"topic_id"`
	// TopicOwnerID discussion owner ID
	TopicOwnerID int64 `json:"topic_owner_id"`
}

// BoardPostEditEvent editing a comment
type BoardPostEditEvent = BoardPostNewEvent

// BoardPostRestoreEvent restoring a comment
type BoardPostRestoreEvent = BoardPostNewEvent

// BoardPostDeleteEvent deleting a comment in a discussion
type BoardPostDeleteEvent struct {

	// TopicID discussion ID
	TopicID int64 `json:"topic_id"`
	// TopicOwnerID discussion owner ID
	TopicOwnerID int64 `json:"topic_owner_id"`
	// ID of the review
	ID int64 `json:"id"`
}

// MarketCommentNewEvent new comment to item
type MarketCommentNewEvent struct {
	WallWallComment
	// ItemID item ID
	ItemID int64 `json:"item_id"`
	// MarketOwnerID ID of the owner of the item
	MarketOwnerID int64 `json:"market_owner_id"`
}

// MarketCommentEditEvent editing a product comment
type MarketCommentEditEvent = MarketCommentNewEvent

// MarketCommentRestoreEvent restore review to the product
type MarketCommentRestoreEvent = MarketCommentNewEvent

// MarketCommentDeleteEvent delete a product comment
type MarketCommentDeleteEvent struct {

	// OwnerID ID of the owner of the item
	OwnerID int64 `json:"owner_id"`
	// ID of the review
	ID int64 `json:"id"`
	// UserID ID of the comment author
	UserID int64 `json:"user_id"`
	// DeleterID ID of the user who deleted the comment
	DeleterID int64 `json:"deleter_id"`
	// ItemID item ID
	ItemID int64 `json:"item_id"`
}

// GroupLeaveEvent remove a member from a community
type GroupLeaveEvent struct {

	// UserID user ID
	UserID int64 `json:"user_id"`
	// Self a value indicating whether the user was deleted or logged out on their own. ([0, 1])
	Self int64 `json:"self"`
}

// GroupJoinEvent adding a member or applying to join a community
type GroupJoinEvent struct {

	// UserID user ID
	UserID int64 `json:"user_id"`
	// JoinType specifies how the member was added.
	JoinType string `json:"join_type"`
}

// UserBlockEvent adding a user to the blacklist
type UserBlockEvent struct {

	// AdminID ID of the administrator who blacklisted the user
	AdminID int64 `json:"admin_id"`
	// UserID user ID
	UserID int64 `json:"user_id"`
	// UnblockDate unblock date (Unixtime)
	UnblockDate int64 `json:"unblock_date"`
	// Reason the reason for blocking
	Reason int64 `json:"reason"`
}

// UserUnblockEvent remove a user from the blacklist
type UserUnblockEvent struct {

	// AdminID ID of the administrator who removed the user from the blacklist
	AdminID int64 `json:"admin_id"`
	// UserID user ID
	UserID int64 `json:"user_id"`
	// ByEndDate value 1, if the blocking period has expired
	ByEndDate int64 `json:"by_end_date"`
}

// PollVoteNewEvent adding a voice to a public poll
type PollVoteNewEvent struct {

	// OwnerID the owner of the survey
	OwnerID int64 `json:"owner_id"`
	// PollID ID of the survey
	PollID int64 `json:"poll_id"`
	// OptionID ID of the answer
	OptionID int64 `json:"option_id"`
	// UserID user ID
	UserID int64 `json:"user_id"`
}

// GroupOfficersEditEvent editing the list of managers
type GroupOfficersEditEvent struct {

	// AdminID ID of the manager who made the changes
	AdminID int64 `json:"admin_id"`
	// UserID ID of the user whose permissions have been changed
	UserID int64 `json:"user_id"`
	// LevelOld old level of authority
	LevelOld int64 `json:"level_old"`
	// LevelNew new level of authority
	LevelNew int64 `json:"level_new"`
}

// GroupChangeSettingsEvent change the settings of the community
type GroupChangeSettingsEvent struct {

	// UserID ID of the user who made the changes
	UserID int64 `json:"user_id"`
	// Changes the description of the changes
	Changes struct {
		// Field the name of the section or section that has been changed
		Field string `json:"{FIELD}"`
		// OldValue old value
		OldValue string `json:"old_value"`
		// NewValue new value
		NewValue string `json:"new_value"`
	} `json:"changes"`
}

// GroupChangePhotoEvent changing the main photo
type GroupChangePhotoEvent struct {

	// UserID ID of the user who made the changes
	UserID int64 `json:"user_id"`
	// Photo the object that describes the photo.
	Photo *PhotosPhoto `json:"photo"`
}

// VkPayTransactionEvent payment via VK Pay
type VkPayTransactionEvent struct {

	// FromID the user ID of the sender
	FromID int64 `json:"from_id"`
	// Amount amount of transfer in thousand rubles
	Amount int64 `json:"amount"`
	// Description comment to translation
	Description string `json:"description"`
	// Date time of transfer to the Unixtime
	Date int64 `json:"date"`
}
