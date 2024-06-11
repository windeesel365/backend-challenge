package main

import (
	"context"
	"countmeat/grpc"
	"countmeat/handler"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// godotenv load จาก .env file
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}
}

func main() {
	loadEnv()
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e := echo.New() //echo instance

	e.GET("/beef/summary", handler.BeefSummary)

	//start grpcserver
	grpcServer, lis, err := grpc.StartGRPCServer()
	if err != nil {
		fmt.Printf("Error starting gRPC server: %v\n", err)
		return
	}

	go func() {
		if err := e.Start(":" + httpPort); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the Echo server")
		}
	}()

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Printf("Failed to serve gRPC server: %v\n", err)
		}
	}()

	//gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	grpcServer.GracefulStop()
}
