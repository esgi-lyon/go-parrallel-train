package lobstrio

import "testing"

// Unit test dor getCluster
func TestGetCluster(t *testing.T) {
	cluster := GetCluster()
	if cluster["id"] != CLUSTER_ID {
		t.Errorf("Cluster id is not correct")
	}
}

// Unit test for createTask
func TestCreateTask(t *testing.T) {
	task := CreateTask("rhone_alpes", "rhone")
	if task["id"] != nil {
		t.Errorf("not created task")
	}
}

// Unit test for getTask
func TestActivateTask(t *testing.T) {
	task := CreateTask("rhone_alpes", "rhone")
	taskId := task["id"].(string)
	task = ActivateTask(taskId)
	if task["is_active"] != "true" {
		t.Errorf("Cluster id is not correct")
	}
}

// Unit test for getTaskResults
// func TestGetTaskResults(t *testing.T) {
// 	task := CreateTask("ile_de_france", "paris")
// 	taskId := task["id"].(string)
// 	task = ActivateTask(taskId)
// 	if task["is_active"] != "false" {
// 		t.Errorf("Cluster id is not correct")
// 	}
// }

func TestRunCluster(t *testing.T) {
	cluster := RunCluster()
	if cluster["id"] != CLUSTER_ID {
		t.Errorf("Cluster id is not correct")
	}

	run := ListRuns()

	if len(run["data"].([]interface{})) == 0 {
		t.Errorf("No run")
	}
}

func TestDownloadRun(t *testing.T) {
	run := ListRuns()
	runId := run["data"].([]interface{})[0].(map[string]interface{})["id"].(string)
	download := DownloadRun(runId)
	if download["id"] != runId {
		t.Errorf("Run id is not correct")
	}
}
