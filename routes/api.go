// Package routes 注册路由
package routes

import (
	controllers "GoLaravel/app/http/controllers/api/v1"
	"GoLaravel/app/http/controllers/api/v1/auth"
	"GoLaravel/app/http/middlewares"
	middleware_app "GoLaravel/app/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 静态图片访问url
	r.StaticFS("/public/uploads", http.Dir("./public/uploads"))

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		v1.Use(middlewares.LimitIP("200-H"))
		{
			authGroup := v1.Group("/auth")
			{
				// suc := new(auth.SignupController)
				suc := auth.SignupController{}
				// 判断手机是否已注册
				authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
				authGroup.POST("/signup/email/exist", suc.IsEmailExist)
				authGroup.POST("/signup/email", suc.SignupUsingEmail)
				authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)

				vcc := new(auth.VerifyCodeController)
				authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)

				user := new(auth.UserController)
				authGroup.POST("/getUser", user.GetUser)

				authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)

				authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)

				// lgc := new(auth.LoginController)
				var lgc auth.LoginController
				// 使用手机号，短信验证码进行登录
				authGroup.POST("/login/using-phone", lgc.LoginByPhone)
				// 使用邮箱登录
				authGroup.POST("/login/using-email", lgc.LoginByEmail)
				// 账号密码登录
				authGroup.POST("/login/account", lgc.LoginByPassword)
				// 刷新 token
				authGroup.POST("/login/refresh-token", lgc.RefreshToken)

				// 手机修改密码
				pwc := auth.PasswordController{}
				authGroup.POST("/password-reset/using-phone", middlewares.LimitPerRoute("20-H"), pwc.PhoneByPassword)
				authGroup.POST("/password-reset/using-email", middlewares.LimitPerRoute("20-H"), pwc.EmailByPassword)

				uc := new(controllers.UsersController)

				// 获取当前用户
				// http://127.0.0.1:3002/v1/user
				v1.GET("/user", middleware_app.AuthJWT(), uc.CurrentUser)

				// 127.0.0.1:3002/v1/users
				usersGroup := v1.Group("/users")
				{
					usersGroup.GET("", uc.Index)
					usersGroup.PUT("", middleware_app.AuthJWT(), uc.UpdateProfile)
					usersGroup.PUT("/email", middleware_app.AuthJWT(), uc.UpdateEmail)
					usersGroup.PUT("/phone", middleware_app.AuthJWT(), uc.UpdatePhone)
					usersGroup.PUT("/update_password", middleware_app.AuthJWT(), uc.UpdatePassword)
					usersGroup.PUT("/avatar", middleware_app.AuthJWT(), uc.UpdateAvatar)
				}

				cgc := controllers.CategoriesController{}
				cgcGroup := v1.Group("/categories")
				{
					cgcGroup.POST("", middleware_app.AuthJWT(), cgc.Store)
					cgcGroup.PUT("/:id", middleware_app.AuthJWT(), cgc.Update) //127.0.0.1:3002/v1/categories/1
					cgcGroup.GET("", middleware_app.AuthJWT(), cgc.Index)
					cgcGroup.DELETE("/:id", middleware_app.AuthJWT(), cgc.Delete)
				}

				tpc := controllers.TopicsController{}
				tpcGroup := v1.Group("/topics")
				{
					tpcGroup.GET("/find", tpc.FindTest)
					tpcGroup.GET("", tpc.Index)
					tpcGroup.GET("/:id", tpc.Show)
					tpcGroup.POST("", middleware_app.AuthJWT(), tpc.Store)
					tpcGroup.PUT("/:id", middleware_app.AuthJWT(), tpc.Update)
					tpcGroup.DELETE("/:id", middleware_app.AuthJWT(), tpc.Delete)
				}

				link := controllers.LinksController{}
				linkGroup := v1.Group("/link")
				{
					linkGroup.GET("", link.Index)
				}
			}

		}
	}
}
