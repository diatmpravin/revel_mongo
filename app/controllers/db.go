package controllers

import (
  	"fmt"
	"labix.org/v2/mgo"
	r "github.com/revel/revel"
)

var db *mgo.Session

func init() {
	var err error
	db, err = mgo.Dial("localhost")
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	} else {
		fmt.Println("Successfully connected to DB: %#v", db)
	}
}

type ModelController struct {
	*r.Controller
}

func (c *ModelController) Begin() r.Result {
  	fmt.Println("DB HERE---->%+v", db)
	return nil
}

func PanicIf(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}


