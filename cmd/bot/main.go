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
package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const token="8124804402:AAExuhpCMJeEpprkxrRi2boWImCeX9LxbRs"

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	// u:=tgbotapi.UpdateConfig{
	// 	Timeout: 60,
	// }
	
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// msg.ReplyToMessageID = update.Message.MessageID
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote: "+update.Message.Text)

			bot.Send(msg)
		}
	}
}
