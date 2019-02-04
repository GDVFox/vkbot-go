package vkbotgo

// GENERATED FILE: DO NOT EDIT

// PhotosPhotoXtrTagInfo type from VK API Schema(photos_photo_xtr_tag_info).
type PhotosPhotoXtrTagInfo struct {

	// AccessKey Access key for the photo
	AccessKey string `json:"access_key,omitempty"`
	// AlbumID Album ID
	AlbumID int64 `json:"album_id"`
	// Date Date when uploaded
	Date int64 `json:"date"`
	// Height Original photo height
	Height int64 `json:"height,omitempty"`
	// ID Photo ID
	ID int64 `json:"id"`
	// Lat Latitude
	Lat float64 `json:"lat,omitempty"`
	// Long Longitude
	Long float64 `json:"long,omitempty"`
	// OwnerID Photo owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// Photo1280 URL of image with 1280 px width
	Photo1280 string `json:"photo_1280,omitempty"`
	// Photo130 URL of image with 130 px width
	Photo130 string `json:"photo_130,omitempty"`
	// Photo2560 URL of image with 2560 px width
	Photo2560 string `json:"photo_2560,omitempty"`
	// Photo604 URL of image with 604 px width
	Photo604 string `json:"photo_604,omitempty"`
	// Photo75 URL of image with 75 px width
	Photo75 string `json:"photo_75,omitempty"`
	// Photo807 URL of image with 807 px width
	Photo807 string `json:"photo_807,omitempty"`
	// PlacerID ID of the tag creator
	PlacerID int64 `json:"placer_id,omitempty"`
	// PostID Post ID
	PostID int64 `json:"post_id,omitempty"`
	// Sizes
	Sizes []*PhotosPhotoSizes `json:"sizes,omitempty"`
	// TagCreated Date when tag has been added in Unixtime
	TagCreated int64 `json:"tag_created,omitempty"`
	// TagID Tag ID
	TagID int64 `json:"tag_id,omitempty"`
	// Text Photo caption
	Text string `json:"text,omitempty"`
	// UserID ID of the user who have uploaded the photo
	UserID int64 `json:"user_id,omitempty"`
	// Width Original photo width
	Width int64 `json:"width,omitempty"`
}

// UsersUserXtrCounters type from VK API Schema(users_user_xtr_counters).
type UsersUserXtrCounters struct {
}

// VideoVideoAlbum type from VK API Schema(video_video_album).
type VideoVideoAlbum struct {

	// ID Album ID
	ID int64 `json:"id"`
	// OwnerID Album owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// Title Album title
	Title string `json:"title"`
}

// BaseSex type from VK API Schema(base_sex).
type BaseSex int64

// GroupsLongPollEvents type from VK API Schema(groups_long_poll_events).
type GroupsLongPollEvents struct {

	// AudioNew
	AudioNew *BaseBoolInt `json:"audio_new"`
	// BoardPostDelete
	BoardPostDelete *BaseBoolInt `json:"board_post_delete"`
	// BoardPostEdit
	BoardPostEdit *BaseBoolInt `json:"board_post_edit"`
	// BoardPostNew
	BoardPostNew *BaseBoolInt `json:"board_post_new"`
	// BoardPostRestore
	BoardPostRestore *BaseBoolInt `json:"board_post_restore"`
	// GroupChangePhoto
	GroupChangePhoto *BaseBoolInt `json:"group_change_photo"`
	// GroupChangeSettings
	GroupChangeSettings *BaseBoolInt `json:"group_change_settings"`
	// GroupJoin
	GroupJoin *BaseBoolInt `json:"group_join"`
	// GroupLeave
	GroupLeave *BaseBoolInt `json:"group_leave"`
	// GroupOfficersEdit
	GroupOfficersEdit *BaseBoolInt `json:"group_officers_edit"`
	// LeadFormsNew
	LeadFormsNew *BaseBoolInt `json:"lead_forms_new,omitempty"`
	// MarketCommentDelete
	MarketCommentDelete *BaseBoolInt `json:"market_comment_delete"`
	// MarketCommentEdit
	MarketCommentEdit *BaseBoolInt `json:"market_comment_edit"`
	// MarketCommentNew
	MarketCommentNew *BaseBoolInt `json:"market_comment_new"`
	// MarketCommentRestore
	MarketCommentRestore *BaseBoolInt `json:"market_comment_restore"`
	// MessageAllow
	MessageAllow *BaseBoolInt `json:"message_allow"`
	// MessageDeny
	MessageDeny *BaseBoolInt `json:"message_deny"`
	// MessageNew
	MessageNew *BaseBoolInt `json:"message_new"`
	// MessageReply
	MessageReply *BaseBoolInt `json:"message_reply"`
	// MessageTypingState
	MessageTypingState *BaseBoolInt `json:"message_typing_state"`
	// MessagesEdit
	MessagesEdit *BaseBoolInt `json:"messages_edit"`
	// PhotoCommentDelete
	PhotoCommentDelete *BaseBoolInt `json:"photo_comment_delete"`
	// PhotoCommentEdit
	PhotoCommentEdit *BaseBoolInt `json:"photo_comment_edit"`
	// PhotoCommentNew
	PhotoCommentNew *BaseBoolInt `json:"photo_comment_new"`
	// PhotoCommentRestore
	PhotoCommentRestore *BaseBoolInt `json:"photo_comment_restore"`
	// PhotoNew
	PhotoNew *BaseBoolInt `json:"photo_new"`
	// PollVoteNew
	PollVoteNew *BaseBoolInt `json:"poll_vote_new"`
	// UserBlock
	UserBlock *BaseBoolInt `json:"user_block"`
	// UserUnblock
	UserUnblock *BaseBoolInt `json:"user_unblock"`
	// VideoCommentDelete
	VideoCommentDelete *BaseBoolInt `json:"video_comment_delete"`
	// VideoCommentEdit
	VideoCommentEdit *BaseBoolInt `json:"video_comment_edit"`
	// VideoCommentNew
	VideoCommentNew *BaseBoolInt `json:"video_comment_new"`
	// VideoCommentRestore
	VideoCommentRestore *BaseBoolInt `json:"video_comment_restore"`
	// VideoNew
	VideoNew *BaseBoolInt `json:"video_new"`
	// WallPostNew
	WallPostNew *BaseBoolInt `json:"wall_post_new"`
	// WallReplyDelete
	WallReplyDelete *BaseBoolInt `json:"wall_reply_delete"`
	// WallReplyEdit
	WallReplyEdit *BaseBoolInt `json:"wall_reply_edit"`
	// WallReplyNew
	WallReplyNew *BaseBoolInt `json:"wall_reply_new"`
	// WallReplyRestore
	WallReplyRestore *BaseBoolInt `json:"wall_reply_restore"`
	// WallRepost
	WallRepost *BaseBoolInt `json:"wall_repost"`
}

// MessagesLastActivity type from VK API Schema(messages_last_activity).
type MessagesLastActivity struct {

	// Online Information whether user is online
	Online *BaseBoolInt `json:"online"`
	// Time Time when user was online in Unixtime
	Time int64 `json:"time"`
}

// NewsfeedNewsfeedItem type from VK API Schema(newsfeed_newsfeed_item).
type NewsfeedNewsfeedItem struct {
}

// NewsfeedItemVideo type from VK API Schema(newsfeed_item_video).
type NewsfeedItemVideo struct {

	// Video
	Video *NewsfeedItemVideoVideo `json:"video,omitempty"`
}

// WidgetsCommentMedia type from VK API Schema(widgets_comment_media).
type WidgetsCommentMedia struct {

	// ItemID Media item ID
	ItemID int64 `json:"item_id,omitempty"`
	// OwnerID Media owner&#39;s ID
	OwnerID int64 `json:"owner_id,omitempty"`
	// ThumbSrc URL of the preview image (type=photo only)
	ThumbSrc string `json:"thumb_src,omitempty"`
	// Type
	Type *WidgetsCommentMediaType `json:"type,omitempty"`
}

// AdsAccountType type from VK API Schema(ads_account_type). Account type
type AdsAccountType string

// DatabaseUniversity type from VK API Schema(database_university).
type DatabaseUniversity struct {

	// ID University ID
	ID int64 `json:"id,omitempty"`
	// Title University title
	Title string `json:"title,omitempty"`
}

// GroupsGroup type from VK API Schema(groups_group).
type GroupsGroup struct {

	// AdminLevel
	AdminLevel *GroupsGroupAdminLevel `json:"admin_level,omitempty"`
	// Deactivated Information whether community is banned
	Deactivated string `json:"deactivated,omitempty"`
	// ID Community ID
	ID int64 `json:"id,omitempty"`
	// IsAdmin Information whether current user is administrator
	IsAdmin *BaseBoolInt `json:"is_admin,omitempty"`
	// IsClosed
	IsClosed *GroupsGroupIsClosed `json:"is_closed,omitempty"`
	// IsMember Information whether current user is member
	IsMember *BaseBoolInt `json:"is_member,omitempty"`
	// Name Community name
	Name string `json:"name,omitempty"`
	// Photo100 URL of square photo of the community with 100 pixels in width
	Photo100 string `json:"photo_100,omitempty"`
	// Photo200 URL of square photo of the community with 200 pixels in width
	Photo200 string `json:"photo_200,omitempty"`
	// Photo50 URL of square photo of the community with 50 pixels in width
	Photo50 string `json:"photo_50,omitempty"`
	// ScreenName Domain of the community page
	ScreenName string `json:"screen_name,omitempty"`
	// Type
	Type *GroupsGroupType `json:"type,omitempty"`
}

// SecureLevel type from VK API Schema(secure_level).
type SecureLevel struct {

	// Level Level
	Level int64 `json:"level,omitempty"`
	// UID User ID
	UID int64 `json:"uid,omitempty"`
}

// AccountOnoffOptions type from VK API Schema(account_onoff_options). Settings parameters
type AccountOnoffOptions string

// AdsCategory type from VK API Schema(ads_category).
type AdsCategory struct {

	// ID Category ID
	ID int64 `json:"id"`
	// Name Category name
	Name string `json:"name"`
	// Subcategories
	Subcategories []*BaseObjectWithName `json:"subcategories,omitempty"`
}

// MessagesHistoryMessageAttachment type from VK API Schema(messages_history_message_attachment).
type MessagesHistoryMessageAttachment struct {

	// Audio
	Audio *AudioAudioFull `json:"audio,omitempty"`
	// Doc
	Doc *DocsDoc `json:"doc,omitempty"`
	// Link
	Link *BaseLink `json:"link,omitempty"`
	// Market
	Market *BaseLink `json:"market,omitempty"`
	// Photo
	Photo *PhotosPhoto `json:"photo,omitempty"`
	// Share
	Share *BaseLink `json:"share,omitempty"`
	// Type
	Type *MessagesHistoryMessageAttachmentType `json:"type"`
	// Video
	Video *VideoVideo `json:"video,omitempty"`
	// Wall
	Wall *BaseLink `json:"wall,omitempty"`
}

// MessagesKeyboardButton type from VK API Schema(messages_keyboard_button).
type MessagesKeyboardButton struct {

	// Action
	Action *MessagesKeyboardButtonAction `json:"action"`
	// Color Button color
	Color string `json:"color"`
}

// StatsViews type from VK API Schema(stats_views). Views stats
type StatsViews struct {

	// Age
	Age []*StatsSexAge `json:"age,omitempty"`
	// Cities
	Cities []*StatsCity `json:"cities,omitempty"`
	// Countries
	Countries []*StatsCountry `json:"countries,omitempty"`
	// MobileViews Number of views from mobile devices
	MobileViews int64 `json:"mobile_views,omitempty"`
	// Sex
	Sex []*StatsSexAge `json:"sex,omitempty"`
	// SexAge
	SexAge []*StatsSexAge `json:"sex_age,omitempty"`
	// Views Views number
	Views int64 `json:"views,omitempty"`
	// Visitors Visitors number
	Visitors int64 `json:"visitors,omitempty"`
}

// AudioAudioFull type from VK API Schema(audio_audio_full).
type AudioAudioFull struct {
}

// GroupsBanInfo type from VK API Schema(groups_ban_info).
type GroupsBanInfo struct {

	// AdminID Administrator ID
	AdminID int64 `json:"admin_id,omitempty"`
	// Comment Comment for a ban
	Comment string `json:"comment,omitempty"`
	// Date Date when user has been added to blacklist in Unixtime
	Date int64 `json:"date,omitempty"`
	// EndDate Date when user will be removed from blacklist in Unixtime
	EndDate int64 `json:"end_date,omitempty"`
	// Reason
	Reason *GroupsBanInfoReason `json:"reason,omitempty"`
}

// GroupsGroupSettings type from VK API Schema(groups_group_settings).
type GroupsGroupSettings struct {

	// Access Community access settings
	Access int64 `json:"access,omitempty"`
	// Address Community&#39;s page domain
	Address string `json:"address,omitempty"`
	// Audio Audio settings
	Audio int64 `json:"audio,omitempty"`
	// Description Community description
	Description string `json:"description,omitempty"`
	// Docs Docs settings
	Docs int64 `json:"docs,omitempty"`
	// ObsceneFilter Information whether the obscene filter is enabled
	ObsceneFilter *BaseBoolInt `json:"obscene_filter,omitempty"`
	// ObsceneStopwords Information whether the stopwords filter is enabled
	ObsceneStopwords *BaseBoolInt `json:"obscene_stopwords,omitempty"`
	// ObsceneWords The list of stop words
	ObsceneWords string `json:"obscene_words,omitempty"`
	// Photos Photos settings
	Photos int64 `json:"photos,omitempty"`
	// Place
	Place *PlacesPlaceMin `json:"place,omitempty"`
	// PublicCategory Information about the group category
	PublicCategory int64 `json:"public_category,omitempty"`
	// PublicCategoryList
	PublicCategoryList []*GroupsGroupPublicCategoryList `json:"public_category_list,omitempty"`
	// PublicSubcategory Information about the group subcategory
	PublicSubcategory int64 `json:"public_subcategory,omitempty"`
	// Rss URL of the RSS feed
	Rss string `json:"rss,omitempty"`
	// Subject Community subject ID
	Subject int64 `json:"subject,omitempty"`
	// SubjectList
	SubjectList []*GroupsSubjectItem `json:"subject_list,omitempty"`
	// Title Community title
	Title string `json:"title,omitempty"`
	// Topics Topics settings
	Topics int64 `json:"topics,omitempty"`
	// Video Video settings
	Video int64 `json:"video,omitempty"`
	// Wall Wall settings
	Wall int64 `json:"wall,omitempty"`
	// Website Community website
	Website string `json:"website,omitempty"`
	// Wiki Wiki settings
	Wiki int64 `json:"wiki,omitempty"`
}

// PhotosPhotoUploadResponse type from VK API Schema(photos_photo_upload_response).
type PhotosPhotoUploadResponse struct {

	// Aid Album ID
	Aid int64 `json:"aid,omitempty"`
	// Hash Uploading hash
	Hash string `json:"hash,omitempty"`
	// PhotosList Uploaded photos data
	PhotosList string `json:"photos_list,omitempty"`
	// Server Upload server number
	Server int64 `json:"server,omitempty"`
}

// UsersSchool type from VK API Schema(users_school).
type UsersSchool struct {

	// City City ID
	City int64 `json:"city,omitempty"`
	// Class School class letter
	Class string `json:"class,omitempty"`
	// Country Country ID
	Country int64 `json:"country,omitempty"`
	// ID School ID
	ID string `json:"id,omitempty"`
	// Name School name
	Name string `json:"name,omitempty"`
	// Type School type ID
	Type int64 `json:"type,omitempty"`
	// TypeStr School type name
	TypeStr string `json:"type_str,omitempty"`
	// YearFrom Year the user started to study
	YearFrom int64 `json:"year_from,omitempty"`
	// YearGraduated Graduation year
	YearGraduated int64 `json:"year_graduated,omitempty"`
	// YearTo Year the user finished to study
	YearTo int64 `json:"year_to,omitempty"`
}

// VideoVideoFull type from VK API Schema(video_video_full).
type VideoVideoFull struct {

	// AccessKey Video access key
	AccessKey string `json:"access_key,omitempty"`
	// AddingDate Date when the video has been added in Unixtime
	AddingDate int64 `json:"adding_date,omitempty"`
	// CanAdd Information whether current user can add the video
	CanAdd *BaseBoolInt `json:"can_add,omitempty"`
	// CanComment Information whether current user can comment the video
	CanComment *BaseBoolInt `json:"can_comment,omitempty"`
	// CanEdit Information whether current user can edit the video
	CanEdit *BaseBoolInt `json:"can_edit,omitempty"`
	// CanRepost Information whether current user can comment the video
	CanRepost *BaseBoolInt `json:"can_repost,omitempty"`
	// Comments Number of comments
	Comments int64 `json:"comments,omitempty"`
	// Date Date when video has been uploaded in Unixtime
	Date int64 `json:"date,omitempty"`
	// Description Video description
	Description string `json:"description,omitempty"`
	// Duration Video duration in seconds
	Duration int64 `json:"duration,omitempty"`
	// Files
	Files *VideoVideoFiles `json:"files,omitempty"`
	// ID Video ID
	ID int64 `json:"id,omitempty"`
	// Likes
	Likes *BaseLikes `json:"likes,omitempty"`
	// Live Returns if the video is live translation
	Live *BasePropertyExists `json:"live,omitempty"`
	// OwnerID Video owner ID
	OwnerID int64 `json:"owner_id,omitempty"`
	// Photo130 URL of the preview image with 130 px in width
	Photo130 string `json:"photo_130,omitempty"`
	// Photo320 URL of the preview image with 320 px in width
	Photo320 string `json:"photo_320,omitempty"`
	// Photo800 URL of the preview image with 800 px in width
	Photo800 string `json:"photo_800,omitempty"`
	// Player URL of the page with a player that can be used to play the video in the browser.
	Player string `json:"player,omitempty"`
	// PrivacyComment
	PrivacyComment []string `json:"privacy_comment,omitempty"`
	// PrivacyView
	PrivacyView []string `json:"privacy_view,omitempty"`
	// Processing Returns if the video is processing
	Processing *BasePropertyExists `json:"processing,omitempty"`
	// Repeat Information whether the video is repeated
	Repeat *BaseBoolInt `json:"repeat,omitempty"`
	// Title Video title
	Title string `json:"title,omitempty"`
	// Views Number of views
	Views int64 `json:"views,omitempty"`
}

// AdsAdApproved type from VK API Schema(ads_ad_approved). Review status
type AdsAdApproved int64

// AdsRejectReason type from VK API Schema(ads_reject_reason).
type AdsRejectReason struct {

	// Comment Comment text
	Comment string `json:"comment,omitempty"`
	// Rules
	Rules []*AdsRules `json:"rules,omitempty"`
}

// FriendsRequests type from VK API Schema(friends_requests).
type FriendsRequests struct {

	// From ID of the user by whom friend has been suggested
	From string `json:"from,omitempty"`
	// Mutual
	Mutual *FriendsRequestsMutual `json:"mutual,omitempty"`
	// UserID User ID
	UserID int64 `json:"user_id,omitempty"`
}

// GroupsOwnerXtrBanInfoType type from VK API Schema(groups_owner_xtr_ban_info_type). Owner type
type GroupsOwnerXtrBanInfoType string

// UtilsLinkChecked type from VK API Schema(utils_link_checked).
type UtilsLinkChecked struct {

	// Link Link URL
	Link string `json:"link,omitempty"`
	// Status
	Status *UtilsLinkCheckedStatus `json:"status,omitempty"`
}

// WidgetsWidgetComment type from VK API Schema(widgets_widget_comment).
type WidgetsWidgetComment struct {

	// Attachments
	Attachments []*WallCommentAttachment `json:"attachments,omitempty"`
	// CanDelete Information whether current user can delete the comment
	CanDelete *BaseBoolInt `json:"can_delete,omitempty"`
	// Comments
	Comments *WidgetsCommentReplies `json:"comments,omitempty"`
	// Date Date when the comment has been added in Unixtime
	Date int64 `json:"date"`
	// FromID Comment author ID
	FromID int64 `json:"from_id"`
	// ID Comment ID
	ID int64 `json:"id"`
	// Likes
	Likes *BaseLikesInfo `json:"likes,omitempty"`
	// Media
	Media *WidgetsCommentMedia `json:"media,omitempty"`
	// PostSource
	PostSource *WallPostSource `json:"post_source,omitempty"`
	// PostType Post type
	PostType int64 `json:"post_type"`
	// Reposts
	Reposts *BaseRepostsInfo `json:"reposts,omitempty"`
	// Text Comment text
	Text string `json:"text"`
	// ToID Wall owner
	ToID int64 `json:"to_id"`
	// User
	User *UsersUserFull `json:"user,omitempty"`
}

// WidgetsWidgetLikes type from VK API Schema(widgets_widget_likes).
type WidgetsWidgetLikes struct {

	// Count Likes number
	Count int64 `json:"count,omitempty"`
}

// AdsStatsFormat type from VK API Schema(ads_stats_format).
type AdsStatsFormat struct {

	// Clicks Clicks number
	Clicks int64 `json:"clicks,omitempty"`
	// Day Day as YYYY-MM-DD
	Day string `json:"day,omitempty"`
	// Impressions Impressions number
	Impressions int64 `json:"impressions,omitempty"`
	// JoinRate Events number
	JoinRate int64 `json:"join_rate,omitempty"`
	// Month Month as YYYY-MM
	Month string `json:"month,omitempty"`
	// Overall 1 if period=overall
	Overall int64 `json:"overall,omitempty"`
	// Reach Reach
	Reach int64 `json:"reach,omitempty"`
	// Spent Spent funds
	Spent int64 `json:"spent,omitempty"`
	// VideoClicksSite Clickthoughs to the advertised site
	VideoClicksSite int64 `json:"video_clicks_site,omitempty"`
	// VideoViews Video views number
	VideoViews int64 `json:"video_views,omitempty"`
	// VideoViewsFull Video views (full video)
	VideoViewsFull int64 `json:"video_views_full,omitempty"`
	// VideoViewsHalf Video views (half of video)
	VideoViewsHalf int64 `json:"video_views_half,omitempty"`
}

// AdsStatsAge type from VK API Schema(ads_stats_age).
type AdsStatsAge struct {

	// ClicksRate Clicks rate
	ClicksRate float64 `json:"clicks_rate,omitempty"`
	// ImpressionsRate Impressions rate
	ImpressionsRate float64 `json:"impressions_rate,omitempty"`
	// Value Age interval
	Value string `json:"value,omitempty"`
}

// BaseLinkProduct type from VK API Schema(base_link_product).
type BaseLinkProduct struct {

	// Price
	Price *MarketPrice `json:"price,omitempty"`
}

// MessagesChatSettingsPhoto type from VK API Schema(messages_chat_settings_photo).
type MessagesChatSettingsPhoto struct {

	// Photo100 URL of the preview image with 100px in width
	Photo100 string `json:"photo_100,omitempty"`
	// Photo200 URL of the preview image with 200px in width
	Photo200 string `json:"photo_200,omitempty"`
	// Photo50 URL of the preview image with 50px in width
	Photo50 string `json:"photo_50,omitempty"`
}

// PagesPrivacySettings type from VK API Schema(pages_privacy_settings).
type PagesPrivacySettings int64

// AccountPushParamsSettings type from VK API Schema(account_push_params_settings). Settings parameters
type AccountPushParamsSettings string

// DocsDocPreview type from VK API Schema(docs_doc_preview).
type DocsDocPreview struct {

	// Photo
	Photo *DocsDocPreviewPhoto `json:"photo,omitempty"`
	// Video
	Video *DocsDocPreviewVideo `json:"video,omitempty"`
}

// AdsAdCostType type from VK API Schema(ads_ad_cost_type). Cost type
type AdsAdCostType int64

// AdsCampaign type from VK API Schema(ads_campaign).
type AdsCampaign struct {

	// AllLimit Campaign&#39;s total limit, rubles
	AllLimit string `json:"all_limit"`
	// DayLimit Campaign&#39;s day limit, rubles
	DayLimit string `json:"day_limit"`
	// ID Campaign ID
	ID int64 `json:"id"`
	// Name Campaign title
	Name string `json:"name"`
	// StartTime Campaign start time, as Unixtime
	StartTime int64 `json:"start_time"`
	// Status
	Status *AdsCampaignStatus `json:"status"`
	// StopTime Campaign stop time, as Unixtime
	StopTime int64 `json:"stop_time"`
	// Type
	Type *AdsCampaignType `json:"type"`
}

// DocsDocTypes type from VK API Schema(docs_doc_types).
type DocsDocTypes struct {

	// Count Number of docs
	Count int64 `json:"count,omitempty"`
	// ID Doc type ID
	ID int64 `json:"id,omitempty"`
	// Title Doc type title
	Title string `json:"title,omitempty"`
}

// NewsfeedNewsfeedNote type from VK API Schema(newsfeed_newsfeed_note).
type NewsfeedNewsfeedNote struct {

	// Comments Comments Number
	Comments int64 `json:"comments,omitempty"`
	// ID Note ID
	ID int64 `json:"id,omitempty"`
	// OwnerID integer
	OwnerID int64 `json:"owner_id,omitempty"`
	// Title Note title
	Title string `json:"title,omitempty"`
}

// VideoVideoFiles type from VK API Schema(video_video_files).
type VideoVideoFiles struct {

	// External URL of the external player
	External string `json:"external,omitempty"`
	// Mp1080 URL of the mpeg4 file with 1080p quality
	Mp1080 string `json:"mp_1080,omitempty"`
	// Mp240 URL of the mpeg4 file with 240p quality
	Mp240 string `json:"mp_240,omitempty"`
	// Mp360 URL of the mpeg4 file with 360p quality
	Mp360 string `json:"mp_360,omitempty"`
	// Mp480 URL of the mpeg4 file with 480p quality
	Mp480 string `json:"mp_480,omitempty"`
	// Mp720 URL of the mpeg4 file with 720p quality
	Mp720 string `json:"mp_720,omitempty"`
}

// SearchHintType type from VK API Schema(search_hint_type). Object type
type SearchHintType string

// StoriesStoryStatsState type from VK API Schema(stories_story_stats_state). Statistic state
type StoriesStoryStatsState string

// FriendsRequestsXtrMessage type from VK API Schema(friends_requests_xtr_message).
type FriendsRequestsXtrMessage struct {

	// From ID of the user by whom friend has been suggested
	From string `json:"from,omitempty"`
	// Message Message sent with a request
	Message string `json:"message,omitempty"`
	// Mutual
	Mutual *FriendsRequestsMutual `json:"mutual,omitempty"`
	// UserID User ID
	UserID int64 `json:"user_id,omitempty"`
}

// GiftsLayout type from VK API Schema(gifts_layout).
type GiftsLayout struct {

	// ID Gift ID
	ID int64 `json:"id,omitempty"`
	// Thumb256 URL of the preview image with 256 px in width
	Thumb256 string `json:"thumb_256,omitempty"`
	// Thumb48 URL of the preview image with 48 px in width
	Thumb48 string `json:"thumb_48,omitempty"`
	// Thumb96 URL of the preview image with 96 px in width
	Thumb96 string `json:"thumb_96,omitempty"`
}

// MessagesMessageAction type from VK API Schema(messages_message_action).
type MessagesMessageAction struct {

	// ConversationMessageID Message ID
	ConversationMessageID int64 `json:"conversation_message_id,omitempty"`
	// Email Email address for chat_invite_user or chat_kick_user actions
	Email string `json:"email,omitempty"`
	// MemberID User or email peer ID
	MemberID int64 `json:"member_id,omitempty"`
	// Message Message body of related message
	Message string `json:"message,omitempty"`
	// Photo NO Struct in JSON SCHEMA
	// Photo *MessagesMessageActionPhoto `json:"photo,omitempty"`
	// Text New chat title for chat_create and chat_title_update actions
	Text string `json:"text,omitempty"`
	// Type
	Type *MessagesMessageActionStatus `json:"type"`
}

// MessagesChatSettingsState type from VK API Schema(messages_chat_settings_state).
type MessagesChatSettingsState string

// PagesWikipageVersion type from VK API Schema(pages_wikipage_version).
type PagesWikipageVersion struct {

	// Edited Date when the page has been edited in Unixtime
	Edited int64 `json:"edited,omitempty"`
	// EditorID Last editor ID
	EditorID int64 `json:"editor_id,omitempty"`
	// EditorName Last editor name
	EditorName string `json:"editor_name,omitempty"`
	// ID Version ID
	ID int64 `json:"id,omitempty"`
	// Length Page size in bytes
	Length int64 `json:"length,omitempty"`
}

// GroupsGroupsArray type from VK API Schema(groups_groups_array).
type GroupsGroupsArray struct {

	// Count Communities number
	Count int64 `json:"count"`
	// Items
	Items []int64 `json:"items"`
}

// WallWallpostAttached type from VK API Schema(wall_wallpost_attached).
type WallWallpostAttached struct {

	// Attachments
	Attachments []*WallWallpostAttachment `json:"attachments,omitempty"`
	// CanDelete Information whether current user can delete the post
	CanDelete *BaseBoolInt `json:"can_delete,omitempty"`
	// Comments
	Comments *BaseCommentsInfo `json:"comments,omitempty"`
	// CopyOwnerID Source post owner&#39;s ID
	CopyOwnerID int64 `json:"copy_owner_id,omitempty"`
	// CopyPostID Source post ID
	CopyPostID int64 `json:"copy_post_id,omitempty"`
	// CopyText Repost comment
	CopyText string `json:"copy_text,omitempty"`
	// Date Date of publishing in Unixtime
	Date int64 `json:"date,omitempty"`
	// FromID Post author ID
	FromID int64 `json:"from_id,omitempty"`
	// Geo
	Geo *BaseGeo `json:"geo,omitempty"`
	// ID Post ID
	ID int64 `json:"id,omitempty"`
	// Likes
	Likes *BaseLikesInfo `json:"likes,omitempty"`
	// PostSource
	PostSource *WallPostSource `json:"post_source,omitempty"`
	// PostType
	PostType *WallPostType `json:"post_type,omitempty"`
	// Reposts
	Reposts *BaseRepostsInfo `json:"reposts,omitempty"`
	// SignerID Post signer ID
	SignerID int64 `json:"signer_id,omitempty"`
	// Text Post text
	Text string `json:"text,omitempty"`
	// ToID Post addresse
	ToID int64 `json:"to_id,omitempty"`
}

// AdsAccesses type from VK API Schema(ads_accesses).
type AdsAccesses struct {

	// ClientID Client ID
	ClientID string `json:"client_id,omitempty"`
	// Role
	Role *AdsAccessRole `json:"role,omitempty"`
}

// MarketSection type from VK API Schema(market_section).
type MarketSection struct {

	// ID Section ID
	ID int64 `json:"id"`
	// Name Section name
	Name string `json:"name"`
}

// NotificationsReply type from VK API Schema(notifications_reply).
type NotificationsReply struct {

	// Date Date when the reply has been created in Unixtime
	Date int64 `json:"date,omitempty"`
	// ID Reply ID
	ID int64 `json:"id,omitempty"`
	// Text Reply text
	Text int64 `json:"text,omitempty"`
}

// PlacesCheckin type from VK API Schema(places_checkin).
type PlacesCheckin struct {

	// Date Date when the checkin has been added in Unixtime
	Date int64 `json:"date"`
	// Distance Distance to the place
	Distance int64 `json:"distance,omitempty"`
	// ID Checkin ID
	ID int64 `json:"id"`
	// Latitude Place latitude
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude Place longitude
	Longitude float64 `json:"longitude,omitempty"`
	// PlaceCity City ID
	PlaceCity int64 `json:"place_city,omitempty"`
	// PlaceCountry Country ID
	PlaceCountry int64 `json:"place_country,omitempty"`
	// PlaceIcon URL of the place&#39;s icon
	PlaceIcon string `json:"place_icon,omitempty"`
	// PlaceID Place ID
	PlaceID int64 `json:"place_id,omitempty"`
	// PlaceTitle Place title
	PlaceTitle string `json:"place_title,omitempty"`
	// PlaceType Place type
	PlaceType string `json:"place_type,omitempty"`
	// Text Comment text
	Text string `json:"text,omitempty"`
	// UserID User ID
	UserID int64 `json:"user_id"`
}

// StatsCity type from VK API Schema(stats_city).
type StatsCity struct {

	// Count Visitors number
	Count int64 `json:"count,omitempty"`
	// Name City name
	Name string `json:"name,omitempty"`
	// Value City ID
	Value int64 `json:"value,omitempty"`
}

// DatabaseSchool type from VK API Schema(database_school).
type DatabaseSchool struct {

	// ID School ID
	ID int64 `json:"id,omitempty"`
	// Title School title
	Title string `json:"title,omitempty"`
}

// AccountInfo type from VK API Schema(account_info).
type AccountInfo struct {

	// 2faRequired Two factor authentication is enabled
	TwofaRequired *BaseBoolInt `json:"2fa_required,omitempty"`
	// Country Country code
	Country string `json:"country,omitempty"`
	// HTTPSRequired Information whether HTTPS-only is enabled
	HTTPSRequired *BaseBoolInt `json:"https_required,omitempty"`
	// Intro Information whether user has been processed intro
	Intro *BaseBoolInt `json:"intro,omitempty"`
	// Lang Language ID
	Lang int64 `json:"lang,omitempty"`
	// NoWallReplies Information whether wall comments should be hidden
	NoWallReplies *BaseBoolInt `json:"no_wall_replies,omitempty"`
	// OwnPostsDefault Information whether only owners posts should be shown
	OwnPostsDefault *BaseBoolInt `json:"own_posts_default,omitempty"`
}

// AdsClient type from VK API Schema(ads_client).
type AdsClient struct {

	// AllLimit Client&#39;s total limit, rubles
	AllLimit string `json:"all_limit"`
	// DayLimit Client&#39;s day limit, rubles
	DayLimit string `json:"day_limit"`
	// ID Client ID
	ID int64 `json:"id"`
	// Name Client name
	Name string `json:"name"`
}

// AdsCriteriaSex type from VK API Schema(ads_criteria_sex). Sex
type AdsCriteriaSex int64

// AdsStatsSexValue type from VK API Schema(ads_stats_sex_value). Sex
type AdsStatsSexValue string

// AdsStatsSex type from VK API Schema(ads_stats_sex).
type AdsStatsSex struct {

	// ClicksRate Clicks rate
	ClicksRate float64 `json:"clicks_rate,omitempty"`
	// ImpressionsRate Impressions rate
	ImpressionsRate float64 `json:"impressions_rate,omitempty"`
	// Value
	Value *AdsStatsSexValue `json:"value,omitempty"`
}

// MessagesMessage type from VK API Schema(messages_message).
type MessagesMessage struct {

	// Action
	Action *MessagesMessageAction `json:"action,omitempty"`
	// Attachments
	Attachments []*MessagesMessageAttachment `json:"attachments,omitempty"`
	// ConversationMessageID Unique auto-incremented number for all messages with this peer
	ConversationMessageID int64 `json:"conversation_message_id,omitempty"`
	// Date Date when the message has been sent in Unixtime
	Date int64 `json:"date"`
	// FromID Message author&#39;s ID
	FromID int64 `json:"from_id"`
	// FwdMessages Forwarded messages
	FwdMessages []*MessagesMessage `json:"fwd_messages,omitempty"`
	// Geo
	Geo *BaseGeo `json:"geo,omitempty"`
	// ID Message ID
	ID int64 `json:"id"`
	// Important Is it an important message
	Important bool `json:"important,omitempty"`
	// Keyboard
	Keyboard *MessagesKeyboard `json:"keyboard,omitempty"`
	// Payload
	Payload string `json:"payload,omitempty"`
	// PeerID Peer ID
	PeerID int64 `json:"peer_id"`
	// RandomID ID used for sending messages. It returned only for outgoing messages
	RandomID int64 `json:"random_id,omitempty"`
	// Text Message text
	Text string `json:"text"`
	// UpdateTime Date when the message has been updated in Unixtime
	UpdateTime int64 `json:"update_time,omitempty"`
}

// UsersOccupation type from VK API Schema(users_occupation).
type UsersOccupation struct {

	// ID ID of school, university, company group
	ID int64 `json:"id,omitempty"`
	// Name Name of occupation
	Name string `json:"name,omitempty"`
	// Type Type of occupation
	Type string `json:"type,omitempty"`
}

// WallWallpostAttachment type from VK API Schema(wall_wallpost_attachment).
type WallWallpostAttachment struct {

	// Album
	Album *PhotosPhotoAlbum `json:"album,omitempty"`
	// App
	App *WallAppPost `json:"app,omitempty"`
	// Audio
	Audio *AudioAudioFull `json:"audio,omitempty"`
	// Doc
	Doc *DocsDoc `json:"doc,omitempty"`
	// Graffiti
	Graffiti *WallGraffiti `json:"graffiti,omitempty"`
	// Link
	Link *BaseLink `json:"link,omitempty"`
	// Market
	Market *MarketMarketItem `json:"market,omitempty"`
	// MarketMarketAlbum
	MarketMarketAlbum *MarketMarketAlbum `json:"market_market_album,omitempty"`
	// Note
	Note *WallAttachedNote `json:"note,omitempty"`
	// Page
	Page *PagesWikipageFull `json:"page,omitempty"`
	// Photo
	Photo *PhotosPhoto `json:"photo,omitempty"`
	// PhotosList
	PhotosList []string `json:"photos_list,omitempty"`
	// Poll
	Poll *PollsPoll `json:"poll,omitempty"`
	// PostedPhoto
	PostedPhoto *WallPostedPhoto `json:"posted_photo,omitempty"`
	// Type
	Type *WallWallpostAttachmentType `json:"type"`
	// Video
	Video *VideoVideo `json:"video,omitempty"`
}

// AccountPushConversationsItem type from VK API Schema(account_push_conversations_item).
type AccountPushConversationsItem struct {

	// DisabledUntil Time until that notifications are disabled in seconds
	DisabledUntil int64 `json:"disabled_until,omitempty"`
	// PeerID Peer ID
	PeerID int64 `json:"peer_id,omitempty"`
	// Sound Information whether the sound are enabled
	Sound *BaseBoolInt `json:"sound,omitempty"`
}

// AdsDemostatsFormat type from VK API Schema(ads_demostats_format).
type AdsDemostatsFormat struct {

	// Age
	Age []*AdsStatsAge `json:"age,omitempty"`
	// Cities
	Cities []*AdsStatsCities `json:"cities,omitempty"`
	// Day Day as YYYY-MM-DD
	Day string `json:"day,omitempty"`
	// Month Month as YYYY-MM
	Month string `json:"month,omitempty"`
	// Overall 1 if period=overall
	Overall int64 `json:"overall,omitempty"`
	// Sex
	Sex []*AdsStatsSex `json:"sex,omitempty"`
	// SexAge
	SexAge []*AdsStatsSexAge `json:"sex_age,omitempty"`
}

// BasePlace type from VK API Schema(base_place).
type BasePlace struct {

	// Address Place address
	Address string `json:"address,omitempty"`
	// Checkins Checkins number
	Checkins int64 `json:"checkins,omitempty"`
	// City City name
	City string `json:"city,omitempty"`
	// Country Country name
	Country string `json:"country,omitempty"`
	// Created Date of the place creation in Unixtime
	Created int64 `json:"created,omitempty"`
	// Icon URL of the place&#39;s icon
	Icon string `json:"icon,omitempty"`
	// ID Place ID
	ID int64 `json:"id,omitempty"`
	// Latitude Place latitude
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude Place longitude
	Longitude float64 `json:"longitude,omitempty"`
	// Title Place title
	Title string `json:"title,omitempty"`
	// Type Place type
	Type string `json:"type,omitempty"`
}

// DocsDocUploadResponse type from VK API Schema(docs_doc_upload_response).
type DocsDocUploadResponse struct {

	// File Uploaded file data
	File string `json:"file,omitempty"`
}

// GroupsMemberRoleStatus type from VK API Schema(groups_member_role_status). User&#39;s credentials as community admin
type GroupsMemberRoleStatus string

// FriendsFriendsList type from VK API Schema(friends_friends_list).
type FriendsFriendsList struct {

	// ID List ID
	ID int64 `json:"id"`
	// Name List title
	Name string `json:"name"`
}

// FriendsFriendStatusStatus type from VK API Schema(friends_friend_status_status). Friend status with the user
type FriendsFriendStatusStatus int64

// NewsfeedItemTopic type from VK API Schema(newsfeed_item_topic).
type NewsfeedItemTopic struct {

	// Comments
	Comments *BaseCommentsInfo `json:"comments,omitempty"`
	// Likes
	Likes *BaseLikesInfo `json:"likes,omitempty"`
	// PostID Topic post ID
	PostID int64 `json:"post_id"`
	// Text Post text
	Text string `json:"text"`
}

// UsersUserMin type from VK API Schema(users_user_min).
type UsersUserMin struct {

	// Deactivated Returns if a profile is deleted or blocked
	Deactivated string `json:"deactivated,omitempty"`
	// FirstName User first name
	FirstName string `json:"first_name"`
	// Hidden Returns if a profile is hidden.
	Hidden int64 `json:"hidden,omitempty"`
	// ID User ID
	ID int64 `json:"id"`
	// LastName User last name
	LastName string `json:"last_name"`
}

// AccountAccountCounters type from VK API Schema(account_account_counters).
type AccountAccountCounters struct {

	// AppRequests New app requests number
	AppRequests int64 `json:"app_requests,omitempty"`
	// Events New events number
	Events int64 `json:"events,omitempty"`
	// Friends New friends requests number
	Friends int64 `json:"friends,omitempty"`
	// FriendsSuggestions New friends suggestions number
	FriendsSuggestions int64 `json:"friends_suggestions,omitempty"`
	// Gifts New gifts number
	Gifts int64 `json:"gifts,omitempty"`
	// Groups New groups number
	Groups int64 `json:"groups,omitempty"`
	// Messages New messages number
	Messages int64 `json:"messages,omitempty"`
	// Notifications New notifications number
	Notifications int64 `json:"notifications,omitempty"`
	// Photos New photo tags number
	Photos int64 `json:"photos,omitempty"`
	// Videos New video tags number
	Videos int64 `json:"videos,omitempty"`
}

// GroupsGroupFullMainSection type from VK API Schema(groups_group_full_main_section). Main section of community
type GroupsGroupFullMainSection int64

// NewsfeedItemFriendFriends type from VK API Schema(newsfeed_item_friend_friends).
type NewsfeedItemFriendFriends struct {

	// Count Number of friends has been added
	Count int64 `json:"count,omitempty"`
	// Items
	Items []*BaseUserID `json:"items,omitempty"`
}

// OauthError type from VK API Schema(oauth_error).
type OauthError struct {

	// Error Error type
	Error string `json:"error"`
	// ErrorDescription Error description
	ErrorDescription string `json:"error_description"`
	// RedirectURI URI for validation
	RedirectURI string `json:"redirect_uri,omitempty"`
}

// StoriesStoryStatsStat type from VK API Schema(stories_story_stats_stat).
type StoriesStoryStatsStat struct {

	// Count Stat value
	Count int64 `json:"count,omitempty"`
	// State
	State *StoriesStoryStatsState `json:"state"`
}

// GroupsGroupFullAgeLimits type from VK API Schema(groups_group_full_age_limits).
type GroupsGroupFullAgeLimits int64

// UsersUniversity type from VK API Schema(users_university).
type UsersUniversity struct {

	// Chair Chair ID
	Chair int64 `json:"chair,omitempty"`
	// ChairName Chair name
	ChairName string `json:"chair_name,omitempty"`
	// City City ID
	City int64 `json:"city,omitempty"`
	// Country Country ID
	Country int64 `json:"country,omitempty"`
	// EducationForm Education form
	EducationForm string `json:"education_form,omitempty"`
	// EducationStatus Education status
	EducationStatus string `json:"education_status,omitempty"`
	// Faculty Faculty ID
	Faculty int64 `json:"faculty,omitempty"`
	// FacultyName Faculty name
	FacultyName string `json:"faculty_name,omitempty"`
	// Graduation Graduation year
	Graduation int64 `json:"graduation,omitempty"`
	// ID University ID
	ID int64 `json:"id,omitempty"`
	// Name University name
	Name string `json:"name,omitempty"`
}

// WallPostSource type from VK API Schema(wall_post_source).
type WallPostSource struct {

	// Data Additional data
	Data string `json:"data,omitempty"`
	// Platform Platform name
	Platform string `json:"platform,omitempty"`
	// Type
	Type *WallPostSourceType `json:"type,omitempty"`
	// URL URL to an external site used to publish the post
	URL string `json:"url,omitempty"`
}

// FriendsFriendStatus type from VK API Schema(friends_friend_status).
type FriendsFriendStatus struct {

	// FriendStatus
	FriendStatus *FriendsFriendStatusStatus `json:"friend_status"`
	// ReadState Information whether request is unviewed
	ReadState *BaseBoolInt `json:"read_state,omitempty"`
	// RequestMessage Message sent with request
	RequestMessage string `json:"request_message,omitempty"`
	// Sign MD5 hash for the result validation
	Sign string `json:"sign,omitempty"`
	// UserID User ID
	UserID int64 `json:"user_id"`
}

// NewsfeedItemPhotoTag type from VK API Schema(newsfeed_item_photo_tag).
type NewsfeedItemPhotoTag struct {

	// PhotoTags
	PhotoTags *NewsfeedItemPhotoTagPhotoTags `json:"photo_tags,omitempty"`
	// PostID Post ID
	PostID int64 `json:"post_id,omitempty"`
}

// VideoCatElement type from VK API Schema(video_cat_element).
type VideoCatElement struct {

	// CanAdd Information whether current user can add the video
	CanAdd *BaseBoolInt `json:"can_add,omitempty"`
	// CanEdit Information whether current user can edit the video
	CanEdit *BaseBoolInt `json:"can_edit,omitempty"`
	// Comments Comments number
	Comments int64 `json:"comments,omitempty"`
	// Count Videos number (for album)
	Count int64 `json:"count,omitempty"`
	// Date Date when the element has been created
	Date int64 `json:"date,omitempty"`
	// Description Element description
	Description string `json:"description,omitempty"`
	// Duration Duration in seconds
	Duration int64 `json:"duration,omitempty"`
	// ID Element ID
	ID int64 `json:"id"`
	// IsPrivate Information whether the video is private
	IsPrivate *BaseBoolInt `json:"is_private,omitempty"`
	// OwnerID Element owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// Photo130 URL of the preview image with 130 px in width
	Photo130 string `json:"photo_130,omitempty"`
	// Photo160 URL of the preview image with 160 px in width
	Photo160 string `json:"photo_160,omitempty"`
	// Photo320 URL of the preview image with 320 px in width
	Photo320 string `json:"photo_320,omitempty"`
	// Photo640 URL of the preview image with 640 px in width
	Photo640 string `json:"photo_640,omitempty"`
	// Photo800 URL of the preview image with 800 px in width
	Photo800 string `json:"photo_800,omitempty"`
	// Title Element title
	Title string `json:"title"`
	// Type
	Type *VideoCatElementType `json:"type"`
	// UpdatedTime Date of last update (for album) in Unixtime
	UpdatedTime int64 `json:"updated_time,omitempty"`
	// Views Views number
	Views int64 `json:"views,omitempty"`
}

// WallViews type from VK API Schema(wall_views).
type WallViews struct {

	// Count Count
	Count int64 `json:"count,omitempty"`
}

// MessagesChatPushSettings type from VK API Schema(messages_chat_push_settings).
type MessagesChatPushSettings struct {

	// DisabledUntil Time until that notifications are disabled
	DisabledUntil int64 `json:"disabled_until,omitempty"`
	// Sound Information whether the sound is on
	Sound *BaseBoolInt `json:"sound,omitempty"`
}

// NewsfeedItemVideoVideo type from VK API Schema(newsfeed_item_video_video).
type NewsfeedItemVideoVideo struct {

	// Count Tags number
	Count int64 `json:"count,omitempty"`
	// Items
	Items []*VideoVideo `json:"items,omitempty"`
}

// UsersUserFull type from VK API Schema(users_user_full).
type UsersUserFull struct {
}

// WallWallpostFull type from VK API Schema(wall_wallpost_full).
type WallWallpostFull struct {
}

// StatsActivity type from VK API Schema(stats_activity). Activity stats
type StatsActivity struct {

	// Comments Comments number
	Comments int64 `json:"comments,omitempty"`
	// Copies Reposts number
	Copies int64 `json:"copies,omitempty"`
	// Hidden Hidden from news count
	Hidden int64 `json:"hidden,omitempty"`
	// Likes Likes number
	Likes int64 `json:"likes,omitempty"`
	// Subscribed New subscribers count
	Subscribed int64 `json:"subscribed,omitempty"`
	// Unsubscribed Unsubscribed count
	Unsubscribed int64 `json:"unsubscribed,omitempty"`
}

// StatsCountry type from VK API Schema(stats_country).
type StatsCountry struct {

	// Code Country code
	Code string `json:"code,omitempty"`
	// Count Visitors number
	Count int64 `json:"count,omitempty"`
	// Name Country name
	Name string `json:"name,omitempty"`
	// Value Country ID
	Value int64 `json:"value,omitempty"`
}

// UtilsStatsCity type from VK API Schema(utils_stats_city).
type UtilsStatsCity struct {

	// CityID City ID
	CityID int64 `json:"city_id,omitempty"`
	// Views Views number
	Views int64 `json:"views,omitempty"`
}

// VideoCatElementType type from VK API Schema(video_cat_element_type). Element type
type VideoCatElementType string

// GroupsTokenPermissions type from VK API Schema(groups_token_permissions).
type GroupsTokenPermissions struct {

	// Mask
	Mask int64 `json:"mask"`
	// Permissions
	Permissions []*GroupsTokenPermissionSetting `json:"permissions,omitempty"`
}

// MarketPrice type from VK API Schema(market_price).
type MarketPrice struct {

	// Amount Amount
	Amount string `json:"amount,omitempty"`
	// Currency
	Currency *MarketCurrency `json:"currency,omitempty"`
	// Text Text
	Text string `json:"text,omitempty"`
}

// NewsfeedItemAudio type from VK API Schema(newsfeed_item_audio).
type NewsfeedItemAudio struct {

	// Audio
	Audio *NewsfeedItemAudioAudio `json:"audio,omitempty"`
	// PostID Post ID
	PostID int64 `json:"post_id,omitempty"`
}

// AccountPushParams type from VK API Schema(account_push_params).
type AccountPushParams struct {

	// AppRequest
	AppRequest []*AccountOnoffOptions `json:"app_request,omitempty"`
	// Birthday
	Birthday []*AccountOnoffOptions `json:"birthday,omitempty"`
	// Chat
	Chat []*AccountPushParamsMode `json:"chat,omitempty"`
	// Comment
	Comment []*AccountPushParamsSettings `json:"comment,omitempty"`
	// EventSoon
	EventSoon []*AccountOnoffOptions `json:"event_soon,omitempty"`
	// Friend
	Friend []*AccountOnoffOptions `json:"friend,omitempty"`
	// FriendAccepted
	FriendAccepted []*AccountOnoffOptions `json:"friend_accepted,omitempty"`
	// FriendFound
	FriendFound []*AccountOnoffOptions `json:"friend_found,omitempty"`
	// GroupAccepted
	GroupAccepted []*AccountOnoffOptions `json:"group_accepted,omitempty"`
	// GroupInvite
	GroupInvite []*AccountOnoffOptions `json:"group_invite,omitempty"`
	// Like
	Like []*AccountPushParamsSettings `json:"like,omitempty"`
	// Mention
	Mention []*AccountPushParamsSettings `json:"mention,omitempty"`
	// Msg
	Msg []*AccountPushParamsMode `json:"msg,omitempty"`
	// NewPost
	NewPost []*AccountOnoffOptions `json:"new_post,omitempty"`
	// PhotosTag
	PhotosTag []*AccountPushParamsSettings `json:"photos_tag,omitempty"`
	// Reply
	Reply []*AccountOnoffOptions `json:"reply,omitempty"`
	// Repost
	Repost []*AccountPushParamsSettings `json:"repost,omitempty"`
	// SdkOpen
	SdkOpen []*AccountOnoffOptions `json:"sdk_open,omitempty"`
	// WallPost
	WallPost []*AccountOnoffOptions `json:"wall_post,omitempty"`
	// WallPublish
	WallPublish []*AccountOnoffOptions `json:"wall_publish,omitempty"`
}

// AdsAdStatus type from VK API Schema(ads_ad_status). Ad atatus
type AdsAdStatus int64

// AdsTargSuggestionsRegions type from VK API Schema(ads_targ_suggestions_regions).
type AdsTargSuggestionsRegions struct {

	// ID Object ID
	ID int64 `json:"id,omitempty"`
	// Name Object name
	Name string `json:"name,omitempty"`
	// Type Object type
	Type string `json:"type,omitempty"`
}

// BaseUploadServer type from VK API Schema(base_upload_server).
type BaseUploadServer struct {

	// UploadURL Upload URL
	UploadURL string `json:"upload_url"`
}

// BaseLinkButtonActionType type from VK API Schema(base_link_button_action_type). Action type
type BaseLinkButtonActionType string

// UsersUserBroadcast type from VK API Schema(users_user_broadcast).
type UsersUserBroadcast struct {
}

// AccountOffer type from VK API Schema(account_offer).
type AccountOffer struct {

	// Description Offer description
	Description string `json:"description,omitempty"`
	// ID Offer ID
	ID int64 `json:"id,omitempty"`
	// Img URL of the preview image
	Img string `json:"img,omitempty"`
	// Instruction Instruction how to process the offer
	Instruction string `json:"instruction,omitempty"`
	// InstructionHTML Instruction how to process the offer (HTML format)
	InstructionHTML string `json:"instruction_html,omitempty"`
	// Price Offer price
	Price int64 `json:"price,omitempty"`
	// ShortDescription Offer short description
	ShortDescription string `json:"short_description,omitempty"`
	// Tag Offer tag
	Tag string `json:"tag,omitempty"`
	// Title Offer title
	Title string `json:"title,omitempty"`
}

// MessagesUserXtrInvitedBy type from VK API Schema(messages_user_xtr_invited_by).
type MessagesUserXtrInvitedBy struct {
}

// NewsfeedItemPhotoTagPhotoTags type from VK API Schema(newsfeed_item_photo_tag_photo_tags).
type NewsfeedItemPhotoTagPhotoTags struct {

	// Count Tags number
	Count int64 `json:"count,omitempty"`
	// Items
	Items []*NewsfeedNewsfeedPhoto `json:"items,omitempty"`
}

// UtilsDomainResolvedType type from VK API Schema(utils_domain_resolved_type). Object type
type UtilsDomainResolvedType string

// WidgetsCommentMediaType type from VK API Schema(widgets_comment_media_type). Media type
type WidgetsCommentMediaType string

// NewsfeedNewsfeedPhoto type from VK API Schema(newsfeed_newsfeed_photo).
type NewsfeedNewsfeedPhoto struct {
}

// StoriesStoryType type from VK API Schema(stories_story_type). Story type.
type StoriesStoryType string

// UsersRelative type from VK API Schema(users_relative).
type UsersRelative struct {

	// ID Relative ID
	ID int64 `json:"id,omitempty"`
	// Type Relative type
	Type string `json:"type,omitempty"`
}

// AccountNameRequest type from VK API Schema(account_name_request).
type AccountNameRequest struct {

	// FirstName First name in request
	FirstName string `json:"first_name,omitempty"`
	// ID Request ID needed to cancel the request
	ID int64 `json:"id,omitempty"`
	// LastName Last name in request
	LastName string `json:"last_name,omitempty"`
	// Status
	Status *AccountNameRequestStatus `json:"status,omitempty"`
}

// AccountPushSettings type from VK API Schema(account_push_settings).
type AccountPushSettings struct {

	// Conversations
	Conversations *AccountPushConversations `json:"conversations,omitempty"`
	// Disabled Information whether notifications are disabled
	Disabled *BaseBoolInt `json:"disabled,omitempty"`
	// DisabledUntil Time until that notifications are disabled in Unixtime
	DisabledUntil int64 `json:"disabled_until,omitempty"`
	// Settings
	Settings *AccountPushParams `json:"settings,omitempty"`
}

// BoardTopic type from VK API Schema(board_topic).
type BoardTopic struct {

	// Comments Comments number
	Comments int64 `json:"comments,omitempty"`
	// Created Date when the topic has been created in Unixtime
	Created int64 `json:"created,omitempty"`
	// CreatedBy Creator ID
	CreatedBy int64 `json:"created_by,omitempty"`
	// ID Topic ID
	ID int64 `json:"id,omitempty"`
	// IsClosed Information whether the topic is closed
	IsClosed *BaseBoolInt `json:"is_closed,omitempty"`
	// IsFixed Information whether the topic is fixed
	IsFixed *BaseBoolInt `json:"is_fixed,omitempty"`
	// Title Topic title
	Title string `json:"title,omitempty"`
	// Updated Date when the topic has been updated in Unixtime
	Updated int64 `json:"updated,omitempty"`
	// UpdatedBy ID of user who updated the topic
	UpdatedBy int64 `json:"updated_by,omitempty"`
}

// FaveFavesLink type from VK API Schema(fave_faves_link).
type FaveFavesLink struct {

	// Description Link description
	Description string `json:"description,omitempty"`
	// ID Link ID
	ID int64 `json:"id,omitempty"`
	// Photo100 URL of the preview image with 100 px in width
	Photo100 string `json:"photo_100,omitempty"`
	// Photo200 URL of the preview image with 200 px in width
	Photo200 string `json:"photo_200,omitempty"`
	// Photo50 URL of the preview image with 50 px in width
	Photo50 string `json:"photo_50,omitempty"`
	// Title Link title
	Title string `json:"title,omitempty"`
	// URL Link URL
	URL string `json:"url,omitempty"`
}

// LeadsCheckedResult type from VK API Schema(leads_checked_result). Information whether user can start the lead
type LeadsCheckedResult string

// MessagesHistoryMessageAttachmentType type from VK API Schema(messages_history_message_attachment_type). Attachments type
type MessagesHistoryMessageAttachmentType string

// WallAttachedNote type from VK API Schema(wall_attached_note).
type WallAttachedNote struct {

	// Comments Comments number
	Comments int64 `json:"comments"`
	// Date Date when the note has been created in Unixtime
	Date int64 `json:"date"`
	// ID Note ID
	ID int64 `json:"id"`
	// OwnerID Note owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// ReadComments Read comments number
	ReadComments int64 `json:"read_comments"`
	// Title Note title
	Title string `json:"title"`
	// ViewURL URL of the page with note preview
	ViewURL string `json:"view_url"`
}

// WallCommentAttachment type from VK API Schema(wall_comment_attachment).
type WallCommentAttachment struct {

	// Audio
	Audio *AudioAudioFull `json:"audio,omitempty"`
	// Doc
	Doc *DocsDoc `json:"doc,omitempty"`
	// Link
	Link *BaseLink `json:"link,omitempty"`
	// Market
	Market *MarketMarketItem `json:"market,omitempty"`
	// MarketMarketAlbum
	MarketMarketAlbum *MarketMarketAlbum `json:"market_market_album,omitempty"`
	// Note
	Note *WallAttachedNote `json:"note,omitempty"`
	// Page
	Page *PagesWikipageFull `json:"page,omitempty"`
	// Photo
	Photo *PhotosPhoto `json:"photo,omitempty"`
	// Sticker
	Sticker *BaseSticker `json:"sticker,omitempty"`
	// Type
	Type *WallCommentAttachmentType `json:"type"`
	// Video
	Video *VideoVideo `json:"video,omitempty"`
}

// UtilsStatsSexAge type from VK API Schema(utils_stats_sex_age).
type UtilsStatsSexAge struct {

	// AgeRange Age denotation
	AgeRange string `json:"age_range,omitempty"`
	// Female  Views by female users
	Female int64 `json:"female,omitempty"`
	// Male  Views by male users
	Male int64 `json:"male,omitempty"`
}

// AccountUserXtrContact type from VK API Schema(account_user_xtr_contact).
type AccountUserXtrContact struct {
}

// BoardTopicPoll type from VK API Schema(board_topic_poll).
type BoardTopicPoll struct {

	// AnswerID Current user&#39;s answer ID
	AnswerID int64 `json:"answer_id"`
	// Answers
	Answers []*PollsAnswer `json:"answers"`
	// Created Date when poll has been created in Unixtime
	Created int64 `json:"created"`
	// IsClosed Information whether the poll is closed
	IsClosed *BaseBoolInt `json:"is_closed,omitempty"`
	// OwnerID Poll owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// PollID Poll ID
	PollID int64 `json:"poll_id"`
	// Question Poll question
	Question string `json:"question"`
	// Votes Votes number
	Votes string `json:"votes"`
}

// GroupsTokenPermissionSetting type from VK API Schema(groups_token_permission_setting).
type GroupsTokenPermissionSetting struct {

	// Name
	Name string `json:"name"`
	// Setting
	Setting int64 `json:"setting"`
}

// GroupsOnlineStatusType type from VK API Schema(groups_online_status_type). Type of online status of group
type GroupsOnlineStatusType string

// MarketMarketItem type from VK API Schema(market_market_item).
type MarketMarketItem struct {

	// Availability
	Availability *MarketMarketItemAvailability `json:"availability"`
	// Category
	Category *MarketMarketCategory `json:"category"`
	// Date Date when the item has been created in Unixtime
	Date int64 `json:"date"`
	// Description Item description
	Description string `json:"description"`
	// ID Item ID
	ID int64 `json:"id"`
	// OwnerID Item owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// Price
	Price *MarketPrice `json:"price"`
	// ThumbPhoto URL of the preview image
	ThumbPhoto string `json:"thumb_photo"`
	// Title Item title
	Title string `json:"title"`
}

// MessagesLongpollMessages type from VK API Schema(messages_longpoll_messages).
type MessagesLongpollMessages struct {

	// Count Total number
	Count int64 `json:"count,omitempty"`
	// Items
	Items []*MessagesMessage `json:"items,omitempty"`
}

// NewsfeedItemWallpost type from VK API Schema(newsfeed_item_wallpost).
type NewsfeedItemWallpost struct {

	// Attachments
	Attachments []*WallWallpostAttachment `json:"attachments,omitempty"`
	// Comments
	Comments *BaseCommentsInfo `json:"comments,omitempty"`
	// CopyHistory
	CopyHistory []*WallWallpost `json:"copy_history,omitempty"`
	// Geo
	Geo *BaseGeo `json:"geo,omitempty"`
	// Likes
	Likes *BaseLikesInfo `json:"likes,omitempty"`
	// PostID Post ID
	PostID int64 `json:"post_id,omitempty"`
	// PostSource
	PostSource *WallPostSource `json:"post_source,omitempty"`
	// PostType
	PostType *NewsfeedItemWallpostType `json:"post_type,omitempty"`
	// Reposts
	Reposts *BaseRepostsInfo `json:"reposts,omitempty"`
	// Text Post text
	Text string `json:"text,omitempty"`
}

// UsersExports type from VK API Schema(users_exports).
type UsersExports struct {

	// Facebook
	Facebook int64 `json:"facebook,omitempty"`
	// Livejournal
	Livejournal int64 `json:"livejournal,omitempty"`
	// Twitter
	Twitter int64 `json:"twitter,omitempty"`
}

// BaseCity type from VK API Schema(base_city).
type BaseCity struct {

	// ID City ID
	ID int64 `json:"id"`
	// Title City title
	Title string `json:"title"`
}

// GroupsBanInfoReason type from VK API Schema(groups_ban_info_reason). Ban reason
type GroupsBanInfoReason int64

// GroupsGroupType type from VK API Schema(groups_group_type). Community type
type GroupsGroupType string

// GroupsGroupXtrInvitedBy type from VK API Schema(groups_group_xtr_invited_by).
type GroupsGroupXtrInvitedBy struct {

	// AdminLevel
	AdminLevel *GroupsGroupXtrInvitedByAdminLevel `json:"admin_level,omitempty"`
	// ID Community ID
	ID string `json:"id,omitempty"`
	// InvitedBy Inviter ID
	InvitedBy int64 `json:"invited_by,omitempty"`
	// IsAdmin Information whether current user is manager
	IsAdmin *BaseBoolInt `json:"is_admin,omitempty"`
	// IsClosed Information whether community is closed
	IsClosed *BaseBoolInt `json:"is_closed,omitempty"`
	// IsMember Information whether current user is member
	IsMember *BaseBoolInt `json:"is_member,omitempty"`
	// Name Community name
	Name string `json:"name,omitempty"`
	// Photo100 URL of square photo of the community with 100 pixels in width
	Photo100 string `json:"photo_100,omitempty"`
	// Photo200 URL of square photo of the community with 200 pixels in width
	Photo200 string `json:"photo_200,omitempty"`
	// Photo50 URL of square photo of the community with 50 pixels in width
	Photo50 string `json:"photo_50,omitempty"`
	// ScreenName Domain of the community page
	ScreenName string `json:"screen_name,omitempty"`
	// Type
	Type *GroupsGroupXtrInvitedByType `json:"type,omitempty"`
}

// MessagesHistoryAttachment type from VK API Schema(messages_history_attachment).
type MessagesHistoryAttachment struct {

	// Attachment
	Attachment *MessagesHistoryMessageAttachment `json:"attachment"`
	// MessageID Message ID
	MessageID int64 `json:"message_id,omitempty"`
}

// WallPostedPhoto type from VK API Schema(wall_posted_photo).
type WallPostedPhoto struct {

	// ID Photo ID
	ID int64 `json:"id,omitempty"`
	// OwnerID Photo owner&#39;s ID
	OwnerID int64 `json:"owner_id,omitempty"`
	// Photo130 URL of the preview image with 130 px in width
	Photo130 string `json:"photo_130,omitempty"`
	// Photo604 URL of the preview image with 604 px in width
	Photo604 string `json:"photo_604,omitempty"`
}

// WallWallpostToID type from VK API Schema(wall_wallpost_to_id).
type WallWallpostToID struct {

	// Attachments
	Attachments []*WallWallpostAttachment `json:"attachments,omitempty"`
	// Comments
	Comments *BaseCommentsInfo `json:"comments,omitempty"`
	// CopyOwnerID ID of the source post owner
	CopyOwnerID int64 `json:"copy_owner_id,omitempty"`
	// CopyPostID ID of the source post
	CopyPostID int64 `json:"copy_post_id,omitempty"`
	// Date Date of publishing in Unixtime
	Date int64 `json:"date,omitempty"`
	// FromID Post author ID
	FromID int64 `json:"from_id,omitempty"`
	// Geo
	Geo *BaseGeo `json:"geo,omitempty"`
	// ID Post ID
	ID int64 `json:"id,omitempty"`
	// Likes
	Likes *BaseLikesInfo `json:"likes,omitempty"`
	// PostID wall post ID (if comment)
	PostID int64 `json:"post_id,omitempty"`
	// PostSource
	PostSource *WallPostSource `json:"post_source,omitempty"`
	// PostType
	PostType *WallPostType `json:"post_type,omitempty"`
	// Reposts
	Reposts *BaseRepostsInfo `json:"reposts,omitempty"`
	// SignerID Post signer ID
	SignerID int64 `json:"signer_id,omitempty"`
	// Text Post text
	Text string `json:"text,omitempty"`
	// ToID Wall owner&#39;s ID
	ToID int64 `json:"to_id,omitempty"`
}

// DocsDoc type from VK API Schema(docs_doc).
type DocsDoc struct {

	// AccessKey Access key for the document
	AccessKey string `json:"access_key,omitempty"`
	// Date Date when file has been uploaded in Unixtime
	Date int64 `json:"date"`
	// Ext File extension
	Ext string `json:"ext"`
	// ID Document ID
	ID int64 `json:"id"`
	// OwnerID Document owner ID
	OwnerID int64 `json:"owner_id"`
	// Preview
	Preview *DocsDocPreview `json:"preview,omitempty"`
	// Size File size in bites
	Size int64 `json:"size"`
	// Title Document title
	Title string `json:"title"`
	// Type Document type
	Type int64 `json:"type"`
	// URL File URL
	URL string `json:"url,omitempty"`
}

// GroupsMemberRole type from VK API Schema(groups_member_role).
type GroupsMemberRole struct {

	// ID User ID
	ID int64 `json:"id,omitempty"`
	// Role
	Role *GroupsMemberRoleStatus `json:"role,omitempty"`
}

// SearchHintSection type from VK API Schema(search_hint_section). Section title
type SearchHintSection string

// AdsCampaignStatus type from VK API Schema(ads_campaign_status). Campaign status
type AdsCampaignStatus int64

// GroupsLongPollServer type from VK API Schema(groups_long_poll_server).
type GroupsLongPollServer struct {

	// Key Long Poll key
	Key string `json:"key"`
	// Server Long Poll server address
	Server string `json:"server"`
	// Ts Number of the last event
	Ts int64 `json:"ts"`
}

// LeadsEntry type from VK API Schema(leads_entry).
type LeadsEntry struct {

	// Aid Application ID
	Aid int64 `json:"aid,omitempty"`
	// Comment Comment text
	Comment string `json:"comment,omitempty"`
	// Date Date when the action has been started in Unixtime
	Date int64 `json:"date,omitempty"`
	// Sid Session string ID
	Sid string `json:"sid,omitempty"`
	// StartDate Start date in Unixtime (for status=2)
	StartDate int64 `json:"start_date,omitempty"`
	// Status Action type
	Status int64 `json:"status,omitempty"`
	// TestMode Information whether test mode is enabled
	TestMode *BaseBoolInt `json:"test_mode,omitempty"`
	// UID User ID
	UID int64 `json:"uid,omitempty"`
}

// UsersUserType type from VK API Schema(users_user_type). Object type
type UsersUserType string

// AccountOtherContact type from VK API Schema(account_other_contact).
type AccountOtherContact struct {

	// CommonCount Mutual friends count
	CommonCount int64 `json:"common_count,omitempty"`
	// Contact Contact
	Contact string `json:"contact,omitempty"`
}

// GiftsGift type from VK API Schema(gifts_gift).
type GiftsGift struct {

	// Date Date when gist has been sent in Unixtime
	Date int64 `json:"date,omitempty"`
	// FromID Gift sender ID
	FromID int64 `json:"from_id,omitempty"`
	// Gift
	Gift *GiftsLayout `json:"gift,omitempty"`
	// GiftHash Hash
	GiftHash string `json:"gift_hash,omitempty"`
	// ID Gift ID
	ID int64 `json:"id,omitempty"`
	// Message Comment text
	Message string `json:"message,omitempty"`
	// Privacy
	Privacy *GiftsGiftPrivacy `json:"privacy,omitempty"`
}

// GroupsGroupLink type from VK API Schema(groups_group_link).
type GroupsGroupLink struct {

	// Desc Link description
	Desc string `json:"desc,omitempty"`
	// EditTitle Information whether the title can be edited
	EditTitle *BaseBoolInt `json:"edit_title,omitempty"`
	// ID Link ID
	ID int64 `json:"id,omitempty"`
	// ImageProcessing Information whether the image on processing
	ImageProcessing *BaseBoolInt `json:"image_processing,omitempty"`
	// URL Link URL
	URL string `json:"url,omitempty"`
}

// NewsfeedItemAudioAudio type from VK API Schema(newsfeed_item_audio_audio).
type NewsfeedItemAudioAudio struct {

	// Count Audios number
	Count int64 `json:"count,omitempty"`
	// Items
	Items []*AudioAudioFull `json:"items,omitempty"`
}

// DatabaseCity type from VK API Schema(database_city).
type DatabaseCity struct {
}

// GroupsGroupBanInfo type from VK API Schema(groups_group_ban_info).
type GroupsGroupBanInfo struct {

	// Comment Ban comment
	Comment string `json:"comment,omitempty"`
	// EndDate End date of ban in Unixtime
	EndDate int64 `json:"end_date,omitempty"`
}

// MarketMarketAlbum type from VK API Schema(market_market_album).
type MarketMarketAlbum struct {

	// Count Items number
	Count int64 `json:"count"`
	// ID Market album ID
	ID int64 `json:"id"`
	// OwnerID Market album owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// Photo
	Photo *PhotosPhoto `json:"photo,omitempty"`
	// Title Market album title
	Title string `json:"title"`
	// UpdatedTime Date when album has been updated last time in Unixtime
	UpdatedTime int64 `json:"updated_time"`
}

// NotificationsNotificationParent type from VK API Schema(notifications_notification_parent).
type NotificationsNotificationParent struct {
}

// PhotosMessageUploadResponse type from VK API Schema(photos_message_upload_response).
type PhotosMessageUploadResponse struct {

	// Hash Uploading hash
	Hash string `json:"hash,omitempty"`
	// Photo Uploaded photo data
	Photo string `json:"photo,omitempty"`
	// Server Upload server number
	Server int64 `json:"server,omitempty"`
}

// PollsVoters type from VK API Schema(polls_voters).
type PollsVoters struct {

	// AnswerID Answer ID
	AnswerID int64 `json:"answer_id,omitempty"`
	// Users
	Users *PollsVotersUsers `json:"users,omitempty"`
}

// AdsPostStats type from VK API Schema(ads_post_stats).
type AdsPostStats struct {
}

// AudioLyrics type from VK API Schema(audio_lyrics).
type AudioLyrics struct {

	// LyricsID Lyrics ID
	LyricsID int64 `json:"lyrics_id"`
	// Text Lyrics text
	Text string `json:"text"`
}

// GroupsMemberStatusFull type from VK API Schema(groups_member_status_full).
type GroupsMemberStatusFull struct {

	// Invitation Information whether user has been invited to the group
	Invitation *BaseBoolInt `json:"invitation,omitempty"`
	// Member Information whether user is a member of the group
	Member *BaseBoolInt `json:"member"`
	// Request Information whether user has send request to the group
	Request *BaseBoolInt `json:"request,omitempty"`
	// UserID User ID
	UserID int64 `json:"user_id"`
}

// MessagesPinnedMessage type from VK API Schema(messages_pinned_message).
type MessagesPinnedMessage struct {

	// Attachments
	Attachments []*MessagesMessageAttachment `json:"attachments,omitempty"`
	// ConversationMessageID Unique auto-incremented number for all messages with this peer
	ConversationMessageID int64 `json:"conversation_message_id,omitempty"`
	// Date Date when the message has been sent in Unixtime
	Date int64 `json:"date"`
	// FromID Message author&#39;s ID
	FromID int64 `json:"from_id"`
	// FwdMessages Forwarded messages
	FwdMessages []*MessagesMessage `json:"fwd_messages,omitempty"`
	// Geo
	Geo *BaseGeo `json:"geo,omitempty"`
	// ID Message ID
	ID int64 `json:"id"`
	// PeerID Peer ID
	PeerID int64 `json:"peer_id"`
	// Text Message text
	Text string `json:"text"`
}

// PhotosOwnerUploadResponse type from VK API Schema(photos_owner_upload_response).
type PhotosOwnerUploadResponse struct {

	// Hash Uploading hash
	Hash string `json:"hash,omitempty"`
	// Photo Uploaded photo data
	Photo string `json:"photo,omitempty"`
	// Server Upload server number
	Server int64 `json:"server,omitempty"`
}

// GroupsGroupAdminLevel type from VK API Schema(groups_group_admin_level). Level of current user&#39;s credentials as manager
type GroupsGroupAdminLevel int64

// GroupsUserXtrRole type from VK API Schema(groups_user_xtr_role).
type GroupsUserXtrRole struct {
}

// PlacesPlaceFull type from VK API Schema(places_place_full).
type PlacesPlaceFull struct {

	// Address Place address
	Address string `json:"address,omitempty"`
	// Checkins Checkins number
	Checkins int64 `json:"checkins,omitempty"`
	// City City ID
	City int64 `json:"city,omitempty"`
	// Country Country ID
	Country int64 `json:"country,omitempty"`
	// Created Date of the place creation in Unixtime
	Created int64 `json:"created,omitempty"`
	// Distance Distance to the place
	Distance int64 `json:"distance,omitempty"`
	// GroupID Community ID
	GroupID int64 `json:"group_id,omitempty"`
	// GroupPhoto URL of the community&#39;s photo
	GroupPhoto string `json:"group_photo,omitempty"`
	// Icon URL of the place&#39;s icon
	Icon string `json:"icon,omitempty"`
	// ID Place ID
	ID int64 `json:"id,omitempty"`
	// Latitude Place latitude
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude Place longitude
	Longitude float64 `json:"longitude,omitempty"`
	// Title Place title
	Title string `json:"title,omitempty"`
	// Type Place type
	Type string `json:"type,omitempty"`
}

// WallWallpostAttachmentType type from VK API Schema(wall_wallpost_attachment_type). Attachment type
type WallWallpostAttachmentType string

// AdsStatsCities type from VK API Schema(ads_stats_cities).
type AdsStatsCities struct {

	// ClicksRate Clicks rate
	ClicksRate float64 `json:"clicks_rate,omitempty"`
	// ImpressionsRate Impressions rate
	ImpressionsRate float64 `json:"impressions_rate,omitempty"`
	// Name City name
	Name string `json:"name,omitempty"`
	// Value City ID
	Value int64 `json:"value,omitempty"`
}

// MessagesMessageActionStatus type from VK API Schema(messages_message_action_status). Action status
type MessagesMessageActionStatus string

// StoriesStory type from VK API Schema(stories_story).
type StoriesStory struct {

	// AccessKey Access key for private object.
	AccessKey string `json:"access_key,omitempty"`
	// CanComment Information whether current user can comment the story (0 - no, 1 - yes).
	CanComment *BaseBoolInt `json:"can_comment,omitempty"`
	// CanReply Information whether current user can reply to the story (0 - no, 1 - yes).
	CanReply *BaseBoolInt `json:"can_reply,omitempty"`
	// CanSee Information whether current user can see the story (0 - no, 1 - yes).
	CanSee *BaseBoolInt `json:"can_see,omitempty"`
	// CanShare Information whether current user can share the story (0 - no, 1 - yes).
	CanShare *BaseBoolInt `json:"can_share,omitempty"`
	// Date Date when story has been added in Unixtime.
	Date int64 `json:"date,omitempty"`
	// ID Story ID.
	ID int64 `json:"id"`
	// IsDeleted Information whether the story is deleted (false - no, true - yes).
	IsDeleted bool `json:"is_deleted,omitempty"`
	// IsExpired Information whether the story is expired (false - no, true - yes).
	IsExpired bool `json:"is_expired,omitempty"`
	// Link
	Link *StoriesStoryLink `json:"link,omitempty"`
	// OwnerID Story owner&#39;s ID.
	OwnerID int64 `json:"owner_id"`
	// ParentStory
	ParentStory *StoriesStory `json:"parent_story,omitempty"`
	// ParentStoryAccessKey Access key for private object.
	ParentStoryAccessKey string `json:"parent_story_access_key,omitempty"`
	// ParentStoryID Parent story ID.
	ParentStoryID int64 `json:"parent_story_id,omitempty"`
	// ParentStoryOwnerID Parent story owner&#39;s ID.
	ParentStoryOwnerID int64 `json:"parent_story_owner_id,omitempty"`
	// Photo
	Photo *PhotosPhoto `json:"photo,omitempty"`
	// Replies Replies to current story.
	Replies []*StoriesReplies `json:"replies,omitempty"`
	// Seen Information whether current user has seen the story or not (0 - no, 1 - yes).
	Seen *BaseBoolInt `json:"seen,omitempty"`
	// Type
	Type *StoriesStoryType `json:"type,omitempty"`
	// Video
	Video *StoriesStoryVideo `json:"video,omitempty"`
	// Views Views number.
	Views int64 `json:"views,omitempty"`
}

// WallGraffiti type from VK API Schema(wall_graffiti).
type WallGraffiti struct {

	// ID Graffiti ID
	ID int64 `json:"id,omitempty"`
	// OwnerID Graffiti owner&#39;s ID
	OwnerID int64 `json:"owner_id,omitempty"`
	// Photo200 URL of the preview image with 200 px in width
	Photo200 string `json:"photo_200,omitempty"`
	// Photo586 URL of the preview image with 586 px in width
	Photo586 string `json:"photo_586,omitempty"`
}

// FriendsMutualFriend type from VK API Schema(friends_mutual_friend).
type FriendsMutualFriend struct {

	// CommonCount Total mutual friends number
	CommonCount int64 `json:"common_count,omitempty"`
	// CommonFriends
	CommonFriends []int64 `json:"common_friends,omitempty"`
	// ID User ID
	ID int64 `json:"id,omitempty"`
}

// GroupsLongPollSettings type from VK API Schema(groups_long_poll_settings).
type GroupsLongPollSettings struct {

	// APIVersion API version used for the events
	APIVersion string `json:"api_version,omitempty"`
	// Events
	Events *GroupsLongPollEvents `json:"events"`
	// IsEnabled Shows whether Long Poll is enabled
	IsEnabled bool `json:"is_enabled"`
}

// PhotosPhotoUpload type from VK API Schema(photos_photo_upload).
type PhotosPhotoUpload struct {

	// AlbumID Album ID
	AlbumID int64 `json:"album_id"`
	// UploadURL URL to upload photo
	UploadURL string `json:"upload_url"`
	// UserID User ID
	UserID int64 `json:"user_id"`
}

// VideoVideoAlbumFull type from VK API Schema(video_video_album_full).
type VideoVideoAlbumFull struct {

	// Count Total number of videos in album
	Count int64 `json:"count"`
	// ID Album ID
	ID int64 `json:"id"`
	// IsSystem Information whether album is system
	IsSystem int64 `json:"is_system,omitempty"`
	// OwnerID Album owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// Photo160 URL of the preview image with 160px in width
	Photo160 string `json:"photo_160,omitempty"`
	// Photo320 URL of the preview image with 320px in width
	Photo320 string `json:"photo_320,omitempty"`
	// Title Album title
	Title string `json:"title"`
	// UpdatedTime Date when the album has been updated last time in Unixtime
	UpdatedTime int64 `json:"updated_time"`
}

// AdsAd type from VK API Schema(ads_ad).
type AdsAd struct {

	// AdFormat Ad format
	AdFormat int64 `json:"ad_format"`
	// AdPlatform Ad platform
	AdPlatform string `json:"ad_platform,omitempty"`
	// AllLimit Total limit
	AllLimit int64 `json:"all_limit"`
	// Approved
	Approved *AdsAdApproved `json:"approved"`
	// CampaignID Campaign ID
	CampaignID int64 `json:"campaign_id"`
	// Category1ID Category ID
	Category1ID int64 `json:"category1_id,omitempty"`
	// Category2ID Additional category ID
	Category2ID int64 `json:"category2_id,omitempty"`
	// CostType
	CostType *AdsAdCostType `json:"cost_type"`
	// Cpc Cost of a click, kopecks
	Cpc int64 `json:"cpc,omitempty"`
	// Cpm Cost of 1000 impressions, kopecks
	Cpm int64 `json:"cpm,omitempty"`
	// DisclaimerMedical Information whether disclaimer is enabled
	DisclaimerMedical *BaseBoolInt `json:"disclaimer_medical,omitempty"`
	// DisclaimerSpecialist Information whether disclaimer is enabled
	DisclaimerSpecialist *BaseBoolInt `json:"disclaimer_specialist,omitempty"`
	// DisclaimerSupplements Information whether disclaimer is enabled
	DisclaimerSupplements *BaseBoolInt `json:"disclaimer_supplements,omitempty"`
	// ID Ad ID
	ID int64 `json:"id"`
	// ImpressionsLimit Impressions limit
	ImpressionsLimit int64 `json:"impressions_limit,omitempty"`
	// ImpressionsLimited Information whether impressions are limited
	ImpressionsLimited *BaseBoolInt `json:"impressions_limited,omitempty"`
	// Name Ad title
	Name string `json:"name"`
	// Status
	Status *AdsAdStatus `json:"status"`
	// Video Information whether the ad is a video
	Video *BaseBoolInt `json:"video,omitempty"`
}

// AdsAdLayout type from VK API Schema(ads_ad_layout).
type AdsAdLayout struct {

	// AdFormat Ad format
	AdFormat int64 `json:"ad_format"`
	// CampaignID Campaign ID
	CampaignID int64 `json:"campaign_id"`
	// CostType
	CostType *AdsAdLayoutCostType `json:"cost_type"`
	// Description Ad description
	Description string `json:"description"`
	// ID Ad ID
	ID int64 `json:"id"`
	// ImageSrc Image URL
	ImageSrc int64 `json:"image_src"`
	// ImageSrc2x URL of the preview image in double size
	ImageSrc2x int64 `json:"image_src_2x,omitempty"`
	// LinkDomain Domain of advertised object
	LinkDomain string `json:"link_domain,omitempty"`
	// LinkURL URL of advertised object
	LinkURL string `json:"link_url"`
	// PreviewLink link to preview an ad as it is shown on the website
	PreviewLink string `json:"preview_link,omitempty"`
	// Title Ad title
	Title string `json:"title"`
	// Video Information whether the ad is a video
	Video *BaseBoolInt `json:"video,omitempty"`
}

// AdsCampaignType type from VK API Schema(ads_campaign_type). Campaign type
type AdsCampaignType string

// BasePropertyExists type from VK API Schema(base_property_exists).
type BasePropertyExists int64

// PhotosPhotoAlbumFull type from VK API Schema(photos_photo_album_full).
type PhotosPhotoAlbumFull struct {

	// CanUpload Information whether current user can upload photo to the album
	CanUpload *BaseBoolInt `json:"can_upload,omitempty"`
	// CommentsDisabled Information whether album comments are disabled
	CommentsDisabled *BaseBoolInt `json:"comments_disabled,omitempty"`
	// Created Date when the album has been created in Unixtime
	Created int64 `json:"created"`
	// Description Photo album description
	Description string `json:"description,omitempty"`
	// ID Photo album ID
	ID int64 `json:"id"`
	// OwnerID Album owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// PrivacyComment
	PrivacyComment []string `json:"privacy_comment,omitempty"`
	// PrivacyView
	PrivacyView []string `json:"privacy_view,omitempty"`
	// Size Photos number
	Size int64 `json:"size"`
	// Sizes
	Sizes []*PhotosPhotoSizes `json:"sizes,omitempty"`
	// ThumbID Thumb photo ID
	ThumbID int64 `json:"thumb_id,omitempty"`
	// ThumbIsLast Information whether the album thumb is last photo
	ThumbIsLast *BaseBoolInt `json:"thumb_is_last,omitempty"`
	// ThumbSrc URL of the thumb image
	ThumbSrc string `json:"thumb_src,omitempty"`
	// Title Photo album title
	Title string `json:"title"`
	// Updated Date when the album has been updated last time in Unixtime
	Updated int64 `json:"updated"`
	// UploadByAdminsOnly Information whether only community administrators can upload photos
	UploadByAdminsOnly *BaseBoolInt `json:"upload_by_admins_only,omitempty"`
}

// SecureTransaction type from VK API Schema(secure_transaction).
type SecureTransaction struct {

	// Date Transaction date in Unixtime
	Date int64 `json:"date,omitempty"`
	// ID Transaction ID
	ID int64 `json:"id,omitempty"`
	// UIDFrom From ID
	UIDFrom int64 `json:"uid_from,omitempty"`
	// UIDTo To ID
	UIDTo int64 `json:"uid_to,omitempty"`
	// Votes Votes number
	Votes int64 `json:"votes,omitempty"`
}

// StoriesStoryVideo type from VK API Schema(stories_story_video).
type StoriesStoryVideo struct {
}

// VideoVideo type from VK API Schema(video_video).
type VideoVideo struct {

	// AccessKey Video access key
	AccessKey string `json:"access_key,omitempty"`
	// AddingDate Date when the video has been added in Unixtime
	AddingDate int64 `json:"adding_date,omitempty"`
	// CanAdd Information whether current user can add the video
	CanAdd *BaseBoolInt `json:"can_add,omitempty"`
	// CanEdit Information whether current user can edit the video
	CanEdit *BaseBoolInt `json:"can_edit,omitempty"`
	// Comments Number of comments
	Comments int64 `json:"comments,omitempty"`
	// Date Date when video has been uploaded in Unixtime
	Date int64 `json:"date,omitempty"`
	// Description Video description
	Description string `json:"description,omitempty"`
	// Duration Video duration in seconds
	Duration int64 `json:"duration,omitempty"`
	// Files
	Files *VideoVideoFiles `json:"files,omitempty"`
	// ID Video ID
	ID int64 `json:"id,omitempty"`
	// Live Returns if the video is live translation
	Live *BasePropertyExists `json:"live,omitempty"`
	// OwnerID Video owner ID
	OwnerID int64 `json:"owner_id,omitempty"`
	// Photo130 URL of the preview image with 130 px in width
	Photo130 string `json:"photo_130,omitempty"`
	// Photo320 URL of the preview image with 320 px in width
	Photo320 string `json:"photo_320,omitempty"`
	// Photo800 URL of the preview image with 800 px in width
	Photo800 string `json:"photo_800,omitempty"`
	// Player URL of the page with a player that can be used to play the video in the browser.
	Player string `json:"player,omitempty"`
	// Processing Returns if the video is processing
	Processing *BasePropertyExists `json:"processing,omitempty"`
	// Title Video title
	Title string `json:"title,omitempty"`
	// Views Number of views
	Views int64 `json:"views,omitempty"`
}

// AdsLinkStatus type from VK API Schema(ads_link_status).
type AdsLinkStatus struct {

	// Description Reject reason
	Description string `json:"description"`
	// RedirectURL URL
	RedirectURL string `json:"redirect_url"`
	// Status Link status
	Status string `json:"status"`
}

// BaseCountry type from VK API Schema(base_country).
type BaseCountry struct {

	// ID Country ID
	ID int64 `json:"id"`
	// Title Country title
	Title string `json:"title"`
}

// BaseError type from VK API Schema(base_error).
type BaseError struct {

	// ErrorCode Error code
	ErrorCode int64 `json:"error_code,omitempty"`
	// ErrorMsg Error message
	ErrorMsg string `json:"error_msg,omitempty"`
	// RequestParams
	RequestParams []*BaseRequestParam `json:"request_params,omitempty"`
}

// GroupsGroupIsClosed type from VK API Schema(groups_group_is_closed). Information whether community is closed
type GroupsGroupIsClosed int64

// GroupsMarketInfo type from VK API Schema(groups_market_info).
type GroupsMarketInfo struct {

	// ContactID Contact person ID
	ContactID int64 `json:"contact_id,omitempty"`
	// Currency
	Currency *MarketCurrency `json:"currency,omitempty"`
	// CurrencyText Currency name
	CurrencyText string `json:"currency_text,omitempty"`
	// Enabled Information whether the market is enabled
	Enabled *BaseBoolInt `json:"enabled,omitempty"`
	// MainAlbumID Main market album ID
	MainAlbumID int64 `json:"main_album_id,omitempty"`
	// PriceMax Maximum price
	PriceMax int64 `json:"price_max,omitempty"`
	// PriceMin Minimum price
	PriceMin int64 `json:"price_min,omitempty"`
}

// UtilsDomainResolved type from VK API Schema(utils_domain_resolved).
type UtilsDomainResolved struct {

	// ObjectID Object ID
	ObjectID int64 `json:"object_id,omitempty"`
	// Type
	Type *UtilsDomainResolvedType `json:"type,omitempty"`
}

// VideoVideoTag type from VK API Schema(video_video_tag).
type VideoVideoTag struct {

	// Date Date when tag has been added in Unixtime
	Date int64 `json:"date"`
	// ID Tag ID
	ID int64 `json:"id"`
	// PlacerID ID of the tag creator
	PlacerID int64 `json:"placer_id"`
	// TaggedName Tag description
	TaggedName string `json:"tagged_name"`
	// UserID Tagged user ID
	UserID int64 `json:"user_id"`
	// Viewed Information whether tha tag is reviewed
	Viewed *BaseBoolInt `json:"viewed"`
}

// AdsTargStats type from VK API Schema(ads_targ_stats).
type AdsTargStats struct {

	// AudienceCount Audience
	AudienceCount int64 `json:"audience_count"`
	// RecommendedCpc Recommended CPC value
	RecommendedCpc float64 `json:"recommended_cpc,omitempty"`
	// RecommendedCpm Recommended CPM value
	RecommendedCpm float64 `json:"recommended_cpm,omitempty"`
}

// GroupsSubjectItem type from VK API Schema(groups_subject_item).
type GroupsSubjectItem struct {

	// ID Subject ID
	ID int64 `json:"id,omitempty"`
	// Name Subject title
	Name string `json:"name,omitempty"`
}

// StatsReach type from VK API Schema(stats_reach). Reach stats
type StatsReach struct {

	// Age
	Age []*StatsSexAge `json:"age,omitempty"`
	// Cities
	Cities []*StatsCity `json:"cities,omitempty"`
	// Countries
	Countries []*StatsCountry `json:"countries,omitempty"`
	// MobileReach Reach count from mobile devices
	MobileReach int64 `json:"mobile_reach,omitempty"`
	// Reach Reach count
	Reach int64 `json:"reach,omitempty"`
	// ReachSubscribers Subscribers reach count
	ReachSubscribers int64 `json:"reach_subscribers,omitempty"`
	// Sex
	Sex []*StatsSexAge `json:"sex,omitempty"`
	// SexAge
	SexAge []*StatsSexAge `json:"sex_age,omitempty"`
}

// StatsWallpostStat type from VK API Schema(stats_wallpost_stat).
type StatsWallpostStat struct {

	// Hide Hidings number
	Hide int64 `json:"hide,omitempty"`
	// JoinGroup People have joined the group
	JoinGroup int64 `json:"join_group,omitempty"`
	// Links Link clickthrough
	Links int64 `json:"links,omitempty"`
	// ReachSubscribers Subscribers reach
	ReachSubscribers int64 `json:"reach_subscribers,omitempty"`
	// ReachTotal Total reach
	ReachTotal int64 `json:"reach_total,omitempty"`
	// Report Reports number
	Report int64 `json:"report,omitempty"`
	// ToGroup Clickthrough to community
	ToGroup int64 `json:"to_group,omitempty"`
	// Unsubscribe Unsubscribed members
	Unsubscribe int64 `json:"unsubscribe,omitempty"`
}

// StoriesStoryLink type from VK API Schema(stories_story_link).
type StoriesStoryLink struct {

	// Text Link text
	Text string `json:"text"`
	// URL Link URL
	URL string `json:"url"`
}

// MessagesChat type from VK API Schema(messages_chat).
type MessagesChat struct {

	// AdminID Chat creator ID
	AdminID int64 `json:"admin_id"`
	// ID Chat ID
	ID int64 `json:"id"`
	// Kicked Shows that user has been kicked from the chat
	Kicked *BaseBoolInt `json:"kicked,omitempty"`
	// Left Shows that user has been left the chat
	Left *BaseBoolInt `json:"left,omitempty"`
	// Photo100 URL of the preview image with 100 px in width
	Photo100 string `json:"photo_100,omitempty"`
	// Photo200 URL of the preview image with 200 px in width
	Photo200 string `json:"photo_200,omitempty"`
	// Photo50 URL of the preview image with 50 px in width
	Photo50 string `json:"photo_50,omitempty"`
	// PushSettings
	PushSettings *MessagesChatPushSettings `json:"push_settings,omitempty"`
	// Title Chat title
	Title string `json:"title,omitempty"`
	// Type Chat type
	Type string `json:"type"`
	// Users
	Users []int64 `json:"users"`
}

// NewsfeedList type from VK API Schema(newsfeed_list).
type NewsfeedList struct {

	// ID List ID
	ID int64 `json:"id"`
	// Title List title
	Title string `json:"title"`
}

// PhotosImageType type from VK API Schema(photos_image_type). Photo&#39;s type.
type PhotosImageType string

// AdsAccount type from VK API Schema(ads_account).
type AdsAccount struct {

	// AccessRole
	AccessRole *AdsAccessRole `json:"access_role"`
	// AccountID Account ID
	AccountID int64 `json:"account_id"`
	// AccountStatus Information whether account is active
	AccountStatus *BaseBoolInt `json:"account_status"`
	// AccountType
	AccountType *AdsAccountType `json:"account_type"`
}

// AdsObjectType type from VK API Schema(ads_object_type). Object type
type AdsObjectType string

// BaseLikes type from VK API Schema(base_likes).
type BaseLikes struct {

	// Count Likes number
	Count int64 `json:"count,omitempty"`
	// UserLikes Information whether current user likes the photo
	UserLikes *BaseBoolInt `json:"user_likes,omitempty"`
}

// BaseLinkApplicationStore type from VK API Schema(base_link_application_store).
type BaseLinkApplicationStore struct {

	// ID Store Id
	ID float64 `json:"id,omitempty"`
	// Name Store name
	Name string `json:"name,omitempty"`
}

// MarketMarketCategory type from VK API Schema(market_market_category).
type MarketMarketCategory struct {

	// ID Category ID
	ID int64 `json:"id"`
	// Name Category name
	Name string `json:"name"`
	// Section
	Section *MarketSection `json:"section"`
}

// UsersUserXtrType type from VK API Schema(users_user_xtr_type).
type UsersUserXtrType struct {
}

// UsersMilitary type from VK API Schema(users_military).
type UsersMilitary struct {

	// CountryID Country ID
	CountryID int64 `json:"country_id,omitempty"`
	// From From year
	From int64 `json:"from,omitempty"`
	// Unit Unit name
	Unit string `json:"unit,omitempty"`
	// UnitID Unit ID
	UnitID int64 `json:"unit_id,omitempty"`
	// Until Till year
	Until int64 `json:"until,omitempty"`
}

// UsersCropPhotoRect type from VK API Schema(users_crop_photo_rect).
type UsersCropPhotoRect struct {

	// X Coordinate X of the left upper corner
	X float64 `json:"x,omitempty"`
	// X2 Coordinate X of the right lower corner
	X2 float64 `json:"x2,omitempty"`
	// Y Coordinate Y of the left upper corner
	Y float64 `json:"y,omitempty"`
	// Y2 Coordinate Y of the right lower corner
	Y2 float64 `json:"y2,omitempty"`
}

// WallCommentAttachmentType type from VK API Schema(wall_comment_attachment_type). Attachment type
type WallCommentAttachmentType string

// AdsParagraphs type from VK API Schema(ads_paragraphs).
type AdsParagraphs struct {

	// Paragraph Rules paragraph
	Paragraph string `json:"paragraph,omitempty"`
}

// NotesNote type from VK API Schema(notes_note).
type NotesNote struct {

	// CanComment Information whether current user can comment the note
	CanComment *BaseBoolInt `json:"can_comment,omitempty"`
	// Comments Comments number
	Comments int64 `json:"comments"`
	// Date Date when the note has been created in Unixtime
	Date int64 `json:"date"`
	// ID Note ID
	ID int64 `json:"id"`
	// OwnerID Note owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// Text Note text
	Text string `json:"text,omitempty"`
	// TextWiki Note text in wiki format
	TextWiki string `json:"text_wiki,omitempty"`
	// Title Note title
	Title string `json:"title"`
	// ViewURL URL of the page with note preview
	ViewURL string `json:"view_url"`
}

// PhotosPhotoSizes type from VK API Schema(photos_photo_sizes).
type PhotosPhotoSizes struct {

	// Height Height in px
	Height int64 `json:"height"`
	// Src URL of the image
	Src string `json:"src"`
	// Type
	Type *PhotosPhotoSizesType `json:"type"`
	// Width Width in px
	Width int64 `json:"width"`
}

// PollsVotersUsers type from VK API Schema(polls_voters_users).
type PollsVotersUsers struct {

	// Count Votes number
	Count int64 `json:"count,omitempty"`
	// Items
	Items []int64 `json:"items,omitempty"`
}

// UsersUser type from VK API Schema(users_user).
type UsersUser struct {
}

// AdsUsers type from VK API Schema(ads_users).
type AdsUsers struct {

	// Accesses
	Accesses []*AdsAccesses `json:"accesses"`
	// UserID User ID
	UserID int64 `json:"user_id"`
}

// NewsfeedItemWallpostType type from VK API Schema(newsfeed_item_wallpost_type). Post type
type NewsfeedItemWallpostType string

// SecureTokenChecked type from VK API Schema(secure_token_checked).
type SecureTokenChecked struct {

	// Date Date when access_token has been generated in Unixtime
	Date int64 `json:"date,omitempty"`
	// Expire Date when access_token will expire in Unixtime
	Expire int64 `json:"expire,omitempty"`
	// Success Returns if successfully processed
	Success *BaseOkResponse `json:"success,omitempty"`
	// UserID User ID
	UserID int64 `json:"user_id,omitempty"`
}

// UsersUsersArray type from VK API Schema(users_users_array).
type UsersUsersArray struct {

	// Count Users number
	Count int64 `json:"count"`
	// Items
	Items []int64 `json:"items"`
}

// PhotosMarketAlbumUploadResponse type from VK API Schema(photos_market_album_upload_response).
type PhotosMarketAlbumUploadResponse struct {

	// Gid Community ID
	Gid int64 `json:"gid,omitempty"`
	// Hash Uploading hash
	Hash string `json:"hash,omitempty"`
	// Photo Uploaded photo data
	Photo string `json:"photo,omitempty"`
	// Server Upload server number
	Server int64 `json:"server,omitempty"`
}

// AdsAdLayoutCostType type from VK API Schema(ads_ad_layout_cost_type). Cost type
type AdsAdLayoutCostType int64

// DocsDocPreviewPhoto type from VK API Schema(docs_doc_preview_photo).
type DocsDocPreviewPhoto struct {

	// Sizes
	Sizes []*PhotosPhotoSizes `json:"sizes,omitempty"`
}

// FriendsUserXtrPhone type from VK API Schema(friends_user_xtr_phone).
type FriendsUserXtrPhone struct {
}

// GroupsGroupCategoryFull type from VK API Schema(groups_group_category_full).
type GroupsGroupCategoryFull struct {

	// ID Category ID
	ID int64 `json:"id"`
	// Name Category name
	Name string `json:"name"`
	// PageCount Pages number
	PageCount int64 `json:"page_count"`
	// PagePreviews
	PagePreviews []*GroupsGroup `json:"page_previews"`
	// Subcategories
	Subcategories []*GroupsGroupCategory `json:"subcategories,omitempty"`
}

// MarketCurrency type from VK API Schema(market_currency).
type MarketCurrency struct {

	// ID Currency ID
	ID int64 `json:"id"`
	// Name Currency sign
	Name string `json:"name"`
}

// UtilsStatsExtended type from VK API Schema(utils_stats_extended).
type UtilsStatsExtended struct {

	// Cities
	Cities []*UtilsStatsCity `json:"cities,omitempty"`
	// Countries
	Countries []*UtilsStatsCountry `json:"countries,omitempty"`
	// SexAge
	SexAge []*UtilsStatsSexAge `json:"sex_age,omitempty"`
	// Timestamp Start time
	Timestamp int64 `json:"timestamp,omitempty"`
	// Views Total views number
	Views int64 `json:"views,omitempty"`
}

// BaseObjectCount type from VK API Schema(base_object_count).
type BaseObjectCount struct {

	// Count Items count
	Count int64 `json:"count,omitempty"`
}

// GroupsContactsItem type from VK API Schema(groups_contacts_item).
type GroupsContactsItem struct {

	// Desc Contact description
	Desc string `json:"desc,omitempty"`
	// Email Contact email
	Email string `json:"email,omitempty"`
	// Phone Contact phone
	Phone string `json:"phone,omitempty"`
	// UserID User ID
	UserID int64 `json:"user_id,omitempty"`
}

// GroupsMemberStatus type from VK API Schema(groups_member_status).
type GroupsMemberStatus struct {

	// Member Information whether user is a member of the group
	Member *BaseBoolInt `json:"member"`
	// UserID User ID
	UserID int64 `json:"user_id"`
}

// UtilsShortLink type from VK API Schema(utils_short_link).
type UtilsShortLink struct {

	// AccessKey Access key for private stats
	AccessKey string `json:"access_key,omitempty"`
	// Key Link key (characters after vk.cc/)
	Key string `json:"key,omitempty"`
	// ShortURL Short link URL
	ShortURL string `json:"short_url,omitempty"`
	// URL Full URL
	URL string `json:"url,omitempty"`
}

// UtilsStats type from VK API Schema(utils_stats).
type UtilsStats struct {

	// Timestamp Start time
	Timestamp int64 `json:"timestamp,omitempty"`
	// Views Total views number
	Views int64 `json:"views,omitempty"`
}

// AccountPushConversations type from VK API Schema(account_push_conversations).
type AccountPushConversations struct {

	// Count Items count
	Count int64 `json:"count,omitempty"`
	// Items
	Items []*AccountPushConversationsItem `json:"items,omitempty"`
}

// StatusStatus type from VK API Schema(status_status).
type StatusStatus struct {

	// Audio
	Audio *AudioAudioFull `json:"audio,omitempty"`
	// Text Status text
	Text string `json:"text,omitempty"`
}

// UtilsLinkStats type from VK API Schema(utils_link_stats).
type UtilsLinkStats struct {

	// Key Link key (characters after vk.cc/)
	Key string `json:"key,omitempty"`
	// Stats
	Stats []*UtilsStats `json:"stats,omitempty"`
}

// VideoVideoTagInfo type from VK API Schema(video_video_tag_info).
type VideoVideoTagInfo struct {

	// AccessKey Video access key
	AccessKey string `json:"access_key,omitempty"`
	// AddingDate Date when the video has been added in Unixtime
	AddingDate int64 `json:"adding_date,omitempty"`
	// CanAdd Information whether current user can add the video
	CanAdd *BaseBoolInt `json:"can_add,omitempty"`
	// CanEdit Information whether current user can edit the video
	CanEdit *BaseBoolInt `json:"can_edit,omitempty"`
	// Comments Number of comments
	Comments int64 `json:"comments,omitempty"`
	// Date Date when video has been uploaded in Unixtime
	Date int64 `json:"date,omitempty"`
	// Description Video description
	Description string `json:"description,omitempty"`
	// Duration Video duration in seconds
	Duration int64 `json:"duration,omitempty"`
	// Files
	Files *VideoVideoFiles `json:"files,omitempty"`
	// ID Video ID
	ID int64 `json:"id,omitempty"`
	// Live Returns if the video is live translation
	Live *BasePropertyExists `json:"live,omitempty"`
	// OwnerID Video owner ID
	OwnerID int64 `json:"owner_id,omitempty"`
	// Photo130 URL of the preview image with 130 px in width
	Photo130 string `json:"photo_130,omitempty"`
	// Photo320 URL of the preview image with 320 px in width
	Photo320 string `json:"photo_320,omitempty"`
	// Photo800 URL of the preview image with 800 px in width
	Photo800 string `json:"photo_800,omitempty"`
	// PlacerID ID of the tag creator
	PlacerID int64 `json:"placer_id,omitempty"`
	// Player URL of the page with a player that can be used to play the video in the browser.
	Player string `json:"player,omitempty"`
	// Processing Returns if the video is processing
	Processing *BasePropertyExists `json:"processing,omitempty"`
	// TagCreated Date when tag has been added in Unixtime
	TagCreated int64 `json:"tag_created,omitempty"`
	// TagID Tag ID
	TagID int64 `json:"tag_id,omitempty"`
	// Title Video title
	Title string `json:"title,omitempty"`
	// Views Number of views
	Views int64 `json:"views,omitempty"`
}

// WidgetsCommentReplies type from VK API Schema(widgets_comment_replies).
type WidgetsCommentReplies struct {

	// CanPost Information whether current user can comment the post
	CanPost *BaseBoolInt `json:"can_post,omitempty"`
	// Count Comments number
	Count int64 `json:"count,omitempty"`
	// Replies
	Replies []*WidgetsCommentRepliesItem `json:"replies,omitempty"`
}

// NewsfeedItemNote type from VK API Schema(newsfeed_item_note).
type NewsfeedItemNote struct {

	// Notes
	Notes *NewsfeedItemNoteNotes `json:"notes,omitempty"`
}

// UtilsLinkCheckedStatus type from VK API Schema(utils_link_checked_status). Link status
type UtilsLinkCheckedStatus string

// BaseRequestParam type from VK API Schema(base_request_param).
type BaseRequestParam struct {

	// Key Parameter name
	Key string `json:"key,omitempty"`
	// Value Parameter value
	Value string `json:"value,omitempty"`
}

// BaseObject type from VK API Schema(base_object).
type BaseObject struct {

	// ID Object ID
	ID int64 `json:"id"`
	// Title Object title
	Title string `json:"title"`
}

// DatabaseRegion type from VK API Schema(database_region).
type DatabaseRegion struct {

	// ID Region ID
	ID int64 `json:"id,omitempty"`
	// Title Region title
	Title string `json:"title,omitempty"`
}

// GroupsGroupFull type from VK API Schema(groups_group_full).
type GroupsGroupFull struct {
}

// LeadsLeadDays type from VK API Schema(leads_lead_days).
type LeadsLeadDays struct {

	// Completed Completed offers number
	Completed int64 `json:"completed,omitempty"`
	// Impressions Impressions number
	Impressions int64 `json:"impressions,omitempty"`
	// Spent Amount of spent votes
	Spent int64 `json:"spent,omitempty"`
	// Started Started offers number
	Started int64 `json:"started,omitempty"`
}

// UtilsStatsCountry type from VK API Schema(utils_stats_country).
type UtilsStatsCountry struct {

	// CountryID Country ID
	CountryID int64 `json:"country_id,omitempty"`
	// Views Views number
	Views int64 `json:"views,omitempty"`
}

// AccountPushParamsMode type from VK API Schema(account_push_params_mode). Settings parameters
type AccountPushParamsMode string

// AdsCriteria type from VK API Schema(ads_criteria).
type AdsCriteria struct {

	// AgeFrom Age from
	AgeFrom int64 `json:"age_from,omitempty"`
	// AgeTo Age to
	AgeTo int64 `json:"age_to,omitempty"`
	// Apps Apps IDs
	Apps string `json:"apps,omitempty"`
	// AppsNot Apps IDs to except
	AppsNot string `json:"apps_not,omitempty"`
	// Birthday Days to birthday
	Birthday int64 `json:"birthday,omitempty"`
	// Cities Cities IDs
	Cities string `json:"cities,omitempty"`
	// CitiesNot Cities IDs to except
	CitiesNot string `json:"cities_not,omitempty"`
	// Country Country ID
	Country int64 `json:"country,omitempty"`
	// Districts Districts IDs
	Districts string `json:"districts,omitempty"`
	// Groups Communities IDs
	Groups string `json:"groups,omitempty"`
	// InterestCategories Interests categories IDs
	InterestCategories string `json:"interest_categories,omitempty"`
	// Interests Interests
	Interests string `json:"interests,omitempty"`
	// Paying Information whether the user has proceeded VK payments before
	Paying *BaseBoolInt `json:"paying,omitempty"`
	// Positions Positions IDs
	Positions string `json:"positions,omitempty"`
	// Religions Religions IDs
	Religions string `json:"religions,omitempty"`
	// RetargetingGroups Retargeting groups IDs
	RetargetingGroups string `json:"retargeting_groups,omitempty"`
	// RetargetingGroupsNot Retargeting groups IDs to except
	RetargetingGroupsNot string `json:"retargeting_groups_not,omitempty"`
	// SchoolFrom School graduation year from
	SchoolFrom int64 `json:"school_from,omitempty"`
	// SchoolTo School graduation year to
	SchoolTo int64 `json:"school_to,omitempty"`
	// Schools Schools IDs
	Schools string `json:"schools,omitempty"`
	// Sex
	Sex *AdsCriteriaSex `json:"sex,omitempty"`
	// Stations Stations IDs
	Stations string `json:"stations,omitempty"`
	// Statuses Relationship statuses
	Statuses string `json:"statuses,omitempty"`
	// Streets Streets IDs
	Streets string `json:"streets,omitempty"`
	// Travellers Travellers only
	Travellers *BasePropertyExists `json:"travellers,omitempty"`
	// UniFrom University graduation year from
	UniFrom int64 `json:"uni_from,omitempty"`
	// UniTo University graduation year to
	UniTo int64 `json:"uni_to,omitempty"`
	// UserBrowsers Browsers
	UserBrowsers string `json:"user_browsers,omitempty"`
	// UserDevices Devices
	UserDevices string `json:"user_devices,omitempty"`
	// UserOs Operating systems
	UserOs string `json:"user_os,omitempty"`
}

// LeadsChecked type from VK API Schema(leads_checked).
type LeadsChecked struct {

	// Reason Reason why user can&#39;t start the lead
	Reason string `json:"reason,omitempty"`
	// Result
	Result *LeadsCheckedResult `json:"result,omitempty"`
	// Sid Session ID
	Sid string `json:"sid,omitempty"`
	// StartLink URL user should open to start the lead
	StartLink string `json:"start_link,omitempty"`
}

// PhotosMarketUploadResponse type from VK API Schema(photos_market_upload_response).
type PhotosMarketUploadResponse struct {

	// CropData Crop data
	CropData string `json:"crop_data,omitempty"`
	// CropHash Crop hash
	CropHash string `json:"crop_hash,omitempty"`
	// GroupID Community ID
	GroupID int64 `json:"group_id,omitempty"`
	// Hash Uploading hash
	Hash string `json:"hash,omitempty"`
	// Photo Uploaded photo data
	Photo string `json:"photo,omitempty"`
	// Server Upload server number
	Server int64 `json:"server,omitempty"`
}

// PlacesTypes type from VK API Schema(places_types).
type PlacesTypes struct {

	// Icon URL of the place&#39;s icon
	Icon string `json:"icon,omitempty"`
	// ID Place type ID
	ID int64 `json:"id,omitempty"`
	// Title Place type title
	Title string `json:"title,omitempty"`
}

// DatabaseFaculty type from VK API Schema(database_faculty).
type DatabaseFaculty struct {

	// ID Faculty ID
	ID int64 `json:"id,omitempty"`
	// Title Faculty title
	Title string `json:"title,omitempty"`
}

// GroupsGroupCategory type from VK API Schema(groups_group_category).
type GroupsGroupCategory struct {

	// ID Category ID
	ID int64 `json:"id"`
	// Name Category name
	Name string `json:"name"`
	// Subcategories
	Subcategories []*BaseObjectWithName `json:"subcategories,omitempty"`
}

// UsersUserCounters type from VK API Schema(users_user_counters).
type UsersUserCounters struct {

	// Albums Albums number
	Albums int64 `json:"albums,omitempty"`
	// Audios Audios number
	Audios int64 `json:"audios,omitempty"`
	// Followers Followers number
	Followers int64 `json:"followers,omitempty"`
	// Friends Friends number
	Friends int64 `json:"friends,omitempty"`
	// Gifts Gifts number
	Gifts int64 `json:"gifts,omitempty"`
	// Groups Communities number
	Groups int64 `json:"groups,omitempty"`
	// Notes Notes number
	Notes int64 `json:"notes,omitempty"`
	// OnlineFriends Online friends number
	OnlineFriends int64 `json:"online_friends,omitempty"`
	// Pages Public pages number
	Pages int64 `json:"pages,omitempty"`
	// Photos Photos number
	Photos int64 `json:"photos,omitempty"`
	// Subscriptions Subscriptions number
	Subscriptions int64 `json:"subscriptions,omitempty"`
	// UserPhotos Number of photos with user
	UserPhotos int64 `json:"user_photos,omitempty"`
	// UserVideos Number of videos with user
	UserVideos int64 `json:"user_videos,omitempty"`
	// Videos Videos number
	Videos int64 `json:"videos,omitempty"`
}

// VideoSaveResult type from VK API Schema(video_save_result).
type VideoSaveResult struct {

	// Description Video description
	Description string `json:"description,omitempty"`
	// OwnerID Video owner ID
	OwnerID int64 `json:"owner_id,omitempty"`
	// Title Video title
	Title string `json:"title,omitempty"`
	// UploadURL URL for the video uploading
	UploadURL string `json:"upload_url,omitempty"`
	// VideoID Video ID
	VideoID int64 `json:"video_id,omitempty"`
}

// AdsTargSuggestions type from VK API Schema(ads_targ_suggestions).
type AdsTargSuggestions struct {

	// ID Object ID
	ID int64 `json:"id,omitempty"`
	// Name Object name
	Name string `json:"name,omitempty"`
}

// GroupsGroupXtrInvitedByType type from VK API Schema(groups_group_xtr_invited_by_type). Community type
type GroupsGroupXtrInvitedByType string

// NewsfeedItemPhoto type from VK API Schema(newsfeed_item_photo).
type NewsfeedItemPhoto struct {

	// Photos
	Photos *NewsfeedItemPhotoPhotos `json:"photos,omitempty"`
	// PostID Post ID
	PostID int64 `json:"post_id,omitempty"`
}

// PhotosWallUploadResponse type from VK API Schema(photos_wall_upload_response).
type PhotosWallUploadResponse struct {

	// Hash Uploading hash
	Hash string `json:"hash,omitempty"`
	// Photo Uploaded photo data
	Photo string `json:"photo,omitempty"`
	// Server Upload server number
	Server int64 `json:"server,omitempty"`
}

// PhotosPhotoAlbum type from VK API Schema(photos_photo_album).
type PhotosPhotoAlbum struct {

	// Created Date when the album has been created in Unixtime
	Created int64 `json:"created"`
	// Description Photo album description
	Description string `json:"description,omitempty"`
	// ID Photo album ID
	ID int64 `json:"id"`
	// OwnerID Album owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// Size Photos number
	Size int64 `json:"size"`
	// Thumb
	Thumb *PhotosPhoto `json:"thumb,omitempty"`
	// Title Photo album title
	Title string `json:"title"`
	// Updated Date when the album has been updated last time in Unixtime
	Updated int64 `json:"updated"`
}

// BaseGeoCoordinates type from VK API Schema(base_geo_coordinates).
type BaseGeoCoordinates struct {

	// Latitude
	Latitude float64 `json:"latitude"`
	// Longitude
	Longitude float64 `json:"longitude"`
}

// BaseImage type from VK API Schema(base_image).
type BaseImage struct {

	// Height Image height
	Height int64 `json:"height"`
	// URL Image url
	URL string `json:"url"`
	// Width Image width
	Width int64 `json:"width"`
}

// GroupsOwnerXtrBanInfo type from VK API Schema(groups_owner_xtr_ban_info).
type GroupsOwnerXtrBanInfo struct {

	// BanInfo
	BanInfo *GroupsBanInfo `json:"ban_info,omitempty"`
	// Group Information about group if type = group
	Group *GroupsGroup `json:"group,omitempty"`
	// Profile Information about group if type = profile
	Profile *UsersUser `json:"profile,omitempty"`
	// Type
	Type *GroupsOwnerXtrBanInfoType `json:"type,omitempty"`
}

// LeadsLead type from VK API Schema(leads_lead).
type LeadsLead struct {

	// Completed Completed offers number
	Completed int64 `json:"completed,omitempty"`
	// Cost Offer cost
	Cost int64 `json:"cost,omitempty"`
	// Days
	Days *LeadsLeadDays `json:"days,omitempty"`
	// Impressions Impressions number
	Impressions int64 `json:"impressions,omitempty"`
	// Limit Lead limit
	Limit int64 `json:"limit,omitempty"`
	// Spent Amount of spent votes
	Spent int64 `json:"spent,omitempty"`
	// Started Started offers number
	Started int64 `json:"started,omitempty"`
}

// NewsfeedItemPhotoPhotos type from VK API Schema(newsfeed_item_photo_photos).
type NewsfeedItemPhotoPhotos struct {

	// Count Photos number
	Count int64 `json:"count,omitempty"`
	// Items
	Items []*NewsfeedNewsfeedPhoto `json:"items,omitempty"`
}

// AdsTargSuggestionsSchools type from VK API Schema(ads_targ_suggestions_schools).
type AdsTargSuggestionsSchools struct {

	// Desc Full school title
	Desc string `json:"desc,omitempty"`
	// ID School ID
	ID int64 `json:"id,omitempty"`
	// Name School title
	Name string `json:"name,omitempty"`
	// Parent City name
	Parent string `json:"parent,omitempty"`
	// Type
	Type *AdsTargSuggestionsSchoolsType `json:"type,omitempty"`
}

// GiftsGiftPrivacy type from VK API Schema(gifts_gift_privacy). Gift privacy
type GiftsGiftPrivacy int64

// SearchHint type from VK API Schema(search_hint).
type SearchHint struct {

	// Description Object description
	Description string `json:"description"`
	// Global Information whether the object has been found globally
	Global *BaseBoolInt `json:"global,omitempty"`
	// Group
	Group *GroupsGroup `json:"group,omitempty"`
	// Profile
	Profile *UsersUserMin `json:"profile,omitempty"`
	// Section
	Section *SearchHintSection `json:"section"`
	// Type
	Type *SearchHintType `json:"type"`
}

// UsersUserLim type from VK API Schema(users_user_lim).
type UsersUserLim struct {

	// ID User ID
	ID int64 `json:"id,omitempty"`
	// Name User name and last name
	Name string `json:"name,omitempty"`
	// NameGen User name in genitive declension
	NameGen string `json:"name_gen,omitempty"`
	// Photo URL of square photo of the user with 50 pixels in width
	Photo string `json:"photo,omitempty"`
}

// VideoCatBlockView type from VK API Schema(video_cat_block_view). Type of view
type VideoCatBlockView string

// VideoCatBlock type from VK API Schema(video_cat_block).
type VideoCatBlock struct {

	// CanHide Information whether the block can be hidden
	CanHide *BaseBoolInt `json:"can_hide"`
	// ID Block ID
	ID int64 `json:"id"`
	// Items
	Items []*VideoCatElement `json:"items"`
	// Name Block name
	Name string `json:"name"`
	// Next New value for _from_ parameter
	Next string `json:"next"`
	// Type
	Type *VideoCatBlockView `json:"type,omitempty"`
	// View
	View *VideoCatBlockView `json:"view"`
}

// AppsLeaderboard type from VK API Schema(apps_leaderboard).
type AppsLeaderboard struct {

	// Level Level
	Level int64 `json:"level,omitempty"`
	// Points Points number
	Points int64 `json:"points,omitempty"`
	// Score Score number
	Score int64 `json:"score,omitempty"`
	// UserID User ID
	UserID int64 `json:"user_id"`
}

// AudioAudioUploadResponse type from VK API Schema(audio_audio_upload_response).
type AudioAudioUploadResponse struct {

	// Audio Uploaded aduio data
	Audio string `json:"audio,omitempty"`
	// Hash Uploading hash
	Hash string `json:"hash,omitempty"`
	// Redirect Redirect URL
	Redirect string `json:"redirect,omitempty"`
	// Server Upload server number
	Server int64 `json:"server,omitempty"`
}

// GroupsGroupFullMemberStatus type from VK API Schema(groups_group_full_member_status).
type GroupsGroupFullMemberStatus int64

// MessagesChatFull type from VK API Schema(messages_chat_full).
type MessagesChatFull struct {

	// AdminID Chat creator ID
	AdminID int64 `json:"admin_id"`
	// ID Chat ID
	ID int64 `json:"id"`
	// Kicked Shows that user has been kicked from the chat
	Kicked *BaseBoolInt `json:"kicked,omitempty"`
	// Left Shows that user has been left the chat
	Left *BaseBoolInt `json:"left,omitempty"`
	// Photo100 URL of the preview image with 100 px in width
	Photo100 string `json:"photo_100,omitempty"`
	// Photo200 URL of the preview image with 200 px in width
	Photo200 string `json:"photo_200,omitempty"`
	// Photo50 URL of the preview image with 50 px in width
	Photo50 string `json:"photo_50,omitempty"`
	// PushSettings
	PushSettings *MessagesChatPushSettings `json:"push_settings,omitempty"`
	// Title Chat title
	Title string `json:"title,omitempty"`
	// Type Chat type
	Type string `json:"type"`
	// Users
	Users []*MessagesUserXtrInvitedBy `json:"users"`
}

// NotificationsFeedback type from VK API Schema(notifications_feedback).
type NotificationsFeedback struct {

	// Attachments
	Attachments []*WallWallpostAttachment `json:"attachments,omitempty"`
	// FromID Reply author&#39;s ID
	FromID int64 `json:"from_id,omitempty"`
	// Geo
	Geo *BaseGeo `json:"geo,omitempty"`
	// ID Item ID
	ID int64 `json:"id,omitempty"`
	// Likes
	Likes *BaseLikesInfo `json:"likes,omitempty"`
	// Text Reply text
	Text string `json:"text,omitempty"`
	// ToID Wall owner&#39;s ID
	ToID int64 `json:"to_id,omitempty"`
}

// NewsfeedNewsfeedItemType type from VK API Schema(newsfeed_newsfeed_item_type). Item type
type NewsfeedNewsfeedItemType string

// NotificationsNotificationsComment type from VK API Schema(notifications_notifications_comment).
type NotificationsNotificationsComment struct {

	// Date Date when the comment has been added in Unixtime
	Date int64 `json:"date,omitempty"`
	// ID Comment ID
	ID int64 `json:"id,omitempty"`
	// OwnerID Author ID
	OwnerID int64 `json:"owner_id,omitempty"`
	// Photo
	Photo *PhotosPhoto `json:"photo,omitempty"`
	// Post
	Post *WallWallpost `json:"post,omitempty"`
	// Text Comment text
	Text string `json:"text,omitempty"`
	// Topic
	Topic *BoardTopic `json:"topic,omitempty"`
	// Video
	Video *VideoVideo `json:"video,omitempty"`
}

// PhotosPhotoFullXtrRealOffset type from VK API Schema(photos_photo_full_xtr_real_offset).
type PhotosPhotoFullXtrRealOffset struct {

	// AccessKey Access key for the photo
	AccessKey string `json:"access_key,omitempty"`
	// AlbumID Album ID
	AlbumID int64 `json:"album_id"`
	// CanComment
	CanComment *BaseBoolInt `json:"can_comment,omitempty"`
	// Comments
	Comments *BaseObjectCount `json:"comments,omitempty"`
	// Date Date when uploaded
	Date int64 `json:"date"`
	// Height Original photo height
	Height int64 `json:"height,omitempty"`
	// Hidden Returns if the photo is hidden above the wall
	Hidden *BasePropertyExists `json:"hidden,omitempty"`
	// ID Photo ID
	ID int64 `json:"id"`
	// Lat Latitude
	Lat float64 `json:"lat,omitempty"`
	// Likes
	Likes *BaseLikes `json:"likes,omitempty"`
	// Long Longitude
	Long float64 `json:"long,omitempty"`
	// OwnerID Photo owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// Photo1280 URL of image with 1280 px width
	Photo1280 string `json:"photo_1280,omitempty"`
	// Photo130 URL of image with 130 px width
	Photo130 string `json:"photo_130,omitempty"`
	// Photo2560 URL of image with 2560 px width
	Photo2560 string `json:"photo_2560,omitempty"`
	// Photo604 URL of image with 604 px width
	Photo604 string `json:"photo_604,omitempty"`
	// Photo75 URL of image with 75 px width
	Photo75 string `json:"photo_75,omitempty"`
	// Photo807 URL of image with 807 px width
	Photo807 string `json:"photo_807,omitempty"`
	// PostID Post ID
	PostID int64 `json:"post_id,omitempty"`
	// RealOffset Real position of the photo
	RealOffset int64 `json:"real_offset,omitempty"`
	// Reposts
	Reposts *BaseObjectCount `json:"reposts,omitempty"`
	// Sizes
	Sizes []*PhotosPhotoSizes `json:"sizes,omitempty"`
	// Tags
	Tags *BaseObjectCount `json:"tags,omitempty"`
	// Text Photo caption
	Text string `json:"text,omitempty"`
	// UserID ID of the user who have uploaded the photo
	UserID int64 `json:"user_id,omitempty"`
	// Width Original photo width
	Width int64 `json:"width,omitempty"`
}

// AccountLookupResult type from VK API Schema(account_lookup_result).
type AccountLookupResult struct {

	// Found
	Found []*AccountUserXtrContact `json:"found,omitempty"`
	// Other
	Other []*AccountOtherContact `json:"other,omitempty"`
}

// BaseLinkButtonAction type from VK API Schema(base_link_button_action).
type BaseLinkButtonAction struct {

	// Type
	Type *BaseLinkButtonActionType `json:"type,omitempty"`
	// URL Action URL
	URL string `json:"url,omitempty"`
}

// BaseUserID type from VK API Schema(base_user_id).
type BaseUserID struct {

	// UserID User ID
	UserID int64 `json:"user_id,omitempty"`
}

// MarketMarketItemAvailability type from VK API Schema(market_market_item_availability). Information whether the item is available
type MarketMarketItemAvailability int64

// MessagesDialog type from VK API Schema(messages_dialog).
type MessagesDialog struct {

	// Important Is it an important dialog
	Important *BaseBoolInt `json:"important,omitempty"`
	// InRead ID of the last message read by current user
	InRead int64 `json:"in_read,omitempty"`
	// Message
	Message *MessagesMessage `json:"message,omitempty"`
	// OutRead ID of the last message read by the others
	OutRead int64 `json:"out_read,omitempty"`
	// Unanswered Is it an unanswered dialog
	Unanswered *BaseBoolInt `json:"unanswered,omitempty"`
	// Unread Information whether unread messages are in the dialog
	Unread int64 `json:"unread,omitempty"`
}

// PollsPoll type from VK API Schema(polls_poll).
type PollsPoll struct {

	// Anonymous Information whether the pole is anonymous
	Anonymous *BaseBoolInt `json:"anonymous"`
	// AnswerID Current user&#39;s answer ID
	AnswerID int64 `json:"answer_id"`
	// Answers
	Answers []*PollsAnswer `json:"answers"`
	// Created Date when poll has been created in Unixtime
	Created int64 `json:"created"`
	// ID Poll ID
	ID int64 `json:"id"`
	// OwnerID Poll owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// Question Poll question
	Question string `json:"question"`
	// Votes Votes number
	Votes string `json:"votes"`
}

// UsersLastSeen type from VK API Schema(users_last_seen).
type UsersLastSeen struct {

	// Platform Type of the platform that used for the last authorization
	Platform int64 `json:"platform,omitempty"`
	// Time Last visit date (in Unix time)
	Time int64 `json:"time,omitempty"`
}

// UsersCropPhoto type from VK API Schema(users_crop_photo).
type UsersCropPhoto struct {

	// Crop
	Crop *UsersCropPhotoCrop `json:"crop,omitempty"`
	// Photo
	Photo *PhotosPhoto `json:"photo,omitempty"`
	// Rect
	Rect *UsersCropPhotoRect `json:"rect,omitempty"`
}

// PollsAnswer type from VK API Schema(polls_answer).
type PollsAnswer struct {

	// ID Answer ID
	ID int64 `json:"id"`
	// Rate Answer rate in percents
	Rate float64 `json:"rate"`
	// Text Answer text
	Text string `json:"text"`
	// Votes Votes number
	Votes int64 `json:"votes"`
}

// SecureSmsNotification type from VK API Schema(secure_sms_notification).
type SecureSmsNotification struct {

	// AppID Application ID
	AppID int64 `json:"app_id,omitempty"`
	// Date Date when message has been sent in Unixtime
	Date int64 `json:"date,omitempty"`
	// ID Notification ID
	ID int64 `json:"id,omitempty"`
	// Message Messsage text
	Message string `json:"message,omitempty"`
	// UserID User ID
	UserID int64 `json:"user_id,omitempty"`
}

// UtilsLinkStatsExtended type from VK API Schema(utils_link_stats_extended).
type UtilsLinkStatsExtended struct {

	// Key Link key (characters after vk.cc/)
	Key string `json:"key,omitempty"`
	// Stats
	Stats []*UtilsStatsExtended `json:"stats,omitempty"`
}

// AdsAccessRole type from VK API Schema(ads_access_role). Current user&#39;s role
type AdsAccessRole string

// BaseLink type from VK API Schema(base_link).
type BaseLink struct {

	// Application
	Application *BaseLinkApplication `json:"application,omitempty"`
	// Button
	Button *BaseLinkButton `json:"button,omitempty"`
	// Caption Link caption
	Caption string `json:"caption,omitempty"`
	// Description Link description
	Description string `json:"description,omitempty"`
	// Photo
	Photo *PhotosPhoto `json:"photo,omitempty"`
	// PreviewPage String ID of the page with article preview
	PreviewPage string `json:"preview_page,omitempty"`
	// PreviewURL URL of the page with article preview
	PreviewURL string `json:"preview_url,omitempty"`
	// Product
	Product *BaseLinkProduct `json:"product,omitempty"`
	// Rating
	Rating *BaseLinkRating `json:"rating,omitempty"`
	// Title Link title
	Title string `json:"title,omitempty"`
	// URL Link URL
	URL string `json:"url"`
}

// FriendsRequestsMutual type from VK API Schema(friends_requests_mutual).
type FriendsRequestsMutual struct {

	// Count Total mutual friends number
	Count int64 `json:"count,omitempty"`
	// Users
	Users []int64 `json:"users,omitempty"`
}

// MessagesMessageAttachmentType type from VK API Schema(messages_message_attachment_type). Attachment type
type MessagesMessageAttachmentType string

// MessagesMessageAttachment type from VK API Schema(messages_message_attachment).
type MessagesMessageAttachment struct {

	// Audio
	Audio *AudioAudioFull `json:"audio,omitempty"`
	// Doc
	Doc *DocsDoc `json:"doc,omitempty"`
	// Gift
	Gift *GiftsLayout `json:"gift,omitempty"`
	// Link
	Link *BaseLink `json:"link,omitempty"`
	// Market
	Market *MarketMarketItem `json:"market,omitempty"`
	// MarketMarketAlbum
	MarketMarketAlbum *MarketMarketAlbum `json:"market_market_album,omitempty"`
	// Photo
	Photo *PhotosPhoto `json:"photo,omitempty"`
	// Sticker
	Sticker *BaseSticker `json:"sticker,omitempty"`
	// Type
	Type *MessagesMessageAttachmentType `json:"type"`
	// Video
	Video *VideoVideo `json:"video,omitempty"`
	// Wall
	Wall *WallWallpostAttached `json:"wall,omitempty"`
	// WallReply
	WallReply *WallWallComment `json:"wall_reply,omitempty"`
}

// UsersUserFullXtrType type from VK API Schema(users_user_full_xtr_type).
type UsersUserFullXtrType struct {
}

// AdsTargSuggestionsCities type from VK API Schema(ads_targ_suggestions_cities).
type AdsTargSuggestionsCities struct {

	// ID Object ID
	ID int64 `json:"id,omitempty"`
	// Name Object name
	Name string `json:"name,omitempty"`
	// Parent Parent object
	Parent string `json:"parent,omitempty"`
}

// BaseLinkButton type from VK API Schema(base_link_button).
type BaseLinkButton struct {

	// Action Button action
	Action *BaseLinkButtonAction `json:"action,omitempty"`
	// Title Button title
	Title string `json:"title,omitempty"`
}

// MessagesLongpollParams type from VK API Schema(messages_longpoll_params).
type MessagesLongpollParams struct {

	// Key Key
	Key string `json:"key,omitempty"`
	// Pts Persistent timestamp
	Pts int64 `json:"pts,omitempty"`
	// Server Server URL
	Server string `json:"server,omitempty"`
	// Ts Timestamp
	Ts int64 `json:"ts,omitempty"`
}

// MessagesConversationWithMessage type from VK API Schema(messages_conversation_with_message).
type MessagesConversationWithMessage struct {

	// Conversation No struct in JSON SCHEMA
	// Conversation *MessagesConversation `json:"conversation,omitempty"`
	// LastMessage
	LastMessage *MessagesMessage `json:"last_message,omitempty"`
}

// MessagesKeyboardButtonAction type from VK API Schema(messages_keyboard_button_action). Description of the action, that should be performed on button click
type MessagesKeyboardButtonAction struct {

	// Label Label for button
	Label string `json:"label,omitempty"`
	// Payload Additional data sent along with message for developer convenience
	Payload string `json:"payload,omitempty"`
	// Type Button type
	Type string `json:"type"`
}

// NewsfeedItemNoteNotes type from VK API Schema(newsfeed_item_note_notes).
type NewsfeedItemNoteNotes struct {

	// Count Notes number
	Count int64 `json:"count,omitempty"`
	// Items
	Items []*NewsfeedNewsfeedNote `json:"items,omitempty"`
}

// PhotosPhotoTag type from VK API Schema(photos_photo_tag).
type PhotosPhotoTag struct {

	// Date Date when tag has been added in Unixtime
	Date int64 `json:"date"`
	// ID Tag ID
	ID int64 `json:"id"`
	// PlacerID ID of the tag creator
	PlacerID int64 `json:"placer_id"`
	// TaggedName Tag description
	TaggedName string `json:"tagged_name"`
	// UserID Tagged user ID
	UserID int64 `json:"user_id"`
	// Viewed Information whether the tag is reviewed
	Viewed *BaseBoolInt `json:"viewed"`
	// X Coordinate X of the left upper corner
	X float64 `json:"x"`
	// X2 Coordinate X of the right lower corner
	X2 float64 `json:"x2"`
	// Y Coordinate Y of the left upper corner
	Y float64 `json:"y"`
	// Y2 Coordinate Y of the right lower corner
	Y2 float64 `json:"y2"`
}

// AccountNameRequestStatus type from VK API Schema(account_name_request_status). Request status
type AccountNameRequestStatus string

// AppsAppLeaderboardType type from VK API Schema(apps_app_leaderboard_type). Leaderboard type
type AppsAppLeaderboardType int64

// BaseLinkRating type from VK API Schema(base_link_rating).
type BaseLinkRating struct {

	// ReviewsCount Count of reviews
	ReviewsCount int64 `json:"reviews_count,omitempty"`
	// Stars Count of stars
	Stars float64 `json:"stars,omitempty"`
}

// BaseSticker type from VK API Schema(base_sticker).
type BaseSticker struct {

	// Images
	Images []*BaseImage `json:"images,omitempty"`
	// ImagesWithBackground
	ImagesWithBackground []*BaseImage `json:"images_with_background,omitempty"`
	// ProductID Collection ID
	ProductID int64 `json:"product_id,omitempty"`
	// StickerID Sticker ID
	StickerID int64 `json:"sticker_id,omitempty"`
}

// GroupsRoleOptions type from VK API Schema(groups_role_options). User&#39;s credentials as community admin
type GroupsRoleOptions string

// StatsSexAge type from VK API Schema(stats_sex_age).
type StatsSexAge struct {

	// Count Visitors number
	Count int64 `json:"count,omitempty"`
	// Value Sex/age value
	Value string `json:"value,omitempty"`
}

// WallPostSourceType type from VK API Schema(wall_post_source_type). Type of post source
type WallPostSourceType string

// WallPostType type from VK API Schema(wall_post_type). Post type
type WallPostType string

// WidgetsWidgetPage type from VK API Schema(widgets_widget_page).
type WidgetsWidgetPage struct {

	// Comments
	Comments *BaseObjectCount `json:"comments,omitempty"`
	// Date Date when widgets on the page has been initialized firstly in Unixtime
	Date int64 `json:"date,omitempty"`
	// Description Page description
	Description string `json:"description,omitempty"`
	// ID Page ID
	ID int64 `json:"id,omitempty"`
	// Likes
	Likes *BaseObjectCount `json:"likes,omitempty"`
	// PageID page_id parameter value
	PageID string `json:"page_id,omitempty"`
	// Photo URL of the preview image
	Photo string `json:"photo,omitempty"`
	// Title Page title
	Title string `json:"title,omitempty"`
	// URL Page absolute URL
	URL string `json:"url,omitempty"`
}

// AudioAudio type from VK API Schema(audio_audio).
type AudioAudio struct {

	// AccessKey Access key for the audio
	AccessKey string `json:"access_key,omitempty"`
	// Artist Artist name
	Artist string `json:"artist"`
	// ID Audio ID
	ID int64 `json:"id"`
	// OwnerID Audio owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// Title Title
	Title string `json:"title"`
	// URL URL of mp3 file
	URL string `json:"url,omitempty"`
}

// GroupsCover type from VK API Schema(groups_cover).
type GroupsCover struct {

	// Enabled Information whether cover is enabled
	Enabled *BaseBoolInt `json:"enabled"`
	// Images
	Images []*BaseImage `json:"images,omitempty"`
}

// GroupsLinksItem type from VK API Schema(groups_links_item).
type GroupsLinksItem struct {

	// Desc Link description
	Desc string `json:"desc,omitempty"`
	// EditTitle Information whether the link title can be edited
	EditTitle *BaseBoolInt `json:"edit_title,omitempty"`
	// ID Link ID
	ID int64 `json:"id,omitempty"`
	// Name Link title
	Name string `json:"name,omitempty"`
	// Photo100 URL of square image of the link with 100 pixels in width
	Photo100 string `json:"photo_100,omitempty"`
	// Photo50 URL of square image of the link with 50 pixels in width
	Photo50 string `json:"photo_50,omitempty"`
	// URL Link URL
	URL string `json:"url,omitempty"`
}

// MarketMarketItemFull type from VK API Schema(market_market_item_full).
type MarketMarketItemFull struct {
}

// UsersCropPhotoCrop type from VK API Schema(users_crop_photo_crop).
type UsersCropPhotoCrop struct {

	// X Coordinate X of the left upper corner
	X float64 `json:"x,omitempty"`
	// X2 Coordinate X of the right lower corner
	X2 float64 `json:"x2,omitempty"`
	// Y Coordinate Y of the left upper corner
	Y float64 `json:"y,omitempty"`
	// Y2 Coordinate Y of the right lower corner
	Y2 float64 `json:"y2,omitempty"`
}

// UtilsLastShortenedLink type from VK API Schema(utils_last_shortened_link).
type UtilsLastShortenedLink struct {

	// AccessKey Access key for private stats
	AccessKey string `json:"access_key,omitempty"`
	// Key Link key (characters after vk.cc/)
	Key string `json:"key,omitempty"`
	// ShortURL Short link URL
	ShortURL string `json:"short_url,omitempty"`
	// Timestamp Creation time in Unixtime
	Timestamp int64 `json:"timestamp,omitempty"`
	// URL Full URL
	URL string `json:"url,omitempty"`
	// Views Total views number
	Views int64 `json:"views,omitempty"`
}

// BoardDefaultOrder type from VK API Schema(board_default_order). Sort type
type BoardDefaultOrder int64

// BoardTopicComment type from VK API Schema(board_topic_comment).
type BoardTopicComment struct {

	// Attachments
	Attachments []*WallCommentAttachment `json:"attachments,omitempty"`
	// Date Date when the comment has been added in Unixtime
	Date int64 `json:"date"`
	// FromID Author ID
	FromID int64 `json:"from_id"`
	// ID Comment ID
	ID int64 `json:"id"`
	// RealOffset Real position of the comment
	RealOffset int64 `json:"real_offset,omitempty"`
	// Text Comment text
	Text string `json:"text"`
}

// NewsfeedItemFriend type from VK API Schema(newsfeed_item_friend).
type NewsfeedItemFriend struct {

	// Friends
	Friends *NewsfeedItemFriendFriends `json:"friends,omitempty"`
}

// NotesNoteComment type from VK API Schema(notes_note_comment).
type NotesNoteComment struct {

	// Date Date when the comment has beed added in Unixtime
	Date int64 `json:"date"`
	// ID Comment ID
	ID int64 `json:"id"`
	// Message Comment text
	Message string `json:"message"`
	// Nid Note ID
	Nid int64 `json:"nid"`
	// Oid Note ID
	Oid int64 `json:"oid"`
	// ReplyTo ID of replied comment
	ReplyTo int64 `json:"reply_to,omitempty"`
	// UID Comment author&#39;s ID
	UID int64 `json:"uid"`
}

// PlacesPlaceMin type from VK API Schema(places_place_min).
type PlacesPlaceMin struct {

	// Address Place address
	Address string `json:"address,omitempty"`
	// Checkins Checkins number
	Checkins int64 `json:"checkins,omitempty"`
	// City City ID
	City int64 `json:"city,omitempty"`
	// Country Country ID
	Country int64 `json:"country,omitempty"`
	// Created Date of the place creation in Unixtime
	Created int64 `json:"created,omitempty"`
	// Icon URL of the place&#39;s icon
	Icon string `json:"icon,omitempty"`
	// ID Place ID
	ID int64 `json:"id,omitempty"`
	// Latitude Place latitude
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude Place longitude
	Longitude float64 `json:"longitude,omitempty"`
	// Title Place title
	Title string `json:"title,omitempty"`
	// Type Place type
	Type string `json:"type,omitempty"`
}

// AdsDemoStats type from VK API Schema(ads_demo_stats).
type AdsDemoStats struct {

	// ID Object ID
	ID int64 `json:"id,omitempty"`
	// Stats
	Stats *AdsDemostatsFormat `json:"stats,omitempty"`
	// Type
	Type *AdsObjectType `json:"type,omitempty"`
}

// AppsAppType type from VK API Schema(apps_app_type). Application type
type AppsAppType string

// LeadsComplete type from VK API Schema(leads_complete).
type LeadsComplete struct {

	// Cost Offer cost
	Cost int64 `json:"cost,omitempty"`
	// Limit Offer limit
	Limit int64 `json:"limit,omitempty"`
	// Spent Amount of spent votes
	Spent int64 `json:"spent,omitempty"`
	// Success
	Success *BaseOkResponse `json:"success,omitempty"`
	// TestMode Information whether test mode is enabled
	TestMode *BaseBoolInt `json:"test_mode,omitempty"`
}

// OrdersOrder type from VK API Schema(orders_order).
type OrdersOrder struct {

	// Amount Amount
	Amount int64 `json:"amount,omitempty"`
	// AppOrderID App order ID
	AppOrderID int64 `json:"app_order_id,omitempty"`
	// CancelTransactionID Cancel transaction ID
	CancelTransactionID int64 `json:"cancel_transaction_id,omitempty"`
	// Date Date of creation in Unixtime
	Date int64 `json:"date,omitempty"`
	// ID Order ID
	ID int64 `json:"id,omitempty"`
	// Item Order item
	Item string `json:"item,omitempty"`
	// ReceiverID Receiver ID
	ReceiverID int64 `json:"receiver_id,omitempty"`
	// Status Order status
	Status string `json:"status,omitempty"`
	// TransactionID Transaction ID
	TransactionID int64 `json:"transaction_id,omitempty"`
	// UserID User ID
	UserID int64 `json:"user_id,omitempty"`
}

// PhotosPhotoSizesType type from VK API Schema(photos_photo_sizes_type). Size type
type PhotosPhotoSizesType string

// GroupsOnlineStatus type from VK API Schema(groups_online_status). Online status of group
type GroupsOnlineStatus struct {

	// Minutes Estimated time of answer (for status = answer_mark)
	Minutes int64 `json:"minutes,omitempty"`
	// Status
	Status *GroupsOnlineStatusType `json:"status"`
}

// PhotosImage type from VK API Schema(photos_image).
type PhotosImage struct {

	// Height Height of the photo in px.
	Height int64 `json:"height,omitempty"`
	// Type
	Type *PhotosImageType `json:"type,omitempty"`
	// URL Photo URL.
	URL string `json:"url,omitempty"`
	// Width Width of the photo in px.
	Width int64 `json:"width,omitempty"`
}

// PhotosPhoto type from VK API Schema(photos_photo).
type PhotosPhoto struct {

	// AccessKey Access key for the photo
	AccessKey string `json:"access_key,omitempty"`
	// AlbumID Album ID
	AlbumID int64 `json:"album_id"`
	// Date Date when uploaded
	Date int64 `json:"date"`
	// Height Original photo height
	Height int64 `json:"height,omitempty"`
	// ID Photo ID
	ID int64 `json:"id"`
	// Images
	Images []*PhotosImage `json:"images,omitempty"`
	// Lat Latitude
	Lat float64 `json:"lat,omitempty"`
	// Long Longitude
	Long float64 `json:"long,omitempty"`
	// OwnerID Photo owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// PostID Post ID
	PostID int64 `json:"post_id,omitempty"`
	// Text Photo caption
	Text string `json:"text,omitempty"`
	// UserID ID of the user who have uploaded the photo
	UserID int64 `json:"user_id,omitempty"`
	// Width Original photo width
	Width int64 `json:"width,omitempty"`
}

// PhotosCommentXtrPid type from VK API Schema(photos_comment_xtr_pid).
type PhotosCommentXtrPid struct {

	// Attachments
	Attachments []*WallCommentAttachment `json:"attachments,omitempty"`
	// Date Date when the comment has been added in Unixtime
	Date int64 `json:"date"`
	// FromID Author ID
	FromID int64 `json:"from_id"`
	// ID Comment ID
	ID int64 `json:"id"`
	// Likes
	Likes *BaseLikesInfo `json:"likes,omitempty"`
	// Pid Photo ID
	Pid int64 `json:"pid"`
	// ReplyToComment Replied comment ID
	ReplyToComment int64 `json:"reply_to_comment,omitempty"`
	// ReplyToUser Replied user ID
	ReplyToUser int64 `json:"reply_to_user,omitempty"`
	// Text Comment text
	Text string `json:"text"`
}

// UsersCareer type from VK API Schema(users_career).
type UsersCareer struct {

	// CityID City ID
	CityID int64 `json:"city_id,omitempty"`
	// Company Company name
	Company string `json:"company,omitempty"`
	// CountryID Country ID
	CountryID int64 `json:"country_id,omitempty"`
	// From From year
	From int64 `json:"from,omitempty"`
	// GroupID Community ID
	GroupID int64 `json:"group_id,omitempty"`
	// Position Position
	Position string `json:"position,omitempty"`
	// Until Till year
	Until int64 `json:"until,omitempty"`
}

// WidgetsCommentRepliesItem type from VK API Schema(widgets_comment_replies_item).
type WidgetsCommentRepliesItem struct {

	// Cid Comment ID
	Cid int64 `json:"cid,omitempty"`
	// Date Date when the comment has been added in Unixtime
	Date int64 `json:"date,omitempty"`
	// Likes
	Likes *WidgetsWidgetLikes `json:"likes,omitempty"`
	// Text Comment text
	Text string `json:"text,omitempty"`
	// UID User ID
	UID int64 `json:"uid,omitempty"`
	// User
	User *UsersUserFull `json:"user,omitempty"`
}

// AdsFloodStats type from VK API Schema(ads_flood_stats).
type AdsFloodStats struct {

	// Left Requests left
	Left int64 `json:"left"`
	// Refresh Time to refresh in seconds
	Refresh int64 `json:"refresh"`
}

// BaseObjectWithName type from VK API Schema(base_object_with_name).
type BaseObjectWithName struct {

	// ID Object ID
	ID int64 `json:"id"`
	// Name Object name
	Name string `json:"name"`
}

// GroupsCountersGroup type from VK API Schema(groups_counters_group).
type GroupsCountersGroup struct {

	// Albums Photo albums number
	Albums int64 `json:"albums,omitempty"`
	// Audios Audios number
	Audios int64 `json:"audios,omitempty"`
	// Docs Docs number
	Docs int64 `json:"docs,omitempty"`
	// Market Market items number
	Market int64 `json:"market,omitempty"`
	// Photos Photos number
	Photos int64 `json:"photos,omitempty"`
	// Topics Topics number
	Topics int64 `json:"topics,omitempty"`
	// Videos Videos number
	Videos int64 `json:"videos,omitempty"`
}

// GroupsGroupPublicCategoryList type from VK API Schema(groups_group_public_category_list).
type GroupsGroupPublicCategoryList struct {

	// ID
	ID int64 `json:"id,omitempty"`
	// Name
	Name string `json:"name,omitempty"`
	// SubtypesList
	SubtypesList []*GroupsGroupCategoryType `json:"subtypes_list,omitempty"`
}

// MessagesKeyboard type from VK API Schema(messages_keyboard).
type MessagesKeyboard struct {

	// AuthorID Community or bot, which set this keyboard
	AuthorID int64 `json:"author_id,omitempty"`
	// Buttons
	Buttons [][]*MessagesKeyboardButton `json:"buttons"`
	// OneTime Should this keyboard disappear on first use
	OneTime bool `json:"one_time"`
}

// AccountUserSettings type from VK API Schema(account_user_settings).
type AccountUserSettings struct {

	// Bdate User&#39;s date of birth
	Bdate string `json:"bdate,omitempty"`
	// BdateVisibility Information whether user&#39;s birthdate are hidden
	BdateVisibility int64 `json:"bdate_visibility,omitempty"`
	// City
	City *BaseObject `json:"city,omitempty"`
	// Country
	Country *BaseCountry `json:"country,omitempty"`
	// FirstName User first name
	FirstName string `json:"first_name,omitempty"`
	// HomeTown User&#39;s hometown
	HomeTown string `json:"home_town,omitempty"`
	// LastName User last name
	LastName string `json:"last_name,omitempty"`
	// MaidenName User maiden name
	MaidenName string `json:"maiden_name,omitempty"`
	// NameRequest
	NameRequest *AccountNameRequest `json:"name_request,omitempty"`
	// Phone User phone number with some hidden digits
	Phone string `json:"phone,omitempty"`
	// Relation User relationship status
	Relation int64 `json:"relation,omitempty"`
	// RelationPartner
	RelationPartner *UsersUserMin `json:"relation_partner,omitempty"`
	// RelationPending Information whether relation status is pending
	RelationPending int64 `json:"relation_pending,omitempty"`
	// RelationRequests
	RelationRequests []*UsersUserMin `json:"relation_requests,omitempty"`
	// ScreenName Domain name of the user&#39;s page
	ScreenName string `json:"screen_name,omitempty"`
	// Sex User sex
	Sex *BaseSex `json:"sex,omitempty"`
	// Status User status
	Status string `json:"status,omitempty"`
}

// GroupsGroupCategoryType type from VK API Schema(groups_group_category_type).
type GroupsGroupCategoryType struct {

	// ID
	ID int64 `json:"id,omitempty"`
	// Name
	Name string `json:"name,omitempty"`
}

// PagesWikipageFull type from VK API Schema(pages_wikipage_full).
type PagesWikipageFull struct {

	// Created Date when the page has been created in Unixtime
	Created int64 `json:"created"`
	// CreatorID Page creator ID
	CreatorID int64 `json:"creator_id,omitempty"`
	// CurrentUserCanEdit Information whether current user can edit the page
	CurrentUserCanEdit *BaseBoolInt `json:"current_user_can_edit,omitempty"`
	// CurrentUserCanEditAccess Information whether current user can edit the page access settings
	CurrentUserCanEditAccess *BaseBoolInt `json:"current_user_can_edit_access,omitempty"`
	// Edited Date when the page has been edited in Unixtime
	Edited int64 `json:"edited"`
	// EditorID Last editor ID
	EditorID int64 `json:"editor_id,omitempty"`
	// GroupID Community ID
	GroupID int64 `json:"group_id"`
	// HTML Page content, HTML
	HTML string `json:"html,omitempty"`
	// ID Page ID
	ID int64 `json:"id"`
	// Source Page content, wiki
	Source string `json:"source,omitempty"`
	// Title Page title
	Title string `json:"title"`
	// ViewURL URL of the page preview
	ViewURL string `json:"view_url"`
	// Views Views number
	Views int64 `json:"views"`
	// WhoCanEdit Edit settings of the page
	WhoCanEdit *PagesPrivacySettings `json:"who_can_edit"`
	// WhoCanView View settings of the page
	WhoCanView *PagesPrivacySettings `json:"who_can_view"`
}

// UsersPersonal type from VK API Schema(users_personal).
type UsersPersonal struct {

	// Alcohol User&#39;s views on alcohol
	Alcohol int64 `json:"alcohol,omitempty"`
	// InspiredBy User&#39;s inspired by
	InspiredBy string `json:"inspired_by,omitempty"`
	// Langs
	Langs []string `json:"langs,omitempty"`
	// LifeMain User&#39;s personal priority in life
	LifeMain int64 `json:"life_main,omitempty"`
	// PeopleMain User&#39;s personal priority in people
	PeopleMain int64 `json:"people_main,omitempty"`
	// Political User&#39;s political views
	Political int64 `json:"political,omitempty"`
	// Religion User&#39;s religion
	Religion string `json:"religion,omitempty"`
	// Smoking User&#39;s views on smoking
	Smoking int64 `json:"smoking,omitempty"`
}

// WallWallComment type from VK API Schema(wall_wall_comment).
type WallWallComment struct {

	// Attachments
	Attachments []*WallCommentAttachment `json:"attachments,omitempty"`
	// Date Date when the comment has been added in Unixtime
	Date int64 `json:"date"`
	// FromID Author ID
	FromID int64 `json:"from_id"`
	// ID Comment ID
	ID int64 `json:"id"`
	// Likes
	Likes *BaseLikesInfo `json:"likes,omitempty"`
	// RealOffset Real position of the comment
	RealOffset int64 `json:"real_offset,omitempty"`
	// ReplyToComment Replied comment ID
	ReplyToComment int64 `json:"reply_to_comment,omitempty"`
	// ReplyToUser Replied user ID
	ReplyToUser int64 `json:"reply_to_user,omitempty"`
	// Text Comment text
	Text string `json:"text"`
}

// PhotosPhotoXtrRealOffset type from VK API Schema(photos_photo_xtr_real_offset).
type PhotosPhotoXtrRealOffset struct {

	// AccessKey Access key for the photo
	AccessKey string `json:"access_key,omitempty"`
	// AlbumID Album ID
	AlbumID int64 `json:"album_id"`
	// Date Date when uploaded
	Date int64 `json:"date"`
	// Height Original photo height
	Height int64 `json:"height,omitempty"`
	// Hidden Returns if the photo is hidden above the wall
	Hidden *BasePropertyExists `json:"hidden,omitempty"`
	// ID Photo ID
	ID int64 `json:"id"`
	// Lat Latitude
	Lat float64 `json:"lat,omitempty"`
	// Long Longitude
	Long float64 `json:"long,omitempty"`
	// OwnerID Photo owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// Photo1280 URL of image with 1280 px width
	Photo1280 string `json:"photo_1280,omitempty"`
	// Photo130 URL of image with 130 px width
	Photo130 string `json:"photo_130,omitempty"`
	// Photo2560 URL of image with 2560 px width
	Photo2560 string `json:"photo_2560,omitempty"`
	// Photo604 URL of image with 604 px width
	Photo604 string `json:"photo_604,omitempty"`
	// Photo75 URL of image with 75 px width
	Photo75 string `json:"photo_75,omitempty"`
	// Photo807 URL of image with 807 px width
	Photo807 string `json:"photo_807,omitempty"`
	// PostID Post ID
	PostID int64 `json:"post_id,omitempty"`
	// RealOffset Real position of the photo
	RealOffset int64 `json:"real_offset,omitempty"`
	// Sizes
	Sizes []*PhotosPhotoSizes `json:"sizes,omitempty"`
	// Text Photo caption
	Text string `json:"text,omitempty"`
	// UserID ID of the user who have uploaded the photo
	UserID int64 `json:"user_id,omitempty"`
	// Width Original photo width
	Width int64 `json:"width,omitempty"`
}

// StoriesStoryStats type from VK API Schema(stories_story_stats).
type StoriesStoryStats struct {

	// Answer
	Answer *StoriesStoryStatsStat `json:"answer"`
	// Bans
	Bans *StoriesStoryStatsStat `json:"bans"`
	// OpenLink
	OpenLink *StoriesStoryStatsStat `json:"open_link"`
	// Replies
	Replies *StoriesStoryStatsStat `json:"replies"`
	// Shares
	Shares *StoriesStoryStatsStat `json:"shares"`
	// Subscribers
	Subscribers *StoriesStoryStatsStat `json:"subscribers"`
	// Views
	Views *StoriesStoryStatsStat `json:"views"`
}

// WallAppPost type from VK API Schema(wall_app_post).
type WallAppPost struct {

	// ID Application ID
	ID int64 `json:"id,omitempty"`
	// Name Application name
	Name string `json:"name,omitempty"`
	// Photo130 URL of the preview image with 130 px in width
	Photo130 string `json:"photo_130,omitempty"`
	// Photo604 URL of the preview image with 604 px in width
	Photo604 string `json:"photo_604,omitempty"`
}

// BaseGeo type from VK API Schema(base_geo).
type BaseGeo struct {

	// Coordinates
	Coordinates *BaseGeoCoordinates `json:"coordinates,omitempty"`
	// Place
	Place *BasePlace `json:"place,omitempty"`
	// Showmap Information whether a map is showed
	Showmap int64 `json:"showmap,omitempty"`
	// Type Place type
	Type string `json:"type,omitempty"`
}

// BaseRepostsInfo type from VK API Schema(base_reposts_info).
type BaseRepostsInfo struct {

	// Count Reposts number
	Count int64 `json:"count,omitempty"`
	// UserReposted Information whether current user has reposted the post
	UserReposted int64 `json:"user_reposted,omitempty"`
}

// DocsDocPreviewVideo type from VK API Schema(docs_doc_preview_video).
type DocsDocPreviewVideo struct {

	// Filesize Video file size in bites
	Filesize int64 `json:"filesize"`
	// Height Video&#39;s height in pixels
	Height int64 `json:"height"`
	// Src Video URL
	Src string `json:"src"`
	// Width Video&#39;s width in pixels
	Width int64 `json:"width"`
}

// FriendsUserXtrLists type from VK API Schema(friends_user_xtr_lists).
type FriendsUserXtrLists struct {
}

// OrdersAmount type from VK API Schema(orders_amount).
type OrdersAmount struct {

	// Amounts
	Amounts []*OrdersAmountItem `json:"amounts,omitempty"`
	// Currency Currency name
	Currency string `json:"currency,omitempty"`
}

// AdsStatsSexAge type from VK API Schema(ads_stats_sex_age).
type AdsStatsSexAge struct {

	// ClicksRate Clicks rate
	ClicksRate float64 `json:"clicks_rate,omitempty"`
	// ImpressionsRate Impressions rate
	ImpressionsRate float64 `json:"impressions_rate,omitempty"`
	// Value Sex and age interval
	Value string `json:"value,omitempty"`
}

// AdsTargSuggestionsSchoolsType type from VK API Schema(ads_targ_suggestions_schools_type). School type
type AdsTargSuggestionsSchoolsType string

// AppsApp type from VK API Schema(apps_app).
type AppsApp struct {

	// AuthorGroup Official community&#39;s ID
	AuthorGroup int64 `json:"author_group,omitempty"`
	// AuthorID Application author&#39;s ID
	AuthorID int64 `json:"author_id,omitempty"`
	// AuthorURL Application author&#39;s URL
	AuthorURL string `json:"author_url,omitempty"`
	// Banner1120 URL of the app banner with 1120 px in width
	Banner1120 string `json:"banner_1120,omitempty"`
	// Banner560 URL of the app banner with 560 px in width
	Banner560 string `json:"banner_560,omitempty"`
	// CatalogPosition Catalog position
	CatalogPosition int64 `json:"catalog_position,omitempty"`
	// Description Application description
	Description string `json:"description,omitempty"`
	// Genre Genre name
	Genre string `json:"genre,omitempty"`
	// GenreID Genre ID
	GenreID int64 `json:"genre_id,omitempty"`
	// Icon139 URL of the app icon with 139 px in width
	Icon139 string `json:"icon_139,omitempty"`
	// Icon150 URL of the app icon with 150 px in width
	Icon150 string `json:"icon_150,omitempty"`
	// Icon278 URL of the app icon with 279 px in width
	Icon278 string `json:"icon_278,omitempty"`
	// Icon75 URL of the app icon with 75 px in width
	Icon75 string `json:"icon_75,omitempty"`
	// ID Application ID
	ID int64 `json:"id"`
	// International Information whether the application is multilanguage
	International int64 `json:"international,omitempty"`
	// IsInCatalog Information whether application is in mobile catalog
	IsInCatalog int64 `json:"is_in_catalog,omitempty"`
	// LeaderboardType
	LeaderboardType *AppsAppLeaderboardType `json:"leaderboard_type,omitempty"`
	// MembersCount Members number
	MembersCount int64 `json:"members_count,omitempty"`
	// PlatformID Application ID in store
	PlatformID int64 `json:"platform_id,omitempty"`
	// PublishedDate Date when the application has been published in Unixtime
	PublishedDate int64 `json:"published_date,omitempty"`
	// ScreenName Screen name
	ScreenName string `json:"screen_name,omitempty"`
	// Screenshots
	Screenshots []*PhotosPhoto `json:"screenshots,omitempty"`
	// Section Application section name
	Section string `json:"section,omitempty"`
	// Title Application title
	Title string `json:"title"`
	// Type
	Type *AppsAppType `json:"type"`
}

// GroupsGroupXtrInvitedByAdminLevel type from VK API Schema(groups_group_xtr_invited_by_admin_level). Level of current user&#39;s credentials as manager
type GroupsGroupXtrInvitedByAdminLevel int64

// PhotosPhotoFull type from VK API Schema(photos_photo_full).
type PhotosPhotoFull struct {

	// AccessKey Access key for the photo
	AccessKey string `json:"access_key,omitempty"`
	// AlbumID Album ID
	AlbumID int64 `json:"album_id"`
	// CanComment Information whether current user can comment the photo
	CanComment *BaseBoolInt `json:"can_comment,omitempty"`
	// Comments
	Comments *BaseObjectCount `json:"comments,omitempty"`
	// Date Date when uploaded
	Date int64 `json:"date"`
	// Height Original photo height
	Height int64 `json:"height,omitempty"`
	// ID Photo ID
	ID int64 `json:"id"`
	// Images
	Images []*PhotosImage `json:"images,omitempty"`
	// Lat Latitude
	Lat float64 `json:"lat,omitempty"`
	// Likes
	Likes *BaseLikes `json:"likes,omitempty"`
	// Long Longitude
	Long float64 `json:"long,omitempty"`
	// OwnerID Photo owner&#39;s ID
	OwnerID int64 `json:"owner_id"`
	// PostID Post ID
	PostID int64 `json:"post_id,omitempty"`
	// Reposts
	Reposts *BaseObjectCount `json:"reposts,omitempty"`
	// Tags
	Tags *BaseObjectCount `json:"tags,omitempty"`
	// Text Photo caption
	Text string `json:"text,omitempty"`
	// UserID ID of the user who have uploaded the photo
	UserID int64 `json:"user_id,omitempty"`
	// Width Original photo width
	Width int64 `json:"width,omitempty"`
}

// AdsRules type from VK API Schema(ads_rules).
type AdsRules struct {

	// Paragraphs
	Paragraphs []*AdsParagraphs `json:"paragraphs,omitempty"`
	// Title Comment
	Title string `json:"title,omitempty"`
}

// GroupsCallbackSettings type from VK API Schema(groups_callback_settings).
type GroupsCallbackSettings struct {

	// APIVersion API version used for the events
	APIVersion string `json:"api_version,omitempty"`
	// Events
	Events *GroupsLongPollEvents `json:"events,omitempty"`
}

// LeadsStart type from VK API Schema(leads_start).
type LeadsStart struct {

	// TestMode Information whether test mode is enabled
	TestMode *BaseBoolInt `json:"test_mode,omitempty"`
	// VkSid Session data
	VkSid string `json:"vk_sid,omitempty"`
}

// NewsfeedListFull type from VK API Schema(newsfeed_list_full).
type NewsfeedListFull struct {
}

// PagesWikipage type from VK API Schema(pages_wikipage).
type PagesWikipage struct {

	// CreatorID Page creator ID
	CreatorID int64 `json:"creator_id,omitempty"`
	// CreatorName Page creator name
	CreatorName int64 `json:"creator_name,omitempty"`
	// EditorID Last editor ID
	EditorID int64 `json:"editor_id,omitempty"`
	// EditorName Last editor name
	EditorName string `json:"editor_name,omitempty"`
	// GroupID Community ID
	GroupID int64 `json:"group_id"`
	// ID Page ID
	ID int64 `json:"id"`
	// Title Page title
	Title string `json:"title"`
	// Views Views number
	Views int64 `json:"views"`
	// WhoCanEdit Edit settings of the page
	WhoCanEdit *PagesPrivacySettings `json:"who_can_edit"`
	// WhoCanView View settings of the page
	WhoCanView *PagesPrivacySettings `json:"who_can_view"`
}

// WallWallpost type from VK API Schema(wall_wallpost).
type WallWallpost struct {

	// AccessKey Access key to private object
	AccessKey string `json:"access_key,omitempty"`
	// Attachments
	Attachments []*WallWallpostAttachment `json:"attachments,omitempty"`
	// Date Date of publishing in Unixtime
	Date int64 `json:"date,omitempty"`
	// FromID Post author ID
	FromID int64 `json:"from_id,omitempty"`
	// Geo
	Geo *BaseGeo `json:"geo,omitempty"`
	// ID Post ID
	ID int64 `json:"id,omitempty"`
	// OwnerID Wall owner&#39;s ID
	OwnerID int64 `json:"owner_id,omitempty"`
	// PostSource
	PostSource *WallPostSource `json:"post_source,omitempty"`
	// PostType
	PostType *WallPostType `json:"post_type,omitempty"`
	// SignerID Post signer ID
	SignerID int64 `json:"signer_id,omitempty"`
	// Text Post text
	Text string `json:"text,omitempty"`
	// Views Count of views
	Views *WallViews `json:"views,omitempty"`
}

// BaseBoolInt type from VK API Schema(base_bool_int).
type BaseBoolInt int64

// BaseLikesInfo type from VK API Schema(base_likes_info).
type BaseLikesInfo struct {

	// CanLike Information whether current user can like the post
	CanLike *BaseBoolInt `json:"can_like"`
	// CanPublish Information whether current user can repost
	CanPublish *BaseBoolInt `json:"can_publish,omitempty"`
	// Count Likes number
	Count int64 `json:"count"`
	// UserLikes Information whether current uer has liked the post
	UserLikes int64 `json:"user_likes"`
}

// OrdersAmountItem type from VK API Schema(orders_amount_item).
type OrdersAmountItem struct {

	// Amount Votes amount in user&#39;s currency
	Amount int64 `json:"amount,omitempty"`
	// Description Amount description
	Description string `json:"description,omitempty"`
	// Votes Votes number
	Votes string `json:"votes,omitempty"`
}

// StoriesReplies type from VK API Schema(stories_replies).
type StoriesReplies struct {
	// Count Replies number.
	Count int64 `json:"count"`
	// New New replies number.
	New int64 `json:"new,omitempty"`
}

// VideoUploadResponse type from VK API Schema(video_upload_response).
type VideoUploadResponse struct {

	// Size Video size
	Size int64 `json:"size,omitempty"`
	// VideoID Video ID
	VideoID int64 `json:"video_id,omitempty"`
}

// BaseOkResponse type from VK API Schema(base_ok_response). Returns 1 if request has been processed successfully
type BaseOkResponse int64

// NotificationsNotification type from VK API Schema(notifications_notification).
type NotificationsNotification struct {

	// Date Date when the event has been occured
	Date int64 `json:"date,omitempty"`
	// Feedback
	Feedback *NotificationsFeedback `json:"feedback,omitempty"`
	// Parent
	Parent *NotificationsNotificationParent `json:"parent,omitempty"`
	// Reply
	Reply *NotificationsReply `json:"reply,omitempty"`
	// Type Notification type
	Type string `json:"type,omitempty"`
}

// StatsPeriod type from VK API Schema(stats_period).
type StatsPeriod struct {

	// Activity
	Activity *StatsActivity `json:"activity,omitempty"`
	// PeriodFrom Day (YYYY-MM-DD)
	PeriodFrom string `json:"period_from,omitempty"`
	// PeriodTo Day (YYYY-MM-DD)
	PeriodTo string `json:"period_to,omitempty"`
	// Reach
	Reach *StatsReach `json:"reach,omitempty"`
	// Visitors
	Visitors *StatsViews `json:"visitors,omitempty"`
}

// AdsStats type from VK API Schema(ads_stats).
type AdsStats struct {

	// ID Object ID
	ID int64 `json:"id,omitempty"`
	// Stats
	Stats *AdsStatsFormat `json:"stats,omitempty"`
	// Type
	Type *AdsObjectType `json:"type,omitempty"`
}

// AdsTargSettings type from VK API Schema(ads_targ_settings).
type AdsTargSettings struct {
}

// AdsTargetGroup type from VK API Schema(ads_target_group).
type AdsTargetGroup struct {

	// AudienceCount Audience
	AudienceCount int64 `json:"audience_count,omitempty"`
	// Domain Site domain
	Domain string `json:"domain,omitempty"`
	// ID Group ID
	ID int64 `json:"id,omitempty"`
	// Lifetime Number of days for user to be in group
	Lifetime int64 `json:"lifetime,omitempty"`
	// Name Group name
	Name string `json:"name,omitempty"`
	// Pixel Pixel code
	Pixel string `json:"pixel,omitempty"`
}

// BaseCommentsInfo type from VK API Schema(base_comments_info).
type BaseCommentsInfo struct {

	// CanPost Information whether current user can comment the post
	CanPost *BaseBoolInt `json:"can_post,omitempty"`
	// Count Comments number
	Count int64 `json:"count,omitempty"`
	// GroupsCanPost Information whether groups can comment the post
	GroupsCanPost *BaseBoolInt `json:"groups_can_post,omitempty"`
}

// BaseLinkApplication type from VK API Schema(base_link_application).
type BaseLinkApplication struct {

	// AppID Application Id
	AppID float64 `json:"app_id,omitempty"`
	// Store
	Store *BaseLinkApplicationStore `json:"store,omitempty"`
}
