package service

import (
	"Psql/model"
)

type Service struct {
	Storage StoragerInterface
}

type StoragerInterface interface {
	SetName(fio model.FIO) (model.FIO, error)
	AddLetter(letter model.Letter) (string, error)
	AddMessage(message model.Message) (string, error)
	GetMessage(message model.AnswMessage) ([]model.AnswMessage, error)
	GetLetter(letterServ model.Letter) ([]model.Letter, error)
}

func New(s StoragerInterface) Service {
	return Service{
		Storage: s,
	}
}

func (s Service) SetNameServ(fio model.FIO) (int, error) {
	// 1.Здесь все проверки
	// 2.Здесь Бизнес-логика
	// var err error
	fio, err := s.Storage.SetName(fio) // уже другое fio
	if err != nil {
		return -1, err
	}

	return fio.UserID, nil
}

func (s Service) AddLetterServ(letter model.Letter) (string, error) {

	letterStr, err := s.Storage.AddLetter(letter)
	if err != nil {
		return "ошибка в сервисном уровне", err
	}
	return letterStr, nil
}

func (s Service) AddMessageServ(messageFromHand model.Message) (string, error) {

	message, err := s.Storage.AddMessage(messageFromHand)
	if err != nil {
		return "ошибка в сервисном уровне", err
	}

	return message, nil
}

func (s Service) GetMessageServ(messageFromHand model.AnswMessage) (model.Response, error) {

	arrMessages, err := s.Storage.GetMessage(messageFromHand)
	if err != nil {
		return model.Response{}, err
	}

	var fullansw []string

	if messageFromHand.Amount != 0 {
		for i := 0; i < messageFromHand.Amount; i++ {
			fullansw = append(fullansw, arrMessages[i].Answer)
		}
	} else {

		for _, m := range arrMessages {
			fullansw = append(fullansw, m.Answer)
		}
	}

	responce := model.Response{
		UserID:        messageFromHand.UserID,
		TotalMessages: len(arrMessages),
		Messages:      fullansw,
	}

	return responce, nil
}

func (s Service) GetLetterServ(letterFromHand model.Letter) (string, error) {

	s.Storage.GetLetter(letterFromHand)

	return "", nil // исправить "" и nil

}
