package lib

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetDriver() neo4j.DriverWithContext {
	// n := config.GetNeo4jConfig()
	// dr, err := neo4j.NewDriver(n.GetString("neo4j.uri"), neo4j.BasicAuth(n.GetString("neo4j.user"), n.GetString("neo4j.password"), ""))
	dr, err := neo4j.NewDriverWithContext("neo4j://localhost:57687", neo4j.BasicAuth("neo4j", "madeintokyo", ""))
	if err != nil {
		log.Fatal(err)
	}
	return dr
}
