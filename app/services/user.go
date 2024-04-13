package services

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"jhart.dev/ctc-coding-challenge-app/db"
	"jhart.dev/ctc-coding-challenge-app/models"
)

type UserExistsError struct{
  Email string
}

func (ue UserExistsError) Error() string {
  return fmt.Sprintf("User %v alredy exists",ue.Email)
}

func dbUserToModel(user db.User) models.User {
  return models.User{
    ID: user.ID,
    Email: user.Email,
    Password: user.Password,
  }
}

func hashPassword(password string) (string,error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password),14)
  return string(bytes), err
}

func checkPassword(password string, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
  return err == nil
}

func (s *Provider) CreateUser(newUser models.User) (models.User, error)  {
  
  _, err := s.GetUserByEmail(newUser.Email)

  if err == nil {
    return models.User{},UserExistsError{Email:newUser.Email}
  }


  hashedPassword, err := hashPassword(newUser.Password)
  if err != nil {
    return models.User{},errors.Join(err,errors.New("Error hashing password"))
  }

  dbUser, err := s.q.CreateUser(context.Background(),db.CreateUserParams{
  	Email:    newUser.Email,
  	Password: hashedPassword,
  })


  if err != nil {
    return models.User{} ,err
  }


  return dbUserToModel(dbUser),nil
}

func (s *Provider) GetUsers() ([]models.User,error){
  dbUsers, err := s.q.GetUsers(context.Background())

  if err != nil {
    return []models.User{},err
  }
  
  users := []models.User{}
  
  for _, dbUser := range dbUsers {
    users = append(users,dbUserToModel(dbUser))
  }

  return users,nil
}

func (s *Provider) GetUser(id int64) (models.User,error)  {
  dbUser, err := s.q.GetUser(context.Background(),id)
  if err != nil {
    return models.User{},err
  }
  return dbUserToModel(dbUser),err
}

func (s *Provider) GetUserByEmail(email string) (models.User,error)  {
  dbUser, err := s.q.GetUserByEmail(context.Background(),email)
  if err != nil {
    return models.User{},err
  }
  return dbUserToModel(dbUser),err
}
