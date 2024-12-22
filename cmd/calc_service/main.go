package main

import (
	"log"
	"github.com/gusev1alex/Calculator_API.git/internal/server"
)

func main() {
	// Создаём экземпляр сервера
	s := server.NewServer()

	// Адрес, на котором будет запущен сервер
	addr := ":8080"

	log.Printf("Starting server on %s", addr)

	// Запускаем сервер и обрабатываем возможные ошибки
	if err := s.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
