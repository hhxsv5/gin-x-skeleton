package test

import (
	"github.com/hhxsv5/gin-x/model/mysql"
)

type Test struct {
	mysql.Model
	Name   string `json:"name"`
	Number uint64 `json:"number"`
	Type   uint16 `json:"type"`
	Status uint16 `json:"status"`
}
