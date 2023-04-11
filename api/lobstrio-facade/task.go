package lobstriofacade

import (
	"context"
	"encoding/json"

	"github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

type TaskResult struct {
	DuplicatedCount int `json:"duplicated_count"`
	Tasks []TaskDto `json:"tasks"`
}
type Params struct {
	URL string `json:"url"`
}
type TaskDto struct {
	ID string `json:"id"`
	CreatedAt string `json:"created_at"`
	IsActive bool `json:"is_active"`
	Params Params `json:"params"`
	Object string `json:"object"`
}

func CreateTask(id string, url string) *TaskResult {

	createtaskNotaddedRequest := lobstrio.NewCreatetaskNotaddedRequestWithDefaults()
	createtaskNotaddedRequest.Cluster = id
	createtaskNotaddedRequest.Tasks = []lobstrio.Task{
		{Url: url},
	}

	data, err := GetClient().MiscApi.
		CreatetaskNotadded(context.Background()).
		CreatetaskNotaddedRequest(*createtaskNotaddedRequest).
		Execute()

	if err != nil {
		panic(err)
	}

	var task *TaskResult

	json.NewDecoder(data.Body).Decode(&task)

	return task
}
