package main

import (
	"fmt"
	"log"
	"machine-coding-basic-jira/internal/handler"
	"machine-coding-basic-jira/internal/model"
	"time"
)

func main() {
	log.Println("\nin main: ")
	taskHandler := handler.GetTaskHandler()

	feature, _ := taskHandler.CreateTask("abc", "s", time.Now(), model.TASK_TYPE_FEATURE)
	fmt.Println(feature)

}
