package vkbotgo

import "encoding/json"

// APIParameters describes API method parameters with validate method
type APIParameters interface {
	Validate() error
}

//Event represents an event coming from the VK server.
type Event struct {
	Type    string          `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int64           `json:"group_id"`
}

//EventsChannel is the channel for getting events.
type EventsChannel chan *Event

// VkResponse is a raw response from the VK API.
type VkResponse struct {
	Response json.RawMessage `json:"response"`
}

//Confirmation contains group and string for Vk conformation request.
type Confirmation struct {
	Pattern       string
	GroupID       int64
	ResponseBytes []byte
}
