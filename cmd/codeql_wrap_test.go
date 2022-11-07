package cmd

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestStatisticalCode(t *testing.T) {
	codeDir := "/Users/owefsad/GoProjects/io.ast.plugins"
	StatisticalCode(codeDir)
}

func TestGetCodeLanguage(t *testing.T) {
	codeDir := "/Users/owefsad/GoProjects/io.ast.plugins"
	codeDir = "/Users/owefsad/IdeaProjects/logging-log4j2"
	codeDir = "/Users/owefsad/IdeaProjects/spark"
	codeDir = "/Users/owefsad/IdeaProjects/SpringCloud"
	codeDir = "/Users/owefsad/IdeaProjects/spark"
	codeStatistic, err := StatisticalCode(codeDir)
	if err != nil {
		return
	}
	languageValue := GetCodeLanguage(codeStatistic)
	logrus.Infof("代码仓库%s的主要语言为：%s", codeDir, languageValue)
}

func TestGetBuildTool(t *testing.T) {
	codeDir := "/Users/owefsad/GoProjects/io.ast.plugins"
	codeDir = "/Users/owefsad/IdeaProjects/logging-log4j2"
	codeDir = "/Users/owefsad/IdeaProjects/spark"
	codeDir = "/Users/owefsad/IdeaProjects/SpringCloud"
	codeDir = "/Users/owefsad/IdeaProjects/spark"
	codeStatistic, err := StatisticalCode(codeDir)
	if err != nil {
		return
	}
	languageValue := GetCodeLanguage(codeStatistic)

	switch languageValue {
	case Java, Groovy, Scala, Kotlin:
		logrus.Infof("代码仓库%s的主要语言为：%s", codeDir, languageValue)
		GetBuildTool(codeStatistic)
		break
	case GoLang:
		logrus.Infof("代码仓库%s的主要语言为：%s", codeDir, languageValue)
		break
	case Python:
		logrus.Infof("代码仓库%s的主要语言为：%s", codeDir, languageValue)
		break
	}
}
