package models

import (
  . "o-confdb/database"
)

type Conference struct{
  Id          int     `json:"id"`
  IdSpeaker   int     `json:"id_speaker"`
  IdAdmin     int     `json:"id_admin"`
  IdPlace     int     `json:"id_pace"`
  Title       string  `json:"title"`
  Description string  `json:"description"`
  CreatedAt   string  `json:"created_at"`
  StartAt     string  `json:"start_at"`
  Rank        float32  `json:"rank"`
}

func GetConfs() (conferences []Conference, err error){
  conferences = make([]Conference, 0)

  rows, err := SqlDB.Query("SELECT * FROM conference")
  defer rows.Close()

  for rows.Next() {
    var conf Conference
    rows.Scan(
      &conf.Id, &conf.IdSpeaker, &conf.IdAdmin,
      &conf.IdPlace, &conf.Title, &conf.Description,
      &conf.CreatedAt, &conf.StartAt, &conf.Rank)

    conferences = append(conferences, conf)
  }
  if err = rows.Err(); err != nil {
    return
  }
  return
}
