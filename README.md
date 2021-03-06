# VkBot-Go

[![GoDoc](https://godoc.org/github.com/GDVFox/vkbot-go?status.svg)](https://godoc.org/github.com/GDVFox/vkbot-go)
[![Travis](https://travis-ci.org/GDVFox/vkbot-go.svg)](https://travis-ci.org/GDVFox/vkbot-go)
[![rcard](https://goreportcard.com/badge/github.com/GDVFox/vkbot-go)](https://goreportcard.com/report/github.com/GDVFox/vkbot-go)

This package provides you bindings for the VK API

## Install & Import

### Install

```bash
go get -u github.com/GDVFox/vkbot-go
```

### Import

```go
import "github.com/GDVFox/vkbot-go"
```

## Example

### Simple request to VK API

Bot gets information about user id304508916

```go
package main

import (
    "log"
    "net/http"
    "net/url"

    vkAPI "github.com/GDVFox/vkbot-go"
)

func main() {
    bot, _ := vkAPI.NewVkBot("AccessToken", "5.92", &http.Client{})

    resp, err := bot.Request("users.get", url.Values{
        "user_ids": []string{"304508916"},
        "fields":   []string{"status"},
    })
    if err != nil {
        log.Panic(err)
    }

    log.Println(string(resp.Response))
}
```

### Simple use of Callback API

Bot gets events of group id1337

```go
package main

import (
    "log"
    "net/http"

    vkAPI "github.com/GDVFox/vkbot-go"
)

func main() {
    bot, _ := vkAPI.NewVkBot("AccessToken", "5.92", &http.Client{})

    bot.SetConfirmation(vkAPI.NewConfirmation("/", 1337, "ConfirmString"))
    events := bot.ListenForEvents()

    go http.ListenAndServe(":8080", nil)
    log.Println("start listen :8080")

    for event := range events {
        log.Printf("[%s] %s from group: %d", event.Type, string(event.Object), event.GroupID)
    }
}
```
