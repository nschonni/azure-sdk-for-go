package insights

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// MetricBaselineClient is the monitor Management Client
type MetricBaselineClient struct {
	BaseClient
}

// NewMetricBaselineClient creates an instance of the MetricBaselineClient client.
func NewMetricBaselineClient(subscriptionID string) MetricBaselineClient {
	return NewMetricBaselineClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMetricBaselineClientWithBaseURI creates an instance of the MetricBaselineClient client.
func NewMetricBaselineClientWithBaseURI(baseURI string, subscriptionID string) MetricBaselineClient {
	return MetricBaselineClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CalculateBaseline **Lists the baseline values for a resource**.
// Parameters:
// resourceURI - the identifier of the resource. It has the following structure:
// subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceName}.
// For example:
// subscriptions/b368ca2f-e298-46b7-b0ab-012281956afa/resourceGroups/vms/providers/Microsoft.Compute/virtualMachines/vm1
// timeSeriesInformation - information that need to be specified to calculate a baseline on a time series.
func (client MetricBaselineClient) CalculateBaseline(ctx context.Context, resourceURI string, timeSeriesInformation TimeSeriesInformation) (result CalculateBaselineResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MetricBaselineClient.CalculateBaseline")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: timeSeriesInformation,
			Constraints: []validation.Constraint{{Target: "timeSeriesInformation.Sensitivities", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "timeSeriesInformation.Values", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("insights.MetricBaselineClient", "CalculateBaseline", err.Error())
	}

	req, err := client.CalculateBaselinePreparer(ctx, resourceURI, timeSeriesInformation)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetricBaselineClient", "CalculateBaseline", nil, "Failure preparing request")
		return
	}

	resp, err := client.CalculateBaselineSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.MetricBaselineClient", "CalculateBaseline", resp, "Failure sending request")
		return
	}

	result, err = client.CalculateBaselineResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetricBaselineClient", "CalculateBaseline", resp, "Failure responding to request")
	}

	return
}

// CalculateBaselinePreparer prepares the CalculateBaseline request.
func (client MetricBaselineClient) CalculateBaselinePreparer(ctx context.Context, resourceURI string, timeSeriesInformation TimeSeriesInformation) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceUri": resourceURI,
	}

	const APIVersion = "2017-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/microsoft.insights/calculatebaseline", pathParameters),
		autorest.WithJSON(timeSeriesInformation),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CalculateBaselineSender sends the CalculateBaseline request. The method will close the
// http.Response Body if it receives an error.
func (client MetricBaselineClient) CalculateBaselineSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CalculateBaselineResponder handles the response to the CalculateBaseline request. The method always
// closes the http.Response Body.
func (client MetricBaselineClient) CalculateBaselineResponder(resp *http.Response) (result CalculateBaselineResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get **Gets the baseline values for a specific metric**.
// Parameters:
// resourceURI - the identifier of the resource. It has the following structure:
// subscriptions/{subscriptionName}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceName}.
// For example:
// subscriptions/b368ca2f-e298-46b7-b0ab-012281956afa/resourceGroups/vms/providers/Microsoft.Compute/virtualMachines/vm1
// metricName - the name of the metric to retrieve the baseline for.
// timespan - the timespan of the query. It is a string with the following format
// 'startDateTime_ISO/endDateTime_ISO'.
// interval - the interval (i.e. timegrain) of the query.
// aggregation - the aggregation type of the metric to retrieve the baseline for.
// sensitivities - the list of sensitivities (comma separated) to retrieve.
// resultType - allows retrieving only metadata of the baseline. On data request all information is retrieved.
func (client MetricBaselineClient) Get(ctx context.Context, resourceURI string, metricName string, timespan string, interval *string, aggregation string, sensitivities string, resultType ResultType) (result BaselineResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MetricBaselineClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceURI, metricName, timespan, interval, aggregation, sensitivities, resultType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetricBaselineClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.MetricBaselineClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetricBaselineClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MetricBaselineClient) GetPreparer(ctx context.Context, resourceURI string, metricName string, timespan string, interval *string, aggregation string, sensitivities string, resultType ResultType) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"metricName":  autorest.Encode("path", metricName),
		"resourceUri": resourceURI,
	}

	const APIVersion = "2017-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(timespan) > 0 {
		queryParameters["timespan"] = autorest.Encode("query", timespan)
	}
	if interval != nil {
		queryParameters["interval"] = autorest.Encode("query", *interval)
	}
	if len(aggregation) > 0 {
		queryParameters["aggregation"] = autorest.Encode("query", aggregation)
	}
	if len(sensitivities) > 0 {
		queryParameters["sensitivities"] = autorest.Encode("query", sensitivities)
	}
	if len(string(resultType)) > 0 {
		queryParameters["resultType"] = autorest.Encode("query", resultType)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/microsoft.insights/baseline/{metricName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MetricBaselineClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MetricBaselineClient) GetResponder(resp *http.Response) (result BaselineResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
