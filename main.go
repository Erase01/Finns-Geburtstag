package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("sites/*.html")

	r.GET("/", index)
	r.GET("/quiz", quiz)
	r.GET("/rdr2", rdr2)

	r.Run()
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"text": "moin",
	})
}

func quiz(c *gin.Context) {
	c.HTML(http.StatusOK, "quiz.html", gin.H{})
}

func rdr2(c *gin.Context) {
	c.HTML(http.StatusOK, "rdr2.html", gin.H{})
}