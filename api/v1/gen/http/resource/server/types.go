// Code generated by goa v3.3.1, DO NOT EDIT.
//
// resource HTTP server types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package server

import (
	resource "github.com/tektoncd/hub/api/v1/gen/resource"
	resourceviews "github.com/tektoncd/hub/api/v1/gen/resource/views"
	goa "goa.design/goa/v3/pkg"
)

// QueryResponseBody is the type of the "resource" service "Query" endpoint
// HTTP response body.
type QueryResponseBody struct {
	Data ResourceDataResponseBodyWithoutVersionCollection `form:"data" json:"data" xml:"data"`
}

// ListResponseBody is the type of the "resource" service "List" endpoint HTTP
// response body.
type ListResponseBody struct {
	Data ResourceDataResponseBodyWithoutVersionCollection `form:"data" json:"data" xml:"data"`
}

// VersionsByIDResponseBody is the type of the "resource" service
// "VersionsByID" endpoint HTTP response body.
type VersionsByIDResponseBody struct {
	Data *VersionsResponseBody `form:"data" json:"data" xml:"data"`
}

// ByCatalogKindNameVersionResponseBody is the type of the "resource" service
// "ByCatalogKindNameVersion" endpoint HTTP response body.
type ByCatalogKindNameVersionResponseBody struct {
	Data *ResourceVersionDataResponseBody `form:"data" json:"data" xml:"data"`
}

// ByVersionIDResponseBody is the type of the "resource" service "ByVersionId"
// endpoint HTTP response body.
type ByVersionIDResponseBody struct {
	Data *ResourceVersionDataResponseBody `form:"data" json:"data" xml:"data"`
}

// ByCatalogKindNameResponseBody is the type of the "resource" service
// "ByCatalogKindName" endpoint HTTP response body.
type ByCatalogKindNameResponseBody struct {
	Data *ResourceDataResponseBody `form:"data" json:"data" xml:"data"`
}

// ByIDResponseBody is the type of the "resource" service "ById" endpoint HTTP
// response body.
type ByIDResponseBody struct {
	Data *ResourceDataResponseBody `form:"data" json:"data" xml:"data"`
}

// QueryInternalErrorResponseBody is the type of the "resource" service "Query"
// endpoint HTTP response body for the "internal-error" error.
type QueryInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// QueryInvalidKindResponseBody is the type of the "resource" service "Query"
// endpoint HTTP response body for the "invalid-kind" error.
type QueryInvalidKindResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// QueryNotFoundResponseBody is the type of the "resource" service "Query"
// endpoint HTTP response body for the "not-found" error.
type QueryNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListInternalErrorResponseBody is the type of the "resource" service "List"
// endpoint HTTP response body for the "internal-error" error.
type ListInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// VersionsByIDInternalErrorResponseBody is the type of the "resource" service
// "VersionsByID" endpoint HTTP response body for the "internal-error" error.
type VersionsByIDInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// VersionsByIDNotFoundResponseBody is the type of the "resource" service
// "VersionsByID" endpoint HTTP response body for the "not-found" error.
type VersionsByIDNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ByCatalogKindNameVersionInternalErrorResponseBody is the type of the
// "resource" service "ByCatalogKindNameVersion" endpoint HTTP response body
// for the "internal-error" error.
type ByCatalogKindNameVersionInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ByCatalogKindNameVersionNotFoundResponseBody is the type of the "resource"
// service "ByCatalogKindNameVersion" endpoint HTTP response body for the
// "not-found" error.
type ByCatalogKindNameVersionNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ByVersionIDInternalErrorResponseBody is the type of the "resource" service
// "ByVersionId" endpoint HTTP response body for the "internal-error" error.
type ByVersionIDInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ByVersionIDNotFoundResponseBody is the type of the "resource" service
// "ByVersionId" endpoint HTTP response body for the "not-found" error.
type ByVersionIDNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ByCatalogKindNameInternalErrorResponseBody is the type of the "resource"
// service "ByCatalogKindName" endpoint HTTP response body for the
// "internal-error" error.
type ByCatalogKindNameInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ByCatalogKindNameNotFoundResponseBody is the type of the "resource" service
// "ByCatalogKindName" endpoint HTTP response body for the "not-found" error.
type ByCatalogKindNameNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ByIDInternalErrorResponseBody is the type of the "resource" service "ById"
// endpoint HTTP response body for the "internal-error" error.
type ByIDInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ByIDNotFoundResponseBody is the type of the "resource" service "ById"
// endpoint HTTP response body for the "not-found" error.
type ByIDNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ResourceDataResponseBodyWithoutVersionCollection is used to define fields on
// response body types.
type ResourceDataResponseBodyWithoutVersionCollection []*ResourceDataResponseBodyWithoutVersion

// ResourceDataResponseBodyWithoutVersion is used to define fields on response
// body types.
type ResourceDataResponseBodyWithoutVersion struct {
	// ID is the unique id of the resource
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of resource
	Name string `form:"name" json:"name" xml:"name"`
	// Type of catalog to which resource belongs
	Catalog *CatalogResponseBodyMin `form:"catalog" json:"catalog" xml:"catalog"`
	// Categories related to resource
	Categories []*CategoryResponseBody `form:"categories" json:"categories" xml:"categories"`
	// Kind of resource
	Kind string `form:"kind" json:"kind" xml:"kind"`
	// Latest version of resource
	LatestVersion *ResourceVersionDataResponseBodyWithoutResource `form:"latestVersion" json:"latestVersion" xml:"latestVersion"`
	// Tags related to resource
	Tags []*TagResponseBody `form:"tags" json:"tags" xml:"tags"`
	// Platforms related to resource
	Platforms []*PlatformResponseBody `form:"platforms" json:"platforms" xml:"platforms"`
	// Rating of resource
	Rating float64 `form:"rating" json:"rating" xml:"rating"`
}

// CatalogResponseBodyMin is used to define fields on response body types.
type CatalogResponseBodyMin struct {
	// ID is the unique id of the catalog
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of catalog
	Name string `form:"name" json:"name" xml:"name"`
	// Type of catalog
	Type string `form:"type" json:"type" xml:"type"`
}

// CategoryResponseBody is used to define fields on response body types.
type CategoryResponseBody struct {
	// ID is the unique id of the category
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of category
	Name string `form:"name" json:"name" xml:"name"`
}

// ResourceVersionDataResponseBodyWithoutResource is used to define fields on
// response body types.
type ResourceVersionDataResponseBodyWithoutResource struct {
	// ID is the unique id of resource's version
	ID uint `form:"id" json:"id" xml:"id"`
	// Version of resource
	Version string `form:"version" json:"version" xml:"version"`
	// Display name of version
	DisplayName string `form:"displayName" json:"displayName" xml:"displayName"`
	// Description of version
	Description string `form:"description" json:"description" xml:"description"`
	// Minimum pipelines version the resource's version is compatible with
	MinPipelinesVersion string `form:"minPipelinesVersion" json:"minPipelinesVersion" xml:"minPipelinesVersion"`
	// Raw URL of resource's yaml file of the version
	RawURL string `form:"rawURL" json:"rawURL" xml:"rawURL"`
	// Web URL of resource's yaml file of the version
	WebURL string `form:"webURL" json:"webURL" xml:"webURL"`
	// Timestamp when version was last updated
	UpdatedAt string `form:"updatedAt" json:"updatedAt" xml:"updatedAt"`
	// Platforms related to resource version
	Platforms []*PlatformResponseBody `form:"platforms" json:"platforms" xml:"platforms"`
}

// PlatformResponseBody is used to define fields on response body types.
type PlatformResponseBody struct {
	// ID is the unique id of platform
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of platform
	Name string `form:"name" json:"name" xml:"name"`
}

// TagResponseBody is used to define fields on response body types.
type TagResponseBody struct {
	// ID is the unique id of tag
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of tag
	Name string `form:"name" json:"name" xml:"name"`
}

// VersionsResponseBody is used to define fields on response body types.
type VersionsResponseBody struct {
	// Latest Version of resource
	Latest *ResourceVersionDataResponseBodyMin `form:"latest" json:"latest" xml:"latest"`
	// List of all versions of resource
	Versions []*ResourceVersionDataResponseBodyMin `form:"versions" json:"versions" xml:"versions"`
}

// ResourceVersionDataResponseBodyMin is used to define fields on response body
// types.
type ResourceVersionDataResponseBodyMin struct {
	// ID is the unique id of resource's version
	ID uint `form:"id" json:"id" xml:"id"`
	// Version of resource
	Version string `form:"version" json:"version" xml:"version"`
	// Raw URL of resource's yaml file of the version
	RawURL string `form:"rawURL" json:"rawURL" xml:"rawURL"`
	// Web URL of resource's yaml file of the version
	WebURL string `form:"webURL" json:"webURL" xml:"webURL"`
	// Platforms related to resource version
	Platforms []*PlatformResponseBody `form:"platforms" json:"platforms" xml:"platforms"`
}

// ResourceVersionDataResponseBody is used to define fields on response body
// types.
type ResourceVersionDataResponseBody struct {
	// ID is the unique id of resource's version
	ID uint `form:"id" json:"id" xml:"id"`
	// Version of resource
	Version string `form:"version" json:"version" xml:"version"`
	// Display name of version
	DisplayName string `form:"displayName" json:"displayName" xml:"displayName"`
	// Description of version
	Description string `form:"description" json:"description" xml:"description"`
	// Minimum pipelines version the resource's version is compatible with
	MinPipelinesVersion string `form:"minPipelinesVersion" json:"minPipelinesVersion" xml:"minPipelinesVersion"`
	// Raw URL of resource's yaml file of the version
	RawURL string `form:"rawURL" json:"rawURL" xml:"rawURL"`
	// Web URL of resource's yaml file of the version
	WebURL string `form:"webURL" json:"webURL" xml:"webURL"`
	// Timestamp when version was last updated
	UpdatedAt string `form:"updatedAt" json:"updatedAt" xml:"updatedAt"`
	// Resource to which the version belongs
	Resource *ResourceDataResponseBodyInfo `form:"resource" json:"resource" xml:"resource"`
	// Platforms related to resource version
	Platforms []*PlatformResponseBody `form:"platforms" json:"platforms" xml:"platforms"`
}

// ResourceDataResponseBodyInfo is used to define fields on response body types.
type ResourceDataResponseBodyInfo struct {
	// ID is the unique id of the resource
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of resource
	Name string `form:"name" json:"name" xml:"name"`
	// Type of catalog to which resource belongs
	Catalog *CatalogResponseBodyMin `form:"catalog" json:"catalog" xml:"catalog"`
	// Categories related to resource
	Categories []*CategoryResponseBody `form:"categories" json:"categories" xml:"categories"`
	// Kind of resource
	Kind string `form:"kind" json:"kind" xml:"kind"`
	// Tags related to resource
	Tags []*TagResponseBody `form:"tags" json:"tags" xml:"tags"`
	// Platforms related to resource
	Platforms []*PlatformResponseBody `form:"platforms" json:"platforms" xml:"platforms"`
	// Rating of resource
	Rating float64 `form:"rating" json:"rating" xml:"rating"`
}

// ResourceDataResponseBody is used to define fields on response body types.
type ResourceDataResponseBody struct {
	// ID is the unique id of the resource
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of resource
	Name string `form:"name" json:"name" xml:"name"`
	// Type of catalog to which resource belongs
	Catalog *CatalogResponseBodyMin `form:"catalog" json:"catalog" xml:"catalog"`
	// Categories related to resource
	Categories []*CategoryResponseBody `form:"categories" json:"categories" xml:"categories"`
	// Kind of resource
	Kind string `form:"kind" json:"kind" xml:"kind"`
	// Latest version of resource
	LatestVersion *ResourceVersionDataResponseBodyWithoutResource `form:"latestVersion" json:"latestVersion" xml:"latestVersion"`
	// Tags related to resource
	Tags []*TagResponseBody `form:"tags" json:"tags" xml:"tags"`
	// Platforms related to resource
	Platforms []*PlatformResponseBody `form:"platforms" json:"platforms" xml:"platforms"`
	// Rating of resource
	Rating float64 `form:"rating" json:"rating" xml:"rating"`
	// List of all versions of a resource
	Versions []*ResourceVersionDataResponseBodyTiny `form:"versions" json:"versions" xml:"versions"`
}

// ResourceVersionDataResponseBodyTiny is used to define fields on response
// body types.
type ResourceVersionDataResponseBodyTiny struct {
	// ID is the unique id of resource's version
	ID uint `form:"id" json:"id" xml:"id"`
	// Version of resource
	Version string `form:"version" json:"version" xml:"version"`
}

// NewQueryResponseBody builds the HTTP response body from the result of the
// "Query" endpoint of the "resource" service.
func NewQueryResponseBody(res *resourceviews.ResourcesView) *QueryResponseBody {
	body := &QueryResponseBody{}
	if res.Data != nil {
		body.Data = make([]*ResourceDataResponseBodyWithoutVersion, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalResourceviewsResourceDataViewToResourceDataResponseBodyWithoutVersion(val)
		}
	}
	return body
}

// NewListResponseBody builds the HTTP response body from the result of the
// "List" endpoint of the "resource" service.
func NewListResponseBody(res *resourceviews.ResourcesView) *ListResponseBody {
	body := &ListResponseBody{}
	if res.Data != nil {
		body.Data = make([]*ResourceDataResponseBodyWithoutVersion, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalResourceviewsResourceDataViewToResourceDataResponseBodyWithoutVersion(val)
		}
	}
	return body
}

// NewVersionsByIDResponseBody builds the HTTP response body from the result of
// the "VersionsByID" endpoint of the "resource" service.
func NewVersionsByIDResponseBody(res *resourceviews.ResourceVersionsView) *VersionsByIDResponseBody {
	body := &VersionsByIDResponseBody{}
	if res.Data != nil {
		body.Data = marshalResourceviewsVersionsViewToVersionsResponseBody(res.Data)
	}
	return body
}

// NewByCatalogKindNameVersionResponseBody builds the HTTP response body from
// the result of the "ByCatalogKindNameVersion" endpoint of the "resource"
// service.
func NewByCatalogKindNameVersionResponseBody(res *resourceviews.ResourceVersionView) *ByCatalogKindNameVersionResponseBody {
	body := &ByCatalogKindNameVersionResponseBody{}
	if res.Data != nil {
		body.Data = marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBody(res.Data)
	}
	return body
}

// NewByVersionIDResponseBody builds the HTTP response body from the result of
// the "ByVersionId" endpoint of the "resource" service.
func NewByVersionIDResponseBody(res *resourceviews.ResourceVersionView) *ByVersionIDResponseBody {
	body := &ByVersionIDResponseBody{}
	if res.Data != nil {
		body.Data = marshalResourceviewsResourceVersionDataViewToResourceVersionDataResponseBody(res.Data)
	}
	return body
}

// NewByCatalogKindNameResponseBody builds the HTTP response body from the
// result of the "ByCatalogKindName" endpoint of the "resource" service.
func NewByCatalogKindNameResponseBody(res *resourceviews.ResourceView) *ByCatalogKindNameResponseBody {
	body := &ByCatalogKindNameResponseBody{}
	if res.Data != nil {
		body.Data = marshalResourceviewsResourceDataViewToResourceDataResponseBody(res.Data)
	}
	return body
}

// NewByIDResponseBody builds the HTTP response body from the result of the
// "ById" endpoint of the "resource" service.
func NewByIDResponseBody(res *resourceviews.ResourceView) *ByIDResponseBody {
	body := &ByIDResponseBody{}
	if res.Data != nil {
		body.Data = marshalResourceviewsResourceDataViewToResourceDataResponseBody(res.Data)
	}
	return body
}

// NewQueryInternalErrorResponseBody builds the HTTP response body from the
// result of the "Query" endpoint of the "resource" service.
func NewQueryInternalErrorResponseBody(res *goa.ServiceError) *QueryInternalErrorResponseBody {
	body := &QueryInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewQueryInvalidKindResponseBody builds the HTTP response body from the
// result of the "Query" endpoint of the "resource" service.
func NewQueryInvalidKindResponseBody(res *goa.ServiceError) *QueryInvalidKindResponseBody {
	body := &QueryInvalidKindResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewQueryNotFoundResponseBody builds the HTTP response body from the result
// of the "Query" endpoint of the "resource" service.
func NewQueryNotFoundResponseBody(res *goa.ServiceError) *QueryNotFoundResponseBody {
	body := &QueryNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListInternalErrorResponseBody builds the HTTP response body from the
// result of the "List" endpoint of the "resource" service.
func NewListInternalErrorResponseBody(res *goa.ServiceError) *ListInternalErrorResponseBody {
	body := &ListInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewVersionsByIDInternalErrorResponseBody builds the HTTP response body from
// the result of the "VersionsByID" endpoint of the "resource" service.
func NewVersionsByIDInternalErrorResponseBody(res *goa.ServiceError) *VersionsByIDInternalErrorResponseBody {
	body := &VersionsByIDInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewVersionsByIDNotFoundResponseBody builds the HTTP response body from the
// result of the "VersionsByID" endpoint of the "resource" service.
func NewVersionsByIDNotFoundResponseBody(res *goa.ServiceError) *VersionsByIDNotFoundResponseBody {
	body := &VersionsByIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewByCatalogKindNameVersionInternalErrorResponseBody builds the HTTP
// response body from the result of the "ByCatalogKindNameVersion" endpoint of
// the "resource" service.
func NewByCatalogKindNameVersionInternalErrorResponseBody(res *goa.ServiceError) *ByCatalogKindNameVersionInternalErrorResponseBody {
	body := &ByCatalogKindNameVersionInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewByCatalogKindNameVersionNotFoundResponseBody builds the HTTP response
// body from the result of the "ByCatalogKindNameVersion" endpoint of the
// "resource" service.
func NewByCatalogKindNameVersionNotFoundResponseBody(res *goa.ServiceError) *ByCatalogKindNameVersionNotFoundResponseBody {
	body := &ByCatalogKindNameVersionNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewByVersionIDInternalErrorResponseBody builds the HTTP response body from
// the result of the "ByVersionId" endpoint of the "resource" service.
func NewByVersionIDInternalErrorResponseBody(res *goa.ServiceError) *ByVersionIDInternalErrorResponseBody {
	body := &ByVersionIDInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewByVersionIDNotFoundResponseBody builds the HTTP response body from the
// result of the "ByVersionId" endpoint of the "resource" service.
func NewByVersionIDNotFoundResponseBody(res *goa.ServiceError) *ByVersionIDNotFoundResponseBody {
	body := &ByVersionIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewByCatalogKindNameInternalErrorResponseBody builds the HTTP response body
// from the result of the "ByCatalogKindName" endpoint of the "resource"
// service.
func NewByCatalogKindNameInternalErrorResponseBody(res *goa.ServiceError) *ByCatalogKindNameInternalErrorResponseBody {
	body := &ByCatalogKindNameInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewByCatalogKindNameNotFoundResponseBody builds the HTTP response body from
// the result of the "ByCatalogKindName" endpoint of the "resource" service.
func NewByCatalogKindNameNotFoundResponseBody(res *goa.ServiceError) *ByCatalogKindNameNotFoundResponseBody {
	body := &ByCatalogKindNameNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewByIDInternalErrorResponseBody builds the HTTP response body from the
// result of the "ById" endpoint of the "resource" service.
func NewByIDInternalErrorResponseBody(res *goa.ServiceError) *ByIDInternalErrorResponseBody {
	body := &ByIDInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewByIDNotFoundResponseBody builds the HTTP response body from the result of
// the "ById" endpoint of the "resource" service.
func NewByIDNotFoundResponseBody(res *goa.ServiceError) *ByIDNotFoundResponseBody {
	body := &ByIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewQueryPayload builds a resource service Query endpoint payload.
func NewQueryPayload(name string, catalogs []string, categories []string, kinds []string, tags []string, limit uint, match string) *resource.QueryPayload {
	v := &resource.QueryPayload{}
	v.Name = name
	v.Catalogs = catalogs
	v.Categories = categories
	v.Kinds = kinds
	v.Tags = tags
	v.Limit = limit
	v.Match = match

	return v
}

// NewListPayload builds a resource service List endpoint payload.
func NewListPayload(limit uint) *resource.ListPayload {
	v := &resource.ListPayload{}
	v.Limit = limit

	return v
}

// NewVersionsByIDPayload builds a resource service VersionsByID endpoint
// payload.
func NewVersionsByIDPayload(id uint) *resource.VersionsByIDPayload {
	v := &resource.VersionsByIDPayload{}
	v.ID = id

	return v
}

// NewByCatalogKindNameVersionPayload builds a resource service
// ByCatalogKindNameVersion endpoint payload.
func NewByCatalogKindNameVersionPayload(catalog string, kind string, name string, version string) *resource.ByCatalogKindNameVersionPayload {
	v := &resource.ByCatalogKindNameVersionPayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name
	v.Version = version

	return v
}

// NewByVersionIDPayload builds a resource service ByVersionId endpoint payload.
func NewByVersionIDPayload(versionID uint) *resource.ByVersionIDPayload {
	v := &resource.ByVersionIDPayload{}
	v.VersionID = versionID

	return v
}

// NewByCatalogKindNamePayload builds a resource service ByCatalogKindName
// endpoint payload.
func NewByCatalogKindNamePayload(catalog string, kind string, name string, pipelinesversion *string) *resource.ByCatalogKindNamePayload {
	v := &resource.ByCatalogKindNamePayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name
	v.Pipelinesversion = pipelinesversion

	return v
}

// NewByIDPayload builds a resource service ById endpoint payload.
func NewByIDPayload(id uint) *resource.ByIDPayload {
	v := &resource.ByIDPayload{}
	v.ID = id

	return v
}
