package router

import (
	"encoding/json"
	"log"
	"os"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
	"jhart.dev/ctc-coding-challenge-app/models"
	"jhart.dev/ctc-coding-challenge-app/services"
)



type registerForm struct {
  Email string `json:"email"`
  Password string `json:"pasword"`
}

type loginForm struct {
  Email string `json:"email"`
  Password string `json:"password"`
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
    jwtSecret := os.Getenv("JWT_SECRET")

    if jwtSecret == "" {
      log.Fatal("Missing JWT_SECRET")
    }
    r.Post("/login",func(c *fiber.Ctx) error {
      body := loginForm{}
      if err := c.BodyParser(&body); err != nil {
        log.Printf("Error parsing body in login: %+v", err.Error())
        return c.JSON(APIMessage{
        	Status:  STATUS_MISSING_VALUES,
        	Message: "The form submitted has invalid data",
        	Data:    nil,
        })
      }
      
      user,err := sp.ValidateUserCredentials(body.Email,body.Password)
      log.Printf("User: %+v", user)

      if err == nil {
        claims := jwt.MapClaims{
          "email":  user.Email,
          "exp":   time.Now().Add(time.Hour * 72).Unix(),
        } 
        token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

        t, err := token.SignedString([]byte(jwtSecret))

        if err != nil {
          return c.SendStatus(fiber.StatusInternalServerError)
        }

        // Yeah... should've made it http only
        cookie := new(fiber.Cookie)
        cookie.Name = "auth_token"
        cookie.Value = t

        c.Cookie(cookie)

        return c.JSON(APIMessage{
          Status:  STATUS_SUCCESS,
          Message: "You have successfully logged in!",
          Data:    t,
        })
      }
      log.Printf("Error: %+v", err.Error())

      switch err.(type) {
        case services.UserDoesNotExistError:
          return c.JSON(APIMessage{
            Status:  STATUS_FAILURE,
            Message: "That user does not exist",
            Data:    nil,
          })
        case services.UserInvalidPasswordError:
          return c.JSON(APIMessage{
            Status:  STATUS_FAILURE,
            Message: "Invalid password",
            Data:    nil,
          })
        default:
          log.Printf("Error: %v", err.Error())
          return c.JSON(APIMessage{
            Status:  STATUS_FAILURE,
            Message: "Unknown Error",
            Data:    nil,
          })
      }
    })
    r.Post("/register",func(c *fiber.Ctx) error{
      body := registerForm{}

      testModeParam := c.Query("test") //Indicates that the user will be deleted after creation

      testMode := testModeParam != ""

      if err := c.BodyParser(&body); err != nil {
        return c.JSON(APIMessage{
        	Status:  STATUS_MISSING_VALUES,
        	Message: "The form submitted has invalid data",
        	Data:    nil,
        })
      }

      user, err := sp.CreateUser(models.User{
        Email:body.Email,
        Password: body.Password,
      })

      if err != nil {
        statusMessage := ""
        switch err.(type){
          case services.UserEmailInvalid:
            statusMessage = "Email entered is invalid"
            break
          case services.UserExistsError: 
            statusMessage = "User already exists with that email"
            break 
          default:
            statusMessage = "Unknown error"
            log.Printf("Error creating user with: %+v err: %+v", body,err.Error())
        }
        return c.JSON(APIMessage{
          Message: statusMessage,
          Status: STATUS_FAILURE,
          Data: nil,
        })
      }
      if testMode {
        err := sp.DeleteUser(user.ID)

        if err != nil {
          statusMessage := ""
          switch err.(type) {
            case services.UserDoesNotExistError:
              statusMessage = "User you are attempting to delete doesn't exist"
            default:
              statusMessage = "Unknown error"
              log.Printf("Error deleting user with: %+v err: %+v", body,err.Error())
          }
          return c.JSON(APIMessage{
            Message: statusMessage,
            Status: STATUS_FAILURE,
            Data: nil,
          })
        }
      }

      return c.JSON(APIMessage{
      	Status:  STATUS_SUCCESS,
      	Message: "success",
      	Data: user,
      })
    })
    r.Get("/check-email",func(c *fiber.Ctx) error{
      emailValue := c.Query("value","")
      if emailValue == ""{
        return c.JSON(APIMessage{
        	Status:  STATUS_MISSING_VALUES,
        	Message: "value not set",
        	Data:    nil,
        })
      }

      emailCheck := sp.CheckEmail(emailValue)

      if emailCheck != nil {
        statusMessage := ""
        switch emailCheck.(type){
          case services.UserEmailInvalid:
            statusMessage = "Email entered is invalid"
            break
          case services.UserExistsError: 
            statusMessage = "User already exists with that email"
            break 
          default:
            statusMessage = "Unknown error"
            log.Printf("Error checking email: %+v", emailCheck)
        }
        return c.JSON(APIMessage{
          Message: statusMessage,
          Status: STATUS_FAILURE,
          Data: nil,
        })
      }

      return c.JSON(APIMessage{
      	Status:  STATUS_SUCCESS,
      	Message: "success",
      	Data: nil,
      })
    })

    r.Use(jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
    }))

    r.Get("/",func(c *fiber.Ctx) error{
      if c.Locals("user") != nil{
        log.Printf("user: %+v", c.Locals("user"))
        user := c.Locals("user").(*jwt.Token)
        claims := user.Claims.(jwt.MapClaims)
        return c.JSON(APIMessage{
          Status: STATUS_SUCCESS,
          Message: "Successfully retrived user",
          Data: claims,
        })
      }else {
        return c.JSON(APIMessage{
          Status: STATUS_FAILURE,
          Message: "Could not retrive user",
          Data: nil,
        })
      }
    })

  }

}
