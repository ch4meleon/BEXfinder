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
var firefoxCmd = &cobra.Command{
	Use:   "firefox",
	Short: "Find and list out FireFox plugins",
	Long:  "Find and list out FireFox plugins",
	Run: func(cmd *cobra.Command, args []string) {
		// Init colors
		// yellow := color.New(color.FgYellow).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()

		// Get flags
		to_check_online, _ := cmd.Flags().GetBool("check")
		to_save, _ := cmd.Flags().GetString("save")
		custom_firefox_profile, _ := cmd.Flags().GetString("custom_firefox_profile")

		fmt.Printf("%s", blue("[FireFox Browsers]\n"))

		if custom_firefox_profile == "" {
			helper.FindFireFoxExtensions(to_save, to_check_online, is_save_all)
		} else {
			helper.FindFireFoxExtensionsFromPath(custom_firefox_profile, to_save, to_check_online, is_save_all)
		}

	},
}

func init() {
	rootCmd.AddCommand(firefoxCmd)
	firefoxCmd.Flags().BoolP("check", "c", false, "Check plugins online")
}
