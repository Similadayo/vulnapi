package discover_test

import (
	"net/http"
	"testing"

	"github.com/cerberauth/vulnapi/internal/auth"
	"github.com/cerberauth/vulnapi/internal/request"
	"github.com/cerberauth/vulnapi/report"
	"github.com/cerberauth/vulnapi/scan/discover"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckSignatureHeaderWithSignatureHeader(t *testing.T) {
	client := request.DefaultClient
	httpmock.ActivateNonDefault(client.Client)
	defer httpmock.DeactivateAndReset()

	token := "token"
	securityScheme := auth.NewAuthorizationBearerSecurityScheme("default", &token)
	operation, _ := request.NewOperation(client, http.MethodGet, "http://localhost:8080/", nil, nil, nil)
	vulnerabilityReport := report.VulnerabilityReport{
		SeverityLevel: discover.ServerSignatureSeverityLevel,

		OWASP2023Category: discover.ServerSignatureOWASP2023Category,

		ID:   discover.ServerSignatureVulnerabilityID,
		Name: discover.ServerSignatureVulnerabilityName,
		URL:  discover.ServerSignatureVulnerabilityURL,
	}

	httpmock.RegisterResponder(operation.Method, operation.Request.URL.String(), httpmock.NewBytesResponder(204, nil).HeaderAdd(http.Header{"Server": []string{"Apache/2.4.29 (Ubuntu)"}}))

	report, err := discover.ServerSignatureScanHandler(operation, securityScheme)

	require.NoError(t, err)
	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.True(t, report.HasFailedVulnerabilityReport())
	assert.Equal(t, report.GetVulnerabilityReports()[0], &vulnerabilityReport)
}

func TestCheckSignatureHeaderWithoutSignatureHeader(t *testing.T) {
	client := request.DefaultClient
	httpmock.ActivateNonDefault(client.Client)
	defer httpmock.DeactivateAndReset()

	token := "token"
	securityScheme := auth.NewAuthorizationBearerSecurityScheme("default", &token)
	operation, _ := request.NewOperation(client, http.MethodGet, "http://localhost:8080/", nil, nil, nil)
	httpmock.RegisterResponder(operation.Method, operation.Request.URL.String(), httpmock.NewBytesResponder(204, nil))

	report, err := discover.ServerSignatureScanHandler(operation, securityScheme)

	require.NoError(t, err)
	assert.Equal(t, 1, httpmock.GetTotalCallCount())
	assert.False(t, report.HasFailedVulnerabilityReport())
}
