package main

import (
  "fmt"
  . "o-sunnahdb/apis"
  "github.com/gin-gonic/gin"
)

func testLog(){
  fmt.Println("hello")
}

func initRouter() *gin.Engine{
  router := gin.Default()

  //speaker
  router.GET("/speakers", GetSpeakersApi)
  router.POST("/speakers", AddSpeakerApi)
  router.PATCH("/speakers", ModSpeakersApi)
  router.DELETE("/speakers", DelSpeakersApi)


  router.GET("/speakers/:id", GetSpeakerApi)
  router.PATCH("/speakers/:id", ModSpeakerApi)
  router.DELETE("/speakers/:id", DelSpeakerApi)

  router.GET("/speakers/:id/conferences", GetSpeakerConfsApi)


  //origin
  router.GET("/origins")
  router.POST("/origins")
  router.PATCH("/origins")
  router.DELETE("/origins")

  router.GET("/origins/:id")
  router.PATCH("/origins/:id")
  router.DELETE("/origins/:id")

  router.GET("/origins/:id/speakers")
  router.GET("/origins/:id/conferences")


  //places
  router.GET("/places")
  router.POST("/places")
  router.PATCH("/places")
  router.DELETE("/places")

  router.GET("/places/:id")
  router.PATCH("/places/:id")
  router.DELETE("/places/:id")

  router.GET("/places/:id/conferences")


  // regions
  router.GET("/regions")
  router.POST("/regions")
  router.PATCH("/regions")
  router.DELETE("/regions")

  router.GET("/regions/:id")
  router.PATCH("/regions/:id")
  router.DELETE("/regions/:id")

  router.GET("/regions/:id/places")
  router.GET("/region/:id/conferences")

  // conferences
  router.GET("/conferences")
  router.POST("/conferences")
  router.PATCH("/conferences")
  router.DELETE("/conferences")

  router.GET("/conferences/:id")
  router.PATCH("/conferences/:id")
  router.DELETE("/conferences/:id")

  router.GET("/conferences/:id/comments")

  // router.GET("/conferences/:id/ranks");

  // comments
  // router.GET("/comments");
  // router.GET("/comments/:id");

  return router
}
