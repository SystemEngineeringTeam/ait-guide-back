package main

import (
	"github.com/SystemEngineeringTeam/ait-guide-back/controller"
	"github.com/SystemEngineeringTeam/ait-guide-back/service"
)

func main() {
	service.Migration()
	controller.ApiHandler()
}
