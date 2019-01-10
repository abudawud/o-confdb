package models

type Region struct{
  Id          int     `id:"id"`
  Name        string  `json:"name"`
  Description string  `json:"description"`
}
