package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-commons-gox/random"
)

type GuidesMockClientV1 struct {
	guides []*GuideV1
}

func NewGuidesMockClientV1() *GuidesMockClientV1 {
	return &GuidesMockClientV1{
		guides: make([]*GuideV1, 0),
	}
}

func (c *GuidesMockClientV1) contains(array1 []string, array2 []string) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if i < len(array2) {
				if array1[i] == array2[i] {
					return true
				}
			} else {
				break
			}
		}
	}

	return false
}

func (c *GuidesMockClientV1) composeFilter(filter *data.FilterParams) func(item *GuideV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	id, idOk := filter.GetAsNullableString("id")
	guideType, guideTypeOk := filter.GetAsNullableString("type")
	app, appOk := filter.GetAsNullableString("app")
	name, nameOk := filter.GetAsNullableString("name")
	status, statusOk := filter.GetAsNullableString("status")
	tagsString := filter.GetAsString("tags")
	tags := make([]string, 0)

	if tagsString != "" {
		tags = data.TagsProcessor.CompressTags([]string{tagsString})
	}

	return func(item *GuideV1) bool {
		if idOk && id != item.Id {
			return false
		}
		if guideTypeOk && guideType != item.Type {
			return false
		}
		if appOk && app != item.App {
			return false
		}
		if nameOk && name != item.Name {
			return false
		}
		if statusOk && status != item.Status {
			return false
		}
		if len(tags) > 0 && !c.contains(item.AllTags, tags) {
			return false
		}
		return true
	}
}

func (c *GuidesMockClientV1) GetGuides(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*GuideV1], error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*GuideV1, 0)
	for _, v := range c.guides {
		item := *v
		if filterFunc(&item) {
			items = append(items, &item)
		}
	}
	return *data.NewDataPage(items, len(c.guides)), nil
}

func (c *GuidesMockClientV1) GetRandomGuide(ctx context.Context, correlationId string, filter *data.FilterParams) (*GuideV1, error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*GuideV1, 0)
	for _, v := range c.guides {
		item := v
		if filterFunc(item) {
			items = append(items, item)
		}
	}

	buf := *items[random.Integer.Next(0, len(items))]
	return &buf, nil
}

func (c *GuidesMockClientV1) GetGuideById(ctx context.Context, correlationId string, guideId string) (result *GuideV1, err error) {
	for _, v := range c.guides {
		if v.Id == guideId {
			buf := *v
			result = &buf
			break
		}
	}
	return result, err
}

func (c *GuidesMockClientV1) CreateGuide(ctx context.Context, correlationId string, guide *GuideV1) (*GuideV1, error) {
	buf := *guide
	c.guides = append(c.guides, &buf)
	return guide, nil
}

func (c *GuidesMockClientV1) UpdateGuide(ctx context.Context, correlationId string, guide *GuideV1) (*GuideV1, error) {
	if guide == nil {
		return nil, nil
	}

	var index = -1
	for i, v := range c.guides {
		if v.Id == guide.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	buf := *guide
	c.guides[index] = &buf
	return guide, nil
}

func (c *GuidesMockClientV1) DeleteGuideById(ctx context.Context, correlationId string, guideId string) (*GuideV1, error) {
	var index = -1
	for i, v := range c.guides {
		if v.Id == guideId {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}
	var item = c.guides[index]
	if index < len(c.guides) {
		c.guides = append(c.guides[:index], c.guides[index+1:]...)
	} else {
		c.guides = c.guides[:index]
	}
	return item, nil
}
