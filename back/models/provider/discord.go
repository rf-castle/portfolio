package provider

import (
	"github.com/bwmarrin/discordgo"
	conf2 "github.com/rf-castle/portfolio/back/conf"
	"github.com/rf-castle/portfolio/back/models"
	"golang.org/x/oauth2"
)

type DiscordProvider struct{}

func _() Provider {
	return &DiscordProvider{}
}

func (d DiscordProvider) Conf() *oauth2.Config {
	return &oauth2.Config{
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
		Scopes:       []string{"identify", "guilds"},
		RedirectURL:  conf2.RedirectUri + "/" + d.Name(),
		ClientID:     conf2.ClientId,
		ClientSecret: conf2.ClientSecret,
	}
}

func (d DiscordProvider) Name() string {
	return "discord"
}

func (d DiscordProvider) getSession(token *oauth2.Token) *discordgo.Session {
	client := getClient(d, token)
	return &discordgo.Session{
		Ratelimiter: discordgo.NewRatelimiter(),
		Client:      client,
	}
}

func (d DiscordProvider) GetUser(token *oauth2.Token) (*models.User, error) {
	session := d.getSession(token)
	dUser, err := session.User("@me")
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Name:  dUser.Username,
		Image: dUser.AvatarURL(""),
	}
	return user, nil
}

func (d DiscordProvider) CanShowSecret(token *oauth2.Token) (bool, error) {
	session := d.getSession(token)
	guilds, err := session.UserGuilds(200, "", "")
	if err != nil {
		return false, err
	}
	for _, guild := range guilds {
		if guild.ID == "0" {
			return true, nil
		}
	}
	return false, nil
}
