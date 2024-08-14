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
	e.POST("/message", handle.Message)

	e.POST("/name", handle.Name)
	e.POST("/name/letter", handle.Letter)

	go e.Start(":8080")

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("GF")

	err := storeDB.Db.Close()
	if err != nil {
		log.Println("Ошибка закрытия БД")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Println("ОШИБКА - остановка сервера")
	}

}
