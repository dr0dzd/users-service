package main

import (
	"github.com/dr0dzd/users-service/internal/database"
	"github.com/dr0dzd/users-service/internal/user"
	"github.com/dr0dzd/users-service/internal/transport/grpc"
	"log"
	"time"
)

func main(){
	database.InitDB() // Инициализация базы данных
	repo := user.NewRepository(database.DB)
	svc  := user.NewService(repo)
  
	go func(){
		if err := grpc.RunGRPC(svc); err != nil {
	  		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
		}
	}()
	for{
		log.Println("server is running on port: 50051")
		time.Sleep(time.Second*10)
	}
}