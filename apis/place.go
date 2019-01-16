package apis

import (
  "log"
  "fmt"
  "net/http"
  "database/sql"
  "github.com/gin-gonic/gin"
  "o-confdb/utils"
  . "o-confdb/helper"
  . "o-confdb/models"
)

func GetPlacesApi(c *gin.Context){
  var apiErr ApiErr

  token := c.Request.FormValue("token")
  apiErr = ValidateToken(token, ROLE_GUEST)
  if(apiErr.Code != 200){
    c.JSON(apiErr.Code, apiErr)
    return
  }

  places, err := GetPlaces();

  if err != nil{
    log.Fatalln(err)
  }

  c.JSON(http.StatusOK, places)
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
  var apiErr ApiErr

  id := utils.Str2Int(c.Param("id"))

  if id == -1{
    invalidRoute(c);
    return
  }

  token := c.Request.FormValue("token")
  apiErr = ValidateToken(token, ROLE_GUEST)
  if(apiErr.Code != 200){
    c.JSON(apiErr.Code, apiErr)
    return
  }

  place, err := GetPlace(id);

  if err != sql.ErrNoRows && err != nil{
    log.Fatalln(err)
  }

  c.JSON(http.StatusOK, place)
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
