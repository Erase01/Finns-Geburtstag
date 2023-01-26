package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func check(err error) {
	if err != nil {
		fmt.Println("Error!!!")
		os.Exit(1)
	}
}

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("sites/*.html")

	r.GET("/", index)
	r.GET("/quiz", quiz)
	r.GET("/rdr2", rdr2)
	r.GET("/rdr2gusser", rdr2gusser)
	r.GET("/download", download)
	r.GET("/image_click", image_click)

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
	c.HTML(http.StatusOK, "rdr2gusser.html", gin.H{
				"img_num": 0,
	})
}

func submit_flag(c *gin.Context) {
	flag := c.Query("flag")
	hash_flag := get_hash(flag)
	// 73c46df076e245e59cfe4e3d362b0c2c is the hash of the `strings` command flag
	session := sessions.Default(c)
	if hash_flag == "73c46df076e245e59cfe4e3d362b0c2c" {
		session.Set("quiz", "Ja")
		c.HTML(http.StatusPermanentRedirect, "index.html", gin.H{"quiz": "Ja!"})
	} else if hash_flag == "dummy" {
		session.Set("rdr2", "Ja")
	} else {
		fmt.Println("nein")
		c.HTML(http.StatusPermanentRedirect, "index.html", gin.H{})
	}
}

func download(c *gin.Context) {
	c.File("files/out/a")
}

func image_click(c *gin.Context) {
	// get query of request
	q := c.Request.URL.Query()
	// key is of type string
	for key := range q {
		xy := strings.Split(key, ",")  // split string into x and y
		x, xerr := strconv.Atoi(xy[0]) // get x as int
		check(xerr)                    // is x not int?
		y, yerr := strconv.Atoi(xy[1]) // get y as int
		check(yerr)                    // is y not int?

		// is the clicked pos close enough?
		d := check_image_click(x, y, 0)
		if d == -1 {
			return
		}
		// convert to percentage of 25 and return it
		perc_d := d / 25 * 100

		c.HTML(http.StatusPermanentRedirect, "rdr2gusser.html", gin.H{
			"distance_percentage": perc_d,
			"img_num":             1,
		})
	}
}

// coordinates of click and no. of image to check, if -1 the user is too far away
func check_image_click(x int, y int, index int) float64 {
	var d float64
	// len between 2 points: https://youtu.be/CWUr6Jo6tag
	switch i := index; i {
	case 0:
		x2, y2 := 100, 100

		xs, ys := math.Pow(float64(x2-x), 2), math.Pow(float64(y2-y), 2)
		d = math.Sqrt(xs + ys)
	}

	if d > 25 {
		return -1
	}
	return d
}
