package main

import (
	"log"
	"math/rand"
	"time"

	tele "gopkg.in/telebot.v3"
)

type Server interface {
    generateCode() int
}

func main() {
    hello := "👋"
    world := " Assalomu alaykum! Doctorlink platformasiga xush kelibsiz"
	pref := tele.Settings{
		Token:  "7098817542:AAGy7oqxpvfpjQ849AuGFgb5Kg6SnIt6WtE",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer b.Start()
    b.Handle("/start", func(c tele.Context) error {
		return c.Send(hello + world)
	})
}

func generate() int {
    rand.New(rand.NewSource(time.Now().UnixNano()))
    return rand.Intn(90000) + 10000
}
