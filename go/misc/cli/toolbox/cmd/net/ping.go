/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package net

import (
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath string
)

func ping(domain string) (int, error) {
	url := "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}

// PingCmd represents the ping command
var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := ping(urlPath); err != nil {
			log.Println(err)
		} else {
			log.Println(resp)
		}
	},
}

func init() {
	PingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to ping")
	if err := PingCmd.MarkFlagRequired("url"); err != nil {
		log.Fatalln(err)
	}
	NetCmd.AddCommand(PingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
