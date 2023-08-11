// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"TMfRuM4AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
