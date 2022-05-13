package init

import (
	"context"

	"github.com/zr-hebo/project-template-go/internal/web"
)

func PrepareEnv(ctx context.Context) {
	web.StartWebService()
}
