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
	notifClient := client.NewNotificationServiceClient(cfg)

	authHandler := &handler.AuthHandler{Client: authClient, Cfg: cfg}
	notifHandler := &handler.NotificationHandler{Client: notifClient, Auth: authClient, Cfg: cfg}

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/register", authHandler.Register)
	apiMux.HandleFunc("/login", authHandler.Login)
	apiMux.Handle("/profile", middleware.AuthMiddleware(cfg, http.HandlerFunc(authHandler.GetProfile)))

	apiMux.Handle("/notifications/subscribe", middleware.AuthMiddleware(cfg, http.HandlerFunc(notifHandler.Subscribe)))
	apiMux.Handle("/notifications/unsubscribe", middleware.AuthMiddleware(cfg, http.HandlerFunc(notifHandler.Unsubscribe)))
	apiMux.Handle("/notifications/create", middleware.AuthMiddleware(cfg, http.HandlerFunc(notifHandler.CreateNotification)))
	apiMux.Handle("/notifications/history", middleware.AuthMiddleware(cfg, http.HandlerFunc(notifHandler.GetHistory)))

	mainMux := http.NewServeMux()
	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	log.Println("API Gateway running on :8080")
	if err := http.ListenAndServe(":8080", mainMux); err != nil {
		log.Fatal(err)
	}
}
