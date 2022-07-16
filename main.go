/*
Copyright © 2022 NAME HERE <sakunia@foxmail.com>

*/
package main

import (
	"demo1/cmd"
	"os"
	"strings"
)

func main() {
	lang := os.Getenv("LANG")
	if strings.Contains(lang, "CN") {
		cmd.CnExecute()
	} else {
		cmd.EnExecute()
	}
}
