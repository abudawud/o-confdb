package models

import (
  . "o-confdb/database"
)

type Origin struct{
  Id          int     `json:"id"`
  Name        string  `json:"name"`
  Address     string  `json:"address"`
  Description string  `json:"description"`
}

func GetOrigins() (origins []Origin, err error){
  origins = make([]Origin, 0)
  rows, err := SqlDB.Query("Select * FROM ms_speaker_origin")
  defer rows.Close()

  for rows.Next() {
		var origin Origin
		rows.Scan(&origin.Id, &origin.Name, &origin.Address, &origin.Description)
		origins = append(origins, origin)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (o *Origin) AddOrigin() (record int64, err error){
  record = 0
  rs, err := SqlDB.Exec("INSERT INTO ms_speaker_origin(name, address, description) VALUES (?, ?, ?)", o.Name, o.Address, o.Description)

  if err != nil{
    return
  }

  record,_ = rs.LastInsertId()
  return
}
