package db

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type StatementID string

type Statement struct {
	ID   StatementID
	Text string
}

type StanceType string

const (
	For     StanceType = "FOR"
	Against StanceType = "AGAINST"
)

type Stance struct {
	StatementID StatementID
	ArgumentID  ArgumentID
	Type        StanceType
}

const (
	upsertStatementQuery = `
    MERGE (n:Statement {id: COALESCE($id, apoc.create.uuid())})
    ON CREATE SET n.text = $text
    ON MATCH SET n.text = $text
    RETURN n
	`

	statementStanceArgumentQuery = `
	MATCH (s:Statement {id: $statementID}), (a:Argument {id: $argumentID})
	MERGE (s)-[:STANCE {type: $stanceType}]->(a)
	RETURN s, a
	`
)

func UpsertStatement(session neo4j.SessionWithContext, statement Statement) (neo4j.ResultWithContext, error) {
	params := map[string]interface{}{
		"id":   statement.ID,
		"text": statement.Text,
	}

	result, err := session.Run(context.Background(), upsertStatementQuery, params)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func StanceArgumentToStatement(session neo4j.SessionWithContext, stance Stance) (neo4j.ResultWithContext, error) {
	params := map[string]interface{}{
		"argumentID":  stance.ArgumentID,
		"statementID": stance.StatementID,
		"stanceType":  stance.Type,
	}

	result, err := session.Run(context.Background(), statementStanceArgumentQuery, params)
	if err != nil {
		return nil, err
	}

	return result, nil
}
