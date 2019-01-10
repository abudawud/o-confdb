package models

type Conference struct{
  Id          int     `json:"id"`
  IdSpeaker   int     `json:"id_speaker"`
  IdAdmin     int     `json:"id_admin"`
  IdPlace     int     `json:"id_pace"`
  Title       string  `json:"title"`
  Description string  `json:"description"`
  Rank        float   `json:"rank"`
}
