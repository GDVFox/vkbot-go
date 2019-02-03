package vkbotgo

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
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

func parseLoginForm(resp *http.Response) (string, *url.Values, error) {
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return "", nil, err
	}

	postVals := &url.Values{}
	loginForm := doc.Find("form").First()
	actionAddr, _ := loginForm.Attr("action")

	loginForm.Find("input").Each(func(_ int, s *goquery.Selection) {
		if attr, _ := s.Attr("class"); attr != "button" {
			inputName, _ := s.Attr("name")
			inputValue, _ := s.Attr("value")
			postVals.Add(inputName, inputValue)
		}
	})

	return actionAddr, postVals, nil
}
