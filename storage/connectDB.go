package storage

import (
	"Psql/model"
	"log"

	"github.com/jmoiron/sqlx"
)

type StorageConnect struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) StorageConnect {
	return StorageConnect{
		Db: db,
	}
}

func ConnectDB() *sqlx.DB {

	connStr := "user=onek dbname=pDB sslmode=disable password=123"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Println("Ошибка подключения к БД")
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (s StorageConnect) SetName(fio model.FIO) (model.FIO, error) {
	err := s.Db.QueryRow(
		"INSERT INTO users (first_name, last_name) VALUES ($1,$2) RETURNING id",
		fio.FirstName, fio.LastName,
	).Scan(&fio.UserID)
	if err != nil {
		log.Println("ошибка вставки в БД")
		fio.UserID = -1
		return fio, err
	}
	log.Println(fio)

	return fio, nil
}

func (s StorageConnect) AddLetter(letter model.Letter) (string, error) {

	_, err := s.Db.Exec(`
	INSERT INTO letters (user_id, item, letter)
	VALUES ($1, $2, $3)
`, letter.UserID, letter.Item, letter.Letter)

	if err != nil {
		log.Println("Error executing SQL query:", err)
		log.Println("Ошибка при вставке письма в базу данных")
		return "Не удалось вставить письмо в БД", err
	}
	return "Успешно добавлено письмо", nil
}

func (s StorageConnect) AddMessage(message model.Message) (string, error) {

	_, err := s.Db.Exec("INSERT INTO messages (user_id, mess) VALUES ($1, $2)",
		message.UserID, message.Message)

	if err != nil {
		log.Println("Error executing SQL query:", err)
		log.Println("Ошибка при вставке сообщения в базу данных")
		return "Не удалось сохранить сообщение", err
	}
	return "Сообщение успешно добавлено", nil
}

func (s StorageConnect) GetMessage(message model.AnswMessage) ([]model.AnswMessage, error) {

	rows, err := s.Db.Query("SELECT mess FROM messages WHERE user_id = $1", message.UserID)
	if err != nil {
		log.Println("Error executing SQL query:", err)
		log.Println("Ошибка при получении сообщений из базы данных")
		return nil, err
	}
	defer rows.Close()

	var Allmessages []model.AnswMessage

	for rows.Next() {
		var message model.AnswMessage
		err := rows.Scan(&message.Answer)
		if err != nil {
			log.Println("Ошибка при сканировании строки:", err)
			return nil, err
		}
		Allmessages = append(Allmessages, message)
	}

	return Allmessages, nil

}

func (s StorageConnect) GetLetter(letterServ model.Letter) ([]model.Letter, error) {

	rows, err := s.Db.Query("SELECT item, letter FROM letters WHERE user_id = $1", letterServ.UserID)
	if err != nil {
		log.Println("Error executing SQL query:", err)
		log.Println("Ошибка при получении сообщений из базы данных")
		return nil, err
	}
	defer rows.Close()

	var AllLetters []model.Letter
	for rows.Next() {
		var letter model.Letter
		err = rows.Scan(&letter.Item, &letter.Letter)
		if err != nil {
			log.Println("Ошибка при сканировании строки:", err)
			return nil, err
		}
		AllLetters = append(AllLetters, letter)
	}

	return AllLetters, nil

}
