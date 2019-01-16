package main

import (
  "fmt"
  . "o-confdb/apis"
  "github.com/gin-gonic/gin"
)

func testLog(){
  fmt.Println("hello")
}

func initRouter() *gin.Engine{
  router := gin.Default()

  router.POST("/users", AddUserApi)
  router.POST("/auth", AddTokenApi)

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
  router.GET("/origins", GetOriginsApi)
  router.POST("/origins", AddOriginApi)
  router.PATCH("/origins", ModOriginsApi)
  router.DELETE("/origins", DelOriginsApi)


  router.GET("/origins/:id", GetOriginApi)
  router.PATCH("/origins/:id", ModOriginApi)
  router.DELETE("/origins/:id", DelOriginApi)

  router.GET("/origins/:id/speakers", GetOriginSpeakersApi)
  router.GET("/origins/:id/conferences", GetOriginConfsApi)


  //places
  router.GET("/places", GetPlacesApi)
  router.POST("/places", AddPlaceApi)
  router.PATCH("/places", ModPlacesApi)
  router.DELETE("/places", DelPlacesApi)

  router.GET("/places/:id", GetPlaceApi)
  router.PATCH("/places/:id", ModPlaceApi)
  router.DELETE("/places/:id", DelPlaceApi)

  router.GET("/places/:id/conferences", GetPlaceConfsApi)


  // regions
  router.GET("/regions", GetRegionsApi)
  router.POST("/regions", AddRegionApi)
  router.PATCH("/regions", ModRegionsApi)
  router.DELETE("/regions", DelRegionsApi)

  router.GET("/regions/:id", GetRegionApi)
  router.PATCH("/regions/:id", ModRegionApi)
  router.DELETE("/regions/:id", DelRegionApi)

  router.GET("/regions/:id/places", GetRegionPlacesApi)
  router.GET("/region/:id/conferences", GetRegionConfsApi)

  // conferences
  router.GET("/conferences", GetConfsApi)
  router.POST("/conferences", AddConfApi)
  router.PATCH("/conferences", ModConfsApi)
  router.DELETE("/conferences", DelConfsApi)

  router.GET("/conferences/:id", GetConfApi)
  router.PATCH("/conferences/:id", ModConfApi)
  router.DELETE("/conferences/:id", DelConfApi)

  router.GET("/conferences/:id/comments", GetConfCommentsApi)

  // router.GET("/conferences/:id/ranks");

  // comments
  // router.GET("/comments");
  // router.GET("/comments/:id");

  return router
}
