package scenario

import (
	"bytes"

	"github.com/cerberauth/vulnapi/internal/auth"
	"github.com/cerberauth/vulnapi/internal/request"
	"github.com/cerberauth/vulnapi/report"
	"github.com/cerberauth/vulnapi/scan"
	discoverablegraphql "github.com/cerberauth/vulnapi/scan/discover/discoverable_graphql"
	discoverableopenapi "github.com/cerberauth/vulnapi/scan/discover/discoverable_openapi"
)

func NewURLScan(method string, url string, data string, client *request.Client, reporter *report.Reporter) (*scan.Scan, error) {
	if client == nil {
		client = request.DefaultClient
	}

	var securitySchemes []auth.SecurityScheme
	securityScheme, err := detectSecurityScheme(client.Header)
	if err != nil {
		return nil, err
	}

	if securityScheme != nil {
		securitySchemes = append(securitySchemes, securityScheme)
	} else {
		securitySchemes = append(securitySchemes, auth.NewNoAuthSecurityScheme())
	}

	body := bytes.NewBuffer([]byte(data))

	url = addDefaultProtocolWhenMissing(url)
	operation, err := request.NewOperation(method, url, body, client)
	operation.SetSecuritySchemes(securitySchemes)
	if err != nil {
		return nil, err
	}

	if err := operation.IsReachable(); err != nil {
		return nil, err
	}

	operations := request.Operations{operation}
	urlScan, err := scan.NewScan(operations, reporter)
	if err != nil {
		return nil, err
	}

	urlScan.AddScanHandler(scan.NewOperationScanHandler(discoverableopenapi.DiscoverableOpenAPIScanID, discoverableopenapi.ScanHandler))
	urlScan.AddScanHandler(scan.NewOperationScanHandler(discoverablegraphql.DiscoverableGraphQLPathScanID, discoverablegraphql.ScanHandler))
	WithAllCommonScans(urlScan)

	return urlScan, nil
}
