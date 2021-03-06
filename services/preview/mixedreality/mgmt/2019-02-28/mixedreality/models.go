package mixedreality

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/mixedreality/mgmt/2019-02-28/mixedreality"

// NameAvailability enumerates the values for name availability.
type NameAvailability string

const (
	// False ...
	False NameAvailability = "false"
	// True ...
	True NameAvailability = "true"
)

// PossibleNameAvailabilityValues returns an array of possible values for the NameAvailability const type.
func PossibleNameAvailabilityValues() []NameAvailability {
	return []NameAvailability{False, True}
}

// NameUnavailableReason enumerates the values for name unavailable reason.
type NameUnavailableReason string

const (
	// AlreadyExists ...
	AlreadyExists NameUnavailableReason = "AlreadyExists"
	// Invalid ...
	Invalid NameUnavailableReason = "Invalid"
)

// PossibleNameUnavailableReasonValues returns an array of possible values for the NameUnavailableReason const type.
func PossibleNameUnavailableReasonValues() []NameUnavailableReason {
	return []NameUnavailableReason{AlreadyExists, Invalid}
}

// AzureEntityResource the resource model definition for a Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityRequest check Name Availability Request
type CheckNameAvailabilityRequest struct {
	// Name - Resource Name To Verify
	Name *string `json:"name,omitempty"`
	// Type - Fully qualified resource type which includes provider namespace
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityResponse check Name Availability Response
type CheckNameAvailabilityResponse struct {
	autorest.Response `json:"-"`
	// NameAvailable - if name Available. Possible values include: 'True', 'False'
	NameAvailable NameAvailability `json:"nameAvailable,omitempty"`
	// Reason - Resource Name To Verify. Possible values include: 'Invalid', 'AlreadyExists'
	Reason NameUnavailableReason `json:"reason,omitempty"`
	// Message - detail message
	Message *string `json:"message,omitempty"`
}

// ErrorResponse response on Error
type ErrorResponse struct {
	// Message - Describes the error in detail and provides debugging information
	Message *string `json:"message,omitempty"`
	// Code - String that can be used to programmatically identify the error.
	Code *string `json:"code,omitempty"`
	// Target - The target of the particular error
	Target *string `json:"target,omitempty"`
	// Details - An array of JSON objects that MUST contain name/value pairs for code and message, and MAY contain a name/value pair for target, as described above.The contents of this section are service-defined but must adhere to the aforementioned schema.
	Details *string `json:"details,omitempty"`
}

// Operation REST API operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.ResourceProvider
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of operation
	Description *string `json:"description,omitempty"`
}

// OperationList result of the request to list Resource Provider operations. It contains a list of
// operations and a URL link to get the next set of results.
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - List of operations supported by the Resource Provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListIterator provides access to a complete listing of Operation values.
type OperationListIterator struct {
	i    int
	page OperationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListIterator) Response() OperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListIterator type.
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return OperationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ol OperationList) IsEmpty() bool {
	return ol.Value == nil || len(*ol.Value) == 0
}

// operationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ol OperationList) operationListPreparer(ctx context.Context) (*http.Request, error) {
	if ol.NextLink == nil || len(to.String(ol.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ol.NextLink)))
}

// OperationListPage contains a page of Operation values.
type OperationListPage struct {
	fn func(context.Context, OperationList) (OperationList, error)
	ol OperationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.ol)
	if err != nil {
		return err
	}
	page.ol = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListPage) NotDone() bool {
	return !page.ol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListPage) Response() OperationList {
	return page.ol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListPage) Values() []Operation {
	if page.ol.IsEmpty() {
		return nil
	}
	return *page.ol.Value
}

// Creates a new instance of the OperationListPage type.
func NewOperationListPage(getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return OperationListPage{fn: getNextPage}
}

// ProxyResource the resource model definition for a ARM proxy resource. It will have everything other than
// required location and tags
type ProxyResource struct {
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// Resource ...
type Resource struct {
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// SpatialAnchorsAccount spatialAnchorsAccount Response.
type SpatialAnchorsAccount struct {
	autorest.Response `json:"-"`
	// SpatialAnchorsAccountProperties - Property bag.
	*SpatialAnchorsAccountProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for SpatialAnchorsAccount.
func (saa SpatialAnchorsAccount) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if saa.SpatialAnchorsAccountProperties != nil {
		objectMap["properties"] = saa.SpatialAnchorsAccountProperties
	}
	if saa.Tags != nil {
		objectMap["tags"] = saa.Tags
	}
	if saa.Location != nil {
		objectMap["location"] = saa.Location
	}
	if saa.ID != nil {
		objectMap["id"] = saa.ID
	}
	if saa.Name != nil {
		objectMap["name"] = saa.Name
	}
	if saa.Type != nil {
		objectMap["type"] = saa.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SpatialAnchorsAccount struct.
func (saa *SpatialAnchorsAccount) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var spatialAnchorsAccountProperties SpatialAnchorsAccountProperties
				err = json.Unmarshal(*v, &spatialAnchorsAccountProperties)
				if err != nil {
					return err
				}
				saa.SpatialAnchorsAccountProperties = &spatialAnchorsAccountProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				saa.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				saa.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				saa.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				saa.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				saa.Type = &typeVar
			}
		}
	}

	return nil
}

// SpatialAnchorsAccountKeyRegenerateRequest spatial Anchors Account Regenerate Key
type SpatialAnchorsAccountKeyRegenerateRequest struct {
	// Serial - serial of key to be regenerated
	Serial *int32 `json:"serial,omitempty"`
}

// SpatialAnchorsAccountKeys spatial Anchors Account Keys
type SpatialAnchorsAccountKeys struct {
	autorest.Response `json:"-"`
	// PrimaryKey - value of primary key.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - value of secondary key.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// SpatialAnchorsAccountList result of the request to get resource collection. It contains a list of
// resources and a URL link to get the next set of results.
type SpatialAnchorsAccountList struct {
	autorest.Response `json:"-"`
	// Value - List of resources supported by the Resource Provider.
	Value *[]SpatialAnchorsAccount `json:"value,omitempty"`
	// NextLink - URL to get the next set of resource list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// SpatialAnchorsAccountListIterator provides access to a complete listing of SpatialAnchorsAccount values.
type SpatialAnchorsAccountListIterator struct {
	i    int
	page SpatialAnchorsAccountListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *SpatialAnchorsAccountListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SpatialAnchorsAccountListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *SpatialAnchorsAccountListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter SpatialAnchorsAccountListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter SpatialAnchorsAccountListIterator) Response() SpatialAnchorsAccountList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter SpatialAnchorsAccountListIterator) Value() SpatialAnchorsAccount {
	if !iter.page.NotDone() {
		return SpatialAnchorsAccount{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the SpatialAnchorsAccountListIterator type.
func NewSpatialAnchorsAccountListIterator(page SpatialAnchorsAccountListPage) SpatialAnchorsAccountListIterator {
	return SpatialAnchorsAccountListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (saal SpatialAnchorsAccountList) IsEmpty() bool {
	return saal.Value == nil || len(*saal.Value) == 0
}

// spatialAnchorsAccountListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (saal SpatialAnchorsAccountList) spatialAnchorsAccountListPreparer(ctx context.Context) (*http.Request, error) {
	if saal.NextLink == nil || len(to.String(saal.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(saal.NextLink)))
}

// SpatialAnchorsAccountListPage contains a page of SpatialAnchorsAccount values.
type SpatialAnchorsAccountListPage struct {
	fn   func(context.Context, SpatialAnchorsAccountList) (SpatialAnchorsAccountList, error)
	saal SpatialAnchorsAccountList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *SpatialAnchorsAccountListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SpatialAnchorsAccountListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.saal)
	if err != nil {
		return err
	}
	page.saal = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *SpatialAnchorsAccountListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page SpatialAnchorsAccountListPage) NotDone() bool {
	return !page.saal.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page SpatialAnchorsAccountListPage) Response() SpatialAnchorsAccountList {
	return page.saal
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page SpatialAnchorsAccountListPage) Values() []SpatialAnchorsAccount {
	if page.saal.IsEmpty() {
		return nil
	}
	return *page.saal.Value
}

// Creates a new instance of the SpatialAnchorsAccountListPage type.
func NewSpatialAnchorsAccountListPage(getNextPage func(context.Context, SpatialAnchorsAccountList) (SpatialAnchorsAccountList, error)) SpatialAnchorsAccountListPage {
	return SpatialAnchorsAccountListPage{fn: getNextPage}
}

// SpatialAnchorsAccountProperties spatial Anchors Account Customize Properties
type SpatialAnchorsAccountProperties struct {
	// AccountID - unique id of certain Spatial Anchors Account data contract.
	AccountID *string `json:"accountId,omitempty"`
	// AccountDomain - Correspond domain name of certain Spatial Anchors Account
	AccountDomain *string `json:"accountDomain,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	if tr.ID != nil {
		objectMap["id"] = tr.ID
	}
	if tr.Name != nil {
		objectMap["name"] = tr.Name
	}
	if tr.Type != nil {
		objectMap["type"] = tr.Type
	}
	return json.Marshal(objectMap)
}
