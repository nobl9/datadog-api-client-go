// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"bytes"
	_context "context"
	_io "io"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorsApi service type
type MonitorsApi datadog.Service

type apiCheckCanDeleteMonitorRequest struct {
	ctx        _context.Context
	monitorIds *[]int64
}

func (a *MonitorsApi) buildCheckCanDeleteMonitorRequest(ctx _context.Context, monitorIds []int64) (apiCheckCanDeleteMonitorRequest, error) {
	req := apiCheckCanDeleteMonitorRequest{
		ctx:        ctx,
		monitorIds: &monitorIds,
	}
	return req, nil
}

// CheckCanDeleteMonitor Check if a monitor can be deleted.
// Check if the given monitors can be deleted.
func (a *MonitorsApi) CheckCanDeleteMonitor(ctx _context.Context, monitorIds []int64) (CheckCanDeleteMonitorResponse, *_nethttp.Response, error) {
	req, err := a.buildCheckCanDeleteMonitorRequest(ctx, monitorIds)
	if err != nil {
		var localVarReturnValue CheckCanDeleteMonitorResponse
		return localVarReturnValue, nil, err
	}

	return a.checkCanDeleteMonitorExecute(req)
}

// checkCanDeleteMonitorExecute executes the request.
func (a *MonitorsApi) checkCanDeleteMonitorExecute(r apiCheckCanDeleteMonitorRequest) (CheckCanDeleteMonitorResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue CheckCanDeleteMonitorResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.MonitorsApi.CheckCanDeleteMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/can_delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.monitorIds == nil {
		return localVarReturnValue, nil, datadog.ReportError("monitorIds is required and must be specified")
	}
	localVarQueryParams.Add("monitor_ids", datadog.ParameterToString(*r.monitorIds, "csv"))
	localVarHeaderParams["Accept"] = "application/json"

	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v CheckCanDeleteMonitorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiCreateMonitorRequest struct {
	ctx  _context.Context
	body *Monitor
}

func (a *MonitorsApi) buildCreateMonitorRequest(ctx _context.Context, body Monitor) (apiCreateMonitorRequest, error) {
	req := apiCreateMonitorRequest{
		ctx:  ctx,
		body: &body,
	}
	return req, nil
}

// CreateMonitor Create a monitor.
// Create a monitor using the specified options.
//
// #### Monitor Types
//
// The type of monitor chosen from:
//
// - anomaly: `query alert`
// - APM: `query alert` or `trace-analytics alert`
// - composite: `composite`
// - custom: `service check`
// - event: `event alert`
// - forecast: `query alert`
// - host: `service check`
// - integration: `query alert` or `service check`
// - live process: `process alert`
// - logs: `log alert`
// - metric: `query alert`
// - network: `service check`
// - outlier: `query alert`
// - process: `service check`
// - rum: `rum alert`
// - SLO: `slo alert`
// - watchdog: `event alert`
// - event-v2: `event-v2 alert`
// - audit: `audit alert`
// - error-tracking: `error-tracking alert`
//
// #### Query Types
//
// **Metric Alert Query**
//
// Example: `time_aggr(time_window):space_aggr:metric{tags} [by {key}] operator #`
//
// - `time_aggr`: avg, sum, max, min, change, or pct_change
// - `time_window`: `last_#m` (with `#` between 1 and 10080 depending on the monitor type) or `last_#h`(with `#` between 1 and 168 depending on the monitor type) or `last_1d`, or `last_1w`
// - `space_aggr`: avg, sum, min, or max
// - `tags`: one or more tags (comma-separated), or *
// - `key`: a 'key' in key:value tag syntax; defines a separate alert for each tag in the group (multi-alert)
// - `operator`: <, <=, >, >=, ==, or !=
// - `#`: an integer or decimal number used to set the threshold
//
// If you are using the `_change_` or `_pct_change_` time aggregator, instead use `change_aggr(time_aggr(time_window),
// timeshift):space_aggr:metric{tags} [by {key}] operator #` with:
//
// - `change_aggr` change, pct_change
// - `time_aggr` avg, sum, max, min [Learn more](https://docs.datadoghq.com/monitors/create/types/#define-the-conditions)
// - `time_window` last\_#m (between 1 and 2880 depending on the monitor type), last\_#h (between 1 and 48 depending on the monitor type), or last_#d (1 or 2)
// - `timeshift` #m_ago (5, 10, 15, or 30), #h_ago (1, 2, or 4), or 1d_ago
//
// Use this to create an outlier monitor using the following query:
// `avg(last_30m):outliers(avg:system.cpu.user{role:es-events-data} by {host}, 'dbscan', 7) > 0`
//
// **Service Check Query**
//
// Example: `"check".over(tags).last(count).by(group).count_by_status()`
//
// - `check` name of the check, for example `datadog.agent.up`
// - `tags` one or more quoted tags (comma-separated), or "*". for example: `.over("env:prod", "role:db")`; `over` cannot be blank.
// - `count` must be at greater than or equal to your max threshold (defined in the `options`). It is limited to 100.
// For example, if you've specified to notify on 1 critical, 3 ok, and 2 warn statuses, `count` should be at least 3.
// - `group` must be specified for check monitors. Per-check grouping is already explicitly known for some service checks.
// For example, Postgres integration monitors are tagged by `db`, `host`, and `port`, and Network monitors by `host`, `instance`, and `url`. See [Service Checks](https://docs.datadoghq.com/api/latest/service-checks/) documentation for more information.
//
// **Event Alert Query**
//
// Example: `events('sources:nagios status:error,warning priority:normal tags: "string query"').rollup("count").last("1h")"`
//
// - `event`, the event query string:
// - `string_query` free text query to match against event title and text.
// - `sources` event sources (comma-separated).
// - `status` event statuses (comma-separated). Valid options: error, warn, and info.
// - `priority` event priorities (comma-separated). Valid options: low, normal, all.
// - `host` event reporting host (comma-separated).
// - `tags` event tags (comma-separated).
// - `excluded_tags` excluded event tags (comma-separated).
// - `rollup` the stats roll-up method. `count` is the only supported method now.
// - `last` the timeframe to roll up the counts. Examples: 45m, 4h. Supported timeframes: m, h and d. This value should not exceed 48 hours.
//
// **NOTE** The Event Alert Query is being deprecated and replaced by the Event V2 Alert Query. For more information, see the [Event Migration guide](https://docs.datadoghq.com/events/guides/migrating_to_new_events_features/).
//
// **Event V2 Alert Query**
//
// Example: `events(query).rollup(rollup_method[, measure]).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `rollup_method` The stats roll-up method - supports `count`, `avg` and `cardinality`.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
//
// **Process Alert Query**
//
// Example: `processes(search).over(tags).rollup('count').last(timeframe) operator #`
//
// - `search` free text search string for querying processes.
// Matching processes match results on the [Live Processes](https://docs.datadoghq.com/infrastructure/process/?tab=linuxwindows) page.
// - `tags` one or more tags (comma-separated)
// - `timeframe` the timeframe to roll up the counts. Examples: 10m, 4h. Supported timeframes: s, m, h and d
// - `operator` <, <=, >, >=, ==, or !=
// - `#` an integer or decimal number used to set the threshold
//
// **Logs Alert Query**
//
// Example: `logs(query).index(index_name).rollup(rollup_method[, measure]).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `index_name` For multi-index organizations, the log index in which the request is performed.
// - `rollup_method` The stats roll-up method - supports `count`, `avg` and `cardinality`.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
//
// **Composite Query**
//
// Example: `12345 && 67890`, where `12345` and `67890` are the IDs of non-composite monitors
//
// * `name` [*required*, *default* = **dynamic, based on query**]: The name of the alert.
// * `message` [*required*, *default* = **dynamic, based on query**]: A message to include with notifications for this monitor.
// Email notifications can be sent to specific users by using the same '@username' notation as events.
// * `tags` [*optional*, *default* = **empty list**]: A list of tags to associate with your monitor.
// When getting all monitor details via the API, use the `monitor_tags` argument to filter results by these tags.
// It is only available via the API and isn't visible or editable in the Datadog UI.
//
// **SLO Alert Query**
//
// Example: `error_budget("slo_id").over("time_window") operator #`
//
// - `slo_id`: The alphanumeric SLO ID of the SLO you are configuring the alert for.
// - `time_window`: The time window of the SLO target you wish to alert on. Valid options: `7d`, `30d`, `90d`.
// - `operator`: `>=` or `>`
//
// **Audit Alert Query**
//
// Example: `audits(query).rollup(rollup_method[, measure]).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `rollup_method` The stats roll-up method - supports `count`, `avg` and `cardinality`.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
//
// **NOTE** Only available on US1-FED and in closed beta on US1, EU, US3, and US5.
//
// **CI Pipelines Alert Query**
//
// Example: `ci-pipelines(query).rollup(rollup_method[, measure]).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `rollup_method` The stats roll-up method - supports `count`, `avg`, and `cardinality`.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
//
// **NOTE** CI Pipeline monitors are in alpha on US1, EU, US3 and US5.
//
// **CI Tests Alert Query**
//
// Example: `ci-tests(query).rollup(rollup_method[, measure]).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `rollup_method` The stats roll-up method - supports `count`, `avg`, and `cardinality`.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
//
// **NOTE** CI Test monitors are available only in closed beta on US1, EU, US3 and US5.
//
// **Error Tracking Alert Query**
//
// Example(RUM): `error-tracking-rum(query).rollup(rollup_method[, measure]).last(time_window) operator #`
// Example(APM Traces): `error-tracking-traces(query).rollup(rollup_method[, measure]).last(time_window) operator #`
//
// - `query` The search query - following the [Log search syntax](https://docs.datadoghq.com/logs/search_syntax/).
// - `rollup_method` The stats roll-up method - supports `count`, `avg`, and `cardinality`.
// - `measure` For `avg` and cardinality `rollup_method` - specify the measure or the facet name you want to use.
// - `time_window` #m (between 1 and 2880), #h (between 1 and 48).
// - `operator` `<`, `<=`, `>`, `>=`, `==`, or `!=`.
// - `#` an integer or decimal number used to set the threshold.
func (a *MonitorsApi) CreateMonitor(ctx _context.Context, body Monitor) (Monitor, *_nethttp.Response, error) {
	req, err := a.buildCreateMonitorRequest(ctx, body)
	if err != nil {
		var localVarReturnValue Monitor
		return localVarReturnValue, nil, err
	}

	return a.createMonitorExecute(req)
}

// createMonitorExecute executes the request.
func (a *MonitorsApi) createMonitorExecute(r apiCreateMonitorRequest) (Monitor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue Monitor
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.MonitorsApi.CreateMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, datadog.ReportError("body is required and must be specified")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = r.body
	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiDeleteMonitorRequest struct {
	ctx       _context.Context
	monitorId int64
	force     *string
}

// DeleteMonitorOptionalParameters holds optional parameters for DeleteMonitor.
type DeleteMonitorOptionalParameters struct {
	Force *string
}

// NewDeleteMonitorOptionalParameters creates an empty struct for parameters.
func NewDeleteMonitorOptionalParameters() *DeleteMonitorOptionalParameters {
	this := DeleteMonitorOptionalParameters{}
	return &this
}

// WithForce sets the corresponding parameter name and returns the struct.
func (r *DeleteMonitorOptionalParameters) WithForce(force string) *DeleteMonitorOptionalParameters {
	r.Force = &force
	return r
}

func (a *MonitorsApi) buildDeleteMonitorRequest(ctx _context.Context, monitorId int64, o ...DeleteMonitorOptionalParameters) (apiDeleteMonitorRequest, error) {
	req := apiDeleteMonitorRequest{
		ctx:       ctx,
		monitorId: monitorId,
	}

	if len(o) > 1 {
		return req, datadog.ReportError("only one argument of type DeleteMonitorOptionalParameters is allowed")
	}

	if o != nil {
		req.force = o[0].Force
	}
	return req, nil
}

// DeleteMonitor Delete a monitor.
// Delete the specified monitor
func (a *MonitorsApi) DeleteMonitor(ctx _context.Context, monitorId int64, o ...DeleteMonitorOptionalParameters) (DeletedMonitor, *_nethttp.Response, error) {
	req, err := a.buildDeleteMonitorRequest(ctx, monitorId, o...)
	if err != nil {
		var localVarReturnValue DeletedMonitor
		return localVarReturnValue, nil, err
	}

	return a.deleteMonitorExecute(req)
}

// deleteMonitorExecute executes the request.
func (a *MonitorsApi) deleteMonitorExecute(r apiDeleteMonitorRequest) (DeletedMonitor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodDelete
		localVarPostBody    interface{}
		localVarReturnValue DeletedMonitor
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.MonitorsApi.DeleteMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/{monitor_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitor_id"+"}", _neturl.PathEscape(datadog.ParameterToString(r.monitorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.force != nil {
		localVarQueryParams.Add("force", datadog.ParameterToString(*r.force, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiGetMonitorRequest struct {
	ctx         _context.Context
	monitorId   int64
	groupStates *string
}

// GetMonitorOptionalParameters holds optional parameters for GetMonitor.
type GetMonitorOptionalParameters struct {
	GroupStates *string
}

// NewGetMonitorOptionalParameters creates an empty struct for parameters.
func NewGetMonitorOptionalParameters() *GetMonitorOptionalParameters {
	this := GetMonitorOptionalParameters{}
	return &this
}

// WithGroupStates sets the corresponding parameter name and returns the struct.
func (r *GetMonitorOptionalParameters) WithGroupStates(groupStates string) *GetMonitorOptionalParameters {
	r.GroupStates = &groupStates
	return r
}

func (a *MonitorsApi) buildGetMonitorRequest(ctx _context.Context, monitorId int64, o ...GetMonitorOptionalParameters) (apiGetMonitorRequest, error) {
	req := apiGetMonitorRequest{
		ctx:       ctx,
		monitorId: monitorId,
	}

	if len(o) > 1 {
		return req, datadog.ReportError("only one argument of type GetMonitorOptionalParameters is allowed")
	}

	if o != nil {
		req.groupStates = o[0].GroupStates
	}
	return req, nil
}

// GetMonitor Get a monitor's details.
// Get details about the specified monitor from your organization.
func (a *MonitorsApi) GetMonitor(ctx _context.Context, monitorId int64, o ...GetMonitorOptionalParameters) (Monitor, *_nethttp.Response, error) {
	req, err := a.buildGetMonitorRequest(ctx, monitorId, o...)
	if err != nil {
		var localVarReturnValue Monitor
		return localVarReturnValue, nil, err
	}

	return a.getMonitorExecute(req)
}

// getMonitorExecute executes the request.
func (a *MonitorsApi) getMonitorExecute(r apiGetMonitorRequest) (Monitor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue Monitor
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.MonitorsApi.GetMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/{monitor_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitor_id"+"}", _neturl.PathEscape(datadog.ParameterToString(r.monitorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.groupStates != nil {
		localVarQueryParams.Add("group_states", datadog.ParameterToString(*r.groupStates, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiListMonitorsRequest struct {
	ctx           _context.Context
	groupStates   *string
	name          *string
	tags          *string
	monitorTags   *string
	withDowntimes *bool
	idOffset      *int64
	page          *int64
	pageSize      *int32
}

// ListMonitorsOptionalParameters holds optional parameters for ListMonitors.
type ListMonitorsOptionalParameters struct {
	GroupStates   *string
	Name          *string
	Tags          *string
	MonitorTags   *string
	WithDowntimes *bool
	IdOffset      *int64
	Page          *int64
	PageSize      *int32
}

// NewListMonitorsOptionalParameters creates an empty struct for parameters.
func NewListMonitorsOptionalParameters() *ListMonitorsOptionalParameters {
	this := ListMonitorsOptionalParameters{}
	return &this
}

// WithGroupStates sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithGroupStates(groupStates string) *ListMonitorsOptionalParameters {
	r.GroupStates = &groupStates
	return r
}

// WithName sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithName(name string) *ListMonitorsOptionalParameters {
	r.Name = &name
	return r
}

// WithTags sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithTags(tags string) *ListMonitorsOptionalParameters {
	r.Tags = &tags
	return r
}

// WithMonitorTags sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithMonitorTags(monitorTags string) *ListMonitorsOptionalParameters {
	r.MonitorTags = &monitorTags
	return r
}

// WithWithDowntimes sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithWithDowntimes(withDowntimes bool) *ListMonitorsOptionalParameters {
	r.WithDowntimes = &withDowntimes
	return r
}

// WithIdOffset sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithIdOffset(idOffset int64) *ListMonitorsOptionalParameters {
	r.IdOffset = &idOffset
	return r
}

// WithPage sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithPage(page int64) *ListMonitorsOptionalParameters {
	r.Page = &page
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListMonitorsOptionalParameters) WithPageSize(pageSize int32) *ListMonitorsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

func (a *MonitorsApi) buildListMonitorsRequest(ctx _context.Context, o ...ListMonitorsOptionalParameters) (apiListMonitorsRequest, error) {
	req := apiListMonitorsRequest{
		ctx: ctx,
	}

	if len(o) > 1 {
		return req, datadog.ReportError("only one argument of type ListMonitorsOptionalParameters is allowed")
	}

	if o != nil {
		req.groupStates = o[0].GroupStates
		req.name = o[0].Name
		req.tags = o[0].Tags
		req.monitorTags = o[0].MonitorTags
		req.withDowntimes = o[0].WithDowntimes
		req.idOffset = o[0].IdOffset
		req.page = o[0].Page
		req.pageSize = o[0].PageSize
	}
	return req, nil
}

// ListMonitors Get all monitor details.
// Get details about the specified monitor from your organization.
func (a *MonitorsApi) ListMonitors(ctx _context.Context, o ...ListMonitorsOptionalParameters) ([]Monitor, *_nethttp.Response, error) {
	req, err := a.buildListMonitorsRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue []Monitor
		return localVarReturnValue, nil, err
	}

	return a.listMonitorsExecute(req)
}

// listMonitorsExecute executes the request.
func (a *MonitorsApi) listMonitorsExecute(r apiListMonitorsRequest) ([]Monitor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []Monitor
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.MonitorsApi.ListMonitors")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.groupStates != nil {
		localVarQueryParams.Add("group_states", datadog.ParameterToString(*r.groupStates, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", datadog.ParameterToString(*r.name, ""))
	}
	if r.tags != nil {
		localVarQueryParams.Add("tags", datadog.ParameterToString(*r.tags, ""))
	}
	if r.monitorTags != nil {
		localVarQueryParams.Add("monitor_tags", datadog.ParameterToString(*r.monitorTags, ""))
	}
	if r.withDowntimes != nil {
		localVarQueryParams.Add("with_downtimes", datadog.ParameterToString(*r.withDowntimes, ""))
	}
	if r.idOffset != nil {
		localVarQueryParams.Add("id_offset", datadog.ParameterToString(*r.idOffset, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", datadog.ParameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", datadog.ParameterToString(*r.pageSize, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiSearchMonitorGroupsRequest struct {
	ctx     _context.Context
	query   *string
	page    *int64
	perPage *int64
	sort    *string
}

// SearchMonitorGroupsOptionalParameters holds optional parameters for SearchMonitorGroups.
type SearchMonitorGroupsOptionalParameters struct {
	Query   *string
	Page    *int64
	PerPage *int64
	Sort    *string
}

// NewSearchMonitorGroupsOptionalParameters creates an empty struct for parameters.
func NewSearchMonitorGroupsOptionalParameters() *SearchMonitorGroupsOptionalParameters {
	this := SearchMonitorGroupsOptionalParameters{}
	return &this
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorGroupsOptionalParameters) WithQuery(query string) *SearchMonitorGroupsOptionalParameters {
	r.Query = &query
	return r
}

// WithPage sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorGroupsOptionalParameters) WithPage(page int64) *SearchMonitorGroupsOptionalParameters {
	r.Page = &page
	return r
}

// WithPerPage sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorGroupsOptionalParameters) WithPerPage(perPage int64) *SearchMonitorGroupsOptionalParameters {
	r.PerPage = &perPage
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorGroupsOptionalParameters) WithSort(sort string) *SearchMonitorGroupsOptionalParameters {
	r.Sort = &sort
	return r
}

func (a *MonitorsApi) buildSearchMonitorGroupsRequest(ctx _context.Context, o ...SearchMonitorGroupsOptionalParameters) (apiSearchMonitorGroupsRequest, error) {
	req := apiSearchMonitorGroupsRequest{
		ctx: ctx,
	}

	if len(o) > 1 {
		return req, datadog.ReportError("only one argument of type SearchMonitorGroupsOptionalParameters is allowed")
	}

	if o != nil {
		req.query = o[0].Query
		req.page = o[0].Page
		req.perPage = o[0].PerPage
		req.sort = o[0].Sort
	}
	return req, nil
}

// SearchMonitorGroups Monitors group search.
// Search and filter your monitor groups details.
func (a *MonitorsApi) SearchMonitorGroups(ctx _context.Context, o ...SearchMonitorGroupsOptionalParameters) (MonitorGroupSearchResponse, *_nethttp.Response, error) {
	req, err := a.buildSearchMonitorGroupsRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue MonitorGroupSearchResponse
		return localVarReturnValue, nil, err
	}

	return a.searchMonitorGroupsExecute(req)
}

// searchMonitorGroupsExecute executes the request.
func (a *MonitorsApi) searchMonitorGroupsExecute(r apiSearchMonitorGroupsRequest) (MonitorGroupSearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MonitorGroupSearchResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.MonitorsApi.SearchMonitorGroups")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/groups/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.query != nil {
		localVarQueryParams.Add("query", datadog.ParameterToString(*r.query, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", datadog.ParameterToString(*r.page, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", datadog.ParameterToString(*r.perPage, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", datadog.ParameterToString(*r.sort, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiSearchMonitorsRequest struct {
	ctx     _context.Context
	query   *string
	page    *int64
	perPage *int64
	sort    *string
}

// SearchMonitorsOptionalParameters holds optional parameters for SearchMonitors.
type SearchMonitorsOptionalParameters struct {
	Query   *string
	Page    *int64
	PerPage *int64
	Sort    *string
}

// NewSearchMonitorsOptionalParameters creates an empty struct for parameters.
func NewSearchMonitorsOptionalParameters() *SearchMonitorsOptionalParameters {
	this := SearchMonitorsOptionalParameters{}
	return &this
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorsOptionalParameters) WithQuery(query string) *SearchMonitorsOptionalParameters {
	r.Query = &query
	return r
}

// WithPage sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorsOptionalParameters) WithPage(page int64) *SearchMonitorsOptionalParameters {
	r.Page = &page
	return r
}

// WithPerPage sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorsOptionalParameters) WithPerPage(perPage int64) *SearchMonitorsOptionalParameters {
	r.PerPage = &perPage
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *SearchMonitorsOptionalParameters) WithSort(sort string) *SearchMonitorsOptionalParameters {
	r.Sort = &sort
	return r
}

func (a *MonitorsApi) buildSearchMonitorsRequest(ctx _context.Context, o ...SearchMonitorsOptionalParameters) (apiSearchMonitorsRequest, error) {
	req := apiSearchMonitorsRequest{
		ctx: ctx,
	}

	if len(o) > 1 {
		return req, datadog.ReportError("only one argument of type SearchMonitorsOptionalParameters is allowed")
	}

	if o != nil {
		req.query = o[0].Query
		req.page = o[0].Page
		req.perPage = o[0].PerPage
		req.sort = o[0].Sort
	}
	return req, nil
}

// SearchMonitors Monitors search.
// Search and filter your monitors details.
func (a *MonitorsApi) SearchMonitors(ctx _context.Context, o ...SearchMonitorsOptionalParameters) (MonitorSearchResponse, *_nethttp.Response, error) {
	req, err := a.buildSearchMonitorsRequest(ctx, o...)
	if err != nil {
		var localVarReturnValue MonitorSearchResponse
		return localVarReturnValue, nil, err
	}

	return a.searchMonitorsExecute(req)
}

// searchMonitorsExecute executes the request.
func (a *MonitorsApi) searchMonitorsExecute(r apiSearchMonitorsRequest) (MonitorSearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MonitorSearchResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.MonitorsApi.SearchMonitors")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.query != nil {
		localVarQueryParams.Add("query", datadog.ParameterToString(*r.query, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", datadog.ParameterToString(*r.page, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", datadog.ParameterToString(*r.perPage, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", datadog.ParameterToString(*r.sort, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiUpdateMonitorRequest struct {
	ctx       _context.Context
	monitorId int64
	body      *MonitorUpdateRequest
}

func (a *MonitorsApi) buildUpdateMonitorRequest(ctx _context.Context, monitorId int64, body MonitorUpdateRequest) (apiUpdateMonitorRequest, error) {
	req := apiUpdateMonitorRequest{
		ctx:       ctx,
		monitorId: monitorId,
		body:      &body,
	}
	return req, nil
}

// UpdateMonitor Edit a monitor.
// Edit the specified monitor.
func (a *MonitorsApi) UpdateMonitor(ctx _context.Context, monitorId int64, body MonitorUpdateRequest) (Monitor, *_nethttp.Response, error) {
	req, err := a.buildUpdateMonitorRequest(ctx, monitorId, body)
	if err != nil {
		var localVarReturnValue Monitor
		return localVarReturnValue, nil, err
	}

	return a.updateMonitorExecute(req)
}

// updateMonitorExecute executes the request.
func (a *MonitorsApi) updateMonitorExecute(r apiUpdateMonitorRequest) (Monitor, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPut
		localVarPostBody    interface{}
		localVarReturnValue Monitor
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.MonitorsApi.UpdateMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/{monitor_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitor_id"+"}", _neturl.PathEscape(datadog.ParameterToString(r.monitorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, datadog.ReportError("body is required and must be specified")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = r.body
	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiValidateExistingMonitorRequest struct {
	ctx       _context.Context
	monitorId int64
	body      *Monitor
}

func (a *MonitorsApi) buildValidateExistingMonitorRequest(ctx _context.Context, monitorId int64, body Monitor) (apiValidateExistingMonitorRequest, error) {
	req := apiValidateExistingMonitorRequest{
		ctx:       ctx,
		monitorId: monitorId,
		body:      &body,
	}
	return req, nil
}

// ValidateExistingMonitor Validate an existing monitor.
// Validate the monitor provided in the request.
func (a *MonitorsApi) ValidateExistingMonitor(ctx _context.Context, monitorId int64, body Monitor) (interface{}, *_nethttp.Response, error) {
	req, err := a.buildValidateExistingMonitorRequest(ctx, monitorId, body)
	if err != nil {
		var localVarReturnValue interface{}
		return localVarReturnValue, nil, err
	}

	return a.validateExistingMonitorExecute(req)
}

// validateExistingMonitorExecute executes the request.
func (a *MonitorsApi) validateExistingMonitorExecute(r apiValidateExistingMonitorRequest) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.MonitorsApi.ValidateExistingMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/{monitor_id}/validate"
	localVarPath = strings.Replace(localVarPath, "{"+"monitor_id"+"}", _neturl.PathEscape(datadog.ParameterToString(r.monitorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, datadog.ReportError("body is required and must be specified")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = r.body
	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiValidateMonitorRequest struct {
	ctx  _context.Context
	body *Monitor
}

func (a *MonitorsApi) buildValidateMonitorRequest(ctx _context.Context, body Monitor) (apiValidateMonitorRequest, error) {
	req := apiValidateMonitorRequest{
		ctx:  ctx,
		body: &body,
	}
	return req, nil
}

// ValidateMonitor Validate a monitor.
// Validate the monitor provided in the request.
func (a *MonitorsApi) ValidateMonitor(ctx _context.Context, body Monitor) (interface{}, *_nethttp.Response, error) {
	req, err := a.buildValidateMonitorRequest(ctx, body)
	if err != nil {
		var localVarReturnValue interface{}
		return localVarReturnValue, nil, err
	}

	return a.validateMonitorExecute(req)
}

// validateMonitorExecute executes the request.
func (a *MonitorsApi) validateMonitorExecute(r apiValidateMonitorRequest) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.MonitorsApi.ValidateMonitor")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/monitor/validate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, datadog.ReportError("body is required and must be specified")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = r.body
	datadog.SetAuthKeys(
		r.ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewMonitorsApi Returns NewMonitorsApi.
func NewMonitorsApi(client *datadog.APIClient) *MonitorsApi {
	return &MonitorsApi{
		Client: client,
	}
}
