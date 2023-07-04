/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bexfinder/helper"
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Version
var version = "0.1.0"

// Global Variables
var is_save_all bool = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "bexfinder",
	Version: version,
	Short:   "Browsers Extensions Finder (BEXfinder)",
	Long:    "\nBrowsers Extensions Finder (BEXfinder) is a cross-platform and portable command-line tool to find all web browsers (Google Chrome, Microsoft Edge, Brave Browser, Mozilla FireFox, Opera, etc.) extensions installed on system.",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {

	// },

	Run: func(cmd *cobra.Command, args []string) {
		// isDebugSet := cmd.Flags().Lookup("save").Changed
		// fmt.Printf("%s", isDebugSet)

		os_type := helper.GetOperatingSystem()

		if os_type == "linux" {
			// fmt.Println(os_type)
		}

		cmd.Help()

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Print banner
	myFigure := figure.NewColorFigure("BEXfinder", "", "green", true)
	myFigure.Print()

	// Init colors
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Printf("\n %s\n", green("Version "+version))
	fmt.Printf("\n %s\n", yellow("BROWSERS EXTENSIONS FINDER"))
	fmt.Printf(" %s\n", red("Developed by CH4MELE0N [ch4meleon@protonmail-dot-com]"))
	fmt.Println("")

	// Flags
	// rootCmd.Flags().BoolP("debug", "d", false, "Debug mode")
	rootCmd.PersistentFlags().String("save", "", "Save to file")
	rootCmd.PersistentFlags().String("custom_firefox_profile", "", "Specify custom FireFox profile directory")
	rootCmd.PersistentFlags().String("custom_chrome_profile", "", "Specify custom Chrome profile directory")

	// No completion
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
