package apis

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "o-confdb/utils"
)

func GetOriginsApi(c *gin.Context){
  c.String(http.StatusOK, "All Origins")
}

func AddOriginApi(c *gin.Context){
  c.String(http.StatusOK, "Add Origins")
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
