package controllers

type LogOutController struct {
	baseController
}

func (c *LogOutController) Get() {
	c.DelSession("token")
}
