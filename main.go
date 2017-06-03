package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	api "power/service"
	"github.com/dgrijalva/jwt-go"

)

func respondWithError(code int, message string,c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt_string:= c.Request.Header.Get("auth")

		token, err := jwt.Parse(jwt_string, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret123"), nil
		})


		if err == nil && token.Valid {
			c.Next()
			return

		}
		respondWithError(401, "Permission denied", c)
		return

	}
}


func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(TokenAuthMiddleware())

	r.GET("/getDevice",api.GetDevicePower)
	r.GET("/getDeviceLimit",api.GetDevicePowerLimit)
	r.GET("/getDeviceByTime",api.GetDevicePowerByTime)

	r.GET("/getDeviceInMonth",api.GetDevicePowerInMonth)
	r.GET("/getDeviceInYear",api.GetDevicePowerInYear)

	r.GET("/getRankInMonth",api.GetRankingDevicePowerInMonth)
	r.GET("/getRankInYear",api.GetRankingDevicePowerInYear)
	r.GET("/getRankAll",api.GetRankingDevicePowerAll)

	r.GET("/getHomeInYear",api.GetHomePowerInYear)
	r.GET("/getHomeInMonth",api.GetHomePowerInMonth)



	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("127.0.0.1:8081") // listen and serve on 0.0.0.0:8080
}