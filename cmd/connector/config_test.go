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
	"fmt"
	"os"
	"reflect"
	"testing"
)

func Test_environmentVariableMapperWithDefaultSupport(t *testing.T) {
	_ = os.Setenv("NAME", "gopher")
	_ = os.Setenv("BURROW", "/usr/gopher")

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))
	tests := []struct {
		placeholderName string
		want            string
	}{
		{placeholderName: "NAME", want: "gopher"},
		{placeholderName: "BURROW", want: "/usr/gopher"},
		{placeholderName: "DEFAULT:default", want: "default"},
		{placeholderName: "ULR:http://default.com", want: "http://default.com"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := environmentVariableMapperWithDefaultSupport(tt.placeholderName); got != tt.want {
				t.Errorf("environmentVariableMapperWithDefaultSupport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadConfig(t *testing.T) {
	_ = os.Setenv("inventory_authorization_header", "ah")
	tests := []struct {
		configFile                       string
		wantWebHookID                    string
		wantInventoryAuthorizationHeader string
		wantErr                          bool
	}{
		{configFile: "./testdata/nofile", wantErr: true},
		{configFile: "./testdata/powerdns.json", wantInventoryAuthorizationHeader: "ah", wantWebHookID: "52acd668-3171-45a3-b23a-05adc76dc809", wantErr: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			got, err := loadConfig(tt.configFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got == nil {
				t.Errorf("loadConfig() got = nil")
				return
			}
			if !reflect.DeepEqual(got.WebHookID, tt.wantWebHookID) {
				t.Errorf("loadConfig() got = %s, want %s", got.WebHookID, tt.wantWebHookID)
			}
			if !reflect.DeepEqual(got.InventoryAuthorizationHeader, tt.wantInventoryAuthorizationHeader) {
				t.Errorf("loadConfig() got = %s, want %s", got.InventoryAuthorizationHeader, tt.wantInventoryAuthorizationHeader)
			}
		})
	}
}
