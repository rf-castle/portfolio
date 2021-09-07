package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/rf-castle/portfolio/back/routers"
)

func main() {
	beego.Run()
}
