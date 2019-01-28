package models

import (
  . "o-confdb/database"
  "fmt"
)

type Conference struct{
  Id          int     `json:"id"`
  IdSpeaker   int     `json:"-"`
  IdAdmin     int     `json:"-"`
  IdPlace     int     `json:"-"`
  Title       string  `json:"title"`
  Description string  `json:"description"`
  CreatedAt   string  `json:"created_at"`
  StartAt     string  `json:"start_at"`
  Rank        float32  `json:"rank"`

  ISpeaker    Speaker `json:"speaker"`
  IPlace      Place   `json:"place"`
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

func GetConfsByPlace(id int) (conferences []Conference, err error){
  conferences = make([]Conference, 0)
  sql := fmt.Sprintf("SELECT * FROM conference WHERE id_place = %d", id)
  rows, err := SqlDB.Query(sql)
  defer rows.Close()

  for rows.Next() {
    var conf Conference
    rows.Scan(
      &conf.Id, &conf.IdSpeaker, &conf.IdAdmin,
      &conf.IdPlace, &conf.Title, &conf.Description,
      &conf.CreatedAt, &conf.StartAt, &conf.Rank)

    conf.ISpeaker, _ = GetSpeaker(conf.IdSpeaker)
    conf.IPlace, _ = GetPlace(conf.IdPlace)

    conferences = append(conferences, conf)
  }
  if err = rows.Err(); err != nil {
    return
  }
  return
}
