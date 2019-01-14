package helper

import(
  "o-confdb/utils"
  "database/sql"
  . "o-confdb/models"
)

const ROLE_ADMIN = 1
const ROLE_GUEST = 2

func ValidateUser(role int, token string) (data User, vErr ApiErr){
  if token == ""{
    vErr.Code = 400
    vErr.Messege = "Permission Denied, Token not provided!"

    return
  }

  var idtoken int = utils.Str2Int(token)
  user, err := GetUser(idtoken)

  if err == sql.ErrNoRows{
    vErr.Code = 400
    vErr.Messege = "Permission Denied, Token Invalid or Expired!"

    return
  }else if err != nil{
    vErr.Code = 500
    vErr.Messege = "Internal Server Error, @getUser"

    return
  }

  switch role {
    case ROLE_ADMIN:{
      if user.Role != ROLE_ADMIN{
        vErr.Code = 400
        vErr.Messege = "Permission Denied, Role Admin Required!"
        return
      }
      break
    }

    case ROLE_GUEST:{
      if user.Role != ROLE_ADMIN || user.Role != ROLE_GUEST{
        vErr.Code = 400
        vErr.Messege = "Permission Denied, At least Role Guest Required"
        return
      }
      break
    }

    default:{
      vErr.Code = 500
      vErr.Messege = "Internal Server Error!, Role Not Found!"
      return
    }
  }

  vErr.Code = 200
  data = user;
  return
}
