package hand

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type Handle struct {
	conn *pgx.Conn
}

func NewHandle(conn *pgx.Conn) *Handle {

	return &Handle{
		conn: conn,
	}
}

func (p *Handle) GetProfileHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}
	var person Person

	err := p.conn.QueryRow(context.Background(), "select id,name,age from profiles where id=$1", id).Scan(&person.ID, &person.Name, &person.Age)

	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile Found",
		"id":      person.ID,
		"name":    person.Name,
		"age":     person.Age,
	})

}
func (p *Handle) CreateProfileHandler(c *gin.Context) {
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

	_, err := p.conn.Exec(context.Background(), "insert into profiles (id,name,age) values($1,$2,$3)", person.ID, person.Name, person.Age)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create profile: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Profile created",
		"id":      person.ID,
		"name":    person.Name,
		"age":     person.Age,
	})
}
func (p *Handle) UpdateProfileHandler(c *gin.Context) {
	var person Person

	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if person.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}
	err := p.conn.QueryRow(context.Background(), "UPDATE profiles SET name = $1, age = $2 WHERE id = $3 RETURNING id, name, age", person.Name, person.Age, person.ID).Scan(&person.ID, &person.Name, &person.Age)

	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Database error: %v", err)})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated",
		"id":      person.ID,
		"name":    person.Name,
		"age":     person.Age,
	})

}
func (p *Handle) DeleteProfileHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}
	var deleteID string
	err := p.conn.QueryRow(context.Background(), "DELETE FROM profiles WHERE id = $1 RETURNING id", id).Scan(&deleteID)

	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Database error: %v", err)})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted", "id": id})
}
