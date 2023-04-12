package lobstriofacade

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

func GetClient() *lobstrio.APIClient {
	var Authorization = "Token " + os.Getenv("LOBSTRIOS_API_TOKEN")
	cli := lobstrio.NewAPIClient(lobstrio.NewConfiguration())
	cli.GetConfig().AddDefaultHeader("Authorization", Authorization)
	
	return cli
}

func checkError(response *http.Response, err error) {
	if err != nil {
		var body map[string]interface{}
		json.NewDecoder(response.Body).Decode(&body)
		log.Panicf(
			"Error in http request with status %d. Response %v", response.StatusCode, body)
	}
}