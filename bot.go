package vkbotgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// VkBot interacts with the VK API.
type VkBot struct {
	AccessToken string
	Buffer      int
	Version     string

	Client *http.Client
	Conf   *Confirmation
}

// NewVkBot creates a new VkBot
// and you can pass a http.Client.
func NewVkBot(accessToken, ver string, client *http.Client) (*VkBot, error) {
	vkBot := &VkBot{
		AccessToken: accessToken,
		Buffer:      128,
		Version:     ver,
		Client:      client,
		Conf:        nil,
	}

	return vkBot, nil
}

// Request makes a request to a method with given params.
func (vkBot *VkBot) Request(method string, params url.Values) (*VkResponse, error) {
	requestURL := fmt.Sprintf(VkAPIurl, method, vkBot.AccessToken, vkBot.Version)

	//Always use only POST request;
	response, err := vkBot.Client.PostForm(requestURL, params)
	if err != nil {
		return &VkResponse{}, err
	}
	defer response.Body.Close()

	vkResp, err := decodeVkResponse(response.Body)
	if err != nil {
		return &VkResponse{}, err
	}

	return vkResp, nil
}

// SetConfirmation sets a conformation response
//
// if conformation is not set, you can not use Callback API.
func (vkBot *VkBot) SetConfirmation(conf *Confirmation) {
	vkBot.Conf = conf
}

// ListenForEvents registers a http handler for a events.
func (vkBot *VkBot) ListenForEvents() EventsChannel {
	events := make(chan *Event, vkBot.Buffer)

	http.HandleFunc(vkBot.Conf.Pattern, func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		event := &Event{}
		json.Unmarshal(b, event)

		if event.Type == ConfirmationType && event.GroupID == vkBot.Conf.GroupID {
			_, err := w.Write(vkBot.Conf.ResponseBytes)
			if err != nil {
				log.Fatalf("can not send conformation string. error: %s", err.Error())
			}
			return
		}

		w.Write([]byte(OkString)) //Make byte constant
		events <- event
	})

	return events
}
