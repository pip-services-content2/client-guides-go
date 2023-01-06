package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type GuidesNullClientV1 struct {
}

func NewGuidesNullClientV1() *GuidesNullClientV1 {
	return &GuidesNullClientV1{}
}

func (c *GuidesNullClientV1) GetGuides(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*GuideV1], error) {
	return *data.NewEmptyDataPage[*GuideV1](), nil
}

func (c *GuidesNullClientV1) GetRandomGuide(ctx context.Context, correlationId string, filter *data.FilterParams) (*GuideV1, error) {
	return nil, nil
}

func (c *GuidesNullClientV1) GetGuideById(ctx context.Context, correlationId string, guideId string) (*GuideV1, error) {
	return nil, nil
}

func (c *GuidesNullClientV1) CreateGuide(ctx context.Context, correlationId string, guide *GuideV1) (*GuideV1, error) {
	return guide, nil
}

func (c *GuidesNullClientV1) UpdateGuide(ctx context.Context, correlationId string, guide *GuideV1) (*GuideV1, error) {
	return guide, nil
}

func (c *GuidesNullClientV1) DeleteGuideById(ctx context.Context, correlationId string, guideId string) (*GuideV1, error) {
	return nil, nil
}
