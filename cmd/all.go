/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bexfinder/helper"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "List all installed web browsers' extensions",

	Run: func(cmd *cobra.Command, args []string) {
		// Init colors
		yellow := color.New(color.FgYellow).SprintFunc()
		// red := color.New(color.FgRed).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()
		// green := color.New(color.FgGreen).SprintFunc()

		// A newline will be appended automatically
		// blue("Prints %s in blue.")

		// Create SprintXxx functions to mix strings with other non-colorized strings:

		// fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red("error"))

		fmt.Printf("%s", blue("*** Finding All Web Browsers Extensions ***"))

		// Get flags
		to_check_online, _ := cmd.Flags().GetBool("check")
		to_save, _ := cmd.Flags().GetString("save")

		// Set 'is_save_all' to true
		is_save_all = true

		// Check if the file exists
		if _, err := os.Stat(to_save); err == nil {
			// File exists, delete it
			err := os.Remove(to_save)
			if err != nil {
				fmt.Printf("Failed to delete the file: %v\n", err)
				return
			}
			fmt.Println("File deleted successfully.")
		}

		// Start
		fmt.Printf("%s", yellow("\n[Google Chrome Browsers]\n"))
		helper.FindChromeBrowserExtensions(to_save, to_check_online, is_save_all)

		fmt.Printf("%s", yellow("\n[Brave Browsers]\n"))
		helper.FindBraveBrowserExtensions(to_save, to_check_online, is_save_all)

		fmt.Printf("%s", yellow("\n[Chromium Browsers]\n"))
		helper.FindChromiumBrowserExtensions(to_save, to_check_online, is_save_all)

		fmt.Printf("%s", yellow("\n[MSEdge Browsers]\n"))
		helper.FindMSEdgeBrowserExtensions(to_save, to_check_online, is_save_all)

		fmt.Printf("%s", yellow("\n[FireFox Browsers]\n"))
		helper.FindFireFoxExtensions(to_save, to_check_online, is_save_all)

		fmt.Printf("%s", yellow("\n[Opera Browsers]\n"))
		helper.FindOperaBrowserExtensions(to_save, to_check_online, is_save_all)

		fmt.Printf("%s", yellow("\n[Vivaldi Browsers]\n"))
		helper.FindVivaldiBrowserExtensions(to_save, to_check_online, is_save_all)

	},
}

func init() {
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().BoolP("check", "c", false, "Check extensions online")
}
