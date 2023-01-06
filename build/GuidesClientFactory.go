package build

import (
	clients1 "github.com/pip-services-content2/client-guides-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type GuidesClientFactory struct {
	*cbuild.Factory
}

func NewGuidesClientFactory() *GuidesClientFactory {
	c := &GuidesClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-guides", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-guides", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-guides", "client", "commandable-http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewGuidesNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewGuidesMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewGuidesCommandableHttpClientV1)

	return c
}
