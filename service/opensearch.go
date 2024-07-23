package service

import (
	"context"
	"encoding/json"
	"fmt"
	"hw-server/common"
	"hw-server/constants"
	"hw-server/model"
	"log"

	"math/rand"

	"github.com/olivere/elastic/v7"
)

func QueryByDocId(id string) (*model.Student, error) {

	client := common.OpensearchClient
	if client == nil {
		log.Println("Client is nil")
	}
	resp, err := client.Get().
		Index(constants.OpensearchIndex).
		Id(id).
		Do(context.Background())
	if err != nil {
		log.Println("Error getting document: %v", err)
		return nil, err
	}

	if resp == nil || resp.Error != nil {
		log.Println("Error when retrieving data for doc id %v, err = %v", id, err)
		return nil, fmt.Errorf(resp.Error.Reason)
	}

	// Deserialize the document
	var student model.Student
	if err := json.Unmarshal(resp.Source, &student); err != nil {
		log.Println("Error deserializing the document: %v", err)
		return nil, err
	}

	return &student, nil

}

func QueryByParameters(req model.QueryByParametersRequest) (*model.Student, error) {

	log.Println("getting response with parameters")
	client := common.OpensearchClient
	if client == nil {
		log.Println("Client is nil")
	}
	query := elastic.NewBoolQuery().Should(
		elastic.NewMatchQuery("House.keyword", req.House),
		elastic.NewMatchQuery("Specialty", req.Specialty),
		elastic.NewMatchQuery("Wand Type", req.WandType),
		elastic.NewMatchQuery("Spell", req.Spell),
	).MinimumShouldMatch("2")

	searchResult, err := client.Search().
		Index("students").
		Query(query).
		Size(20).              // Return top 20 results
		Sort("_score", false). // Sort by score in descending order
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	if searchResult.TotalHits() == 0 {
		return nil, fmt.Errorf("no matching records found")
	}
	var students []model.Student
	for _, hit := range searchResult.Hits.Hits {
		var student model.Student
		err := json.Unmarshal(hit.Source, &student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	// Select a random student from the top 100 results
	randomIndex := rand.Intn(len(students))
	randomStudent := students[randomIndex]

	return &randomStudent, nil
}
