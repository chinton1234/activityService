package main

import ("github.com/gin-gonic/gin"
		"activity/configs"
		"activity/routes"
	)

func main() {
	r := gin.Default()

	configs.ConnectDB()

	routes.ActivityRoute(r)

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 			"data": "Hello from Gin-gonic & mongoDB",
	// 	})
	// })

	r.Run("localhost:6050")
}