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

// DO NOT EDIT. This is an auto-generated file registering the REST handlers.
// for the various Showcase services.

package genrest

import (
	"fmt"
	"net/http"

	"core.wcloud.io/server/services"
	"github.com/googleapis/gapic-showcase/util/genrest/resttools"

	gmux "github.com/gorilla/mux"
	"google.golang.org/grpc/status"
)

type RESTBackend services.Backend

func RegisterHandlers(router *gmux.Router, backend *services.Backend) {
	rest := (*RESTBackend)(backend)
	router.HandleFunc("/v1/servers", rest.HandleCreateServer).Methods("POST")
	router.HandleFunc("/v1/{name:servers/[^:]+}", rest.HandleGetServer).Methods("GET")
	router.HandleFunc("/v1/servers", rest.HandleListservers).Methods("GET")
	router.HandleFunc("/v1/{name:servers/[^:]+}", rest.HandleDeleteServer).Methods("DELETE")
	router.HandleFunc("/v1/{name:servers/[^:]+}:merge", rest.HandleMergeServers).Methods("POST")
	router.HandleFunc("/v1/{parent:servers/[^:]+}/disks", rest.HandleCreateDisk).Methods("POST")
	router.HandleFunc("/v1/{name:servers/[^:]+/disks/[^:]+}", rest.HandleGetDisk).Methods("GET")
	router.HandleFunc("/v1/{parent:servers/[^:]+}/disks", rest.HandleListDisks).Methods("GET")
	router.HandleFunc("/v1/{name:servers/[^:]+/disks/[^:]+}", rest.HandleDeleteDisk).Methods("DELETE")
	router.HandleFunc("/v1/{book.name:servers/[^:]+/disks/[^:]+}", rest.HandleUpdateDisk).Methods("PATCH")
	router.HandleFunc("/v1/{book.name:servers/[^:]+/disks/[^:]+}", rest.HandleUpdateDisk).HeadersRegexp("X-HTTP-Method-Override", "^PATCH$").Methods("POST")
	router.HandleFunc("/v1/{name:servers/[^:]+/disks/[^:]+}:move", rest.HandleMoveDisk).Methods("POST")
	router.HandleFunc("/v1beta1/operations", rest.HandleListOperations).Methods("GET")
	router.HandleFunc("/v1beta1/{name:operations/[^:]+}", rest.HandleGetOperation).Methods("GET")
	router.HandleFunc("/v1beta1/{name:operations/[^:]+}", rest.HandleDeleteOperation).Methods("DELETE")
	router.HandleFunc("/v1beta1/{name:operations/[^:]+}:cancel", rest.HandleCancelOperation).Methods("POST")
	router.PathPrefix("/").HandlerFunc(rest.catchAllHandler)
}

func (backend *RESTBackend) catchAllHandler(w http.ResponseWriter, r *http.Request) {
	backend.Error(w, http.StatusBadRequest, "unrecognized request: %s %q", r.Method, r.URL)
}

func (backend *RESTBackend) Error(w http.ResponseWriter, httpStatus int, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	backend.ErrLog.Print(message)
	resttools.ErrorResponse(w, httpStatus, message)
}
func (backend *RESTBackend) ReportGRPCError(w http.ResponseWriter, err error) {
	st, ok := status.FromError(err)
	if !ok {
		backend.Error(w, http.StatusInternalServerError, "server error: %s", err.Error())
		return
	}

	backend.ErrLog.Print(st.Message())
	code := resttools.GRPCToHTTP(st.Code())
	resttools.ErrorResponse(w, code, st.Message(), st.Details()...)
}
