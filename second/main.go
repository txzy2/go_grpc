package main

import (
	"log"
	"os"
	"second/handler/http/v1"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	conn, err := grpc.NewClient(
		os.Getenv("CHAT_SERVICE_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("не удалось подключиться: %v", err)
	}
	log.Printf("Успешное подключение GRPC сервера")
	defer conn.Close()

	v1.SetupRoutes(r)

	log.Println("Сервер запущен на http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
