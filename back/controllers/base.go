package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type baseController struct {
	beego.Controller
	body map[string]interface{}
}

func (c *baseController) methodNotAllowed() {
	c.body["error"] = "Method Not Allowed"
	c.Ctx.Output.SetStatus(http.StatusMethodNotAllowed)
}

// Get adds a request function to handle GET request.
func (c *baseController) Get() {
	c.methodNotAllowed()
}

// Post adds a request function to handle POST request.
func (c *baseController) Post() {
	c.methodNotAllowed()
}

// Delete adds a request function to handle DELETE request.
func (c *baseController) Delete() {
	c.methodNotAllowed()
}

// Put adds a request function to handle PUT request.
func (c *baseController) Put() {
	c.methodNotAllowed()
}

// Head adds a request function to handle HEAD request.
func (c *baseController) Head() {
	c.methodNotAllowed()
}

// Patch adds a request function to handle PATCH request.
func (c *baseController) Patch() {
	c.methodNotAllowed()
}

// Options adds a request function to handle OPTIONS request.
func (c *baseController) Options() {
	c.methodNotAllowed()
}

func (c *baseController) Prepare() {
	c.body = map[string]interface{}{
		"error": nil,
	}
	c.Data["json"] = c.body
}

func (c *baseController) Render() error {
	if !c.EnableRender {
		return nil
	}
	return c.ServeJSON(true)
}
