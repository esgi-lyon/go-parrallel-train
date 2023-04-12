package lobstriofacade

import (
	"context"
	"encoding/json"
	// "fmt"

	"github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

type Events struct {
	RunCreated bool `json:"Run.created"`
	RunSuccess bool `json:"Run.success"`
	RunExportDone bool `json:"Run.export_done"`
	RunEndedWithError bool `json:"Run.ended_with_error"`
	RunUnexpectedError bool `json:"Run.unexpected_error"`
}

type WebhookFields struct {
	URL interface{} `json:"url"`
	Events Events `json:"events"`
	IsActive bool `json:"is_active"`
}

type ClusterDtoParams struct {
	HoursBack interface{} `json:"hours_back"`
	MaxDate interface{} `json:"max_date"`
	MaxPages int `json:"max_pages"`
	MaxResults int `json:"max_results"`
	OnlineShop bool `json:"online_shop"`
}

type ClusterDto struct {
	Accounts interface{} `json:"accounts"`
	Concurrency int `json:"concurrency"`
	ExportUniqueResults bool `json:"export_unique_results"`
	Name string `json:"name"`
	NoLineBreaks bool `json:"no_line_breaks"`
	Params ClusterDtoParams `json:"params"`
	RunNotify string `json:"run_notify"`
	ToComplete bool `json:"to_complete"`
	WebhookFields WebhookFields `json:"webhook_fields"`
}

type ClusterDtoPages struct {
	Count int `json:"count"`
	Limit int `json:"limit"`
	Page int `json:"page"`
	TotalPages int `json:"total_pages"`
	ResultFrom int `json:"result_from"`
	ResultTo int `json:"result_to"`
	Data []ClusterDto `json:"data"`
}

func ListCluster(id string) *ClusterDtoPages {
	data, err := GetClient().MiscApi.
		Listcluster(context.Background()).
		Id(id).
		ContentType("application/json").
		Execute()

	if err != nil {
		panic(err)
	}

	var clusterPages *ClusterDtoPages

	json.NewDecoder(data.Body).Decode(&clusterPages)

	return clusterPages
}

func GetClusterWebHook(id string) *WebhookFields {

	if len(ListCluster(id).Data) == 0 {
		return &WebhookFields{}
	}

	a := &ListCluster(id).Data[0]

	return &a.WebhookFields
}

func Getcluster(id string) *ClusterDto {
	data, err := GetClient().MiscApi.
		Getcluster(context.Background(), id).
		ContentType("application/json").
		Execute()

	if err != nil {
		panic(err)
	}

	var cluster *ClusterDto

	json.NewDecoder(data.Body).Decode(&cluster)

	return cluster
}

func UpdateCluster(id string, cluster lobstrio.UpdateclusterRequest) *ClusterDto {
	data, err := GetClient().MiscApi.
		Updatecluster(context.Background(), id).
		UpdateclusterRequest(cluster).
		Execute()

	if err != nil {
		panic(err)
	}

	var clusterResult *ClusterDto

	json.NewDecoder(data.Body).Decode(&clusterResult)

	return clusterResult
}
