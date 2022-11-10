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
```go
go run main.go make model [file_name]
```

- 生成controllerapi
```go
go run main.go make apicontroller [file_name]
```


---


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

---


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

### 3.操作数据库命令
- 生成创建数据表迁移文件
```go
go run main.go make migration add_users_table
```

- 生成修改表字段迁移文件
```go
go run main.go make migration add_info_fields_to_link
```

- 执行迁移
```go
go run main.go migrate up
```

- 执行迁移回滚到上一次
```go
go run main.go migrate down
```

- 回溯所有迁移
```go
go run main.go migrate reset
```
--- 
### 集成的api
| **请求方法** | **API 地址** | **说明** |
| --- | --- | --- |
| POST | /api/v1/auth/login/using-phone | 短信 + 手机号登录 |
| POST | /api/v1/auth/login/using-password | 手机号、用户名、邮箱 + 密码 |
| POST | /api/v1/auth/login/refresh-token | 刷下 Token |
| POST | /api/v1/auth/password-reset/using-email | 邮件密码重置 |
| POST | /api/v1/auth/password-reset/using-phone | 短信验证码密码重置 |
| POST | /api/v1/auth/signup/using-phone | 使用手机号注册 |
| POST | /api/v1/auth/signup/using-email | 使用邮箱注册 |
| POST | /api/v1/auth/signup/phone/exist | 手机号是否已注册 |
| POST | /api/v1/auth/signup/email/exist | email 是否已支持 |
| POST | /api/v1/auth/verify-codes/phone | 发送短信验证码 |
| POST | /api/v1/auth/verify-codes/email | 发送邮件验证码 |
| POST | /api/v1/auth/verify-codes/captcha | 获取图片验证码 |
| GET | /api/v1/user | 获取当前用户 |
| GET | /api/v1/users | 用户列表 |
| PUT | /api/v1/users | 修改个人资料 |
| PUT | /api/v1/users/email | 修改邮箱 |
| PUT | /api/v1/users/phone | 修改手机号 |
| PUT | /api/v1/users/password | 修改密码 |
| PUT | /api/v1/users/avatar | 上传头像 |
| GET | /api/v1/categories | 分类列表 |
| POST | /api/v1/categories | 创建分类 |
| PUT | /api/v1/categories/:id | 更新分类 |
| DELETE | /api/v1/categories/:id | 删除分类 |
| GET | /api/v1/topics | 话题列表 |
| POST | /api/v1/topics | 创建话题 |
| PUT | /api/v1/topics/:id | 更新话题 |
| DELETE | /api/v1/topics/:id | 删除话题 |
| GET | /api/v1/topics/:id | 获取话题 |
| GET | /api/v1/links | 友情链接列表 |


