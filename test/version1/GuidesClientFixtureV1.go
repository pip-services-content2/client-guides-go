package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-content2/client-guides-go/version1"
	"github.com/stretchr/testify/assert"
)

type GuidesClientFixtureV1 struct {
	Client version1.IGuidesClientV1
	GUIDE1 *version1.GuideV1
	GUIDE2 *version1.GuideV1
}

func NewGuidesClientFixtureV1(client version1.IGuidesClientV1) *GuidesClientFixtureV1 {
	return &GuidesClientFixtureV1{
		Client: client,
		GUIDE1: &version1.GuideV1{
			Id:     "1",
			Name:   "Name1",
			Type:   "introduction",
			App:    "Test App 1",
			MaxVer: 0,
			MinVer: 0,
			Status: "new",
		},
		GUIDE2: &version1.GuideV1{
			Id:      "2",
			Name:    "Name2",
			Tags:    []string{"TAG 1"},
			AllTags: []string{"tag1"},
			Type:    "new release",
			App:     "Test App 1",
			MaxVer:  1,
			MinVer:  2,
			Status:  "new",
		},
	}
}

func (c *GuidesClientFixtureV1) TestCrudOperations(t *testing.T) {
	// Create one guide
	guide1, err := c.Client.CreateGuide(context.Background(), "123", c.GUIDE1)
	assert.Nil(t, err)

	assert.NotNil(t, guide1)
	assert.Equal(t, guide1.Type, c.GUIDE1.Type)
	assert.Equal(t, guide1.App, c.GUIDE1.App)

	// Create another guide
	guide2, err := c.Client.CreateGuide(context.Background(), "123", c.GUIDE2)
	assert.Nil(t, err)

	assert.NotNil(t, guide2)
	assert.Equal(t, guide2.Type, c.GUIDE2.Type)
	assert.Equal(t, guide2.App, c.GUIDE2.App)

	// Get all guides
	page, err := c.Client.GetGuides(context.Background(), "123", nil, nil)
	assert.Nil(t, err)

	assert.NotNil(t, page)
	assert.Len(t, len(page.Data), 2)

	// Update the guide
	guide1.App = "New App 1"

	guide, err := c.Client.UpdateGuide(context.Background(), "123", guide1)
	assert.Nil(t, err)

	assert.NotNil(t, guide)
	assert.Equal(t, guide.App, "New App 1")
	assert.Equal(t, guide.Type, c.GUIDE1.Type)

	guide1 = guide

	// Delete guide
	_, err = c.Client.DeleteGuideById(context.Background(), "123", guide1.Id)
	assert.Nil(t, err)

	// Try to get delete guide
	guide, err = c.Client.GetGuideById(context.Background(), "123", guide1.Id)
	assert.Nil(t, err)
	assert.Nil(t, guide)
}
