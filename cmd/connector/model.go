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

type inventoryDNSRecordEventRequest struct {
	//EventID     string                  `json:"event_id"`
	//EventName   string                  `json:"event_name"`
	//TopicName   string                  `json:"topic_name"`
	//DateCreated time.Time               `json:"date_created"`
	Message inventoryDNSRecordEvent `json:"message"`
}
type inventoryDNSRecordEvent struct {
	//GroupID string `json:"group_id"`
	//GroupName    string `json:"group_name"`
	//GroupType    string `json:"group_type"`
	//ElementID    string `json:"element_id"`
	//ElementName  string `json:"element_name"`
	//ElementRole  string `json:"element_role"`
	RecordSet inventoryDNSRecordSet `json:"dns_recordset"`
}
type inventoryDNSRecordSet struct {
	ZoneName     string                `json:"dns_zone_name"`
	Name         *string               `json:"dns_name"`
	WithDrawName *string               `json:"dns_withdrawn_name"`
	Type         string                `json:"dns_type"`
	TTL          int                   `json:"dns_ttl"`
	Records      []inventoryDNSRecords `json:"dns_records"`
	//DNSRecordsetID string `json:"dns_recordset_id"`
	//DNSZoneID      string `json:"dns_zone_id"`
}

type inventoryDNSRecords struct {
	Disabled bool   `json:"disabled"`
	SetPTR   bool   `json:"dns_setptr"`
	Value    string `json:"dns_value"`
}
type inventoryDNSZoneEventRequest struct {
	//EventID     string                  `json:"event_id"`
	//EventName   string                  `json:"event_name"`
	//TopicName   string                  `json:"topic_name"`
	//DateCreated time.Time               `json:"date_created"`
	Message inventoryDNSZone `json:"message"`
}
type inventoryDNSZone struct {
	ZoneName  string `json:"dns_zone_name"`
	DNSZoneID string `json:"dns_zone_id"`
	//DNSZoneDescription string `json:"dns_zone_description"`
}
