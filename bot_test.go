package vkbotgo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

const (
	TestAccessToken = "e381389be7a9110bc1aa638c267380ef5229bc068d64b6bba1d10690f7d690431277e43ab6b237c378f29"
	TestVersion     = "5.92"
)

func TestRequest(t *testing.T) {
	expected := `[{"id":210700286,"first_name":"Lindsey","last_name":"Stirling","is_closed":false,"can_access_closed":true}]`

	bot, _ := NewVkBot(TestAccessToken, TestVersion, &http.Client{})
	resp, err := bot.Request("users.get", url.Values{
		"user_ids": []string{"210700286"},
	})

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if string(resp.Response) != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", string(resp.Response), expected)
		t.Fail()
	}
}

func TestListenForEventsConfirmation(t *testing.T) {
	http.DefaultServeMux = &http.ServeMux{} //ListenForEvents sets HandleFunc on DefaultServeMux
	bot, _ := NewVkBot(TestAccessToken, TestVersion, &http.Client{})

	bot.SetConfirmation(NewConfirmation("/test_endpoint", 177502772, "13f19c78"))
	bot.ListenForEvents()

	request := httptest.NewRequest("GET", bot.Conf.Pattern, strings.NewReader(`{ "type": "confirmation", "group_id": 177502772 }`))
	recorder := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Fail()
	}

	expected := "13f19c78"
	if recorder.Body.String() != expected {
		t.Errorf("results not match\nGot: %v\nExpected: %v", string(recorder.Body.String()), expected)
		t.Fail()
	}
}

type TestEventCase struct {
	EventString string
	Evnt        *Event
}

func TestListenForEventsChannel(t *testing.T) {
	cases := []TestEventCase{
		{
			EventString: `{"type": "group_join", "object": {"user_id": 1, "join_type" : "approved"}, "group_id": 1}`,
			Evnt: &Event{
				Type:    "group_join",
				Object:  json.RawMessage(`{"user_id": 1, "join_type" : "approved"}`),
				GroupID: 1,
			},
		},
		{
			EventString: `{"type": "group_leave", "object": {"user_id": 1, "self" : 1}, "group_id": 1}`,
			Evnt: &Event{
				Type:    "group_leave",
				Object:  json.RawMessage(`{"user_id": 1, "self" : 1}`),
				GroupID: 1,
			},
		},
	}

	http.DefaultServeMux = &http.ServeMux{} //ListenForEvents sets HandleFunc on DefaultServeMux
	bot, _ := NewVkBot(TestAccessToken, TestVersion, &http.Client{})

	bot.SetConfirmation(NewConfirmation("/test_endpoint", 177502772, "13f19c78"))
	events := bot.ListenForEvents()

	go func() {
		for _, newCase := range cases {
			request := httptest.NewRequest("GET", bot.Conf.Pattern, strings.NewReader(newCase.EventString))
			recorder := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(recorder, request)
			if status := recorder.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
				t.Fail()
			}

			if recorder.Body.String() != OkString {
				t.Errorf("results not match\nGot: %v\nExpected: %v", string(recorder.Body.String()), OkString)
				t.Fail()
			}
		}

		close(events)
	}()

	eventCounter := 0
	for event := range events {
		if !reflect.DeepEqual(event, cases[eventCounter].Evnt) {
			t.Errorf("results not match\nGot: %v\nExpected: %v", event, cases[eventCounter].Evnt)
			t.Fail()
		}

		eventCounter++
	}
}
