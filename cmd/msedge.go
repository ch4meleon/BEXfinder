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
var msedgeCmd = &cobra.Command{
	Use:   "msedge",
	Short: "Find and list out MS Edge browser extensions",
	Long:  "Find and list out MS Edge browser extensions",
	Run: func(cmd *cobra.Command, args []string) {
		// Init colors
		// yellow := color.New(color.FgYellow).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()

		// Get flags
		to_check_online, _ := cmd.Flags().GetBool("check")
		to_save, _ := cmd.Flags().GetString("save")
		custom_chrome_profile, _ := cmd.Flags().GetString("custom_chrome_profile")

		fmt.Printf("%s", blue("[MS Edge Browsers]\n"))

		if custom_chrome_profile == "" {
			helper.FindMSEdgeBrowserExtensions(to_save, to_check_online, is_save_all)
		} else {
			helper.FindChromeBrowserExtensionsFromPath(custom_chrome_profile, to_save, to_check_online, true, false, is_save_all)
		}

	},
}

func init() {
	rootCmd.AddCommand(msedgeCmd)
	msedgeCmd.Flags().BoolP("check", "c", false, "Check extensions online")
}
