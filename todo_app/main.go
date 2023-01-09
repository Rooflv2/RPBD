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

var name = os.Getenv("NAME")
var priority = os.Getenv("PRIORITY")
var timer = os.Getenv("TIMER")

var dbInfo = fmt.Sprintf("name=%s priority=%s timer=%s", name, priority, timer)

// Создаем таблицу tasks в БД при подключении к ней
func createTable() error {
	// Подключаемся к БД
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	// Создаем таблицу tasks
	if _, err = db.Exec(`CREATE TABLE "tasks" (
		"id" serial primary key,
		"name_task" VARCHAR NOT NULL,
		"priority" INTEGER NOT NULL,
		"timer" date
		);`); err != nil {
		return err
	}
	return nil
}

// Добавляем данные в БД при подключении к ней
// func AddData(name_task string, priority string, timer string) error {
// 	// Подключаемся к БД
// 	db, err := sql.Open("postgres", dbInfo)
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	// Добавляем данные
// 	data := `INSERT INTO tasks (name_task, priority, timer)
// 	VALUES (name_task, priority, timer);`
// 	// Выполняем наш SQL запрос
// 	if _, err = db.Exec(data, `@`+name_task, priority, timer); err != nil {
// 		return err
// 	}

// 	return nil
// }

// Обновляем данные в БД при подключении к ней
// func updateData(id int, name_task string, priority string, timer string, complete_task string) error {
// 	// Подключаемся к БД
// 	db, err := sql.Open("postgres", dbInfo)
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	// Добавляем данные
// 	data := `UPDATE tasks
// 	SET name_task = COALESCE('данные', name_task),
// 	SET priority = COALESCE('данные', priority),
// 	SET timer = COALESCE('данные', timer),
// 	SET complete_task = COALESCE('данные', complete_task)
// 	WHERE id = $id`
// 	// Выполняем наш SQL запрос
// 	if _, err = db.Exec(data, `@`+name_task, priority, timer, complete_task); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func viewactData(name_task string, priority string, timer string, complete_task string) {
// 	// Подключаемся к БД
// 	db, err := sql.Open("postgres", dbInfo)
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	// Получаем данные
// 	data := `SELECT * FROM tasks WHERE complete_task = 'no'`
//	// Выполняем наш SQL запрос
// 	if _, err = db.Exec(data, `@`+name_task, priority, timer, complete_task); err != nil {
// 		return err
// 	}
//	// ??????????
// 	// Отправлем сообщение
// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, data)
// 	if _, err := bot.Send(msg); err != nil {
// 		log.Panic(err)
// 	}
// 	continue

// 	return
// }

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
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, I'm a TODO bot, I can help you with controlling your tasks. Let's start!")
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
			msg.Text = "I can add new task. To do this, enter the command /new_task \n I can do your task with you ^_^. To do this, enter the command /take_task \n I can match your task complete. To do this, enter the command /complete_task \n I can show your actual tasks. To do this, enter the command /viewactual_tasks \n I can show your complete tasks. To do this, enter the command /viewcomplete_tasks \n I can show your statistic. To do this, enter the command /statistics"
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		case "new_task":
			msg.Text = "Write name, priority (number), time of your task"
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			// if !update.Message.IsCommand() {
			// 	data := update.Message.From.String()
			// 	AddData(data, data, data)
			// }
			continue
		case "take_task":
			msg.Text = "I need to do smth take_task"
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			// if !update.Message.IsCommand() {
			// 	id := update.Message.From.ID
			// 	data := update.Message.From.String()
			// 	updateData(int(id), data, data, data, data)
			// }
			continue
		case "complete_task":
			msg.Text = "I need to do smth complete_task"
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			// if !update.Message.IsCommand() {
			// 	id := update.Message.From.ID
			// 	data := update.Message.From.String()
			// 	updateData(int(id), data, data, data, data)
			// }
			continue
		case "viewactual_tasks":
			msg.Text = "I need to do smth viewactual_tasks"
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		case "viewcomplete_tasks":
			msg.Text = "I need to do smth viewcomplete_tasks"
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		case "statistics":
			msg.Text = "I need to do smth statistics"
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

func main() {
	time.Sleep(30 * time.Second)
	// Создаем таблицу
	if os.Getenv("CREATE_TABLE") == "yes" {
		if os.Getenv("DB_SWITCH") == "on" {
			if err := createTable(); err != nil {
				panic(err)
			}
		}
	}
	time.Sleep(30 * time.Second)
	// Вызываем бота
	telegramBot()
}
