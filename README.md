# codeql-docker
Docker for the CodeQL CLI

## Feature

- Standardized scan engine
- Convenient scanning parameters

## Scenario-based Usage

- Scan Java projects for security risks
- Scan Go projects for security risks
- Scan Python projects for security risks
- Scan Java+Kotlin projects for security risks
- Scan Java+Groovy projects for security risks，etc.

### Scan Java projects for security risks
1. **(precondition)** upload project to oss、git or anywhere
2. run docker with project url and report name
3. download report from oss

```shell
$ docker run codeql-docker:latest "http://oss.cn-hangzhou.aliyuncs.com/vulscan/java/project.tar.gz" "project-202211051439001.sarif"
$ wget http://oss.cn-hangzhou.aliyuncs.com/vulscan/report/project-202211051439001.sarif
```

### Scan Go projects for security risks

### Scan Python projects for security risks

### Scan Java+Kotlin projects for security risks

### Scan Java+Groovy projects for security risks


## ROADMAP

- Intelligent (recognition language、build tool, etc.）

