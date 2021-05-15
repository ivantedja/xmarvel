package characters

import (
	"context"
	"github.com/ivantedja/xmarvel/entity"
)

type Usecase interface {
	Search(ctx context.Context) ([]uint, error)
	Show(ctx context.Context, ID int) (*entity.Character, error)
}
