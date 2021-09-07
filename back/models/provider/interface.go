package provider

import (
	"context"
	"github.com/rf-castle/portfolio/back/models"
	"golang.org/x/oauth2"
	"net/http"
)

type Provider interface {
	Conf() *oauth2.Config
	Name() string
	GetUser(*oauth2.Token) (*models.User, error)
	CanShowSecret(*oauth2.Token) (bool, error)
}

func getClient(prov Provider, token *oauth2.Token, _ctx ...context.Context) *http.Client {
	var ctx context.Context
	if len(_ctx) == 0 {
		ctx = context.Background()
	} else {
		ctx = _ctx[0]
	}
	return prov.Conf().Client(ctx, token)
}
