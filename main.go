// VERSION 1
// package main

// import (
// 	"fmt"
// 	"os"
// )

//	func main() {
//		token:=os.Getenv("TOKEN")
//		fmt.Println(token)
//	}
//
// -------------------------------------------------------------------------------------------------
// VERSION 2
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err:=godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	token := os.Getenv("TOKEN")
    fmt.Println("Bot token:", token) // Вывод токена
}
