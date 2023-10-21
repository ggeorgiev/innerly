package main

import (
	"context"
	"log"
	"os"

	"github.com/ggeorgiev/neo/src/db"
	"github.com/ggeorgiev/neo/src/sample"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	ctx := context.Background()
	dbUri := os.Getenv("NEO4J_DATABASE")
	dbUser := os.Getenv("NEO4J_USER")
	dbPassword := os.Getenv("NEO4J_PASSWORD")
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	if err != nil {
		log.Panicf("%+v, %+v", dbUri, err)
	}
	defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new session
	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}

	session := driver.NewSession(context.Background(), sessionConfig)
	defer session.Close(context.Background())

	for _, topic := range sample.Topics {
		_, err = db.UpsertTopic(session, topic)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, statement := range sample.Statements {
		_, err = db.UpsertStatement(session, statement)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, argument := range sample.Arguments {
		_, err = db.UpsertArgument(session, argument)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, partain := range sample.Partains {
		_, err = db.PartainTopicStatement(session, partain)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, stance := range sample.Stances {
		_, err = db.StanceArgumentToStatement(session, stance)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, evidence := range sample.Evidences {
		_, err = db.EvidenceStatementForArgument(session, evidence)
		if err != nil {
			log.Fatal(err)
		}
	}
}
