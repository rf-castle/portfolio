package conf

import (
	beego "github.com/beego/beego/v2/server/web"
	"log"
)

var (
	ClientSecret string
	ClientId     string
	RedirectUri  string
)

func init() {
	var err error
	ClientSecret, err = beego.AppConfig.String("discord_client_secret")
	if err != nil {
		log.Fatal(err)
	}
	if ClientSecret == "" {
		log.Fatal("discord_client_secret is empty! please set discord_client_secret.")
	}
	ClientId, err = beego.AppConfig.String("discord_client_id")
	if err != nil {
		log.Fatal(err)
	}
	if ClientId == "" {
		log.Fatal("discord_client_id is empty! please set discord_client_id.")
	}
	RedirectUri, err = beego.AppConfig.String("redirect_uri")
	if err != nil {
		log.Fatal(err)
	}
	if RedirectUri == "" {
		log.Fatal("redirect_uri is empty! please set redirect_uri")
	}
}
