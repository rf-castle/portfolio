package controllers

import (
	"github.com/rf-castle/portfolio/back/models"
	"github.com/rf-castle/portfolio/back/models/provider"
	"net/http"
)

type LoginController struct {
	baseController
	provider string
}

func (c *LoginController) Prepare() {
	c.baseController.Prepare()
	c.provider = c.Ctx.Input.Param(":provider")
}

func (c *LoginController) Get() {
	token := c.GetSession("token")
	if token != nil {
		// Todo: You are already logged in
		return
	}
	prov, ok := provider.AllProvider[c.provider]
	if !ok {
		return
	}
	conf := prov.Conf()
	_state := models.MakeRandomStr(32)
	if _state == nil {
		// Todo: MakeRandomStr Error
		return
	}
	state := *_state
	_ = c.SetSession("login", state)
	c.Redirect(conf.AuthCodeURL(state), http.StatusTemporaryRedirect)
}
