package models

import "golang.org/x/oauth2"

type Token struct {
	Token    *oauth2.Token
	Provider string
}
