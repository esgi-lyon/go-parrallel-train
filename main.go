package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	AUTH_TOKEN = "Token 85740612a05ece90fb7bbbf7b507947ffd530cf8"
	CLUSTER_ID = "a3f12e8688de73f1f6bbd0fc8c87354a"
)

func main() {
	client := http.Client{}
	res, err := client.Get("https://api.lobstr.io/v1/clusters/" + CLUSTER_ID)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

}

func getCluster() map[string]interface{} {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.lobstr.io/v1/clusters/"+CLUSTER_ID, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", AUTH_TOKEN)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	return data
}

func createTask(region, department string) map[string]interface{} {
	client := http.Client{}

	bodyMap := map[string]interface{}{
		"cluster": CLUSTER_ID,
		"tasks": []map[string]interface{}{
			{
				"url": "https://www.leboncoin.fr/ventes_immobilieres/offres/" + region + "/" + department,
			},
		},
	}

	//bodyMap to ioreader
	body, err := json.Marshal(bodyMap)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "https://api.lobstr.io/v1/clusters/tasks", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", AUTH_TOKEN)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	return data
}

func updatCluster(max_results, max_pages string) map[string]interface{} {
	client := http.Client{}

	bodyMap := map[string]interface{}{

		"name":                  "Leboncoin Listings Search Export",
		"concurrency":           1,
		"export_unique_results": true,
		"no_line_breaks":        true,
		"to_complete":           false,
		"params": map[string]interface{}{
			"max_date":    nil,
			"max_results": max_results,
			"max_pages":   max_pages,
			"hours_back":  nil,
			"online_shop": false,
		},
	}

	//bodyMap to ioreader
	body, err := json.Marshal(bodyMap)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", "https://api.lobstr.io/v1/clusters/"+CLUSTER_ID, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", AUTH_TOKEN)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	return data
}

func activateTask(task_id string) map[string]interface{} {
	client := http.Client{}

	req, err := http.NewRequest("POST", "https://api.lobstr.io/v1/clusters/tasks/"+task_id+"/activate", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", AUTH_TOKEN)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	return data
}

func runCluster() map[string]interface{} {
	client := http.Client{}

	bodyMap := map[string]interface{}{
		"cluster": CLUSTER_ID,
	}

	//bodyMap to ioreader
	body, err := json.Marshal(bodyMap)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "https://api.lobstr.io/v1/clusters/runs", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", AUTH_TOKEN)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	return data
}

func listRuns() map[string]interface{} {
	client := http.Client{}

	req, err := http.NewRequest("GET", "https://api.lobstr.io/v1/clusters/runs", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", AUTH_TOKEN)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	return data
}
