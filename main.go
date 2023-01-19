package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Printf("%x\n", string((get_hash([]byte("test")))))
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
	c.HTML(http.StatusPermanentRedirect, "index.html", gin.H{"msg": "flag richtig!"})
}

func download(c *gin.Context) {
	c.File("files/out/a")
}