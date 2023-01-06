package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-guides-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type GuidesCommandableHttpClientV1 struct {
	client  *version1.GuidesCommandableHttpClientV1
	fixture *GuidesClientFixtureV1
}

func newGuidesCommandableHttpClientV1() *GuidesCommandableHttpClientV1 {
	return &GuidesCommandableHttpClientV1{}
}

func (c *GuidesCommandableHttpClientV1) setup(t *testing.T) *GuidesClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewGuidesCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewGuidesClientFixtureV1(c.client)

	return c.fixture
}

func (c *GuidesCommandableHttpClientV1) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newGuidesCommandableHttpClientV1()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
