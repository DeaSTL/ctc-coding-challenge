package models


type User struct {
  ID int64
  Email string
  Password string //bcrypt+salt
}

func (u *User) MakePublic()  {
  u.Password = ""
}

  


