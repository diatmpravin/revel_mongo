package controllers

import (
  "github.com/revel/revel"
  "revel_mysql/app/models"
  "labix.org/v2/mgo/bson"
  "fmt"
)

type App struct {
	ModelController
}

func (c App) New() revel.Result {
  greeting := "Hello, Revel MongoDB!"
  return c.Render(greeting)
}

func (c App) Index(myName string) revel.Result {
  fmt.Println(myName)
  
  coll := db.DB("revel_mongo").C("user")
  fmt.Println("Collection: ", coll)

  err := coll.Insert(models.User{myName})
  PanicIf(err)
	
  users := []models.User{}
  err = coll.Find(bson.M{}).All(&users)
  PanicIf(err)
  fmt.Println("Users--->", users)
	
  return c.Render(users)
}

