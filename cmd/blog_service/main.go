package main

import (
	"Blogs/config"
	"Blogs/http"
	"Blogs/http/handlers"
	"Blogs/http/middleware"
	"Blogs/repositories"
	"Blogs/service"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func main() {
	// Load configuration from config.yml
	cfg := config.LoadConfig()

	// Initialize MongoDB client with URI from config
	clientOptions := options.Client().ApplyURI(cfg.Database.URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ensure MongoDB connection is established
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
	defer client.Disconnect(context.TODO())

	fmt.Println("Connected to MongoDB successfully!")

	// Initialize logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(client, cfg.Database.Name, "Users")
	postRepo := repositories.NewPostRepository(client, cfg.Database.Name, "Posts")
	commentRepo := repositories.NewCommentRepository(client, cfg.Database.Name, "Comments")

	// Initialize services
	userService := service.NewUserService(userRepo)
	postService := service.NewpostService(postRepo)
	commentService := service.NewCommentService(*commentRepo)
	// Initialize handlers
	userHandler := handlers.NewUserHandlers(userService)
	postHandler := handlers.NewpostHandlers(postService)
	commentHandler := handlers.NewCommentHandlers(commentService)
	// middleware
	middleware:=middleware.CreateJwtToken(repo)

	// Create the HTTP server instance
	server := http.NewServer(cfg, logger, userHandler, postHandler, commentHandler)

	// Start the server
	addr := cfg.Server.Port
	fmt.Printf("Server started at %s\n", addr)
	if err := server.Listen(context.TODO(), addr); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
