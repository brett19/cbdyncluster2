package dockerdeploy

import (
	"context"

	"github.com/couchbaselabs/cbdinocluster/deployment"
	"golang.org/x/mod/semver"
)

type ImageDef struct {
	Version             string
	BuildNo             int
	UseCommunityEdition bool
	UseServerless       bool
	UseColumnar         bool
}

type ImageRef struct {
	ImagePath string
}

type ImageProvider interface {
	GetImage(ctx context.Context, def *ImageDef) (*ImageRef, error)
	ListImages(ctx context.Context) ([]deployment.Image, error)
	SearchImages(ctx context.Context, version string) ([]deployment.Image, error)
	GetImageRaw(ctx context.Context, imagePath string) (*ImageRef, error)
}

func CompareImageDefs(a, b *ImageDef) int {
	c := semver.Compare("v"+a.Version, "v"+b.Version)
	if c != 0 {
		return c
	}

	if a.BuildNo < b.BuildNo {
		return -1
	} else if a.BuildNo > b.BuildNo {
		return +1
	}

	if a.UseCommunityEdition && !b.UseCommunityEdition {
		return -1
	} else if !a.UseCommunityEdition && b.UseCommunityEdition {
		return +1
	}

	if !a.UseServerless && b.UseServerless {
		return -1
	} else if a.UseServerless && !b.UseServerless {
		return +1
	}

	if !a.UseColumnar && b.UseColumnar {
		return -1
	} else if a.UseColumnar && !b.UseColumnar {
		return +1
	}

	return 0
}
