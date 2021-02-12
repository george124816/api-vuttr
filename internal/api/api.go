package api

import (
	"github.com/george124816/api-vuttr/internal/api/router"
)

func Run(){
	web := router.Setup()
	web.Run(":3000")
}