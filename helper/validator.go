package helper

import(
  "fmt"
  "crypto/md5"
  "database/sql"
  . "o-confdb/models"
)

const ROLE_ADMIN = 1
const ROLE_GUEST = 2

func ValidateUser(username string, password string) (user User, vErr ApiMsg){
  if username == "" || password == ""{
    vErr.Code = 400
    vErr.Message = "Permission Denied!, Username/Password not provided"

    return
  }

  pass := []byte(password)
  hash := fmt.Sprintf("%x", md5.Sum(pass))
  user, err := GetUserByAuth(username, hash)

  if err == sql.ErrNoRows{
    vErr.Code = 400
    vErr.Message = "Permission Denied, Invalid Username/Password!"

    return
  }else if err != nil{
    vErr.Code = 500
    vErr.Message = "Internal Server Error, @validateUser"
    fmt.Println(err)

    return
  }

  vErr.Code = 200;

  return
}

func ValidateToken(rtoken string, role int) (vErr ApiMsg){
  vErr.Code = 200

  if rtoken == ""{
    vErr.Code = 400
    vErr.Message = "Permission Denied, Token not provided!"

    return
  }

  token, err := GetAuth(rtoken)

  if err == sql.ErrNoRows{
    vErr.Code = 400
    vErr.Message = "Permission Denied, Token Invalid or Expired!"

    return
  }else if err != nil{
    vErr.Code = 500
    vErr.Message = "Internal Server Error, @ValidateToken"

    return
  }

  switch(role){
    case ROLE_ADMIN:{
      if token.Role != ROLE_ADMIN{
        vErr.Code = 400
        vErr.Message = "Permission Denied, Admin Role Required!"
      }
      break
    }

    case ROLE_GUEST:{
      if (token.Role & (ROLE_GUEST|ROLE_ADMIN)) == 0{
        vErr.Code = 400
        vErr.Message = "Permission Denied, At least Guest Role Required"
      }
      break
    }

    default:{
      vErr.Code = 500
      vErr.Message = "Internal Server Error!, Role Not Found!"
    }
  }

  return
}
