package models

import (
  . "o-confdb/database"
  "fmt"
)

type User struct{
  Id          int     `json:"id"`
  Username    string  `json:"username"`
  Password    string  `json:"password"`
  Role        int     `json:"role"`
  Email       string  `json:"email"`
  FirstName   string  `json:"first_name"`
  LastName    string  `json:"last_name"`
}

func GetUserById(uid int) (user User, err error){
  err = SqlDB.QueryRow("SELECT * FROM ms_user WHERE id = ?", uid).
    Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.FirstName, &user.LastName)

  user.Password = "*"
  return
}

func GetUserByAuth(username string, password string) (user User, err error){
  sql := fmt.Sprintf("SELECT * FROM ms_user WHERE username = '%s' AND password = '%s'", username, password)
  fmt.Println(sql)
  err = SqlDB.QueryRow(sql).
    Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.Email, &user.FirstName, &user.LastName)

  return
}

func (u *User) AddUser() (err error){
  rs, err := SqlDB.Exec("INSERT INTO ms_user(username, password, role, email, first_name, last_name) VALUES (?, ?, ?, ?, ?, ?)",
    u.Username, u.Password, u.Role, u.Email, u.FirstName, u.LastName)

  if err != nil{
    return
  }

  rs.LastInsertId()
  return
}
