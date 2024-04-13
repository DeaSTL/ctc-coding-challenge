package db

import (
	"context"
	"errors"
	"log"
	"os"

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
    connectionError := errors.New("Could not connect to database")
    return nil, errors.Join(err, connectionError)
  }
  queries := New(connection)
  
  err = runMigrations("./db/schema.sql",queries,&ctx)

  if err != nil {
    return nil,err
  }

  return queries,nil
} 
