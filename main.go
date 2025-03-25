package main

import (
	hand "example.com/myapp/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	Handle := hand.NewPerson()
	router := gin.Default()
	router.GET("/profile", Handle.GetProfileHandler)
	router.POST("/profile", Handle.CreateProfileHandler)
	router.PUT("/profile", Handle.UpdateProfileHandler)
	router.DELETE("/profile", Handle.DeleteProfileHandler)

	router.Run(":8080")

}
