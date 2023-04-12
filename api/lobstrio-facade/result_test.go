package lobstriofacade

import (
	"log"
	"testing"

	assetmanager "github.com/esgi-lyon/go-parrallel-train/pkg/asset-manager"
	"github.com/stretchr/testify/assert"
)

func TestGetResult(t *testing.T) {
	viewRunResult := GetResult[assetmanager.Asset](CLUSTER_ID, RUN_ID)
	
	log.Println(viewRunResult)

	assert.NotEmpty(t, viewRunResult.Data)
}
