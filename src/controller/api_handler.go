package controller

import (
	"io"
	"os"
	"time"

	"github.com/SystemEngineeringTeam/ait-guide-back/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Route []service.Coordinate `json:"route"`
}

func ApiHandler() {
	// サーバーログの出力先を設定
	gin.DisableConsoleColor()
	f, _ := os.Create("../server.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// ルーティングの設定
	r := gin.Default()

	// ここからCorsの設定
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"https://ait-guide.sysken.net",
			"http://localhost:3000",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"GET",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Accept",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	r.GET("api/current", func(c *gin.Context) {
		lat := c.Query("lat")
		lng := c.Query("lng")
		end := c.Query("end")

		route := service.GetShortestDistanceFromCurrentLocation(lat, lng, end)
		c.JSON(200, route)
	})

	r.Run()
}
