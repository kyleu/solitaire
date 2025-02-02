package app

import (
	"context"

	"github.com/kyleu/solitaire/app/util"
)

type GeneratedServices struct{}

func initGeneratedServices(ctx context.Context, st *State, logger util.Logger) GeneratedServices {
	return GeneratedServices{}
}
