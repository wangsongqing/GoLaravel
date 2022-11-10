# GoLaravel
## 基于gin搭建的类似于Laravel的api框架

### 1. 运行项目

- 拉取并安装
```go
git clone 
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


