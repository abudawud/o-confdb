package models

import (
 "fmt"
 . "o-confdb/database"
)

type Token struct{
  Id          int     `json:"id"`
  Token       string  `json:"token"`
  Role        int     `json:"role"`
  Issued      string  `json:"issued"`
  Expired     string  `json:"expired"`
  User        User    `json:"user"`
  Status      ApiMsg  `json:"status"`
}

func GetAuth(rtoken string) (token Token, err error){
  sql := fmt.Sprintf("SELECT password, role FROM ms_user WHERE password = '%s'", rtoken)

  err = SqlDB.QueryRow(sql).
    Scan(&token.Token, &token.Role)

  return
}
