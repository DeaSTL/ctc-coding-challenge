package router

import (
	"encoding/json"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"jhart.dev/ctc-coding-challenge-app/models"
	"jhart.dev/ctc-coding-challenge-app/services"
)

const _EXP_EMAIL = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$` 


type registerForm struct {
  Email string `json:"email"`
  Password string `json:"pasword"`
}

func sendMessage(c *fiber.Ctx,msg APIMessage) error{
  messageStr , err := json.Marshal(msg)

  if err != nil {
    return err
  }
  err = c.SendString(string(messageStr))

  if err != nil {
    return nil
  }
  return nil
}


func UserAPI(sp *services.Provider) func(fiber.Router){

  return func(r fiber.Router){
    r.Post("/register",func(c *fiber.Ctx) error{
      body := registerForm{}

      if err := c.BodyParser(&body); err != nil {
        return c.JSON(APIMessage{
        	Status:  STATUS_MISSING_VALUES,
        	Message: "The form submitted has invalid data",
        	Data:    nil,
        })
      }

      emailExp,err := regexp.Compile(_EXP_EMAIL)

      if err != nil {
        return err
      }

      if !emailExp.MatchString(body.Email) {
        return c.JSON(APIMessage{
        	Status:  STATUS_FAILURE,
        	Message: "Email format is invalid",
        	Data:    nil,
        }) 
      }


      user,err := sp.GetUserByEmail(body.Email)

      if err == nil {
        return c.JSON(APIMessage{
          Status: STATUS_FAILURE,
          Message: "User with that email already exists",
          Data: nil,
        })
      }

      user, err = sp.CreateUser(models.User{
        Email:body.Email,
        Password: body.Password,
      })

      return c.JSON(APIMessage{
      	Status:  STATUS_SUCCESS,
      	Message: "success",
      	Data: user,
      })
    })
    r.Get("/check-email",func(c *fiber.Ctx) error{
      emailValue := c.Query("value","")
      if len(emailValue) < 1{
        return c.JSON(APIMessage{
        	Status:  STATUS_MISSING_VALUES,
        	Message: "value not set",
        	Data:    nil,
        })
      }


      emailExp,err := regexp.Compile(_EXP_EMAIL)

      if err != nil {
        return err
      }

      if !emailExp.MatchString(emailValue) {
        return c.JSON(APIMessage{
        	Status:  STATUS_FAILURE,
        	Message: "Invalid email format",
        	Data:    nil,
        })
      }
      user,err := sp.GetUserByEmail(emailValue)

      if err == nil {
        return c.JSON(APIMessage{
          Status: STATUS_FAILURE,
          Message: "User with that email already exists",
          Data: nil,
        })
      }
      return c.JSON(APIMessage{
      	Status:  STATUS_SUCCESS,
      	Message: "success",
      	Data: user,
      })
    })

  }

}
