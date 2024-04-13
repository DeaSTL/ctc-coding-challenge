package router

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"jhart.dev/ctc-coding-challenge-app/services"
)

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
    r.Get("/check-email",func(c *fiber.Ctx) error{
      emailValue := c.Query("value","")
      if len(emailValue) < 1{
        err := sendMessage(c,APIMessage{
        	Status:  STATUS_MISSING_VALUES,
        	Message: "value not set",
        	Data:    nil,
        })

        if err != nil {
          return err
        }
        return err
      }
      user,err := sp.GetUserByEmail(emailValue)


      log.Printf("user: %+v", user)



      if err == nil {
        err := sendMessage(c,APIMessage{
          Status: STATUS_FAILURE,
          Message: "User with that email already exists",
          Data: nil,
        })
        return err
      }

      sendMessage(c,APIMessage{
      	Status:  STATUS_SUCCESS,
      	Message: "success",
      	Data: user,
      })

      return nil 
    })

  }

}
