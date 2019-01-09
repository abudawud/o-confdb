package apis

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func invalidRoute(c *gin.Context){
  c.String(http.StatusOK, "Invalid Route")
}
