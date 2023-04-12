package lobstriofacade

import (
	"log"
	"os"
	"testing"
)

const (
	AUTH_TOKEN = "dc71bf7d4c072a9565c06d62b0945f55cb1f35f7"
	CLUSTER_ID = "9fb56bc8d697ff45d47ad9ed332e262b"
	RUN_ID     = "e2538ac3aa9147a4911f685a2821648c"
)

func setupTest(tb testing.TB) func(tb testing.TB) {
	log.Println("Setup test")
	os.Setenv("LOBSTRIOS_API_TOKEN", AUTH_TOKEN)

	return func(_ testing.TB) {
		log.Println("teardown test")
	}
}
