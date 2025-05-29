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

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/register", authHandler.Register)
	apiMux.HandleFunc("/login", authHandler.Login)
	apiMux.Handle("/profile", middleware.AuthMiddleware(cfg, http.HandlerFunc(authHandler.GetProfile)))
	apiMux.Handle("/update-profile", middleware.AuthMiddleware(cfg, http.HandlerFunc(authHandler.UpdateProfile)))

	reportClient := client.NewReportServiceClient(cfg)
	reportHandler := &handler.ReportHandler{Client: reportClient}

	apiMux.Handle("/report/create", middleware.AuthMiddleware(cfg, http.HandlerFunc(reportHandler.CreateReport)))
	apiMux.Handle("/report/user", middleware.AuthMiddleware(cfg, http.HandlerFunc(reportHandler.GetReportsByUser)))
	apiMux.Handle("/report/edit", middleware.AuthMiddleware(cfg, http.HandlerFunc(reportHandler.EditReport)))

	mainMux := http.NewServeMux()

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	mainMux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend-main/login.html")
	})

	fs := http.FileServer(http.Dir("../frontend-main"))
	mainMux.Handle("/", fs)

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
