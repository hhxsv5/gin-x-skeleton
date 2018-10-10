package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hhxsv5/gin-x/response"
	"github.com/hhxsv5/gin-x/router"
)

type Test struct {
	GetUser    func(c *gin.Context) `request:"GET /users"`
	CreateUser func(c *gin.Context) `request:"POST /users"`
}

func (t Test) NewController() router.Controller {
	t.GetUser = getUser
	t.CreateUser = createUser
	return t
}

func getUser(c *gin.Context) {
	// panic(errors.New("test error"))
	response.SuccessJSON(c, "GET /users")
}

func createUser(c *gin.Context) {
	response.SuccessJSON(c, "POST /users")
}
