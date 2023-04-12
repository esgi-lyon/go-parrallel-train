package lobstriofacade

import (
	"log"
	"testing"

	"github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
	"github.com/stretchr/testify/assert"
)

func TestGetcluster(t *testing.T) {
	
	cluster := Getcluster(CLUSTER_ID)
	log.Println(cluster)

	assert.NotNil(t, cluster.Name)
}

func TestGetClusterWebHook(t *testing.T) {
	clusterWebhook := GetClusterWebHook(CLUSTER_ID)
	log.Println(clusterWebhook)

	assert.NotNil(t, clusterWebhook.IsActive)
}

func TestUpdateCluster(t *testing.T) {
	log.Println("Update cluster")
	updCluster := UpdateCluster(
		CLUSTER_ID,
		lobstrio.UpdateClusterRequest{
			Params: lobstrio.UpdateClusterRequestParams{
				MaxResults: 100,
				MaxPages:   5,
			},
			RunNotify: "on_error",
		})

	log.Println(updCluster)

	assert.NotNil(t, updCluster.Name)
}

func TestRunCluster(t *testing.T) {
	log.Println("Run cluster")

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	run := RunCluster(CLUSTER_ID)
	log.Println(run)

	assert.NotNil(t, run.EndedAt)
}

func ViewRunTest(t *testing.T) {
	viewRun := ViewRun(RUN_ID)

	log.Println("View cluster Run")
	log.Println(viewRun)

	assert.NotNil(t, viewRun.CreatedAt)
}
