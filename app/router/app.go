package router

import (
	"github.com/gofiber/fiber/v2"
	"jhart.dev/ctc-coding-challenge-app/services"
)


const STATUS_SUCCESS = "success"
const STATUS_FAILURE = "failure"
const STATUS_MISSING_VALUES = "missing-values"

type APIMessage struct {
  Status string `json:"status"`
  Message string `json:"message"`
  Data any `json:"data"`
}


func NewApp(sp *services.Provider) *fiber.App {
  app := fiber.New()

  app.Static("/","frontend/dist")

  app.Route("/api/user/",UserAPI(sp))

  return app
}
