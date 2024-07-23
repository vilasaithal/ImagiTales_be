package main

import (
	"fmt"
	"hw-server/common"
	"net/http"
)

func main() {
	// OpenSearch InitOpenSearch
	common.InitOpenSearch()

	http.HandleFunc("/", PingPong)

	http.HandleFunc("/v1/query_by_id", QueryByDocIdHandler)
	http.HandleFunc("/v1/query_by_parameters", QueryByParametersHandler)

	fmt.Println("Server is listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
