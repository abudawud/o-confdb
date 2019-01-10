package models

type Place struct{
  Id            int     `json:"id"`
  IdRegion      int     `json:"id_region"`
  Name          string  `json:"name"`
  Address       string  `json:"address"`
  Description   string  `json:"description"`
  Rank          float   `json:"rank"`
}
