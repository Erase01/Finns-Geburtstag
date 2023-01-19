package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// 73c46df076e245e59cfe4e3d362b0c2c is the hash of the `strings` command flag
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("sites/*.html")

	r.GET("/", index)
	r.GET("/quiz", quiz)
	r.GET("/rdr2", rdr2)
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

func submit_flag(c *gin.Context) {
	flag := c.Query("flag")
	fmt.Println(flag)
	hash_of_flag := string((get_hash([]byte(flag))))
	fmt.Println(hash_of_flag)
	if flag == "73c46df076e245e59cfe4e3d362b0c2c" {
		c.HTML(http.StatusPermanentRedirect, "index.html", gin.H{"quiz": "Ja!"})
	} else {
		c.HTML(http.StatusPermanentRedirect, "index.html", gin.H{})
	}
}

func download(c *gin.Context) {
	c.File("files/out/a")
}