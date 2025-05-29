package main

import (
	"api-gateway-service/config"
	"api-gateway-service/internal/client"
	"api-gateway-service/internal/handler"
	"api-gateway-service/internal/middleware"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	authClient := client.NewAuthServiceClient(cfg)
	authHandler := &handler.AuthHandler{
		Client: authClient,
		Cfg:    cfg,
	}

	// API routes
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/register", authHandler.Register)
	apiMux.HandleFunc("/login", authHandler.Login)
	apiMux.Handle("/profile", middleware.AuthMiddleware(cfg, http.HandlerFunc(authHandler.GetProfile)))
	apiMux.Handle("/update-profile", middleware.AuthMiddleware(cfg, http.HandlerFunc(authHandler.UpdateProfile)))

	// Main router
	mainMux := http.NewServeMux()

	// Serve API under /api
	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	// Serve static frontend files
	fs := http.FileServer(http.Dir("../frontend-main")) // адаптируй путь под свою структуру
	mainMux.Handle("/", fs)

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mainMux); err != nil {
		log.Fatal(err)
	}
}
