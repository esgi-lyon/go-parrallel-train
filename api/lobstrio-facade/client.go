package lobstriofacade

import (
	"fmt"
	"os"

	"github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

func GetClient() *lobstrio.APIClient {
	cnf := lobstrio.NewConfiguration()
	cnf.AddDefaultHeader("Authorization", fmt.Sprintf("Token %s", os.Getenv("LOBSTRIOS_API_TOKEN")))
	cli := lobstrio.NewAPIClient(cnf)
	
	return cli
}