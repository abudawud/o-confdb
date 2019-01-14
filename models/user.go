package models

import (
  . "o-confdb/database"
)

type User struct{
  Id          int     `json:"id"`
  Username    string  `json:"username"`
  Password    string  `json:"password"`
  Role        int     `json:"role"`
  FirstName   string  `json:"first_name"`
  LastName    string  `json:"last_name"`
}

func GetUser(uid int) (user User, err error){
  err = SqlDB.QueryRow("SELECT * FROM ms_user WHERE id = ?", uid).
    Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.FirstName, &user.LastName)

  user.Password = "*"
  return
}
