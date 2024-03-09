package main

import (
	"doctorlink/bot"
	"doctorlink/database"
	"doctorlink/server"
	"fmt"
	"log"
	"os"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(os.Getenv("BOT_TOKEN"))

	repo := database.NewPostgresRepository(db)
	router := server.NewRouter(repo)
	bot := bot.NewBot(repo)
	bot.Start()
	router.Logger.Fatal(router.Start(":1323"))
}
