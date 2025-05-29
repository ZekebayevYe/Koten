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
	authHandler := &handler.AuthHandler{Client: authClient, Cfg: cfg}

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/register", authHandler.Register)
	apiMux.HandleFunc("/login", authHandler.Login)
	apiMux.Handle("/profile", middleware.AuthMiddleware(cfg, http.HandlerFunc(authHandler.GetProfile)))
	//apiMux.Handle("/update-profile", middleware.AuthMiddleware(cfg, http.HandlerFunc(authHandler.UpdateProfile)))

	reportClient := client.NewReportServiceClient(cfg)
	reportHandler := &handler.ReportHandler{Client: reportClient}

	apiMux.Handle("/report/create", middleware.AuthMiddleware(cfg, http.HandlerFunc(reportHandler.CreateReport)))
	apiMux.Handle("/report/user", middleware.AuthMiddleware(cfg, http.HandlerFunc(reportHandler.GetReportsByUser)))
	apiMux.Handle("/report/edit", middleware.AuthMiddleware(cfg, http.HandlerFunc(reportHandler.EditReport)))

	mainMux := http.NewServeMux()
	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	log.Println("API Gateway running on :8080")
	if err := http.ListenAndServe(":8080", mainMux); err != nil {
		log.Fatal(err)
	}
}
