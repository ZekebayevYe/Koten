package main

import (
	"api-gateway-service/config"
	"api-gateway-service/internal/client"
	"api-gateway-service/internal/handler"
	"api-gateway-service/internal/middleware"
	"log"
	"net/http"

	"github.com/rs/cors"
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

	// Serve login page at /login
	mainMux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend-main/login.html")
	})

	// Serve other static files
	fs := http.FileServer(http.Dir("../frontend-main"))
	mainMux.Handle("/", fs)

	// Add CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(mainMux)

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		log.Fatal(err)
	}
}