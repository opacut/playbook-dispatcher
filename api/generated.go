// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// Account defines model for Account.
type Account string

// Labels defines model for Labels.
type Labels struct {
	AdditionalProperties map[string]string `json:"-"`
}

// Run defines model for Run.
type Run struct {
	Account   Account      `json:"account"`
	Id        RunId        `json:"id"`
	Labels    Labels       `json:"labels"`
	Recipient RunRecipient `json:"recipient"`
	Status    RunStatus    `json:"status"`

	// Amount of seconds after which the run is considered failed due to timeout
	Timeout RunTimeout `json:"timeout"`
	Url     Url        `json:"url"`
}

// RunCreated defines model for RunCreated.
type RunCreated struct {

	// status code of the request
	Code int    `json:"code"`
	Id   *RunId `json:"id,omitempty"`
}

// RunId defines model for RunId.
type RunId string

// RunInput defines model for RunInput.
type RunInput struct {
	Account   Account      `json:"account"`
	Labels    *Labels      `json:"labels,omitempty"`
	Recipient RunRecipient `json:"recipient"`

	// Amount of seconds after which the run is considered failed due to timeout
	Timeout *RunTimeout `json:"timeout,omitempty"`
	Url     Url         `json:"url"`
}

// RunRecipient defines model for RunRecipient.
type RunRecipient string

// RunStatus defines model for RunStatus.
type RunStatus string

// List of RunStatus
const (
	RunStatus_failure RunStatus = "failure"
	RunStatus_running RunStatus = "running"
	RunStatus_success RunStatus = "success"
	RunStatus_timeout RunStatus = "timeout"
)

// RunTimeout defines model for RunTimeout.
type RunTimeout int

// Runs defines model for Runs.
type Runs struct {
	Data []Run `json:"data"`
}

// RunsCreated defines model for RunsCreated.
type RunsCreated []RunCreated

// Url defines model for Url.
type Url string

// RunsFilter defines model for RunsFilter.
type RunsFilter struct {
	Labels *Labels `json:"labels,omitempty"`
	Status *string `json:"status,omitempty"`
}

// ApiRunsGetParams defines parameters for ApiRunsGet.
type ApiRunsGetParams struct {
	Filter *RunsFilter `json:"filter,omitempty"`
}

// ApiRunsGetParams_Filter_Labels defines parameters for ApiRunsGet.
type ApiRunsGetParams_Filter_Labels struct {
	AdditionalProperties map[string]string `json:"-"`
}

// ApiInternalRunsCreateJSONBody defines parameters for ApiInternalRunsCreate.
type ApiInternalRunsCreateJSONBody []RunInput

// ApiInternalRunsCreateRequestBody defines body for ApiInternalRunsCreate for application/json ContentType.
type ApiInternalRunsCreateJSONRequestBody ApiInternalRunsCreateJSONBody

// Getter for additional properties for ApiRunsGetParams_Filter_Labels. Returns the specified
// element and whether it was found
func (a ApiRunsGetParams_Filter_Labels) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ApiRunsGetParams_Filter_Labels
func (a *ApiRunsGetParams_Filter_Labels) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ApiRunsGetParams_Filter_Labels to handle AdditionalProperties
func (a *ApiRunsGetParams_Filter_Labels) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ApiRunsGetParams_Filter_Labels to handle AdditionalProperties
func (a ApiRunsGetParams_Filter_Labels) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for Labels. Returns the specified
// element and whether it was found
func (a Labels) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Labels
func (a *Labels) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Labels to handle AdditionalProperties
func (a *Labels) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Labels to handle AdditionalProperties
func (a Labels) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /api/playbook-dispatcher/v1/runs)
	ApiRunsGet(ctx echo.Context, params ApiRunsGetParams) error

	// (POST /internal/dispatch)
	ApiInternalRunsCreate(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ApiRunsGet converts echo context to params.
func (w *ServerInterfaceWrapper) ApiRunsGet(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ApiRunsGetParams
	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("deepObject", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiRunsGet(ctx, params)
	return err
}

// ApiInternalRunsCreate converts echo context to params.
func (w *ServerInterfaceWrapper) ApiInternalRunsCreate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiInternalRunsCreate(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/playbook-dispatcher/v1/runs", wrapper.ApiRunsGet)
	router.POST(baseURL+"/internal/dispatch", wrapper.ApiInternalRunsCreate)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7yWzW7bOBDHX0WY3SPXkhNgF9At26KF0QIJnOQU5MCIo5ipRDL8SGMEevdiKFmyIzl2",
	"0CI3kx7Ox39+5OgFCl0brVB5B/kLGG55jR5tXC2Dcl9k5dHSCp9NpQVC7m1ABlJBDo8B7RoYKF4j5FC2",
	"xgxcscKaR49WG7ReYvRY8Tus4i8uhPRSK15d7Fj4tSFPzlup7qFhmw1994CFpw3nuQ9TttPG64p2BKI5",
	"73abTX7RyVlR6KD8ZOzvB9MVWPJQ0ekh+hu5L4Mai8KHDP62WEIOf6VDW9Iu13STaMNAikPGy6AWgkwH",
	"xd8y7wptGFgspJF4OJ1lUMvedqctB05dtoYkj6xRh2MiXXWWDYNgq0MHrm0Vu2zxMUiLAvIbkoz1Sm+X",
	"2XrshRqy6ku6nW7kJ4vcoxj3s4i3hNhwhZWGsIlUkLOE/kx0mfgVJpQfOoqEz7w2BOpJNu+jSeXxHu27",
	"Gv6q6pjKnvwX0Wepbc095BBClGjEMFkq03bpd7n9MBg/kqxtlAbAyNke4ZfbVR2j/2V/tVCFOsYMStH/",
	"DFwoCnTEbcllFSxuEXw77e5qEKd/vk7/zTL2itizmmohWB0WWgmX8NKjTX6uZLFqAQ4qkcS0clKgRZFQ",
	"EigSETDxOhmu0hhpGi5jpgT3cWxIj/Uxj8nWK8ut5etRd6LDPY1wW1f42ICbI6O4DK5bfoaOxnfl9ZCi",
	"26xKHeeN9HE4XVR8faf1j+SzdIb7YhVn6BNa1zYim2WzOUXQBhU3EnI4nWWzU2BguF/FtFNuZGo6R/+I",
	"3lH6NE9tJ/U9xqaT3px6TC8AnBlJUnxFH90N8/9mWovBJN36PmhuSXZntHJtI0+yrH0Lle9I58ZUsoiB",
	"0wen4yAcvhMOyO5a5XYBPf9Guw2DlLiyilfppvBIlnbT9S466wEBaKFB5//XYv2uxI8Fp31EGwY1f160",
	"Z+YZg1qqzWpEcsvyjqj//VFRe5j3ats0vwIAAP//rWbEHiMKAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
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

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}