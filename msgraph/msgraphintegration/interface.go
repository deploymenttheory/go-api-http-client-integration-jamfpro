// msgraph/msgraphintegration/interface.go
package msgraphintegration

import (
	"net/http"

	"github.com/deploymenttheory/go-api-http-client/logger"
)

// MSGraphAPIHandler implements the APIHandler interface for the Microsoft Graph API.
type Integration struct {
	BaseDomain           string
	AuthMethodDescriptor string
	Logger               logger.Logger
	auth                 authInterface

	TenantID   string // TenantID used for constructing the authentication endpoint.
	TenantName string // TenantName used for constructing the authentication endpoint.
}

// Info

// TODO migrate strings
func (m *Integration) Domain() string {
	return m.BaseDomain
}

// TODO migrate strings
func (m *Integration) GetAuthMethodDescriptor() string {
	return m.AuthMethodDescriptor
}

// Utilities

// TODO migrate strings
func (m *Integration) CheckRefreshToken() error {
	return m.checkRefreshToken()
}

// TODO migrate strings
func (m *Integration) PrepRequestParamsAndAuth(req *http.Request) error {
	err := m.prepRequest(req)
	return err
}

// TODO migrate strings
func (m *Integration) PrepRequestBody(body interface{}, method string, endpoint string) ([]byte, error) {
	return m.marshalRequest(body, method, endpoint)
}

// TODO migrate strings
func (m *Integration) MarshalMultipartRequest(fields map[string]string, files map[string]string) ([]byte, string, error) {
	return m.marshalMultipartRequest(fields, files)
}

// TODO migrate strings
func (m *Integration) GetSessionCookies() ([]*http.Cookie, error) {
	domain := m.Domain()
	return m.getSessionCookies(domain)
}