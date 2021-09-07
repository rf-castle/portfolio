package controllers

import (
	"context"
	"github.com/rf-castle/portfolio/back/models"
	"github.com/rf-castle/portfolio/back/models/provider"
	"net/http"
)

type AuthController struct {
	baseController
	provider string
}

func (c *AuthController) Prepare() {
	c.baseController.Prepare()
	c.provider = c.Ctx.Input.Param(":provider")
}

func (c *AuthController) Get() {
	prov, ok := provider.AllProvider[c.provider]
	if !ok {
		return
	}
	conf := prov.Conf()
	state := c.GetString("state")
	v := c.GetSession("login")
	if state == "" || v == nil || state != v.(string) {
		c.body["error"] = "State is wrong! it might be csrf attack."
		c.body["token"] = struct{}{}
		c.Ctx.Output.SetStatus(http.StatusForbidden)
	} else {
		_ = c.DelSession("login")
		code := c.GetString("code")
		tok, err := conf.Exchange(context.Background(), code)
		if err != nil {
			// Todo: Error Handler
		}
		c.SetSession("token", &models.Token{
			Token:    tok,
			Provider: c.provider,
		})
	}
	_ = c.ServeJSON(true)

}
