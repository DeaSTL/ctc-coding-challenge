package db

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func runMigrations(migrationFile string,q *Queries, ctx *context.Context) (error) {
  fileContents, err := os.ReadFile(migrationFile)
  if err != nil {
    return err
  } 

  _, err = q.db.Exec(*ctx,string(fileContents))

  if err != nil {
    return err
  }

  return nil
}


func NewDBConnection() (*Queries,error) {
  ctx := context.Background() 

  connectionString := os.Getenv("DB_URI")

  log.Printf("Connecting to %v", connectionString)

  if len(connectionString) == 0 {
    return nil,errors.New("DB_URL was empty")
  }

  connection, err := pgx.Connect(ctx,connectionString)
  if err != nil {
    log.Printf("Failed to connect to database we will re-attempt to establish a connection in 10s")
    time.Sleep(time.Second * 10)
    for i := 0; i < 4; i++ {
      log.Printf("re-attempt to connect to database attempt number %d/4 error: %+v", i,err.Error())
      connection, err = pgx.Connect(ctx,connectionString) 
      if err == nil {
        log.Println("Successfully connected to database")
        break
      }
      time.Sleep(time.Second * 10)
      log.Printf("Failed to connect to database %+v", err.Error())
    }
  }
  queries := New(connection)
  
  err = runMigrations("./db/schema.sql",queries,&ctx)

  if err != nil {
    return nil,err
  }

  return queries,nil
} 
