package internal

import (
	"context"
	"fmt"

	"github.com/zr-hebo/project-template-go/config"
)

func ExecMain(_ context.Context) {
	controller()
}

func controller() {
	fmt.Printf("hello, welcome to %s!\n", config.ProjectName)
}
