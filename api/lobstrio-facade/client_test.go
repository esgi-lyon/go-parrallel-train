package lobstriofacade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClient(t *testing.T) {

	cli := GetClient()

	assert.EqualValues(t,
		"Token " + AUTH_TOKEN, 
		cli.GetConfig().DefaultHeader["Authorization"],
	)
}
