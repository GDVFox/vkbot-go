package vkbotgo

const (
	// VkAPIurl is the url template for all API methods(for Sprintf)
	VkAPIurl = "https://api.vk.com/method/%s?access_token=%s&v=%s"

	//ConfirmationType value of "type" field of JSON struct with confirmation from VK server
	ConfirmationType = "confirmation"

	//OkString a string that is sent in response to each request from the Callback APU
	OkString = "ok"
)
