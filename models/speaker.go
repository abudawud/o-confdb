package models

type Speaker struct {
	Id         int    `json:"id" form:"id"`
	IdOrigin   int    `json:"id_origin"`
	FirstName  string `json:"first_name"`
	LastName   int    `json:"last_name"`
  Rank       float32  `json:"rank"`
}
