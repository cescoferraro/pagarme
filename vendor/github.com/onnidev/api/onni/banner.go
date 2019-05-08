package onni

import (
	"context"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
)

// PublishedBanners asdfnjkas
func PublishedBanners(ctx context.Context) ([]types.Banner, error) {
	bannerRepo := ctx.Value(middlewares.BannerRepoKey).(interfaces.BannerRepo)
	banners, err := bannerRepo.GetPublishedBanners()
	if err != nil {
		return banners, err
	}
	return banners, nil
}

// AllBanners asdfnjkas
func AllBanners(ctx context.Context) ([]types.Banner, error) {
	bannerRepo := ctx.Value(middlewares.BannerRepoKey).(interfaces.BannerRepo)
	banners, err := bannerRepo.AllBanners()
	if err != nil {
		return banners, err
	}
	return banners, nil
}
