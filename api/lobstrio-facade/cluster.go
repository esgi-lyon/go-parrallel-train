package lobstriofacade

import (
	"context"
	"encoding/json"
)

type ClusterDto struct {
	Count int `json:"count"`
	Limit int `json:"limit"`
	Page int `json:"page"`
	TotalPages int `json:"total_pages"`
	ResultFrom int `json:"result_from"`
	ResultTo int `json:"result_to"`
	Data []interface{} `json:"data"`
}

func Getcluster(id string) *ClusterDto {
	data, err := GetClient().MiscApi.
		Getcluster(context.Background()).
		ContentType("application/json").
		Id(id).
		Execute()

	if err != nil {
		panic(err)
	}

	var cluster *ClusterDto

	json.NewDecoder(data.Body).Decode(&cluster)

	return cluster
}
