package lobstriofacade

import (
	"context"
	"encoding/json"
)

type ResultDtoResult[T any] struct{
	Count int `json:"count"`
	Page int `json:"page"`
	Limit int `json:"limit"`
	Data []T `json:"data"`
}

func GetResult[T any](clusterId string, runId string) *ResultDtoResult[T] {
	data, err := GetClient().MiscApi.
		ListResult(context.Background()).
		Cluster(clusterId).
		Run(runId).
		ContentType("application/json").
		Execute()

	checkError(data, err)

	var result *ResultDtoResult[T]

	json.NewDecoder(data.Body).Decode(&result)

	return result
}
