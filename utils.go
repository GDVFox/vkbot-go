package vkbotgo

import (
	"encoding/json"
	"io"
)

// NewConfirmation creates new conformation struct
//
// pattern contains endpoint for server events
// groupID contains group_id(to compare with a key from a query)
// confStr contains string that the server must return for confirmation
func NewConfirmation(pattern string, groupID int64, confStr string) *Confirmation {
	return &Confirmation{
		Pattern:       pattern,
		GroupID:       groupID,
		ResponseBytes: []byte(confStr),
	}
}

func decodeVkResponse(body io.Reader) (*VkResponse, error) {
	vkResp := &VkResponse{}

	d := json.NewDecoder(body)
	err := d.Decode(vkResp)
	return vkResp, err
}
