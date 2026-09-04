package main

import (
	"example/first/common"
	"example/first/proto"
	"example/first/repository"
	"example/first/server"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// TODO: Перенести весь в DI в отдельный файл
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := common.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")

	repo := repository.NewUserRepository(db)
	chatRepo := repository.NewChatRepository(db)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(
		grpcServer,
		&server.UserServer{Repo: repo},
	)
	proto.RegisterChatServiceServer(
		grpcServer,
		&server.ChatServer{Repo: chatRepo, UserRepo: repo},
	)

	log.Println("gRPC server started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
