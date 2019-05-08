package onni

import (
	"context"
	"errors"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
)

// MusicStyles TODO: NEEDS COMMENT INFO
func MusicStyles(ctx context.Context, names []string) ([]types.Style, error) {
	musicrepo, ok := ctx.Value(middlewares.MusicStylesRepoKey).(interfaces.MusicStylesRepo)
	if !ok {
		err := errors.New("assert bug")
		return []types.Style{}, err
	}
	all, err := musicrepo.ByNames(names)
	if err != nil {
		return []types.Style{}, err
	}
	return all, nil
}
