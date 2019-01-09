package utils

import(
  "strconv"
)

func Str2Int(str string) int{
  res, err := strconv.Atoi(str)

  if err != nil{
    return -1
  }

  return res
}
