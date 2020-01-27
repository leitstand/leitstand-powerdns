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
	"log"
	"net/http"

	"github.com/leitstand/leitstand-powerdns/pkg/rtbhttp"

	"github.com/mittwald/go-powerdns/apis/zones"
)

// @Summary webhook listener for leitstand-inventory
// @Description listens to events from leitstand-inventory
// @Description **Characteristics:**
// @Description * Operation: **synchronous**
// @Accept  json
// @Produce  json
// @Param body body main.inventoryDNSRecordEventRequest true "body"
// @Success 204 "No Content"
// @Failure 422 {object} rtbhttp.Message
// @Failure 500 {object} rtbhttp.Message
// @Router /api/v1/events/{event_name} [POST]
func (app *application) rbmsEvent(res http.ResponseWriter, req *http.Request) {
	eventName, ok := validateAndGetConfigTypeFromPath(res, req)
	if !ok {
		return
	}
	switch eventName {
	case ElementDnsRecordSetModifiedEvent:
		app.handleElementDnsRecordSetChangedEvent(res, req)
		return
	case DnsZoneCreatedEvent:
		app.handleDnsZoneCreatedEvent(res, req)
		return
	case DnsZoneRemovedEvent:
		app.handleDnsZoneRemovedEvent(res, req)
		return
	}

}

func (app *application) handleDnsZoneCreatedEvent(res http.ResponseWriter, req *http.Request) {
	requestBody := &inventoryDNSZoneEventRequest{}
	err := rtbhttp.ReadJSON(req, requestBody)
	if err != nil {
		log.Printf("Error: %s\n", err)
		rtbhttp.WriteMessage(res, http.StatusBadRequest, fmt.Sprintf("error %v", err))
		return
	}
	dnsZone := requestBody.Message
	zone := zones.Zone{
		Name:        dnsZone.ZoneName,
		Nameservers: app.config.Nameservers,
	}
	_, err = app.powerdnsClient.Zones().CreateZone(req.Context(), app.config.PowerdnsServerID, zone)
	if err != nil {
		message := fmt.Sprintf("error %v", err)
		log.Printf("Error: %s\n", message)
		rtbhttp.WriteMessage(res, http.StatusInternalServerError, message)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}
func (app *application) handleDnsZoneRemovedEvent(res http.ResponseWriter, req *http.Request) {
	requestBody := &inventoryDNSZoneEventRequest{}
	err := rtbhttp.ReadJSON(req, requestBody)
	if err != nil {
		log.Printf("Error: %s\n", err)
		rtbhttp.WriteMessage(res, http.StatusBadRequest, fmt.Sprintf("error %v", err))
		return
	}
	dnsZone := requestBody.Message

	err = app.powerdnsClient.Zones().DeleteZone(req.Context(), app.config.PowerdnsServerID, dnsZone.ZoneName)
	if err != nil {
		message := fmt.Sprintf("error %v", err)
		log.Printf("Error: %s\n", message)
		rtbhttp.WriteMessage(res, http.StatusInternalServerError, message)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}

func (app *application) handleElementDnsRecordSetChangedEvent(res http.ResponseWriter, req *http.Request) {
	requestBody := &inventoryDNSRecordEventRequest{}
	err := rtbhttp.ReadJSON(req, requestBody)
	if err != nil {
		log.Printf("Error: %s\n", err)
		rtbhttp.WriteMessage(res, http.StatusBadRequest, fmt.Sprintf("error %v", err))
		return
	}
	recordSet := requestBody.Message.RecordSet
	if recordSet.WithDrawName != nil {
		err = app.powerdnsClient.Zones().RemoveRecordSetFromZone(req.Context(), app.config.PowerdnsServerID, recordSet.ZoneName, *recordSet.WithDrawName, recordSet.Type)
		if err != nil {
			message := fmt.Sprintf("error %v", err)
			log.Printf("Error: %s\n", message)
			rtbhttp.WriteMessage(res, http.StatusInternalServerError, message)
			return
		}
	}
	if recordSet.Name != nil {
		rrset := zones.ResourceRecordSet{
			Name: *recordSet.Name,
			Type: recordSet.Type,
			TTL:  recordSet.TTL,
		}
		for _, record := range recordSet.Records {
			rrset.Records = append(rrset.Records, zones.Record{
				Content:  record.Value,
				Disabled: record.Disabled,
				SetPTR:   record.SetPTR,
			})
		}
		err = app.powerdnsClient.Zones().AddRecordSetToZone(req.Context(), app.config.PowerdnsServerID, recordSet.ZoneName, rrset)
		if err != nil {
			message := fmt.Sprintf("error %v", err)
			log.Printf("Error: %s\n", message)
			rtbhttp.WriteMessage(res, http.StatusInternalServerError, message)
			return
		}
	}
	res.WriteHeader(http.StatusNoContent)
}
