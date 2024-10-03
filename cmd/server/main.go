package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/gabriel01-jpg/go-grpc/internal/database"
	"github.com/gabriel01-jpg/go-grpc/internal/pb"
	"github.com/gabriel01-jpg/go-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDb := database.NewCategory(db)

	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()

	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	fmt.Println("Before Serve")

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

	fmt.Println("Server running at :50051")
}
