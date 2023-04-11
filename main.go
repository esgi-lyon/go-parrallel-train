package main

import (
	"fmt"
	"os"

	"github.com/esgi-lyon/go-parrallel-train/api/lobstrio-facade"
)

const (
	AUTH_TOKEN = "85740612a05ece90fb7bbbf7b507947ffd530cf8"
	CLUSTER_ID = "a3f12e8688de73f1f6bbd0fc8c87354a"
)

func main() {
	os.Setenv("LOBSTRIOS_API_TOKEN", AUTH_TOKEN)

	cluster := lobstriofacade.Getcluster(CLUSTER_ID)

	fmt.Println(cluster)

	task := lobstriofacade.
		CreateTask(CLUSTER_ID, "https://www.leboncoin.fr/ventes_immobilieres/offres/auvergne_rhone_alpes/ain")

	fmt.Println(task)
}

