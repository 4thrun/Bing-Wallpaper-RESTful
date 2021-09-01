package main

import (
	"Bing-Wallpaper-RESTful/command"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var argVerbose bool
var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "bw",
		Short: "Bing wallpaper API",
		Long:  "Top level command for Bing wallpaper API service",
	}
	rootCmd.PersistentFlags().BoolVarP(&argVerbose, "verbose", "v", false, "verbose output")
	rootCmd.AddCommand(
		command.Cmd,
	)
}

func main() {
	if err1 := rootCmd.Execute(); err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
}
