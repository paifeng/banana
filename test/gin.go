package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OojStudent struct {
	Id     int      `form:"id"`
	Name   string   `form:"name"`
	Age    int      `form:"age"`
	Gender string   `form:"gender"`
	Hobby  []string `form:"hobby"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	router.GET("/1", func(c *gin.Context) {
		t := &OojStudent{}
		err := c.Bind(t)
		if err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}
		c.JSON(http.StatusOK, t)
	})
	router.Run(":7777")
}
