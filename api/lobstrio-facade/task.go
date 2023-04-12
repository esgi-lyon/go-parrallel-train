package lobstriofacade

import (
	"context"
	"encoding/json"

	"github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

type TaskDtoResult struct {
	DuplicatedCount int       `json:"duplicated_count"`
	Tasks           []TaskDto `json:"tasks"`
}

type Params struct {
	URL string `json:"url"`
}

type TaskDto struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	IsActive  bool   `json:"is_active"`
	Params    Params `json:"params"`
	Object    string `json:"object"`
}

func CreateTask(id string, url string) *TaskDtoResult {
	createTaskRequest := lobstrio.NewCreateTaskRequest(id, []lobstrio.Task{
		{Url: url},
	})

	data, err := GetClient().MiscApi.
		CreateTask(context.Background()).
		CreateTaskRequest(*createTaskRequest).
		Execute()

	checkError(data, err)

	var task *TaskDtoResult

	json.NewDecoder(data.Body).Decode(&task)

	return task
}
