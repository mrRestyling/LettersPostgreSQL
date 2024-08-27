package main

import (
	"Psql/handlers"
	"Psql/service"
	"Psql/storage"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

func main() {

	db := storage.ConnectDB()
	storeDB := storage.New(db)

	serv := service.New(storeDB)
	handle := handlers.New(serv)

	e := echo.New()

	e.GET("/", handlers.Hello)
	e.POST("/name", handle.Name)

	e.GET("/message", handle.MessageReturn)
	e.POST("/message", handle.Message)

	e.GET("/letter", handle.LetterReturn)
	e.POST("/letter", handle.Letter)

	go e.Start(":8080")

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("GF")

	err := storeDB.Db.Close()
	if err != nil {
		log.Println("Ошибка закрытия БД")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// context.WithDeadline()
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Println("ОШИБКА - остановка сервера")
	}

}
