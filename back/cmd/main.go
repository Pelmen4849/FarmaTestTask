package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"medical_farm/back/internal/config"
	"medical_farm/back/internal/db"
	"medical_farm/back/internal/handler"
	"medical_farm/back/internal/repository"
	"medical_farm/back/internal/router"
	"medical_farm/back/internal/service"
)

func main() {
	// Загрузка конфигурации
	cfg := config.Load()

	// Подключение к БД
	pool, err := db.NewPostgresPool(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()
	log.Println("Database connected successfully")

	// Инициализация репозиториев
	drugRepo := repository.NewDrugRepository(pool)
	orderRepo := repository.NewOrderRepository(pool)

	// Инициализация сервисов
	drugService := service.NewDrugService(drugRepo)
	orderService := service.NewOrderService(orderRepo, drugRepo)

	// Инициализация хендлеров
	drugHandler := handler.NewDrugHandler(drugService)
	orderHandler := handler.NewOrderHandler(orderService)

	// Настройка маршрутов
	r := router.SetupRouter(drugHandler, orderHandler)
	

	// Запуск HTTP-сервера с graceful shutdown
	srv := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.Timeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Timeout) * time.Second,
	}

	go func() {
		log.Printf("Server is running on port %s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exited")
}
