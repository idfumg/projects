/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package info

import (
	"fmt"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
)

// DiskUsageCmd represents the diskUsage command
var DiskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		usage := du.NewDiskUsage(".")
		fmt.Printf("Available: %d, Free: %d, Size: %d, Usage: %f, Used: %d\n",
			usage.Available(), usage.Free(), usage.Size(), usage.Usage(), usage.Used())
	},
}

func init() {
	InfoCmd.AddCommand(DiskUsageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
