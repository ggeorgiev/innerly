package db

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type TopicID string

type Topic struct {
	ID   TopicID
	Text string
}

type Partain struct {
	TopicID     TopicID
	StatementID StatementID
}

const (
	upsertTopicQuery = `
    MERGE (n:Topic {id: COALESCE($id, apoc.create.uuid())})
    ON CREATE SET n.text = $text
    ON MATCH SET n.text = $text
    RETURN n
	`

	topicPartainStatementQuery = `
	MATCH (t:Topic {id: $topicID}), (s:Statement {id: $statementID})
	MERGE (t)-[:PERTAIN]->(s)
	RETURN t, s
	`
)

func UpsertTopic(session neo4j.SessionWithContext, topic Topic) (neo4j.ResultWithContext, error) {
	params := map[string]interface{}{
		"id":   topic.ID,
		"text": topic.Text,
	}

	result, err := session.Run(context.Background(), upsertTopicQuery, params)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func PartainTopicStatement(session neo4j.SessionWithContext, partain Partain) (neo4j.ResultWithContext, error) {
	params := map[string]interface{}{
		"statementID": partain.StatementID,
		"topicID":     partain.TopicID,
	}

	result, err := session.Run(context.Background(), topicPartainStatementQuery, params)
	if err != nil {
		return nil, err
	}

	return result, nil
}
