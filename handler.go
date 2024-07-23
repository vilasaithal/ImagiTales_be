package main

import (
	"encoding/json"
	"fmt"
	"hw-server/model"
	"hw-server/service"
	"io"
	"log"
	"net/http"
)

func QueryByDocIdHandler(w http.ResponseWriter, r *http.Request) {
	req, statusCode, err := validateRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Request has error %v", err), statusCode)
		return
	}
	res, err := service.QueryByDocId(req.DocId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when querying ES %v", err), statusCode)
	}

	writeResponse(w, err, res.Story)
}

func PingPong(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling basic function")
	fmt.Fprintf(w, "Hello, World!")
}

func QueryByParametersHandler(w http.ResponseWriter, r *http.Request) {

	req, statusCode, err := validateParameterRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Request has error %v", err), statusCode)
		return
	}
	res, err := service.QueryByParameters(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when querying ES %v", err), statusCode)
	}
	writeResponse(w, err, res.Story)
}

func writeResponse(w http.ResponseWriter, err error, res string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprint(w, res)
	if err != nil {
		log.Println("Failed to write response back to writer")
	}
}

func validateRequest(r *http.Request) (*model.QueryByIdRequest, int, error) {
	// check if request is null - check

	// check if request is not POST , then return appropriate status code and error - check

	// check if request body can be desrialized to request model - done by Suhas - check
	// example below

	if r.Method != http.MethodPost {
		return nil, http.StatusBadRequest, fmt.Errorf("Not POST, needs to be POST")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("Unable to read request body")
	}

	if len(body) == 0 {
		return nil, http.StatusBadRequest, fmt.Errorf("Request body is empty")
	}

	var tempMap map[string]interface{}
	err = json.Unmarshal(body, &tempMap)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("Invalid JSON")
	}

	// Check if there are any fields other than "id"
	if len(tempMap) != 1 || tempMap["id"] == nil {
		return nil, http.StatusBadRequest, fmt.Errorf("JSON body must contain only the 'id' field")
	}

	var request model.QueryByIdRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("The request body does not match the expected schema, err = %v", err)
	}
	return &request, 0, nil
}

func validateParameterRequest(r *http.Request) (model.QueryByParametersRequest, int, error) {
	var req model.QueryByParametersRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		return req, http.StatusBadRequest, err
	}
	return req, http.StatusOK, nil
}
