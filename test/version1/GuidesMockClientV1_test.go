package test_version1

import (
	"testing"

	"github.com/pip-services-content2/client-guides-go/version1"
)

type GuidesMockClientV1 struct {
	client  *version1.GuidesMockClientV1
	fixture *GuidesClientFixtureV1
}

func newGuidesMockClientV1() *GuidesMockClientV1 {
	return &GuidesMockClientV1{}
}

func (c *GuidesMockClientV1) setup(t *testing.T) *GuidesClientFixtureV1 {
	c.client = version1.NewGuidesMockClientV1()
	c.fixture = NewGuidesClientFixtureV1(c.client)
	return c.fixture
}

func (c *GuidesMockClientV1) teardown(t *testing.T) {
	c.client = nil
}

func TestMockCrudOperations(t *testing.T) {
	c := newGuidesMockClientV1()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
