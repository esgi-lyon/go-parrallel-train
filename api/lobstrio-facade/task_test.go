package lobstriofacade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ExpectedURL = "https://www.leboncoin.fr/ventes_immobilieres/offres/auvergne_rhone_alpes/ain"
)

func TestCreateTask(t *testing.T) {
	
	tasksResult := CreateTask(CLUSTER_ID, ExpectedURL)

	task := tasksResult.Tasks[len(tasksResult.Tasks) - 1]
	assert.NotNil(t, task)
	assert.Equal(t, 
		ExpectedURL, task.Params.URL)
}

