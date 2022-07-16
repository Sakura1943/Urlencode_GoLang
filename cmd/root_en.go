/*
Copyright © 2022 NAME HERE <sakunia@foxmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var url_utf8 string
var url_utf8_encode string

var rootCmdEn = &cobra.Command{
	Use:   "urlencode_go",
	Short: "URL text encryption and decryption tool.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(url_utf8) > 0 && len(url_utf8_encode) == 0 {
			fmt.Println(UrlEscape(url_utf8))
		} else if len(url_utf8_encode) > 0 && len(url_utf8) == 0 {
			fmt.Println(UrlUnescape(url_utf8_encode))
		} else {
			cmd.Help()
			return
		}
	},
}

func EnExecute() {
	err := rootCmdEn.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmdEn.Flags().StringVarP(&url_utf8, "encode", "e", "", "Encode url text.")
	rootCmdEn.Flags().StringVarP(&url_utf8_encode, "decode", "d", "", "Decode url text.")
}


