package characters

import (
	"context"

	entity "github.com/ivantedja/xmarvel/entity"
)

type Usecase interface {
	Search(ctx context.Context) ([]*entity.Character, error)
}
