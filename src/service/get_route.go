package service

import (
	"context"

	"github.com/SystemEngineeringTeam/ait-guide-back/lib"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var dr = lib.GetDriver()

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func FindShortestRoute(ctx context.Context, fr, to string) []*Coordinate {
	ses := dr.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer ses.Close(ctx)
	var c []*Coordinate
	var cyp = `
		MATCH (from:Point {pointId: $fr}), (to:Point {pointId: $to}),
			path=shortestPath ((from)-[distance:Distance*]-(to))
		WITH
			[point in nodes(path) | point.lat] as lat,
			[point in nodes(path) | point.lng] as lng,
		REDUCE(totalMinutes = 0, d in distance | totalMinutes + d.cost) as Distance
		RETURN lat, lng
		ORDER BY Distance
		LIMIT 1;
	`
	route, _ := ses.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, _ := tx.Run(ctx, cyp, map[string]any{"fr": fr, "to": to})
		// fmt.Println(result, "result")
		records, _ := result.Collect(ctx)
		// fmt.Println(records, "records")
		return records, nil
	})

	// fmt.Println(route, "route")
	for _, point := range route.([]*neo4j.Record) {
		lat, _ := point.Get("lat")
		lng, _ := point.Get("lng")
		latAr := lat.([]interface{})
		lngAr := lng.([]interface{})
		for i := 0; i < len(latAr); i++ {
			c = append(c, &Coordinate{latAr[i].(float64), lngAr[i].(float64)})
		}
	}

	return c
}
