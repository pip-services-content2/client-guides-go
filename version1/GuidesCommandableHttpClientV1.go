package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type GuidesCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewGuidesCommandableHttpClientV1() *GuidesCommandableHttpClientV1 {
	return &GuidesCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/guides"),
	}
}

func (c *GuidesCommandableHttpClientV1) GetGuides(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*GuideV1], error) {
	params := data.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(ctx, "get_guides", correlationId, params)
	if err != nil {
		return *data.NewEmptyDataPage[*GuideV1](), err
	}

	return clients.HandleHttpResponse[data.DataPage[*GuideV1]](res, correlationId)
}

func (c *GuidesCommandableHttpClientV1) GetRandomGuide(ctx context.Context, correlationId string, filter *data.FilterParams) (*GuideV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"filter", filter,
	)

	res, err := c.CallCommand(ctx, "get_random_guide", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*GuideV1](res, correlationId)
}

func (c *GuidesCommandableHttpClientV1) GetGuideById(ctx context.Context, correlationId string, guideId string) (*GuideV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"guide_id", guideId,
	)

	res, err := c.CallCommand(ctx, "get_guide_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*GuideV1](res, correlationId)
}

func (c *GuidesCommandableHttpClientV1) CreateGuide(ctx context.Context, correlationId string, guide *GuideV1) (*GuideV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"guide", guide,
	)

	res, err := c.CallCommand(ctx, "create_guide", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*GuideV1](res, correlationId)
}

func (c *GuidesCommandableHttpClientV1) UpdateGuide(ctx context.Context, correlationId string, guide *GuideV1) (*GuideV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"guide", guide,
	)

	res, err := c.CallCommand(ctx, "update_guide", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*GuideV1](res, correlationId)
}

func (c *GuidesCommandableHttpClientV1) DeleteGuideById(ctx context.Context, correlationId string, guideId string) (*GuideV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"guide_id", guideId,
	)

	res, err := c.CallCommand(ctx, "delete_guide_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*GuideV1](res, correlationId)
}
