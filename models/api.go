package models

const API_ERR = 1

type ApiMsg struct{
  Code      int     `json:"code"`
  Message   string  `json:"message"`
}
