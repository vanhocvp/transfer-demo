package routes

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/vanhocvp/junctionx-hackathon/transfer-demo/controllers"
	"github.com/vanhocvp/junctionx-hackathon/transfer-demo/setting"
	// "go.elastic.co/apm"
)

// CORS ...
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", setting.ServerSetting.CorsWhitelist)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORS())
	r.Use(location.Default())

	r.GET("/heath", controllers.HealthCheck)
	// Nhập thông tin số tài khoản, tên ngân hàng --> Tên khách hàng | So sánh tên khách hàng nếu có
	r.POST("/checkInfo", controllers.CheckInfoReceiver)
	// Lấy thông tin chuyển tiền dựa vào người thụ hưởng:
	// Nhập số thẻ --> Thông tin ngân hàng + họ tên

	// Nhập số điện thoại --> Họ tên

	// ==> Dựa vào thông giao dịch cũ để điền thông tin: Họ tên, stk, so the,

	// Tạo draft transaction
	r.POST("/createTransaction")
	// Authen transaction
	r.POST("/checkOTPAuthed")
	// Lấy danh sách người thụ hưởng gần nhất của userID: Giao dịch gần nhất
	r.GET("/getListRecipient")
	return r
}
