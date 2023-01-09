package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	TelegramBotToken string
}

// Добавление данных для БД
var user_name = os.Getenv("USER_NAME")
var user_surname = os.Getenv("USER_SURNAME")
var login = os.Getenv("LOGIN")
var password = os.Getenv("PASSWORD")
var mail = os.Getenv("MAIL")
var role = os.Getenv("ROLE")

// Соединение их в одну переменную
var dbUs = fmt.Sprintf("user_name=%s user_surname=%s login=%s password=%s mail=%s role=%s", user_name, user_surname, login, password, mail, role)

// Создаем таблицу tasks в БД при подключении к ней
func createTableusers() error {
	// Подключаемся к БД
	db, err := sql.Open("postgres", dbUs)
	if err != nil {
		return err
	}
	defer db.Close()

	// Создаем таблицу пользователей
	if _, err = db.Exec(`CREATE TABLE users(
		id SERIAL PRIMARY KEY,
		user_name VARCHAR(255),
		user_surname VARCHAR(255),
		login VARCHAR(255)	NOT NULL,
		password VARCHAR(255)	NOT NULL,
		mail VARCHAR(255),
		role VARCHAR(255)	NOT NULL
		);`); err != nil {
		return err
	}
	return nil
}

// Добавление данных для БД
var products_name = os.Getenv("PRODUCT_NAME")
var weight = os.Getenv("WEIGHT")
var amount = os.Getenv("AMOINT")
var volume = os.Getenv("VOLUME")
var from_who = os.Getenv("FROM_WHO")

// Соединение их в одну переменную
var dbProd = fmt.Sprintf("products_name=%s weight=%s amount=%s volume=%s from_who=%s ", products_name, weight, amount, volume, from_who)

// Создаем таблицу tasks в БД при подключении к ней
func createTableprod() error {
	// Подключаемся к БД
	db, err := sql.Open("postgres", dbProd)
	if err != nil {
		return err
	}
	defer db.Close()

	// Создаем таблицу пользователей
	if _, err = db.Exec(`CREATE TABLE products(
		id SERIAL PRIMARY KEY,
		products_name VARCHAR(255)	NOT NULL,
		weight VARCHAR(255),
		amount VARCHAR(255)	NOT NULL,
		volume VARCHAR(255),
		from_who VARCHAR(255)	NOT NULL
		);`); err != nil {
		return err
	}
	return nil
}

// Добавление данных для БД
var delivery_date = os.Getenv("DELIVERY_DATE")
var delivery_point = os.Getenv("DELIVERY_POINT")

// Соединение их в одну переменную
var dbDel = fmt.Sprintf("user_name=%s user_surname=%s products_name=%s amount=%s delivery_date=%s delivery_point=%s", user_name, user_surname, products_name, amount, delivery_date, delivery_point)

// Создаем таблицу tasks в БД при подключении к ней
func createTabledelivery() error {
	// Подключаемся к БД
	db, err := sql.Open("postgres", dbDel)
	if err != nil {
		return err
	}
	defer db.Close()

	// Создаем таблицу пользователей
	if _, err = db.Exec(`CREATE TABLE delivery(
		id SERIAL PRIMARY KEY,
		user_name VARCHAR(255) NOT NULL,
		user_surname VARCHAR(255) NOT NULL,
		products_name VARCHAR(255) NOT NULL,
		amount VARCHAR(255) NOT NULL,
		delivery_date VARCHAR(255) NOT NULL,
		delivery_point VARCHAR(255) NOT NULL
		);`); err != nil {
		return err
	}
	return nil
}

// user_name, user_surname, login, password, mail, role

func addUs(user_name string, user_surname string, login string, password string, mail string, role string) error {
	// Подключаемся к БД
	db, err := sql.Open("postgres", dbUs)
	if err != nil {
		return err
	}
	defer db.Close()

	// Добавляем данные
	data := `INSERT INTO users(user_name, user_surname, login, password, mail, role)
	VALUES ('Михаил', 'Баданянин', 'MishaBad', 'Misha123', 'mihael@mail.ru', 'Работник');`
	// Выполняем наш SQL запрос
	if _, err = db.Exec(data, `@`+user_name, user_surname, login, password, mail, role); err != nil {
		return err
	}

	return nil
}

func addProd(products_name string, weight string, amount string, volume string, from_who string) error {
	// Подключаемся к БД
	db, err := sql.Open("postgres", dbUs)
	if err != nil {
		return err
	}
	defer db.Close()

	// Добавляем данные
	data := `INSERT INTO products(product_name, weight, amount, volume, from_who)
	VALUES ('бананы', '2', '1000', NULL, 'Макар');`
	// Выполняем наш SQL запрос
	if _, err = db.Exec(data, `@`+products_name, weight, amount, volume, from_who); err != nil {
		return err
	}

	return nil
}

// Создаем таблицу tasks в БД при подключении к ней
func createViewnt() error {
	// Подключаемся к БД
	db, err := sql.Open("postgres", dbDel)
	if err != nil {
		return err
	}
	defer db.Close()

	// Создаем таблицу пользователей
	if _, err = db.Exec(`CREATE OR REPLACE VIEW new_tasks AS
	SELECT * FROM delivery`); err != nil {
		return err
	}
	return nil
}

// Добавление новой доставки
func addDel(user_name string, user_surname string, products_name string, amount string, delivery_date string, delivery_point string) error {
	// Подключаемся к БД
	db, err := sql.Open("postgres", dbDel)
	if err != nil {
		return err
	}
	defer db.Close()

	// Добавляем данные
	data := `INSERT INTO delivery(user_name, user_surname, products_name, amount, delivery_date, delivery_point)
	VALUES ('имя-пользователя', 'фамилия-пользователя', 'название-продукта', 'кол-во', 'дата', 'место');`
	// Выполняем наш SQL запрос
	if _, err = db.Exec(data, `@`+user_name, user_surname, products_name, amount, delivery_date, delivery_point); err != nil {
		return err
	}

	return nil
}

// Блокировка пользователя
func newTask(user_name string, user_surname string, login string, password string, mail string, role string) error {
	// Подключаемся к БД
	db, err := sql.Open("postgres", dbDel)
	if err != nil {
		return err
	}
	defer db.Close()

	data := `INSERT INTO new_tasks(amount, delivery_date, delivery_point, products_name, user_name, user_surname)
	VALUES ('100', '2022-12-31', 'г.Дубна', 'мороженное', 'Элеонора', 'Матроновна')`
	// Выполняем наш SQL запрос
	if _, err = db.Exec(data, `@`+user_name, user_surname, login, password, mail, role); err != nil {
		return err
	}

	return nil
}

// Блокировка пользователя
func banUs(user_name string, user_surname string, login string, password string, mail string, role string) error {
	// Подключаемся к БД
	db, err := sql.Open("postgres", dbUs)
	if err != nil {
		return err
	}
	defer db.Close()

	data := `UPDATE users
	SET role = COALESCE('заблокирован', role)
	WHERE id = 'номер'`
	// Выполняем наш SQL запрос
	if _, err = db.Exec(data, `@`+user_name, user_surname, login, password, mail, role); err != nil {
		return err
	}

	return nil
}

func updateDel(user_name string, user_surname string, products_name string, amount string, delivery_date string, delivery_point string) error {
	// Подключаемся к БД
	db, err := sql.Open("postgres", dbDel)
	if err != nil {
		return err
	}
	defer db.Close()

	data := `UPDATE delivery
	SET 
	user_name = COALESCE('имя', user_name),
	user_surname = COALESCE('фамилия', user_surname),
	products_name = COALESCE('продукт', products_name),
	amount = COALESCE('кол-во', amount),
	delivery_date = COALESCE('2022-12-31', delivery_date),
	delivery_point = COALESCE('адрес', delivery_point)
	WHERE id = 'номер'`
	// Выполняем наш SQL запрос
	if _, err = db.Exec(data, `@`+user_name, user_surname, products_name, amount, delivery_date, delivery_point); err != nil {
		return err
	}

	return nil
}

func telegramBot() {
	// Деводироавние ключа
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}

	// Подключение ключа
	bot, err := tgbotapi.NewBotAPI(os.Getenv("configuration.TelegramBotToken"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // Если пришло пустое сообщение
			// Отправлем сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Текст приветствия")
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		}

		if !update.Message.IsCommand() { // Если пришло сообщение не команда
			// Отправлем сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Maybe you need help? Type /help to get list of commands")
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		}

		// Создаём новое сообщение. У нас нет текста, поэтому оставляем его пустым.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Вытаскиваем команду из сообщения.
		switch update.Message.Command() {
		case "help":
			msg.Text = ""
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		default:
			msg.Text = "I don't know that command"
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		}

	}
}

var y = "yes"

func main() {
	time.Sleep(30 * time.Second)
	// Создаем таблицу
	if os.Getenv("createTableusers") == y {
		if os.Getenv("createTableprod") == y {
			if os.Getenv("createTabledelivery") == y {
				if os.Getenv("createViewnt") == y {
					if os.Getenv("DB_SWITCH") == "on" {
						if err := createTableusers(); err != nil {
							panic(err)
						}
						if err := createTableprod(); err != nil {
							panic(err)
						}
						if err := createTabledelivery(); err != nil {
							panic(err)
						}
						if err := createViewnt(); err != nil {
							panic(err)
						}
					}
				}
			}
		}
	}
	time.Sleep(30 * time.Second)
	// Вызываем бота
	telegramBot()
}
