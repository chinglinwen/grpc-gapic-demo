// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// DO NOT EDIT. This is an auto-generated file containing the REST handlers
// for service #0: "ServerService" (.core.wcloud.io.v1.ServerService).

package genrest

import (
	gmux "github.com/gorilla/mux"
	serverpb "core.wcloud.io/generated/grpcgen"
	"bytes"
	"io"
	"context"
	"net/http"
	"github.com/googleapis/gapic-showcase/util/genrest/resttools"
)

// HandleCreateServer translates REST requests/responses on the wire to internal proto messages for CreateServer
//
//	Generated for HTTP binding pattern: POST "/v1/servers"
func (backend *RESTBackend) HandleCreateServer(w http.ResponseWriter, r *http.Request) {
	urlPathParams := gmux.Vars(r)
	numUrlPathParams := len(urlPathParams)

	backend.StdLog.Printf("Received %s request matching '/v1/servers': %q", r.Method, r.URL)
	backend.StdLog.Printf("  urlPathParams (expect 0, have %d): %q", numUrlPathParams, urlPathParams)
	backend.StdLog.Printf("  urlRequestHeaders:\n%s", resttools.PrettyPrintHeaders(r, "    "))

	resttools.IncludeRequestHeadersInResponse(w, r)

	if numUrlPathParams != 0 {
		backend.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numUrlPathParams, urlPathParams)
		return
	}

	systemParameters, queryParams, err := resttools.GetSystemParameters(r)
	if err != nil {
		backend.Error(w, http.StatusBadRequest, "error in query string: %s", err)
		return
	}

	request := &serverpb.CreateServerRequest{}
	// Intentional: Field values in the URL path override those set in the body.
	var bodyField serverpb.Server
	var jsonReader bytes.Buffer
	bodyReader := io.TeeReader(r.Body, &jsonReader)
	rBytes := make([]byte, r.ContentLength)
	if _, err := bodyReader.Read(rBytes); err != nil && err != io.EOF {
		backend.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
		return
	}

	if err := resttools.FromJSON().Unmarshal(rBytes, &bodyField); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading body into request field 'server': %s", err)
		return
	}

	if err := resttools.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
		backend.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
		return
	}
	request.Server = &bodyField

	if err := resttools.PopulateSingularFields(request, urlPathParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
		return
	}

	// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
	excludedQueryParams := []string{"server"}
	if duplicates := resttools.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
		backend.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
		return
	}
	if err := resttools.PopulateFields(request, queryParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
		return
	}

	marshaler := resttools.ToJSON()
	marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
	requestJSON, _ := marshaler.Marshal(request)
	backend.StdLog.Printf("  request: %s", requestJSON)

	ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/servers")
	response, err := backend.ServerServiceServer.CreateServer(ctx, request)
	if err != nil {
		backend.ReportGRPCError(w, err)
		return
	}

	json, err := marshaler.Marshal(response)
	if err != nil {
		backend.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
		return
	}

	w.Write(json)
}

// HandleGetServer translates REST requests/responses on the wire to internal proto messages for GetServer
//
//	Generated for HTTP binding pattern: GET "/v1/{name=servers/*}"
func (backend *RESTBackend) HandleGetServer(w http.ResponseWriter, r *http.Request) {
	urlPathParams := gmux.Vars(r)
	numUrlPathParams := len(urlPathParams)

	backend.StdLog.Printf("Received %s request matching '/v1/{name=servers/*}': %q", r.Method, r.URL)
	backend.StdLog.Printf("  urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams)
	backend.StdLog.Printf("  urlRequestHeaders:\n%s", resttools.PrettyPrintHeaders(r, "    "))

	resttools.IncludeRequestHeadersInResponse(w, r)

	if numUrlPathParams != 1 {
		backend.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
		return
	}

	systemParameters, queryParams, err := resttools.GetSystemParameters(r)
	if err != nil {
		backend.Error(w, http.StatusBadRequest, "error in query string: %s", err)
		return
	}

	request := &serverpb.GetServerRequest{}
	if err := resttools.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
		backend.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
		return
	}
	if err := resttools.PopulateSingularFields(request, urlPathParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
		return
	}

	// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
	excludedQueryParams := []string{"name"}
	if duplicates := resttools.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
		backend.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
		return
	}
	if err := resttools.PopulateFields(request, queryParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
		return
	}

	marshaler := resttools.ToJSON()
	marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
	requestJSON, _ := marshaler.Marshal(request)
	backend.StdLog.Printf("  request: %s", requestJSON)

	ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/{name=servers/*}")
	response, err := backend.ServerServiceServer.GetServer(ctx, request)
	if err != nil {
		backend.ReportGRPCError(w, err)
		return
	}

	json, err := marshaler.Marshal(response)
	if err != nil {
		backend.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
		return
	}

	w.Write(json)
}

// HandleListservers translates REST requests/responses on the wire to internal proto messages for Listservers
//
//	Generated for HTTP binding pattern: GET "/v1/servers"
func (backend *RESTBackend) HandleListservers(w http.ResponseWriter, r *http.Request) {
	urlPathParams := gmux.Vars(r)
	numUrlPathParams := len(urlPathParams)

	backend.StdLog.Printf("Received %s request matching '/v1/servers': %q", r.Method, r.URL)
	backend.StdLog.Printf("  urlPathParams (expect 0, have %d): %q", numUrlPathParams, urlPathParams)
	backend.StdLog.Printf("  urlRequestHeaders:\n%s", resttools.PrettyPrintHeaders(r, "    "))

	resttools.IncludeRequestHeadersInResponse(w, r)

	if numUrlPathParams != 0 {
		backend.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 0, have %d: %#v", numUrlPathParams, urlPathParams)
		return
	}

	systemParameters, queryParams, err := resttools.GetSystemParameters(r)
	if err != nil {
		backend.Error(w, http.StatusBadRequest, "error in query string: %s", err)
		return
	}

	request := &serverpb.ListServersRequest{}
	if err := resttools.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
		backend.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
		return
	}
	if err := resttools.PopulateSingularFields(request, urlPathParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
		return
	}

	// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
	if err := resttools.PopulateFields(request, queryParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
		return
	}

	marshaler := resttools.ToJSON()
	marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
	requestJSON, _ := marshaler.Marshal(request)
	backend.StdLog.Printf("  request: %s", requestJSON)

	ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/servers")
	response, err := backend.ServerServiceServer.Listservers(ctx, request)
	if err != nil {
		backend.ReportGRPCError(w, err)
		return
	}

	json, err := marshaler.Marshal(response)
	if err != nil {
		backend.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
		return
	}

	w.Write(json)
}

// HandleDeleteServer translates REST requests/responses on the wire to internal proto messages for DeleteServer
//
//	Generated for HTTP binding pattern: DELETE "/v1/{name=servers/*}"
func (backend *RESTBackend) HandleDeleteServer(w http.ResponseWriter, r *http.Request) {
	urlPathParams := gmux.Vars(r)
	numUrlPathParams := len(urlPathParams)

	backend.StdLog.Printf("Received %s request matching '/v1/{name=servers/*}': %q", r.Method, r.URL)
	backend.StdLog.Printf("  urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams)
	backend.StdLog.Printf("  urlRequestHeaders:\n%s", resttools.PrettyPrintHeaders(r, "    "))

	resttools.IncludeRequestHeadersInResponse(w, r)

	if numUrlPathParams != 1 {
		backend.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
		return
	}

	systemParameters, queryParams, err := resttools.GetSystemParameters(r)
	if err != nil {
		backend.Error(w, http.StatusBadRequest, "error in query string: %s", err)
		return
	}

	request := &serverpb.DeleteServerRequest{}
	if err := resttools.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
		backend.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
		return
	}
	if err := resttools.PopulateSingularFields(request, urlPathParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
		return
	}

	// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
	excludedQueryParams := []string{"name"}
	if duplicates := resttools.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
		backend.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
		return
	}
	if err := resttools.PopulateFields(request, queryParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
		return
	}

	marshaler := resttools.ToJSON()
	marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
	requestJSON, _ := marshaler.Marshal(request)
	backend.StdLog.Printf("  request: %s", requestJSON)

	ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/{name=servers/*}")
	response, err := backend.ServerServiceServer.DeleteServer(ctx, request)
	if err != nil {
		backend.ReportGRPCError(w, err)
		return
	}

	json, err := marshaler.Marshal(response)
	if err != nil {
		backend.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
		return
	}

	w.Write(json)
}

// HandleMergeServers translates REST requests/responses on the wire to internal proto messages for MergeServers
//
//	Generated for HTTP binding pattern: POST "/v1/{name=servers/*}:merge"
func (backend *RESTBackend) HandleMergeServers(w http.ResponseWriter, r *http.Request) {
	urlPathParams := gmux.Vars(r)
	numUrlPathParams := len(urlPathParams)

	backend.StdLog.Printf("Received %s request matching '/v1/{name=servers/*}:merge': %q", r.Method, r.URL)
	backend.StdLog.Printf("  urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams)
	backend.StdLog.Printf("  urlRequestHeaders:\n%s", resttools.PrettyPrintHeaders(r, "    "))

	resttools.IncludeRequestHeadersInResponse(w, r)

	if numUrlPathParams != 1 {
		backend.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
		return
	}

	systemParameters, queryParams, err := resttools.GetSystemParameters(r)
	if err != nil {
		backend.Error(w, http.StatusBadRequest, "error in query string: %s", err)
		return
	}

	request := &serverpb.MergeserversRequest{}
	// Intentional: Field values in the URL path override those set in the body.
	var jsonReader bytes.Buffer
	bodyReader := io.TeeReader(r.Body, &jsonReader)
	rBytes := make([]byte, r.ContentLength)
	if _, err := bodyReader.Read(rBytes); err != nil && err != io.EOF {
		backend.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
		return
	}

	if err := resttools.FromJSON().Unmarshal(rBytes, request); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading body params '*': %s", err)
		return
	}

	if err := resttools.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
		backend.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
		return
	}

	if len(queryParams) > 0 {
		backend.Error(w, http.StatusBadRequest, "encountered unexpected query params: %v", queryParams)
		return
	}
	if err := resttools.PopulateSingularFields(request, urlPathParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
		return
	}

	marshaler := resttools.ToJSON()
	marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
	requestJSON, _ := marshaler.Marshal(request)
	backend.StdLog.Printf("  request: %s", requestJSON)

	ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/{name=servers/*}:merge")
	response, err := backend.ServerServiceServer.MergeServers(ctx, request)
	if err != nil {
		backend.ReportGRPCError(w, err)
		return
	}

	json, err := marshaler.Marshal(response)
	if err != nil {
		backend.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
		return
	}

	w.Write(json)
}

// HandleCreateDisk translates REST requests/responses on the wire to internal proto messages for CreateDisk
//
//	Generated for HTTP binding pattern: POST "/v1/{parent=servers/*}/disks"
func (backend *RESTBackend) HandleCreateDisk(w http.ResponseWriter, r *http.Request) {
	urlPathParams := gmux.Vars(r)
	numUrlPathParams := len(urlPathParams)

	backend.StdLog.Printf("Received %s request matching '/v1/{parent=servers/*}/disks': %q", r.Method, r.URL)
	backend.StdLog.Printf("  urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams)
	backend.StdLog.Printf("  urlRequestHeaders:\n%s", resttools.PrettyPrintHeaders(r, "    "))

	resttools.IncludeRequestHeadersInResponse(w, r)

	if numUrlPathParams != 1 {
		backend.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
		return
	}

	systemParameters, queryParams, err := resttools.GetSystemParameters(r)
	if err != nil {
		backend.Error(w, http.StatusBadRequest, "error in query string: %s", err)
		return
	}

	request := &serverpb.CreateDiskRequest{}
	// Intentional: Field values in the URL path override those set in the body.
	var bodyField serverpb.Disk
	var jsonReader bytes.Buffer
	bodyReader := io.TeeReader(r.Body, &jsonReader)
	rBytes := make([]byte, r.ContentLength)
	if _, err := bodyReader.Read(rBytes); err != nil && err != io.EOF {
		backend.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
		return
	}

	if err := resttools.FromJSON().Unmarshal(rBytes, &bodyField); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading body into request field 'book': %s", err)
		return
	}

	if err := resttools.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
		backend.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
		return
	}
	request.Book = &bodyField

	if err := resttools.PopulateSingularFields(request, urlPathParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
		return
	}

	// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
	excludedQueryParams := []string{"book", "parent"}
	if duplicates := resttools.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
		backend.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
		return
	}
	if err := resttools.PopulateFields(request, queryParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
		return
	}

	marshaler := resttools.ToJSON()
	marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
	requestJSON, _ := marshaler.Marshal(request)
	backend.StdLog.Printf("  request: %s", requestJSON)

	ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/{parent=servers/*}/disks")
	response, err := backend.ServerServiceServer.CreateDisk(ctx, request)
	if err != nil {
		backend.ReportGRPCError(w, err)
		return
	}

	json, err := marshaler.Marshal(response)
	if err != nil {
		backend.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
		return
	}

	w.Write(json)
}

// HandleGetDisk translates REST requests/responses on the wire to internal proto messages for GetDisk
//
//	Generated for HTTP binding pattern: GET "/v1/{name=servers/*/disks/*}"
func (backend *RESTBackend) HandleGetDisk(w http.ResponseWriter, r *http.Request) {
	urlPathParams := gmux.Vars(r)
	numUrlPathParams := len(urlPathParams)

	backend.StdLog.Printf("Received %s request matching '/v1/{name=servers/*/disks/*}': %q", r.Method, r.URL)
	backend.StdLog.Printf("  urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams)
	backend.StdLog.Printf("  urlRequestHeaders:\n%s", resttools.PrettyPrintHeaders(r, "    "))

	resttools.IncludeRequestHeadersInResponse(w, r)

	if numUrlPathParams != 1 {
		backend.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
		return
	}

	systemParameters, queryParams, err := resttools.GetSystemParameters(r)
	if err != nil {
		backend.Error(w, http.StatusBadRequest, "error in query string: %s", err)
		return
	}

	request := &serverpb.GetDiskRequest{}
	if err := resttools.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
		backend.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
		return
	}
	if err := resttools.PopulateSingularFields(request, urlPathParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
		return
	}

	// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
	excludedQueryParams := []string{"name"}
	if duplicates := resttools.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
		backend.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
		return
	}
	if err := resttools.PopulateFields(request, queryParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
		return
	}

	marshaler := resttools.ToJSON()
	marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
	requestJSON, _ := marshaler.Marshal(request)
	backend.StdLog.Printf("  request: %s", requestJSON)

	ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/{name=servers/*/disks/*}")
	response, err := backend.ServerServiceServer.GetDisk(ctx, request)
	if err != nil {
		backend.ReportGRPCError(w, err)
		return
	}

	json, err := marshaler.Marshal(response)
	if err != nil {
		backend.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
		return
	}

	w.Write(json)
}

// HandleListDisks translates REST requests/responses on the wire to internal proto messages for ListDisks
//
//	Generated for HTTP binding pattern: GET "/v1/{parent=servers/*}/disks"
func (backend *RESTBackend) HandleListDisks(w http.ResponseWriter, r *http.Request) {
	urlPathParams := gmux.Vars(r)
	numUrlPathParams := len(urlPathParams)

	backend.StdLog.Printf("Received %s request matching '/v1/{parent=servers/*}/disks': %q", r.Method, r.URL)
	backend.StdLog.Printf("  urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams)
	backend.StdLog.Printf("  urlRequestHeaders:\n%s", resttools.PrettyPrintHeaders(r, "    "))

	resttools.IncludeRequestHeadersInResponse(w, r)

	if numUrlPathParams != 1 {
		backend.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
		return
	}

	systemParameters, queryParams, err := resttools.GetSystemParameters(r)
	if err != nil {
		backend.Error(w, http.StatusBadRequest, "error in query string: %s", err)
		return
	}

	request := &serverpb.ListDisksRequest{}
	if err := resttools.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
		backend.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
		return
	}
	if err := resttools.PopulateSingularFields(request, urlPathParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
		return
	}

	// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
	excludedQueryParams := []string{"parent"}
	if duplicates := resttools.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
		backend.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
		return
	}
	if err := resttools.PopulateFields(request, queryParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
		return
	}

	marshaler := resttools.ToJSON()
	marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
	requestJSON, _ := marshaler.Marshal(request)
	backend.StdLog.Printf("  request: %s", requestJSON)

	ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/{parent=servers/*}/disks")
	response, err := backend.ServerServiceServer.ListDisks(ctx, request)
	if err != nil {
		backend.ReportGRPCError(w, err)
		return
	}

	json, err := marshaler.Marshal(response)
	if err != nil {
		backend.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
		return
	}

	w.Write(json)
}

// HandleDeleteDisk translates REST requests/responses on the wire to internal proto messages for DeleteDisk
//
//	Generated for HTTP binding pattern: DELETE "/v1/{name=servers/*/disks/*}"
func (backend *RESTBackend) HandleDeleteDisk(w http.ResponseWriter, r *http.Request) {
	urlPathParams := gmux.Vars(r)
	numUrlPathParams := len(urlPathParams)

	backend.StdLog.Printf("Received %s request matching '/v1/{name=servers/*/disks/*}': %q", r.Method, r.URL)
	backend.StdLog.Printf("  urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams)
	backend.StdLog.Printf("  urlRequestHeaders:\n%s", resttools.PrettyPrintHeaders(r, "    "))

	resttools.IncludeRequestHeadersInResponse(w, r)

	if numUrlPathParams != 1 {
		backend.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
		return
	}

	systemParameters, queryParams, err := resttools.GetSystemParameters(r)
	if err != nil {
		backend.Error(w, http.StatusBadRequest, "error in query string: %s", err)
		return
	}

	request := &serverpb.DeleteDiskRequest{}
	if err := resttools.CheckRequestFormat(nil, r, request.ProtoReflect()); err != nil {
		backend.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
		return
	}
	if err := resttools.PopulateSingularFields(request, urlPathParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
		return
	}

	// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
	excludedQueryParams := []string{"name"}
	if duplicates := resttools.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
		backend.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
		return
	}
	if err := resttools.PopulateFields(request, queryParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
		return
	}

	marshaler := resttools.ToJSON()
	marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
	requestJSON, _ := marshaler.Marshal(request)
	backend.StdLog.Printf("  request: %s", requestJSON)

	ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/{name=servers/*/disks/*}")
	response, err := backend.ServerServiceServer.DeleteDisk(ctx, request)
	if err != nil {
		backend.ReportGRPCError(w, err)
		return
	}

	json, err := marshaler.Marshal(response)
	if err != nil {
		backend.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
		return
	}

	w.Write(json)
}

// HandleUpdateDisk translates REST requests/responses on the wire to internal proto messages for UpdateDisk
//
//	Generated for HTTP binding pattern: PATCH "/v1/{book.name=servers/*/disks/*}"
func (backend *RESTBackend) HandleUpdateDisk(w http.ResponseWriter, r *http.Request) {
	urlPathParams := gmux.Vars(r)
	numUrlPathParams := len(urlPathParams)

	backend.StdLog.Printf("Received %s request matching '/v1/{book.name=servers/*/disks/*}': %q", r.Method, r.URL)
	backend.StdLog.Printf("  urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams)
	backend.StdLog.Printf("  urlRequestHeaders:\n%s", resttools.PrettyPrintHeaders(r, "    "))

	resttools.IncludeRequestHeadersInResponse(w, r)

	if numUrlPathParams != 1 {
		backend.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
		return
	}

	systemParameters, queryParams, err := resttools.GetSystemParameters(r)
	if err != nil {
		backend.Error(w, http.StatusBadRequest, "error in query string: %s", err)
		return
	}

	request := &serverpb.UpdateDiskRequest{}
	// Intentional: Field values in the URL path override those set in the body.
	var bodyField serverpb.Disk
	var jsonReader bytes.Buffer
	bodyReader := io.TeeReader(r.Body, &jsonReader)
	rBytes := make([]byte, r.ContentLength)
	if _, err := bodyReader.Read(rBytes); err != nil && err != io.EOF {
		backend.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
		return
	}

	if err := resttools.FromJSON().Unmarshal(rBytes, &bodyField); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading body into request field 'book': %s", err)
		return
	}

	if err := resttools.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
		backend.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
		return
	}
	request.Book = &bodyField

	if err := resttools.PopulateSingularFields(request, urlPathParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
		return
	}

	// TODO: Decide whether query-param value or URL-path value takes precedence when a field appears in both
	excludedQueryParams := []string{"book", "book.name"}
	if duplicates := resttools.KeysMatchPath(queryParams, excludedQueryParams); len(duplicates) > 0 {
		backend.Error(w, http.StatusBadRequest, "(QueryParamsInvalidFieldError) found keys that should not appear in query params: %v", duplicates)
		return
	}
	if err := resttools.PopulateFields(request, queryParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading query params: %s", err)
		return
	}

	marshaler := resttools.ToJSON()
	marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
	requestJSON, _ := marshaler.Marshal(request)
	backend.StdLog.Printf("  request: %s", requestJSON)

	ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/{book.name=servers/*/disks/*}")
	response, err := backend.ServerServiceServer.UpdateDisk(ctx, request)
	if err != nil {
		backend.ReportGRPCError(w, err)
		return
	}

	json, err := marshaler.Marshal(response)
	if err != nil {
		backend.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
		return
	}

	w.Write(json)
}

// HandleMoveDisk translates REST requests/responses on the wire to internal proto messages for MoveDisk
//
//	Generated for HTTP binding pattern: POST "/v1/{name=servers/*/disks/*}:move"
func (backend *RESTBackend) HandleMoveDisk(w http.ResponseWriter, r *http.Request) {
	urlPathParams := gmux.Vars(r)
	numUrlPathParams := len(urlPathParams)

	backend.StdLog.Printf("Received %s request matching '/v1/{name=servers/*/disks/*}:move': %q", r.Method, r.URL)
	backend.StdLog.Printf("  urlPathParams (expect 1, have %d): %q", numUrlPathParams, urlPathParams)
	backend.StdLog.Printf("  urlRequestHeaders:\n%s", resttools.PrettyPrintHeaders(r, "    "))

	resttools.IncludeRequestHeadersInResponse(w, r)

	if numUrlPathParams != 1 {
		backend.Error(w, http.StatusBadRequest, "found unexpected number of URL variables: expected 1, have %d: %#v", numUrlPathParams, urlPathParams)
		return
	}

	systemParameters, queryParams, err := resttools.GetSystemParameters(r)
	if err != nil {
		backend.Error(w, http.StatusBadRequest, "error in query string: %s", err)
		return
	}

	request := &serverpb.MoveDiskRequest{}
	// Intentional: Field values in the URL path override those set in the body.
	var jsonReader bytes.Buffer
	bodyReader := io.TeeReader(r.Body, &jsonReader)
	rBytes := make([]byte, r.ContentLength)
	if _, err := bodyReader.Read(rBytes); err != nil && err != io.EOF {
		backend.Error(w, http.StatusBadRequest, "error reading body content: %s", err)
		return
	}

	if err := resttools.FromJSON().Unmarshal(rBytes, request); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading body params '*': %s", err)
		return
	}

	if err := resttools.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
		backend.Error(w, http.StatusBadRequest, "REST request failed format check: %s", err)
		return
	}

	if len(queryParams) > 0 {
		backend.Error(w, http.StatusBadRequest, "encountered unexpected query params: %v", queryParams)
		return
	}
	if err := resttools.PopulateSingularFields(request, urlPathParams); err != nil {
		backend.Error(w, http.StatusBadRequest, "error reading URL path params: %s", err)
		return
	}

	marshaler := resttools.ToJSON()
	marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
	requestJSON, _ := marshaler.Marshal(request)
	backend.StdLog.Printf("  request: %s", requestJSON)

	ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/{name=servers/*/disks/*}:move")
	response, err := backend.ServerServiceServer.MoveDisk(ctx, request)
	if err != nil {
		backend.ReportGRPCError(w, err)
		return
	}

	json, err := marshaler.Marshal(response)
	if err != nil {
		backend.Error(w, http.StatusInternalServerError, "error json-encoding response: %s", err.Error())
		return
	}

	w.Write(json)
}
