package vkbotgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/google/go-querystring/query"
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

// NewVkBotWithAuth creates a new VkBot with login/pass authorization
// and you can pass a http.Client.
func NewVkBotWithAuth(login, pass, scope, ver string, clientID int64, client *http.Client) (*VkBot, error) {
	// We need to remember cookies
	tmpClient := *client
	tmpClient.Jar, _ = cookiejar.New(nil)

	var resp *http.Response

	// User sends login init request
	requestURL := fmt.Sprintf(AuthAPIurl, clientID, scope, ver)
	resp, err := tmpClient.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// User eneters email and pass
	actionAddr, postVals, err := parseLoginForm(resp)
	if err != nil {
		return nil, err
	}

	postVals.Set("email", login)
	postVals.Set("pass", pass)

	resp, err = tmpClient.PostForm(actionAddr, *postVals)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.Request.URL.Path != "/blank.html" {
		queryVals := resp.Request.URL.Query()
		switch {
		case queryVals.Get("act") == "authcheck":
			// User uses Two-factor authentication
			// Enter confirmation code for VK message
			var confirmString string
			_, err := fmt.Scanf("%s", &confirmString)
			if err != nil {
				return nil, err
			}

			confirmAction, confirmValues, err := parseLoginForm(resp)
			if err != nil {
				return nil, err
			}

			confirmValues.Set("code", confirmString)

			resp, err = tmpClient.PostForm(resp.Request.URL.Scheme+"://"+resp.Request.URL.Host+confirmAction, *confirmValues)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			// Still needs to press confirm
			fallthrough
		default:
			// User press "Confirm" button
			grantAction, grantValues, err := parseLoginForm(resp)
			resp, err = tmpClient.PostForm(grantAction, *grantValues)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
		}

		if resp.Request.URL.Path != "/blank.html" {
			return nil, errors.New("login error")
		}

	}

	respArgs, err := url.ParseQuery(resp.Request.URL.Fragment)
	if err != nil {
		return nil, err
	}

	return NewVkBot(respArgs.Get("access_token"), ver, client)
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

// CallMethod calls VK API method with given params(with validation).
//
// For params: CamelCase method name + Params.
// Example: MessagesSendParams (for messages.send)
func (vkBot *VkBot) CallMethod(method string, params APIParameters) (*VkResponse, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	vals, err := query.Values(params)

	if err != nil {
		return nil, err
	}

	return vkBot.Request(method, vals)
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
