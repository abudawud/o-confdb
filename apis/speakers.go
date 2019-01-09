package apis


import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "o-sunnahdb/utils"
)

func GetSpeakersApi(c *gin.Context){
  c.String(http.StatusOK, "All Speakers")
}

func AddSpeakerApi(c *gin.Context){
  c.String(http.StatusOK, "Add Speakers")
}

func ModSpeakersApi(c *gin.Context){
  c.String(http.StatusOK, "Mod All Speakers")
}

func DelSpeakersApi(c *gin.Context){
  c.String(http.StatusOK, "Del All Speakers")
}

func GetSpeakerApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("Speaker id %d", id))
  }else{
    invalidRoute(c);
  }
}

func ModSpeakerApi(c *gin.Context){
  c.String(http.StatusOK, "Mod All Speakers")
}

func DelSpeakerApi(c *gin.Context){
  c.String(http.StatusOK, "Del All Speakers")
}

func GetSpeakerConfsApi(c *gin.Context){
  c.String(http.StatusOK, "Del All Speakers")
}
