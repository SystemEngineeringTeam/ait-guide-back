package lib

import (
	"log"

	conf "github.com/SystemEngineeringTeam/ait-guide-back/conf"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetDriver() neo4j.DriverWithContext {
	n := conf.GetNeo4jConfig()
	dr, err := neo4j.NewDriverWithContext(n.GetString("neo4j.uri"), neo4j.BasicAuth(n.GetString("neo4j.user"), n.GetString("neo4j.password"), ""))
	if err != nil {
		log.Fatal(err)
	}
	return dr
}
