package command

import (
	"Bing-Wallpaper-RESTful/roothandler"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/thinkerou/favicon"
)

// command-line
var Cmd = &cobra.Command{
	Use:   "run",
	Short: "Run this API service",
	Long:  `Run this API service`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(favicon.New("./image/favicon.ico"))
	router.GET("/", roothandler.RootHandler)
	router.Run(":9002") // default port: 9002
	fmt.Println("API is running...")
}
