package apis

import (
  "log"
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "o-confdb/utils"
  . "o-confdb/models"
)

func GetOriginsApi(c *gin.Context){
  origins, err := GetOrigins();

  if err != nil{
    log.Fatalln(err)
  }

  c.JSON(http.StatusOK, origins)
}

func AddOriginApi(c *gin.Context){
  name := c.Request.FormValue("name")
  address := c.Request.FormValue("address")
  description := c.Request.FormValue("description")

  origin := Origin{Name: name, Address: address, Description: description}

  record, err := origin.AddOrigin()

  if err != nil{
    log.Fatalln(err)
  }

  var msg string
  var code int
  if record > 0 {
    code = 200
    msg = "Origin Added"
  }else{
    code = 400
    msg = "Add origin Failed"
  }

  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "messege": msg,
  })
}

func ModOriginsApi(c *gin.Context){
  c.String(http.StatusOK, "Mod All Origins")
}

func DelOriginsApi(c *gin.Context){
  c.String(http.StatusOK, "Del All Origins")
}

func GetOriginApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("Origin id %d", id))
  }else{
    invalidRoute(c);
  }
}

func ModOriginApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("Mod Origin id %d", id))
  }else{
    invalidRoute(c);
  }
}

func DelOriginApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("Del Origin id %d", id))
  }else{
    invalidRoute(c);
  }
}

func GetOriginSpeakersApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("All speakers Origin id %d", id))
  }else{
    invalidRoute(c);
  }
}

func GetOriginConfsApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("All confs Origin id %d", id))
  }else{
    invalidRoute(c);
  }
}
