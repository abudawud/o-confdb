package models

import (
  . "o-confdb/database"
)

type Place struct{
  Id            int       `json:"id"`
  IdRegion      int       `json:"id_region"`
  Name          string    `json:"name"`
  Address       string    `json:"address"`
  Description   string    `json:"description"`
  Rank          float32   `json:"rank"`
}

func GetPlaces() (places []Place, err error){
  places = make([]Place, 0)
  rows, err := SqlDB.Query("Select * FROM ms_place")
  defer rows.Close()

  for rows.Next() {
		var place Place
		rows.Scan(
      &place.Id, &place.IdRegion,
      &place.Name, &place.Address,
      &place.Description, &place.Rank)

		places = append(places, place)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func GetPlace(id int) (place Place, err error){
  err = SqlDB.QueryRow("SELECT * FROM ms_place WHERE id = ?", id).
    Scan( &place.Id, &place.IdRegion,
          &place.Name, &place.Address,
          &place.Description, &place.Rank)
  return
}
