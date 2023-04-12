package lobstriofacade

import (
	"context"
	"encoding/json"
	"time"

	"github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

type Events struct {
	RunCreated         bool `json:"Run.created"`
	RunSuccess         bool `json:"Run.success"`
	RunExportDone      bool `json:"Run.export_done"`
	RunEndedWithError  bool `json:"Run.ended_with_error"`
	RunUnexpectedError bool `json:"Run.unexpected_error"`
}

type WebhookFields struct {
	URL      interface{} `json:"url"`
	Events   Events      `json:"events"`
	IsActive bool        `json:"is_active"`
}

type ClusterDtoParams struct {
	HoursBack  interface{} `json:"hours_back"`
	MaxDate    interface{} `json:"max_date"`
	MaxPages   int         `json:"max_pages"`
	MaxResults int         `json:"max_results"`
	OnlineShop bool        `json:"online_shop"`
}

type ClusterDto struct {
	Accounts            interface{}      `json:"accounts"`
	Concurrency         int              `json:"concurrency"`
	ExportUniqueResults bool             `json:"export_unique_results"`
	Name                string           `json:"name"`
	NoLineBreaks        bool             `json:"no_line_breaks"`
	Params              ClusterDtoParams `json:"params"`
	RunNotify           string           `json:"run_notify"`
	ToComplete          bool             `json:"to_complete"`
	WebhookFields       WebhookFields    `json:"webhook_fields"`
}

type ClusterDtoPages struct {
	Count      int          `json:"count"`
	Limit      int          `json:"limit"`
	Page       int          `json:"page"`
	TotalPages int          `json:"total_pages"`
	ResultFrom int          `json:"result_from"`
	ResultTo   int          `json:"result_to"`
	Data       []ClusterDto `json:"data"`
}

type RunClusterDto struct {
	Id             string    `json:"id"`
	Object         string    `json:"object"`
	StartedAt      time.Time `json:"started_at"`
	CreatedAt      time.Time `json:"created_at"`
	IsDone         bool      `json:"is_done"`
	Duration       int       `json:"duration"`
	CreditUsed     int       `json:"credit_used"`
	EndedAt        time.Time `json:"ended_at"`
	Origin         string    `json:"origin"`
	Status         string    `json:"status"`
	TotalResults   int       `json:"total_results"`
	DoneReason     string    `json:"done_reason"`
	DoneReasonDesc string    `json:"done_reason_desc"`
	ExportCount    int       `json:"export_count"`
	Cluster        string    `json:"cluster"`
}

type RunClusterDtoPages struct {
	Count      int          `json:"count"`
	Limit      int          `json:"limit"`
	Page       int          `json:"page"`
	TotalPages int          `json:"total_pages"`
	ResultFrom int          `json:"result_from"`
	ResultTo   int          `json:"result_to"`
	Data       []RunClusterDto `json:"data"`
}

func ListCluster(id string) *ClusterDtoPages {
	data, err := GetClient().MiscApi.
		ListCluster(context.Background()).
		Id(id).
		ContentType("application/json").
		Execute()

	checkError(data, err)

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
		GetCluster(context.Background(), id).
		ContentType("application/json").
		Execute()

	checkError(data, err)

	var cluster *ClusterDto

	json.NewDecoder(data.Body).Decode(&cluster)

	return cluster
}

func UpdateCluster(id string, cluster lobstrio.UpdateClusterRequest) *ClusterDto {
	data, err := GetClient().MiscApi.
		UpdateCluster(context.Background(), id).
		UpdateClusterRequest(cluster).
		Execute()

	checkError(data, err)

	var clusterResult *ClusterDto

	json.NewDecoder(data.Body).Decode(&clusterResult)

	return clusterResult
}

func RunCluster(id string) *RunClusterDto {
	data, err := GetClient().MiscApi.
		RunCluster(context.Background()).
		RunClusterRequest(lobstrio.RunClusterRequest{Cluster: id}).
		Execute()

	checkError(data, err)

	var clusterRun *RunClusterDto

	json.NewDecoder(data.Body).Decode(&clusterRun)

	return clusterRun
}

func ViewRun(runId string) *RunClusterDto {
	data, err := GetClient().MiscApi.
		ViewRun(context.Background(), runId).
		ContentType("application/json").
		Execute()

	checkError(data, err)

	var run *RunClusterDto

	json.NewDecoder(data.Body).Decode(&run)

	return run
}


func ListRun(limit int32, page int32) *RunClusterDtoPages {
	data, err := GetClient().MiscApi.
		ListRun(context.Background()).Limit(limit).Page(page).
		ContentType("application/json").
		Execute()

	checkError(data, err)

	var run *RunClusterDtoPages

	json.NewDecoder(data.Body).Decode(&run)

	return run
}
