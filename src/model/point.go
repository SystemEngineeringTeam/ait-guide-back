package model

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SystemEngineeringTeam/ait-guide-back/lib"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/ewkb"
)

var db *sql.DB

func init() {
	db = lib.SqlConnect()
	q := `
		CREATE TABLE IF NOT EXISTS points(
			id INT AUTO_INCREMENT PRIMARY KEY,
			point_id INT,
			location GEOMETRY NOT NULL
		) ENGINE = InnoDB;
	`
	if _, err := db.Query(q); err != nil {
		panic(err)
	} else {
		fmt.Println("table created!!")
	}
}

func GetAllPoints() []orb.Point {
	// var s orb.Point
	var locations []orb.Point
	q := `
		SELECT
			ST_AsBinary(location)
		FROM
			points
		;`
	if rows, err := db.Query(q); err != nil {
		fmt.Println(err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var location orb.Point
			rows.Scan(ewkb.Scanner(&location))
			// fmt.Println(location)
			locations = append(locations, location)
		}
		// fmt.Println(s)
	}
	// fmt.Println(locations)
	return locations
}

func GetClosestPoint(lat, lng string) int {
	var id, point_id int
	var distance float64
	q := `
		SELECT
			id, point_id, ST_Distance(location, ST_GeomFromText('POINT(` + lat + " " + lng + `)', 4326)) AS distance
		FROM
			points
		ORDER BY
			distance ASC
		LIMIT 1
		;
	`
	// fmt.Println(q)
	if rows, err := db.Query(q); err != nil {
		fmt.Println(err)
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&id, &point_id, &distance)
		}
	}
	// fmt.Println(id)
	return point_id
}

func InsertPoint(point_id, lat, lng string) {
	q := `
		INSERT INTO
			points (point_id, location)
		VALUES (
			` + point_id + `,
			ST_GeomFromText('POINT(` + lat + " " + lng + `)', 4326)
		);
	`
	if _, err := db.Exec(q); err != nil {
		log.Println(err)
	}
}
