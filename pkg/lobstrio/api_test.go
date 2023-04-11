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
func TestGetTaskResults(t *testing.T) {
	task := CreateTask("ile_de_france", "paris")
	taskId := task["id"].(string)
	task = ActivateTask(taskId)
	if task["is_active"] != "false" {
		t.Errorf("Cluster id is not correct")
	}
}
