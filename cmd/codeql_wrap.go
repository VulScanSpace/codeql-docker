/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os/exec"

	"github.com/spf13/cobra"
)

// codeqlWrapCmd represents the codeqlWrap command
var codeqlWrapCmd = &cobra.Command{
	Use:   "codeqlWrap",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("codeqlWrap called")
	},
}

func init() {
	rootCmd.AddCommand(codeqlWrapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeqlWrapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeqlWrapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type CodeStatistic struct {
	Total struct {
		Blanks   int `yaml:"blanks"`
		Children struct {
			Java []struct {
				Name  string `yaml:"name"`
				Stats struct {
					Blanks int `yaml:"blanks"`
					Blobs  struct {
					} `yaml:"blobs"`
					Code     int `yaml:"code"`
					Comments int `yaml:"comments"`
				} `yaml:"stats"`
			} `yaml:"Java"`
			Go []struct {
				Name  string `yaml:"name"`
				Stats struct {
					Blanks int `yaml:"blanks"`
					Blobs  struct {
					} `yaml:"blobs"`
					Code     int `yaml:"code"`
					Comments int `yaml:"comments"`
				} `yaml:"stats"`
			} `yaml:"Go"`
			Python []struct {
				Name  string `yaml:"name"`
				Stats struct {
					Blanks int `yaml:"blanks"`
					Blobs  struct {
					} `yaml:"blobs"`
					Code     int `yaml:"code"`
					Comments int `yaml:"comments"`
				} `yaml:"stats"`
			} `yaml:"Python"`
			Markdown []struct {
				Name  string `yaml:"name"`
				Stats struct {
					Blanks int `yaml:"blanks"`
					Blobs  struct {
					} `yaml:"blobs"`
					Code     int `yaml:"code"`
					Comments int `yaml:"comments"`
				} `yaml:"stats"`
			} `yaml:"Markdown"`
			Groovy []struct {
				Name  string `yaml:"name"`
				Stats struct {
					Blanks int `yaml:"blanks"`
					Blobs  struct {
					} `yaml:"blobs"`
					Code     int `yaml:"code"`
					Comments int `yaml:"comments"`
				} `yaml:"stats"`
			} `yaml:"Groovy"`
			Kotlin []struct {
				Name  string `yaml:"name"`
				Stats struct {
					Blanks int `yaml:"blanks"`
					Blobs  struct {
					} `yaml:"blobs"`
					Code     int `yaml:"code"`
					Comments int `yaml:"comments"`
				} `yaml:"stats"`
			} `yaml:"Kotlin"`
			Scala []struct {
				Name  string `yaml:"name"`
				Stats struct {
					Blanks int `yaml:"blanks"`
					Blobs  struct {
					} `yaml:"blobs"`
					Code     int `yaml:"code"`
					Comments int `yaml:"comments"`
				} `yaml:"stats"`
			} `yaml:"Scala"`
			Xml []struct {
				Name  string `yaml:"name"`
				Stats struct {
					Blanks int `yaml:"blanks"`
					Blobs  struct {
					} `yaml:"blobs"`
					Code     int `yaml:"code"`
					Comments int `yaml:"comments"`
				} `yaml:"stats"`
			} `yaml:"Xml"`
		} `yaml:"children"`
		Code       int           `yaml:"code"`
		Comments   int           `yaml:"comments"`
		Inaccurate bool          `yaml:"inaccurate"`
		Reports    []interface{} `yaml:"reports"`
	} `yaml:"Total"`
}

// StatisticalCode 获取代码分析模型
func StatisticalCode(codeDir string) (codeStatistic *CodeStatistic, err error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("tokei", codeDir, "--output", "yaml")
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	cmd.Run()

	codeStatistic = &CodeStatistic{}
	yaml.Unmarshal(stdout.Bytes(), codeStatistic)
	return
}

const (
	Java   = "Java"
	Groovy = "Groovy"
	Kotlin = "Kotlin"
	Scala  = "Scala"
	Python = "Python"
	GoLang = "Go"
)

// GetCodeLanguage 获取主要编程语言
func GetCodeLanguage(codeStatistic *CodeStatistic) (languageValue string) {
	var fileCount int
	if len(codeStatistic.Total.Children.Java) > fileCount {
		languageValue = Java
		fileCount = len(codeStatistic.Total.Children.Java)
	}
	if len(codeStatistic.Total.Children.Go) > fileCount {
		languageValue = GoLang
		fileCount = len(codeStatistic.Total.Children.Go)
	}
	if len(codeStatistic.Total.Children.Python) > fileCount {
		languageValue = Python
		fileCount = len(codeStatistic.Total.Children.Python)
	}
	if len(codeStatistic.Total.Children.Groovy) > fileCount {
		languageValue = Groovy
		fileCount = len(codeStatistic.Total.Children.Groovy)
	}
	if len(codeStatistic.Total.Children.Kotlin) > fileCount {
		languageValue = Kotlin
		fileCount = len(codeStatistic.Total.Children.Kotlin)
	}
	if len(codeStatistic.Total.Children.Scala) > fileCount {
		languageValue = Scala
		fileCount = len(codeStatistic.Total.Children.Scala)
	}
	return
}

// GetBuildTool TODO 识别构建工具
func GetBuildTool(codeStatistic *CodeStatistic) {
	xmlItems := codeStatistic.Total.Children.Xml
	for _, xmlItem := range xmlItems {
		logrus.Info(xmlItem.Name)
	}
}

// 识别构建语句
// 调用 CodeQL 构建 DB
// 调用 CodeQL 扫描漏洞
