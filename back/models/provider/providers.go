package provider

var AllProvider map[string]Provider

func init() {
	providers := []Provider{DiscordProvider{}}
	AllProvider = make(map[string]Provider, len(providers))
	for _, provider := range providers {
		AllProvider[provider.Name()] = provider
	}
}
