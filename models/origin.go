package models

type Origin struct{
  Id          int     `json:"id"`
  Name        string  `json:"name"`
  Address     string  `json:"address"`
  Description string  `json:"description"`
}
