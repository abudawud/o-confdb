package apis

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "o-sunnahdb/utils"
)

func GetRegionsApi(c *gin.Context){
  c.String(http.StatusOK, "All Regions")
}

func AddRegionApi(c *gin.Context){
  c.String(http.StatusOK, "Add Regions")
}

func ModRegionsApi(c *gin.Context){
  c.String(http.StatusOK, "Mod All Regions")
}

func DelRegionsApi(c *gin.Context){
  c.String(http.StatusOK, "Del All Regions")
}

func GetRegionApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("Region id %d", id))
  }else{
    invalidRoute(c);
  }
}

func ModRegionApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("Mod Region id %d", id))
  }else{
    invalidRoute(c);
  }
}

func DelRegionApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("Del Region id %d", id))
  }else{
    invalidRoute(c);
  }
}

func GetRegionPlacesApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("All confs Region id %d", id))
  }else{
    invalidRoute(c);
  }
}

func GetRegionConfsApi(c *gin.Context){
  id := utils.Str2Int(c.Param("id"))

  if id != -1{
    c.String(http.StatusOK, fmt.Sprintf("All confs Region id %d", id))
  }else{
    invalidRoute(c);
  }
}
