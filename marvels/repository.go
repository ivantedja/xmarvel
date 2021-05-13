package marvels

import (
	"context"

	entity "github.com/ivantedja/xmarvel/entity"
)

type MarvelsRepository interface {
	Search(ctx context.Context, filter map[string]string) (*entity.CharacterCollection, error)
}
