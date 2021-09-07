package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/rf-castle/portfolio/back/controllers"
	"io/fs"
	"path/filepath"
	"strings"
)

func init() {
	_ = filepath.Walk("./static", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		routePath := strings.TrimPrefix(path, "static")
		if filepath.Ext(routePath) != ".html" {
			beego.SetStaticPath(routePath, path)
			return nil
		}
		routePath = strings.TrimSuffix(routePath, ".html")
		routePath = strings.TrimSuffix(routePath, "index")
		beego.SetStaticPath(routePath, path)
		return nil
	})
	beego.Router("/api/general", &controllers.GeneralController{})
	beego.Router("/api/login/:provider", &controllers.LoginController{})
	beego.Router("/api/auth/:provider", &controllers.AuthController{})
}
