/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"ast.framework.sast.codeql/cmd"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	if os.Getenv("sast.codeql.debug") == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	cmd.Execute()
}
