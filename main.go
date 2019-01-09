package main

func main(){
  router := initRouter()
  router.Run(":3000")
}
