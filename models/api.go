package models

const API_ERR = 1

type ApiErr struct{
  Code      int     `json:"statuscode"`
  Messege   string  `json:"messege"`
}
