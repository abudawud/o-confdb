package apis

import (
  "fmt"
  "crypto/md5"
  "github.com/gin-gonic/gin"
  "o-confdb/utils"
  . "o-confdb/models"
)

func GetUserApi(c *gin.Context){
  ids := c.Param("id")
  id := utils.Str2Int(ids)

  user, _ := GetUserById(int64(id))
  c.JSON(200, user)
}

func AddUserApi(c *gin.Context){
  var aerr ApiMsg

  username := c.Request.FormValue("username")
  password := c.Request.FormValue("password")
  trole := c.Request.FormValue("role")
  email := c.Request.FormValue("email")
  firstName := c.Request.FormValue("first_name")
  lastName := c.Request.FormValue("last_name")

  if username == "" || password == "" || firstName == ""{
      aerr.Code = 400
      aerr.Message = "Permission Denied!, Some Field Required!"

      c.JSON(aerr.Code, aerr)
      return
    }


  role := utils.Str2Int(trole)
  if role == -1 {
    aerr.Code = 400
    aerr.Message = "Permission Denied!, Invalid Role!"

    c.JSON(aerr.Code, aerr)
    return
  }

  password = fmt.Sprintf("%x", md5.Sum([]byte(password)))

  user := User{
    Username: username,
    Password: password,
    Role: role,
    Email: email,
    FirstName: firstName,
    LastName: lastName }

  index, err := user.AddUser()

  if err != nil{
    fmt.Println(err)
    aerr.Code = 400
    aerr.Message = "Permission Denied, Username already exist"
    c.JSON(aerr.Code, aerr)
  }else{
    var user User
    user, err = GetUserById(index)
    c.JSON(200, user)
  }
}
