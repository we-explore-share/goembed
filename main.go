package main

import (
	"io/fs"
	"net/http"
	"newtest/frontend"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := fs.Sub(frontend.Fs, "build")
	fss := http.FS(f)

	r := gin.Default()
	r.StaticFS("frontend", fss)

	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/frontend") {
			c.FileFromFS("/", fss) //try...file
		} else {
			c.Redirect(http.StatusMovedPermanently, "/frontend")
		}
	})

	svr := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	err := svr.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
