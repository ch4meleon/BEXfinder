/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bexfinder/helper"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// firefoxCmd represents the firefox command
var chromeCmd = &cobra.Command{
	Use:   "chrome",
	Short: "Find and list out Google Chrome/Brave extensions",
	Long:  "Find and list out Google Chrome/Brave extensions",
	Run: func(cmd *cobra.Command, args []string) {
		// Init colors
		blue := color.New(color.FgBlue).SprintFunc()

		// Get flags
		to_check_online, _ := cmd.Flags().GetBool("check")
		to_save, _ := cmd.Flags().GetString("save")
		custom_chrome_profile, _ := cmd.Flags().GetString("custom_chrome_profile")

		// Start
		fmt.Printf("%s", blue("[Google Chrome Browsers / Brave Browsers]\n"))

		if custom_chrome_profile == "" {
			helper.FindChromeBrowserExtensions(to_save, to_check_online, is_save_all)
		} else {
			helper.FindChromeBrowserExtensionsFromPath(custom_chrome_profile, to_save, to_check_online, false, false, is_save_all)
		}

		// Debug
		// tmp := helper.GetBrowserPaths()
		// for key, value := range tmp {
		// 	fmt.Println(key, ":", value)

		// }
		// fmt.Println(tmp)

	},
}

func init() {
	rootCmd.AddCommand(chromeCmd)
	chromeCmd.Flags().BoolP("check", "c", false, "Check extensions online")
}
