package greeting

import (
	"context"
	"fmt"
)

func Greet(ctx context.Context, name string) (string, error) {
	_ = ctx
	return fmt.Sprintf("Hello %s", name), nil
}

