package apis

import (
  "github.com/gin-gonic/gin"
  . "o-confdb/helper"
  . "o-confdb/models"
)

func AddTokenApi(c *gin.Context){
  username := c.Request.FormValue("username")
  password := c.Request.FormValue("password")

  user, mErr := ValidateUser(username, password)
  if mErr.Code != 200{
    c.JSON(mErr.Code, mErr)
    return
  }

  mErr.Code = 200
  mErr.Message = "OK"
  var token = Token{Token: user.Password, User: user, Status: mErr}
  c.JSON(200, token)
}
