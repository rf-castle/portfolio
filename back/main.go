package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"io/fs"
	"path/filepath"
	"strings"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	_ = filepath.Walk("./static", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		routePath := strings.TrimPrefix(path, "static")
		if filepath.Ext(routePath) != ".html"{
			r.StaticFile(routePath, path)
			return nil
		}
		routePath = strings.TrimSuffix(routePath, ".html")
		routePath = strings.TrimSuffix(routePath, "index")
		r.StaticFile(routePath, path)
		return nil
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}