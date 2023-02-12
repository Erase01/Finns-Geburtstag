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

var _index = 0

const ERR_VALUE = -1

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func main() {
	//fmt.Println(get_hash("udidit"))

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
	r.GET("/kopf", kopf)

	r.GET("/submit_flag", submit_flag)

	r.POST("/start_rdr2gusser", start_rdr2gusser)

	r.Run(":8080")
}

func index(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.HTML(http.StatusOK, "index.html", gin.H{"rdr2": "Nein", "quiz": "Nein", "kopf": "Nein"})
}

func kopf(c *gin.Context) {
	if c.Request.Header["Accept-Language"][0][:5] == "en-US" {
		fmt.Println("du musst deutsch sein")
		c.HTML(http.StatusPermanentRedirect, "index.html", gin.H{})
	} else if c.Request.Header["Accept-Language"][0][:5] == "de-DE" {
		c.HTML(http.StatusOK, "kopf.html", gin.H{})
	}

	c.HTML(http.StatusOK, "burb.html", gin.H{})
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
	switch hash_flag {
	case "73c46df076e245e59cfe4e3d362b0c2c":
		session.Set("quiz", "Ja")
		c.HTML(http.StatusPermanentRedirect, "index.html", gin.H{"quiz": "Ja!"})
	case "dummy":
		session.Set("rdr2", "Ja")
	case "cac2549b310b664bf3143d888bcf74bc":
		session.Set("gusser", "Ja")
	case "ee4dea37bfe4fb2d92d5b8186e762285":
		session.Set("burp", "Ja")
	default:
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
			c.HTML(http.StatusPermanentRedirect, "rdr2.html", gin.H{
				"message": "Falsch! Du bist zu weit weg!",
			})
			return
		}
		// convert to percentage of 25 and return it
		perc_d := d / 25 * 100

		session := sessions.Default(c)

		img_num := session.Get("img_num")
		if img_num == nil {

			c.HTML(http.StatusPermanentRedirect, "rdr2.html", gin.H{
				"message": "du musst erst das spiel starten",
			})
			return
		}
		session.Set("img_num", img_num.(int)+1)

		session.Set("img_num_perc", perc_d)

		session.Save()

		c.HTML(http.StatusPermanentRedirect, "rdr2gusser.html", gin.H{
			"distance_percentage": perc_d,
			"img_num":             img_num.(int) + 1,
		})
	}
}

// coordinates of click and no. of image to check
func check_image_click(x int, y int, index int) float64 {
	var d float64           // distance
	var x2, y2 int          // x and y of the image the user should have clicked close to
	const MAX_DISTANCE = 50 // max distance between click and wanted pos

	switch i := _index; i {
	case 0:
		x2, y2 = 1088, 656
	case 1:
		x2, y2 = 1030, 385
	case 2:
		x2, y2 = 1060, 370
	case 3:
		x2, y2 = 1080, 325
	case 4:
		x2, y2 = 1100, 344
	case 5:
		x2, y2 = 1162, 333
	case 6:
		x2, y2 = 1167, 338
	case 7:
		x2, y2 = 1188, 269
	case 8:
		x2, y2 = 1211, 267
	case 9:
		x2, y2 = 1223, 377
	case 10:
		x2, y2 = 1141, 264
	default:
		return ERR_VALUE
	}

	d = get_distance(x, x2, y, y2)
	fmt.Println(d)
	if d > MAX_DISTANCE {
		return ERR_VALUE
	}
	_index = _index + 1
	return d
}

// len between 2 points: https://youtu.be/CWUr6Jo6tag
func get_distance(x1 int, x2 int, y1 int, y2 int) float64 {
	d := math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
	return d
}

func start_rdr2gusser(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("img_num", 0)
	session.Save()
	c.HTML(http.StatusPermanentRedirect, "rdr2gusser.html", gin.H{
		"img_num": 0,
	})
}
