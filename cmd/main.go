package main

import (
	"context"

	"github.com/zr-hebo/project-template-go/init"
	"github.com/zr-hebo/project-template-go/internal"
)

func main() {
	showCommand()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	init.PrepareEnv(ctx)
	internal.ExecMain(ctx)
}
