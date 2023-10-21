package db

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type ArgumentID string

type Argument struct {
	ID   ArgumentID
	Text string
}

type EvidenceType string

const (
	Supports    EvidenceType = "SUPPORTS"
	Contradicts EvidenceType = "CONTRADICTS"
)

type Evidence struct {
	StatementID StatementID
	ArgumentID  ArgumentID
	Type        EvidenceType
}

const (
	upsertArgumentQuery = `
    MERGE (n:Argument {id: COALESCE($id, apoc.create.uuid())})
    ON CREATE SET n.text = $text
    ON MATCH SET n.text = $text
    RETURN n
	`

	argumentEvidenceStatementQuery = `
	MATCH (a:Argument {id: $argumentID}), (s:Statement {id: $statementID})
	MERGE (a)-[:EVIDENCE {type: $evidenceType}]->(s)
	RETURN a, s
	`
)

func UpsertArgument(session neo4j.SessionWithContext, Argument Argument) (neo4j.ResultWithContext, error) {
	params := map[string]interface{}{
		"id":   Argument.ID,
		"text": Argument.Text,
	}

	result, err := session.Run(context.Background(), upsertArgumentQuery, params)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func EvidenceStatementForArgument(session neo4j.SessionWithContext, evidence Evidence) (neo4j.ResultWithContext, error) {
	params := map[string]interface{}{
		"statementID":  evidence.StatementID,
		"argumentID":   evidence.ArgumentID,
		"evidenceType": evidence.Type,
	}

	result, err := session.Run(context.Background(), argumentEvidenceStatementQuery, params)
	if err != nil {
		return nil, err
	}

	return result, nil
}
