// microsoft/msgraphintegration/header_exceptions.go
package msgraphintegration

import (
	_ "embed"

	"encoding/json"
	"log"
)

// EndpointConfig is a struct that holds configuration details for a specific API endpoint.
// It includes what type of content it can accept and what content type it should send.
type EndpointConfig struct {
	Accept      string  `json:"accept"`       // Accept specifies the MIME type the endpoint can handle in responses.
	ContentType *string `json:"content_type"` // ContentType, if not nil, specifies the MIME type to set for requests sent to the endpoint. A pointer is used to distinguish between a missing field and an empty string.
}

// ConfigMap is a map that associates endpoint URL patterns with their corresponding configurations.
// The map's keys are strings that identify the endpoint, and the values are EndpointConfig structs
// that hold the configuration for that endpoint.
type ConfigMap map[string]EndpointConfig

// Variables
var configMap ConfigMap

// Embedded Resources
//
//go:embed msgraph_api_exceptions_configuration.json
var graph_api_exceptions_configuration []byte

func init() {
	err := loadAPIExceptionsConfiguration()
	if err != nil {
		log.Fatalf("Error loading Microsoft Graph API exceptions configuration: %s", err)
	}
}

// loadAPIExceptionsConfiguration reads and unmarshals the graph_api_exceptions_configuration JSON data from an embedded file
// into the configMap variable, which holds the exceptions configuration for endpoint-specific headers.
func loadAPIExceptionsConfiguration() error {
	return json.Unmarshal(graph_api_exceptions_configuration, &configMap)
}
