package models

import (
  . "o-confdb/database"
)

type Speaker struct {
	Id         int    `json:"id" form:"id"`
	IdOrigin   int    `json:"-"`
	FirstName  string `json:"first_name"`
	LastName   int    `json:"last_name"`
  Rank       float32  `json:"rank"`
}

func GetSpeaker(id int) (speaker Speaker, err error){
  err = SqlDB.QueryRow("SELECT * FROM ms_speaker WHERE id = ?", id).
    Scan( &speaker.Id, &speaker.IdOrigin,
          &speaker.FirstName, &speaker.LastName,
          &speaker.Rank)
  return
}
