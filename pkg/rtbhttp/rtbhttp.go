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

package rtbhttp

import (
	"encoding/json"
	"net/http"
)

const (
	//POST method definition
	POST = "POST"
)

// Message represents the json result message used in ctrld
type Message struct {
	Message string `json:"message"`
}

// WriteMessage writes the particular message to the response
func WriteMessage(w http.ResponseWriter, statuscode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	jsonEncoder := json.NewEncoder(w)
	_ = jsonEncoder.Encode(&Message{Message: message})
}

// WriteAsJSON write interface as data
func WriteAsJSON(w http.ResponseWriter, statuscode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	jsonEncoder := json.NewEncoder(w)
	_ = jsonEncoder.Encode(data)
}

// ReadJSON write interface as data
func ReadJSON(req *http.Request, data interface{}) error {
	decoder := json.NewDecoder(req.Body)
	return decoder.Decode(data)
}
