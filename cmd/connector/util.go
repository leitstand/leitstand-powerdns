/*
 * Copyright 2020 RtBrick Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License.  You may obtain a copy
 * of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func validateAndGetVariableFromPath(w http.ResponseWriter, req *http.Request, variableName string) (string, bool) {
	vars := mux.Vars(req)
	name, set := vars[variableName]
	if !set {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error: %s: There is an misconfiguration in the path variables!\n", variableName)
		return name, false
	}
	return name, true
}
func validateAndGetConfigTypeFromPath(w http.ResponseWriter, req *http.Request) (string, bool) {
	return validateAndGetVariableFromPath(w, req, "event_name")
}
