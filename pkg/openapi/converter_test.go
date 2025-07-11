package openapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverter_Convert(t *testing.T) {
	converter := NewConverter()

	// Test with a simple OpenAPI specification
	spec := `{
		"openapi": "3.0.0",
		"info": {
			"title": "Test API",
			"description": "Test API description",
			"version": "1.0.0"
		},
		"servers": [
			{
				"url": "https://api.example.com/v1"
			}
		],
		"paths": {
			"/test": {
				"get": {
					"summary": "Test endpoint",
					"responses": {
						"200": {
							"description": "Successful response"
						}
					}
				},
				"options": {
					"summary": "CORS options",
					"responses": {
						"200": {
							"description": "CORS headers"
						}
					}
				}
			}
		}
	}`

	config, err := converter.ConvertFromJSON([]byte(spec))
	assert.NoError(t, err)
	assert.NotNil(t, config)

	// Verify the converted configuration
	assert.Equal(t, "Test API", config.Name)
	assert.Equal(t, 1, len(config.Routers))
	assert.Equal(t, "/test", config.Routers[0].Prefix)
	assert.Equal(t, "Test API", config.Routers[0].Server)
	assert.NotNil(t, config.Routers[0].CORS)

	assert.Equal(t, 1, len(config.Servers))
	assert.Equal(t, "Test API", config.Servers[0].Name)
	assert.Equal(t, "Test API description", config.Servers[0].Description)
	assert.Equal(t, "https://api.example.com/v1", config.Servers[0].Config["url"])
}

func TestConverter_ConvertFromYAML(t *testing.T) {
	converter := NewConverter()

	// Test with a simple OpenAPI specification in YAML
	spec := `openapi: 3.0.0
info:
  title: Test API
  description: Test API description
  version: 1.0.0
servers:
  - url: https://api.example.com/v1
paths:
  /test:
    get:
      summary: Test endpoint
      responses:
        200:
          description: Successful response
    options:
      summary: CORS options
      responses:
        200:
          description: CORS headers`

	config, err := converter.ConvertFromYAML([]byte(spec))
	assert.NoError(t, err)
	assert.NotNil(t, config)

	// Verify the converted configuration
	assert.Equal(t, "Test API", config.Name)
	assert.Equal(t, 1, len(config.Routers))
	assert.Equal(t, "/test", config.Routers[0].Prefix)
	assert.Equal(t, "Test API", config.Routers[0].Server)
	assert.NotNil(t, config.Routers[0].CORS)

	assert.Equal(t, 1, len(config.Servers))
	assert.Equal(t, "Test API", config.Servers[0].Name)
	assert.Equal(t, "Test API description", config.Servers[0].Description)
	assert.Equal(t, "https://api.example.com/v1", config.Servers[0].Config["url"])
}
