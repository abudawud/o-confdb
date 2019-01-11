package apis

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "o-confdb/utils"
)

func GetPlacesApi(c *gin.Context){
  c.String(http.StatusOK, "All Places")
}

func AddPlaceApi(c *gin.Context){
  c.String(http.StatusOK, "Add Places")
}

func ModPlacesApi(c *gin.Context){
  c.String(http.StatusOK, "Mod All Places")
}

func DelPlacesApi(c *gin.Context){
  c.String(http.StatusOK, "Del All Places")
}

func GetPlaceApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("Place id %d", id))
  }else{
    invalidRoute(c);
  }
}

func ModPlaceApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("Mod Place id %d", id))
  }else{
    invalidRoute(c);
  }
}

func DelPlaceApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("Del Place id %d", id))
  }else{
    invalidRoute(c);
  }
}

func GetPlaceConfsApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("All confs Place id %d", id))
  }else{
    invalidRoute(c);
  }
}
