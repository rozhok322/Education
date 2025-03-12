package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var profiles = make(map[string]Person)

func main() {
	router := gin.Default()

	router.GET("/profile", getProfileHandler)
	router.POST("/profile", createProfileHandler)
	router.PUT("/profile", updateProfileHandler)
	router.DELETE("/profile", deleteProfileHandler)

	router.Run(":8080")
}

func getProfileHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}

	person, exists := profiles[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile Found",
		"id":      person.ID,
		"name":    person.Name,
		"age":     person.Age,
	})

}
func createProfileHandler(c *gin.Context) {
	var person Person

	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if person.Age <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ti debil, age doljen bit greater than zero"})
		return
	}

	person.ID = uuid.New().String()

	profiles[person.ID] = person

	c.JSON(http.StatusCreated, gin.H{
		"message": "Profile created",
		"id":      person.ID,
		"name":    person.Name,
		"age":     person.Age,
	})
}
func updateProfileHandler(c *gin.Context) {
	var person Person

	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if person.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}
	_, exists := profiles[person.ID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	profiles[person.ID] = person

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated",
		"id":      person.ID,
		"name":    person.Name,
		"age":     person.Age,
	})

}
func deleteProfileHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}
	_, exists := profiles[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	delete(profiles, id)

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted", "id": id})
}
