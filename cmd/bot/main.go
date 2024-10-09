// VERSION 1
// package main

// import (
// 	"log" // Импортируем пакет для логирования ошибок и информации

// 	// Импортируем библиотеку для работы с API Telegram
// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// func main() {
// 	// Аргумент 0 в вызове tgbotapi.NewUpdate(0) обозначает номер последнего обновления, которое бот будет запрашивать.
// 	// Создаем новый экземпляр бота с указанным токеном
// 	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
// 	if err != nil {
// 		// Если произошла ошибка при создании бота, логируем ошибку и завершаем программу
// 		log.Panic(err)
// 	}

// 	// Включаем режим отладки для получения более подробной информации в логах
// 	bot.Debug = true

// 	// Логируем имя аккаунта бота для подтверждения авторизации
// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	// Создаем новую конфигурацию обновлений с тайм-аутом 60 секунд
// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	// Получаем канал для получения обновлений от Telegram
// 	updates:= bot.GetUpdatesChan(u)

// 	// Цикл для обработки входящих обновлений
// 	for update := range updates {
// 		// Проверяем, содержит ли обновление сообщение
// 		if update.Message != nil { // Если мы получили сообщение
// 			// Логируем имя отправителя и текст сообщения
// 			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

// 			// Создаем новое сообщение для отправки в чат
// 			//------------------------(в какой чать отправим , текст сообщение)
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote: "+update.Message.Text)
// 			// Устанавливаем сообщение как ответ на оригинальное сообщение
// 			// msg.ReplyToMessageID = update.Message.MessageID//подиетим ее как raplay к тому сообщению что тебе написали

//				// Отправляем ответное сообщение обратно в чат
//				bot.Send(msg)
//			}
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// VERSION 2
// package main

// import (
// 	"log"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// const token="8124804402:AAExuhpCMJeEpprkxrRi2boWImCeX9LxbRs"

// func main() {
// 	bot, err := tgbotapi.NewBotAPI(token)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	bot.Debug = true

// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60
// 	// u:=tgbotapi.UpdateConfig{
// 	// 	Timeout: 60,
// 	// }

// 	updates := bot.GetUpdatesChan(u)
// 	for update := range updates {
// 		if update.Message != nil { // If we got a message
// 			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

// 			// msg.ReplyToMessageID = update.Message.MessageID
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote: "+update.Message.Text)

//				bot.Send(msg)
//			}
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// VERSION 3
// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	token := os.Getenv("TOKEN") // откроем командную строку, намбираем команду: TOKEN=8124804402:AAExuhpCMJeEpprkxrRi2boWImCeX9LxbRs так мы задаем токен
// 	fmt.Println("Bot token:", token)
// 	bot, err := tgbotapi.NewBotAPI(token)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	bot.Debug = true

// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60
// 	// u:=tgbotapi.UpdateConfig{
// 	// 	Timeout: 60,
// 	// }

// 	updates := bot.GetUpdatesChan(u)
// 	for update := range updates {
// 		if update.Message != nil { // If we got a message
// 			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

// 			// msg.ReplyToMessageID = update.Message.MessageID
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote: "+update.Message.Text)

//				bot.Send(msg)
//			}
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// VERSION 4
// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	token := os.Getenv("TOKEN")
// 	fmt.Println("Bot token:", token)
// 	bot, err := tgbotapi.NewBotAPI(token)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	bot.Debug = true

// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates := bot.GetUpdatesChan(u)
// 	for update := range updates {
// 		if update.Message == nil {
// 			continue
// 		}

// 		if update.Message.Command() == "help" {
// 			helpCommand(bot, update.Message)
// 			continue
// 		}
// 		defaultBehavior(bot, update.Message)
// 	}
// }

// func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
// 	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help")
// 	bot.Send(msg)
// }

// func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
// 	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
// 	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
// 	bot.Send(msg)
// }
// -------------------------------------------------------------------------------------------------
// VERSION 5 с комментарием
// package main

// import (
// 	"fmt" // Импортируем пакет для работы с форматированным вводом-выводом
// 	"log"  // Импортируем пакет для логирования ошибок
// 	"os"   // Импортируем пакет для работы с окружением

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5" // Импортируем библиотеку для работы с Telegram API
// 	"github.com/joho/godotenv" // Импортируем библиотеку для загрузки переменных окружения из .env файла
// )

// func main() {
// 	// Загружаем переменные окружения из файла .env
// 	err := godotenv.Load()
// 	if err != nil {
// 		// Если произошла ошибка при загрузке, выводим сообщение и завершаем программу
// 		log.Fatal("Error loading .env file")
// 	}

// 	// Получаем токен бота из переменных окружения
// 	token := os.Getenv("TOKEN")
// 	fmt.Println("Bot token:", token) // Выводим токен в консоль
// 	// Создаем нового бота API с использованием токена
// 	bot, err := tgbotapi.NewBotAPI(token)
// 	if err != nil {
// 		// Если произошла ошибка при создании бота, выводим сообщение и завершаем программу
// 		log.Panic(err)
// 	}

// 	bot.Debug = true // Включаем режим отладки

// 	// Логируем имя аккаунта, под которым авторизован бот
// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	u := tgbotapi.NewUpdate(0) // Создаем новый объект обновления с нулевым ID
// 	u.Timeout = 60 // Устанавливаем тайм-аут в 60 секунд для получения обновлений

// 	// Получаем канал обновлений от бота
// 	updates := bot.GetUpdatesChan(u)
// 	for update := range updates {
// 		if update.Message == nil { // Проверяем, есть ли сообщение в обновлении
// 			continue // Если сообщения нет, переходим к следующему обновлению / переходим к следующему итерацию
// 		}

// 		// Если пользователь отправил команду "help", обрабатываем ее
// 		if update.Message.Command() == "help" {
// 			helpCommand(bot, update.Message) // Вызываем функцию для обработки команды help
// 			continue // Переходим к следующему обновлению / переходим к следующему итерацию
// 		}
// 		defaultBehavior(bot, update.Message) // Обрабатываем все остальные сообщения по умолчанию
// 	}
// }

// // Функция для обработки команды help
// func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
// 	// Создаем новое сообщение с текстом помощи
// 	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help")
// 	bot.Send(msg) // Отправляем сообщение пользователю
// }

// // Функция для обработки всех остальных сообщений
// func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
// 	// Логируем имя пользователя и текст сообщения
// 	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
// 	// Создаем новое сообщение с текстом, который отправил пользователь
// 	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
// 	bot.Send(msg) // Отправляем сообщение обратно пользователю
// }
// -------------------------------------------------------------------------------------------------
// VERSION 6 отрефакторим, используем конструкция switch
package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")
	fmt.Println("Bot token:", token)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Command(){
		case "help":
			helpCommand(bot, update.Message)
		default:
			defaultBehavior(bot, update.Message)
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help")
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	bot.Send(msg)
}
