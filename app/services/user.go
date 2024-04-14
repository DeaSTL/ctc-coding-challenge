package services

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"jhart.dev/ctc-coding-challenge-app/db"
	"jhart.dev/ctc-coding-challenge-app/models"
)

const _EXP_EMAIL = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$` 

// Error types
type UserExistsError struct{
  Email string
}

func (ue UserExistsError) Error() string {
  return fmt.Sprintf("User %v already exists",ue.Email)
}

type UserEmailInvalid struct {
  Email string
}

func (ue UserEmailInvalid) Error() string {
  return fmt.Sprintf("Email: %v is invalid",ue.Email)
}

type UserDoesNotExistError struct {
  Email string
}

func (ue UserDoesNotExistError) Error() string {
  return fmt.Sprintf("User %v does not exist",ue.Email)
}

type UserInvalidPasswordError struct {
  Email string
}

func (ue UserInvalidPasswordError) Error() string {
  return fmt.Sprintf("User %v provided a invalid password",ue.Email)
}

// Helpers
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

//CRUD operations
func (s *Provider) CreateUser(newUser models.User) (models.User, error)  {
  

  emailCheck := s.CheckEmail(newUser.Email)

  if emailCheck != nil {
    return models.User{},emailCheck
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

func (s *Provider) DeleteUser(id int64) error{
  err := s.q.DeleteUser(context.Background(),id) 

  if err != nil {
    return err
  }
  return nil
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

//Business logic

func (s *Provider) ValidateUserCredentials(email string, password string) (models.User, error) {
  user, err := s.GetUserByEmail(email)

  if err != nil {
    return models.User{},UserDoesNotExistError{Email: email}
  }

  if !checkPassword(password,user.Password) {
    return models.User{},UserInvalidPasswordError{Email: email} 
  }

  return user, nil
}

func (s *Provider) CheckEmail(email string) error {
    
  emailExp,err := regexp.Compile(_EXP_EMAIL)

  if err != nil {
    return err
  }

  if !emailExp.MatchString(email) {
    return UserEmailInvalid{Email:email}
  }
  _,err = s.GetUserByEmail(email)

  if err == nil {
    return UserExistsError{Email:email} 
  }

  return nil
}



