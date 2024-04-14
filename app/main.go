package main

import (
	"errors"
	"log"

	"jhart.dev/ctc-coding-challenge-app/db"
	"jhart.dev/ctc-coding-challenge-app/router"
	"jhart.dev/ctc-coding-challenge-app/services"
)


func main()  {
  log.Printf("Attempting to start go server")


  queries,err := db.NewDBConnection()


  if err != nil {
    log.Panic(err.Error())
  }

  serviceProvider := services.NewProvider(queries)
  err = serviceProvider.SeedDB()

  app := router.NewApp(&serviceProvider)

  if err != nil {
    log.Print(errors.Join(err,errors.New("Failed to seed database")).Error())
  }

  log.Fatal(app.Listen(":3000"))
}
