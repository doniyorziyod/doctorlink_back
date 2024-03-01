package main

import (
	"doctorlink/database"
    "doctorlink/server"
	"log"
)

func main() {
    db, err := database.Connect()
    if err != nil {
        log.Panic(err)
    }

    repo := database.NewPostgresRepository(db)
    router := server.NewRouter(repo)

    router.Logger.Fatal(router.Start(":1323"))
}
