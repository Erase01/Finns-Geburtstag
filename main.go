package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("sites/*.html")

	r.GET("/", index)
	r.GET("/quiz", quiz)
	r.GET("/rdr2", rdr2)
	r.GET("/rdr2gusser", rdr2gusser)
	r.GET("/download", download)

	r.GET("/submit_flag", submit_flag)

	r.Run()
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"rdr2": "Nein", "quiz": "Nein"})
}

func quiz(c *gin.Context) {
	c.HTML(http.StatusOK, "quiz.html", gin.H{})
}

func rdr2(c *gin.Context) {
	c.HTML(http.StatusOK, "rdr2.html", gin.H{})
}

func rdr2gusser(c *gin.Context) {
	c.HTML(http.StatusOK, "rdr2gusser.html", gin.H{})
}

func submit_flag(c *gin.Context) {
	flag := c.Query("flag")
	hash_flag := get_hash(flag)
	// 73c46df076e245e59cfe4e3d362b0c2c is the hash of the `strings` command flag
	if hash_flag == "73c46df076e245e59cfe4e3d362b0c2c" {
		fmt.Println("ja")
		c.HTML(http.StatusPermanentRedirect, "index.html", gin.H{"quiz": "Ja!"})
	} else {
		fmt.Println("nein")
		c.HTML(http.StatusPermanentRedirect, "index.html", gin.H{})
	}
}

func download(c *gin.Context) {
	c.File("files/out/a")
}
