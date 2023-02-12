package main

import (
	"fmt"
	"log"
	"ocitaskclient"
	"ocitaskprovider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

const OciTaskServHostUrl string = "http://192.9.237.204:8081/v1/ocitaskserv"
const OciTaskRestServHostUrl string = "http://138.2.233.236:8080/application/v1/ocitaskrestservice"

func main() {
	log.Println("OCI Task Management Service Terraform Provider Start")

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			prov := ocitaskprovider.MakeOciTaskServProvider()
			return prov.Provider()
		},
	})

	// taskId := CreateTask()
	// UpdateTask(taskId)

	// var taskId int64 = 22
	// GetTask(&taskId)
	// DeleteTask(&taskId)
	// taskId = 23
	// GetTask(&taskId)
	// DeleteTask(&taskId)

	log.Println("All Done")
}

func GetTask(taskId *int64) {
	var hostUrl = OciTaskRestServHostUrl
	ociClient := ocitaskclient.MakeOciTaskServClient(&hostUrl)
	resp, _ := ociClient.GetTask(taskId)
	if resp != nil {
		if resp.Err != nil {
			ociErr, err := resp.Err.Serialize()
			if err != nil {
				log.Println(fmt.Sprintf("Get Task Failed - Unknown Error - TaskId=%d", *taskId))
			} else {
				log.Println(fmt.Sprintf("Get Task Failed - TaskId=%d - error=%s", *taskId, ociErr))
			}
		} else {
			strTask, _ := resp.Task.Serialize()
			log.Println(fmt.Sprintf("Get Task Succeeded - Task=%s", strTask))
		}
	} else {
		log.Println(fmt.Sprintf("Get Task Failed - Unknown Error - TaskId=%d", *taskId))
	}
}

func DeleteTask(taskId *int64) {
	var hostUrl = OciTaskRestServHostUrl
	ociClient := ocitaskclient.MakeOciTaskServClient(&hostUrl)
	resp, _ := ociClient.DeleteTask(taskId)
	if resp != nil {
		if resp.Err != nil {
			ociErr, err := resp.Err.Serialize()
			if err != nil {
				log.Println(fmt.Sprintf("Delete Task Failed - Unknown Error - TaskId=%d", *taskId))
			} else {
				log.Println(fmt.Sprintf("Delete Task Failed - TaskId=%d - error=%s", *taskId, ociErr))
			}
		} else {
			log.Println(fmt.Sprintf("Delete Task Succeeded - TaskId=%d", *taskId))
		}
	} else {
		log.Println(fmt.Sprintf("Delete Task Failed - Unknown Error - TaskId=%d", *taskId))
	}
}

func CreateTask() *int64 {
	var taskId *int64 = nil
	taskReq := make(map[string]interface{})
	taskReq["title"] = "Task from oci task terraform provider"
	taskReq["description"] = "Task description"
	taskReq["priority"] = 1
	taskReq["completed"] = false
	taskReq["startDate"] = "2023-02-10"
	taskReq["dueDate"] = "2023-02-20"

	taskReqs := make([]interface{}, 0)
	taskReqs = append(taskReqs, taskReq)

	ociReq, _ := ocitaskclient.MakeOciTaskServRequest(&taskReqs[0])

	var hostUrl = OciTaskRestServHostUrl
	ociClient := ocitaskclient.MakeOciTaskServClient(&hostUrl)
	resp, _ := ociClient.CreateTask(ociReq)
	if resp != nil {
		if resp.Err != nil {
			ociErr, err := resp.Err.Serialize()
			if err != nil {
				log.Println("Create Task Failed - Unknown Error")
			} else {
				log.Println(fmt.Sprintf("Create Task Failed - error=%s", ociErr))
			}
		} else {
			log.Println(fmt.Sprintf("Create Task Succeeded - TaskId=%d", *resp.TaskId))
			taskId = resp.TaskId
		}
	} else {
		log.Println("Create Task Failed - Unknown Error")
	}

	return taskId
}

func UpdateTask(taskId *int64) {
	taskReq := make(map[string]interface{})
	taskReq["title"] = "Updated: Task from oci task terraform provider"
	taskReq["description"] = "Updated: Task description"
	taskReq["priority"] = 5
	taskReq["completed"] = false
	taskReq["startDate"] = "2023-02-11"
	taskReq["dueDate"] = "2023-02-21"

	taskReqs := make([]interface{}, 0)
	taskReqs = append(taskReqs, taskReq)

	ociReq, _ := ocitaskclient.MakeOciTaskServRequest(&taskReqs[0])

	var hostUrl = OciTaskRestServHostUrl
	ociClient := ocitaskclient.MakeOciTaskServClient(&hostUrl)
	resp, _ := ociClient.UpdateTask(taskId, ociReq)
	if resp != nil {
		if resp.Err != nil {
			ociErr, err := resp.Err.Serialize()
			if err != nil {
				log.Println("Update Task Failed - Unknown Error")
			} else {
				log.Println(fmt.Sprintf("Update Task Failed - error=%s", ociErr))
			}
		} else {
			log.Println(fmt.Sprintf("Update Task Succeeded - TaskId=%d", *resp.TaskId))
		}
	} else {
		log.Println("Update Task Failed - Unknown Error")
	}
}
