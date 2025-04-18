package main

import (
	"context"
	"log"

	hand "example.com/myapp/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func getConnect(url string) (*pgx.Conn, error) {

	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		return nil, err
	}
	return conn, nil
}
func main() {
	connString := "postgres://postgres:postgres@localhost:5432/postgres"
	conn, err := getConnect(connString)
	if err != nil {
		log.Fatalf("Error : %v, conString: %s", err, connString)
	}
	handle := hand.NewHandle(conn)
	router := gin.Default()
	router.GET("/profile/", handle.GetProfileHandler)
	router.POST("/profile", handle.CreateProfileHandler)
	router.PUT("/profile", handle.UpdateProfileHandler)
	router.DELETE("/profile", handle.DeleteProfileHandler)

	router.Run(":8080")

}
