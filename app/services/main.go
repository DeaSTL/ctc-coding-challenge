package services

import (

	"jhart.dev/ctc-coding-challenge-app/db"
)

type Provider struct{
  q *db.Queries
}

func NewProvider(q *db.Queries) Provider{
  return Provider{
    q: q,
  }
}
