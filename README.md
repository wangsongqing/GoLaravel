# GoLaravel
## 基于gin搭建的类似于Laravel的api框架

### 1. 运行项目

- 拉取并安装
```go
git clone git@github.com:wangsongqing/GoLaravel.git
go mod tidy
```
- 修改 .env.example 为 .env
```azure
cp .env.example .env 
```
- 迁移数据库
```azure
go run main.go migrate up
```
- 运行项目
```azure
go run main.go
```
---
### 2.类artisan创建命令
- 生成model
```azure
go run main.go make model [file_name]
```

- 生成controllerapi
```
go run main.go make apicontroller [file_name]
```

- 生成job文件
```azure
go run main.go make job job_name
```
- 添加子命令到 GoLaravel/app/cmd/job/job.go 
```go
func init() {
	Job.AddCommand(
		CmdJobName,
		)
}
```
- 执行job命令
```azure
go run main.go job job_name
```

- 生成console文件
```go
go run main.go make console console_name
```

- 添加子命令到 GoLaravel/app/cmd/console/console.go 
```go
func init() {
	Console.AddCommand(
		CmdConsoleName,
	)
}
```
- 执行console命令
```go
go run main.go console console_name
```


