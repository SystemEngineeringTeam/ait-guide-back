package service

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateGraph(ctx context.Context, node, rel string) {
	ses := dr.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer ses.Close(ctx)
	var cyp = `
	CALL gds.graph.project(
		'myGraph',
		$node,
		$rel,
		{
			relationshipProperties: 'cost'
		}
	)
	`
	if _, err := ses.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if _, err := tx.Run(ctx, cyp, map[string]any{"node": node, "rel": rel}); err != nil {
			return nil,err
		}
		return nil, nil
	}); err != nil {
		panic(err)
	}
}
