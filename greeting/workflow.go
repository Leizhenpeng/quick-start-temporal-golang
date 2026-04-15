package greeting

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func SayHelloWorkflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	if err := workflow.ExecuteActivity(ctx, Greet, name).Get(ctx, &result); err != nil {
		return "", err
	}
	return result, nil
}

