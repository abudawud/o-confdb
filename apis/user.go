package apis

import (
  "github.com/gin-gonic/gin"
  "o-confdb/utils"
  . "o-confdb/models"
)

func AddUserApi(c *gin.Context){
  var aerr ApiErr

  username := c.Request.FormValue("username")
  password := c.Request.FormValue("password")
  trole := c.Request.FormValue("role")
  email := c.Request.FormValue("email")
  firstName := c.Request.FormValue("first_name")
  lastName := c.Request.FormValue("last_name")

  if username == "" || password == "" || firstName == ""{
      aerr.Code = 400
      aerr.Messege = "Permission Denied!, Some Field Required!"

      c.JSON(aerr.Code, aerr)
      return
    }


  role := utils.Str2Int(trole)
  if role == -1 {
    aerr.Code = 400
    aerr.Messege = "Permission Denied!, Invalid Role!"

    c.JSON(aerr.Code, aerr)
    return
  }

  user := User{
    Username: username,
    Password: password,
    Role: role,
    Email: email,
    FirstName: firstName,
    LastName: lastName }

  err := user.AddUser()

  if err != nil{
    aerr.Code = 500
    aerr.Messege = "Internal Server Error!, @AddUser"
  }else{
    aerr.Code = 200
    aerr.Messege = "OK"
  }

  c.JSON(aerr.Code, aerr)
}
