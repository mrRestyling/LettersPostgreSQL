package handlers

import (
	"Psql/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Service ServiceInterface
}

type ServiceInterface interface {
	SetNameServ(fio model.FIO) (int, error)
	AddLetterServ(letter model.Letter) (string, error)
	AddMessageServ(message model.Message) (string, error)
}

func New(s ServiceInterface) Handlers {
	return Handlers{
		Service: s,
	}
}

func (h Handlers) Name(c echo.Context) error {

	var fio model.FIO

	err := c.Bind(&fio)
	if err != nil {
		log.Println("Неверный запрос")
		return err
	}

	//
	numberFIO, err := h.Service.SetNameServ(fio)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, numberFIO)
}

func (h Handlers) Letter(c echo.Context) error {
	var text model.Letter
	err := c.Bind(&text)
	if err != nil {
		log.Println("Error binding JSON data:", err)
		log.Println("Неверный запрос")
		return err
	}

	responceTo, err := h.Service.AddLetterServ(text)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	//
	return c.JSON(http.StatusOK, responceTo)
}

func (h Handlers) Message(c echo.Context) error {

	var messageClient model.Message

	err := c.Bind(&messageClient)
	if err != nil {
		log.Println("Неверный запрос")
		return err
	}
	log.Println("Получено сообщение:", messageClient)

	messageClientStr, err := h.Service.AddMessageServ(messageClient)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, messageClientStr)
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Приветствую на сервере!")
}
