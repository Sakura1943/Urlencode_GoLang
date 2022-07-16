/*
Copyright © 2022 NAME HERE <sakunia@foxmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var _url_utf8 string
var _url_utf8_encode string

var rootCmdCn = &cobra.Command{
	Use:   "urlencode_go",
	Short: "URL文字解密以及加密工具。",
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

func CnExecute() {
	err := rootCmdCn.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmdCn.Flags().StringVarP(&url_utf8, "encode", "e", "", "加密URL文字。")
	rootCmdCn.Flags().StringVarP(&url_utf8_encode, "decode", "d", "", "解密URL文字.")
}


