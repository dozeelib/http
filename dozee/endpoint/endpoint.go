package endpoint

import (
	"fmt"
	"strings"
)

type endPoint struct {
	tls     bool
	service string
	region  string
	stage   string
	domain  string
	tld     string
}

type endPointBuilder struct {
	endPoint endPoint
}

func EndPointBuilder() *endPointBuilder {
	return &endPointBuilder{
		endPoint: endPoint{},
	}
}
func (builder *endPointBuilder) SetTls() *endPointBuilder {
	builder.endPoint.tls = true
	return builder
}

func (builder *endPointBuilder) SetService(service string) *endPointBuilder {
	builder.endPoint.service = service
	return builder
}

func (builder *endPointBuilder) SetRegion(region string) *endPointBuilder {
	builder.endPoint.region = region
	return builder
}

func (builder *endPointBuilder) SetStage(stage string) *endPointBuilder {
	builder.endPoint.stage = stage
	return builder
}

func (builder *endPointBuilder) SetDomain(domain string) *endPointBuilder {
	builder.endPoint.domain = domain
	return builder
}

func (builder *endPointBuilder) SetTld(tld string) *endPointBuilder {
	builder.endPoint.tld = tld
	return builder
}

func (builder *endPointBuilder) Build() string {
	var endpoint strings.Builder
	if builder.endPoint.tls {
		endpoint.WriteString("https://")
	} else {
		endpoint.WriteString("http://")
	}
	if builder.endPoint.service == "" {
		panic("service name is required")
	} else {
		endpoint.WriteString(builder.endPoint.service)
	}
	if builder.endPoint.region != "" {
		endpoint.WriteString(fmt.Sprintf("-%s", builder.endPoint.region))
	}
	if builder.endPoint.stage != "" {
		endpoint.WriteString(fmt.Sprintf("-%s", builder.endPoint.stage))
	}
	if builder.endPoint.domain == "" {
		panic("domain name is required")
	} else {
		endpoint.WriteString(fmt.Sprintf(".%s", builder.endPoint.domain))
	}
	if builder.endPoint.tld == "" {
		panic("tld name is required")
	} else {
		endpoint.WriteString(fmt.Sprintf(".%s", builder.endPoint.tld))
	}
	return endpoint.String()
}
