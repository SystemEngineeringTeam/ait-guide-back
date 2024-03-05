package main

import (
	"github.com/SystemEngineeringTeam/ait-guide-back/service"
)

func main() {
	service.Migration("240212")
	// for _, v := range model.GetAllPoints() {
	// 	fmt.Println(v)
	// }
	// model.GetClosestPoint(35.18411661111426, 137.1110111474991)
}
