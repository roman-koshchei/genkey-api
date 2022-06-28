package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	genkey "github.com/roman-koshchei/genkey-api/pkg"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.StaticFile("/output.svg", "./guide/output.svg")
	router.LoadHTMLGlob("guide/*.html")

	router.Use(gin.Recovery()) // if panic return 500

	router.GET("/", getGuide)
	router.GET("/together/", getTogether)
	router.GET("/divided/", getDivided)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
	//router.Run("localhost:8080")
}

func getGuide(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// qwerty ask example
// /divided/?topKeys=qwertyuiop[]\&homeKeys=asdfghjkl%3B%27&botKeys=zxcvbnm,./&topFingers=0123344567777&homeFingers=01233445677&botFingers=0123344567
// /divided/?topKeys=qwertyuiop&homeKeys=asdfghjkl%3B%27&botKeys=zxcvbnm,./&topFingers=0123344567&homeFingers=01233445677&botFingers=0123344567
func getDivided(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	topKeys := c.Query("topKeys")
	homeKeys := c.Query("homeKeys")
	botKeys := c.Query("botKeys")
	topFingers := c.Query("topFingers")
	homeFingers := c.Query("homeFingers")
	botFingers := c.Query("botFingers")

	if topKeys == "" || homeKeys == "" || botKeys == "" {
		c.String(http.StatusBadRequest, "Some of query parameter 'row keys' doesn't exist")
		return
	}

	if topFingers == "" || homeFingers == "" || botFingers == "" {
		c.String(http.StatusBadRequest, "Some of query parameter 'row fingers' doesn't exist")
		return
	}

	analysis := genkey.Analyze(rowsFormDivided(topKeys, homeKeys, botKeys, topFingers, homeFingers, botFingers))

	c.JSON(http.StatusOK, analysis)
}

// c = context
// query need to be encoded to utf-8
// ; must be %3B
// example: ...?keys=qwertyuiop%5B%5Dasdfghjkl%3B%27zxcvbnm%2C.%2F&fingers=...
func getTogether(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	keys := c.Query("keys")
	fingers := c.Query("fingers")

	if keys == "" {
		c.String(http.StatusBadRequest, "Query parameter 'keys' doesn't exist")
		return
	}

	if fingers == "" {
		c.String(http.StatusBadRequest, "Query parameter 'fingers' doesn't exist")
		return
	}

	analysis := genkey.Analyze(rowsFromTogether(keys, fingers))

	c.JSON(http.StatusOK, analysis)
}
