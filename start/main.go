package main

import (
	"context"
	"log"
	"os"

	"quick-start-temporal-golang/greeting"

	"go.temporal.io/sdk/client"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run start/main.go <name>")
	}

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: "my-task-queue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, greeting.SayHelloWorkflow, os.Args[1])
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result string
	if err := we.Get(context.Background(), &result); err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}

