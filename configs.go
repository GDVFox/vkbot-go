package vkbotgo

const (
	// VkAPIurl is the url template for all API methods(for Sprintf)
	VkAPIurl = "https://api.vk.com/method/%s?access_token=%s&v=%s"

	// AuthAPIurl is the url template for user authorization(for Sprintf)
	AuthAPIurl = "https://oauth.vk.com/authorize?client_id=%d&scope=%s" +
		"&redirect_uri=https://oauth.vk.com/blank.html&display=mobile&v=%s" +
		"&response_type=token"

	// ConfirmationType value of "type" field of
	// JSON struct with confirmation from VK server
	ConfirmationType = "confirmation"

	// OkString a string that is sent in response to each request from the Callback APU
	OkString = "ok"
)
