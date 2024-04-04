package service

import (
	"context"
	"fmt"

	"github.com/SystemEngineeringTeam/ait-guide-back/model"
)

func GetShortestDistanceFromCurrentLocation(lat, lng, end string) []*Coordinate {
	ctx := context.Background()
	p := model.GetClosestPoint(lat, lng)
	if fmt.Sprint(p) == end {
		op := model.GetPoint(p)
		return []*Coordinate{{op.Lat(), op.Lon()}}
	}
	route := FindShortestRoute(ctx, fmt.Sprint(p), end)
	return route
}
