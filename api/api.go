package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	genkey "github.com/roman-koshchei/genkey-api/pkg"
)

// c = context
// query need to be encoded to utf-8
// ; must be %3B
// example: ...?keys=qwertyuiop%5B%5Dasdfghjkl%3B%27zxcvbnm%2C.%2F&fingers=...
func getAnalysis(c *gin.Context) {
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

	analysis := genkey.Analyze(keys, fingers)

	c.JSON(http.StatusOK, analysis)
}

func Run() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery()) // if panic return 500

	router.GET("/", getAnalysis)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
