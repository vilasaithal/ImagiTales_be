package common

import (
	"hw-server/constants"
	"log"

	"github.com/olivere/elastic/v7"
)

var (
	OpensearchClient *elastic.Client
)

func InitOpenSearch() {
	// Create a client with basic authentication
	var err error
	OpensearchClient, err = elastic.NewClient(
		elastic.SetURL(constants.OpensearchUrl),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(constants.Username, constants.Password),
	)
	if err != nil {
		log.Fatalf("Error creating the client: %v", err)
	}

	log.Println("Initialized opensearch client")

	if OpensearchClient == nil {
		panic("Opensearch client is nil")
	}
}
