package apis

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "o-confdb/utils"
  . "o-confdb/helper"
  . "o-confdb/models"
)

func GetConfsApi(c *gin.Context){
  var apiErr ApiMsg

  token := c.Request.FormValue("token")
  apiErr = ValidateToken(token, ROLE_GUEST)
  if(apiErr.Code != 200){
    c.JSON(apiErr.Code, apiErr)
    return
  }

  conferences, err := GetConfs()

  if(err != nil){
    apiErr.Code = 500
    apiErr.Message = "Internal Server Error, @GetConfApi!"
    c.JSON(apiErr.Code, apiErr)
  }

  c.JSON(http.StatusOK, conferences)
}

func AddConfApi(c *gin.Context){
  c.String(http.StatusOK, "Add Confs")
}

func ModConfsApi(c *gin.Context){
  c.String(http.StatusOK, "Mod All Confs")
}

func DelConfsApi(c *gin.Context){
  c.String(http.StatusOK, "Del All Confs")
}

func GetConfApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("Conf id %d", id))
  }else{
    invalidRoute(c);
  }
}

func ModConfApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("Mod Conf id %d", id))
  }else{
    invalidRoute(c);
  }
}

func DelConfApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("Del Conf id %d", id))
  }else{
    invalidRoute(c);
  }
}

func GetConfCommentsApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("All comments Conf id %d", id))
  }else{
    invalidRoute(c);
  }
}
