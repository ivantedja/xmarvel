package characters

import (
	"context"
)

type Usecase interface {
	Search(ctx context.Context) ([]uint, error)
}
