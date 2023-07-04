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

// braveCmd represents the brave command
var braveCmd = &cobra.Command{
	Use:   "brave",
	Short: "Find and list out Opera Browser extensions",
	Long:  "Find and list out Opera Browser extensions",
	Run: func(cmd *cobra.Command, args []string) {
		// Init colors
		blue := color.New(color.FgBlue).SprintFunc()

		// Get flags
		to_check_online, _ := cmd.Flags().GetBool("check")
		to_save, _ := cmd.Flags().GetString("save")
		custom_chrome_profile, _ := cmd.Flags().GetString("custom_chrome_profile")

		// Start
		fmt.Printf("%s", blue("[Brave Browsers]\n"))

		if custom_chrome_profile == "" {
			helper.FindBraveBrowserExtensions(to_save, to_check_online, is_save_all)
		} else {
			helper.FindChromeBrowserExtensionsFromPath(custom_chrome_profile, to_save, to_check_online, false, true, is_save_all)
		}

	},
}

func init() {
	rootCmd.AddCommand(braveCmd)
	braveCmd.Flags().BoolP("check", "c", false, "Check extensions online")
}
