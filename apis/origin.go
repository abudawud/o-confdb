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

func GetOriginsApi(c *gin.Context){
  var apiErr ApiMsg

  token := c.Request.FormValue("token")
  apiErr = ValidateToken(token, ROLE_GUEST)
  if(apiErr.Code != 200){
    c.JSON(apiErr.Code, apiErr)
    return
  }

  origins, err := GetOrigins();

  if err != nil{
    log.Fatalln(err)
  }

  c.JSON(http.StatusOK, origins)
}

func AddOriginApi(c *gin.Context){
  // _, vErr := ValidateUser(ROLE_ADMIN, c.Request.FormValue("token"))
  //
  // if vErr.Code != 200{
  //   c.JSON(vErr.Code, vErr)
  //
  //   return
  // }

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
  var apiErr ApiMsg

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

  origin, err := GetOrigin(id);

  if err != sql.ErrNoRows && err != nil{
    log.Fatalln(err)
  }

  c.JSON(http.StatusOK, origin)
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
