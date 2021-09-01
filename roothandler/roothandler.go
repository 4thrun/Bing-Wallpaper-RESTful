package roothandler

import (
	"Bing-Wallpaper-RESTful/xmlhandler"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RootHandler: serve as default API using gin
func RootHandler(c *gin.Context) {
	// set default query params
	resolution := c.DefaultQuery("resolution", "1920")
	format := c.DefaultQuery("format", "json")
	index := c.DefaultQuery("index", "0")
	market := c.DefaultQuery("mkt", "zh-CN")

	// check index
	uIndex, err := strconv.ParseUint(index, 10, 64)
	if err != nil {
		// index invalid
		c.String(http.StatusInternalServerError, "image index invalid")
		return
	}

	// check format
	if format != "json" && format != "image" {
		c.String(http.StatusInternalServerError, "format invalid, only `json` or `image` available")
		return
	}

	// fetch info from Bing using xmlhandler
	response, err := xmlhandler.Get(uint(uIndex), market, resolution)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// redirect to image URL directly
	if format == "image" && response.URL != "" {
		c.Redirect(http.StatusTemporaryRedirect, response.URL)
		return
	}
	// render response as JSON
	c.JSON(http.StatusOK, response)
}
