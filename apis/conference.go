package apis

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "o-confdb/utils"
)

func GetConfsApi(c *gin.Context){
  c.String(http.StatusOK, "All Confs")
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
