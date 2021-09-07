package controllers

import (
	"github.com/rf-castle/portfolio/back/models"
	provider2 "github.com/rf-castle/portfolio/back/models/provider"
)

type GeneralController struct {
	baseController
}

func (c *GeneralController) Prepare() {
	c.baseController.Prepare()
	c.body["user"] = nil
	c.body["page"] = []models.Page{}
}

func (c *GeneralController) Get() {
	_token := c.GetSession("token")
	if _token == nil {
		return
	}
	token, ok := _token.(*models.Token)
	if !ok {
		return
	}
	var err error
	_ = err
	prov := provider2.AllProvider[token.Provider]
	// Todo: Error Logging
	c.body["user"], _ = prov.GetUser(token.Token)
	ok, _ = prov.CanShowSecret(token.Token)
	if ok {
		c.body["page"] = models.SecretPage
	}
}
