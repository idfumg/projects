/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package net

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NetCmd represents the net command
var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "Net command",
	Long: `Net command long description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("net command called")
	},
}

func init() {
	// rootCmd.AddCommand(netCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// netCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
