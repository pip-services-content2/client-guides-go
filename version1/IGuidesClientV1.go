package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IGuidesClientV1 interface {
	GetGuides(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*GuideV1], error)

	GetRandomGuide(ctx context.Context, correlationId string, filter *data.FilterParams) (*GuideV1, error)

	GetGuideById(ctx context.Context, correlationId string, guideId string) (*GuideV1, error)

	CreateGuide(ctx context.Context, correlationId string, guide *GuideV1) (*GuideV1, error)

	UpdateGuide(ctx context.Context, correlationId string, guide *GuideV1) (*GuideV1, error)

	DeleteGuideById(ctx context.Context, correlationId string, guideId string) (*GuideV1, error)
}
