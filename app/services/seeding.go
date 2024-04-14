package services



import (

	"jhart.dev/ctc-coding-challenge-app/models"
)

// Seeds test users
func (p *Provider) SeedDB() error{
  mockUsers := []models.User{
    {
      Email:    "someuser1@gmail.com",
      Password: "password",
    },
    {
      Email:    "someuser2@gmail.com",
      Password: "password",
    },
    {
      Email:    "someuser3@gmail.com",
      Password: "password",
    },
  }

  for _, user := range mockUsers {
    _,err := p.CreateUser(user)

    if err != nil {
      return err
    }
  }

  return nil

}
