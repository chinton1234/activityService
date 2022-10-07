package routes

import (
    "github.com/gin-gonic/gin"
    "activity/controllers" 
)

func ActivityRoute(router *gin.Engine)  {
    //All routes related to users comes here

    router.POST("/activity", controllers.CreateActivity())
	router.GET("/activity/:activityId", controllers.GetActivity())
	router.PUT("/activity/:activityId", controllers.EditActivity())
	router.DELETE("/activity/:activityId", controllers.DeleteActivity())
	router.GET("/activity", controllers.GetAllActivity())

}