package controllers

import (
	"github.com/revel/revel"
)

func init() {
	revel.InterceptMethod((*ModelController).Begin, revel.BEFORE)
}
