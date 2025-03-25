package hand

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Person struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Profiles map[string]Person
}

func NewPerson() *Person {
	return &Person{
		Profiles: make(map[string]Person),
	}

}

func (p *Person) GetProfileHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}

	person, exists := p.Profiles[id]
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
func (p *Person) CreateProfileHandler(c *gin.Context) {
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

	p.Profiles[person.ID] = person

	c.JSON(http.StatusCreated, gin.H{
		"message": "Profile created",
		"id":      person.ID,
		"name":    person.Name,
		"age":     person.Age,
	})
}
func (p *Person) UpdateProfileHandler(c *gin.Context) {
	var person Person

	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if person.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}
	_, exists := p.Profiles[person.ID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	p.Profiles[person.ID] = person

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated",
		"id":      person.ID,
		"name":    person.Name,
		"age":     person.Age,
	})

}
func (p *Person) DeleteProfileHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}
	_, exists := p.Profiles[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	delete(p.Profiles, id)

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted", "id": id})
}
