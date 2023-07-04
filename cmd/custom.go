/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bexfinder/helper"
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// customCmd represents the custom command
var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "Find and list out extensions and plugins for custom locations",
	Long:  "Find and list out extensions and plugins for custom locations",
	Run: func(cmd *cobra.Command, args []string) {
		// Init colors
		yellow := color.New(color.FgYellow).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()

		// Get flags
		to_check_online, _ := cmd.Flags().GetBool("check")
		to_save, _ := cmd.Flags().GetString("save")
		custom_profile_path, _ := cmd.Flags().GetString("file")

		if to_check_online == true {
			fmt.Printf("%s\n\n", yellow("INFO: Check online feature is not supported in custom profile path mode."))
		}

		fmt.Printf("%s", blue("[Custom Locations]\n"))

		if custom_profile_path != "" {

			// Open the file
			file, err := os.Open(custom_profile_path)
			if err != nil {
				fmt.Println("ERROR: Unable to open file:", err)
				return
			}
			defer file.Close()

			// Create a new scanner to read the file line by line
			scanner := bufio.NewScanner(file)

			// Read and print each line
			for scanner.Scan() {
				line := scanner.Text()
				// fmt.Println(line)

				helper.FindChromeBrowserExtensionsFromPath(line, to_save, false, false, false, true)
			}

			// Check for any errors during scanning
			if err := scanner.Err(); err != nil {
				fmt.Println("ERROR: Unable to scan file:", err)
			}

		} else {
			fmt.Println("ERROR: No input file. Use '--file' to specify one.")
		}

	},
}

func init() {
	rootCmd.AddCommand(customCmd)
	customCmd.Flags().BoolP("check", "c", false, "Check extensions online")
	customCmd.PersistentFlags().String("file", "", "Specify a file containing custom browser profiles locations")
}
