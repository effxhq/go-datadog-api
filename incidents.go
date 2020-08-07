/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import (
	"fmt"
	"net/url"
	"strings"
)

type IncidentUser struct {
	Data struct {
		Attributes struct {
			Email string `json:"email,omitempty"`
			Name  string `json:"name,omitempty"`
		}
		ID   string `json:"id,omitempty"`
		Type string `json:"type,omitempty"`
	}
}

type Attributes struct {
	Commander        IncidentUser `json:"commander,omitempty"`
	CreatedBy        IncidentUser `json:"created_by,omitempty"`
	LastModifiedBy   IncidentUser `json:"last_modified_by,omitempty"`
	CreatedAt        string       `json:"created,omitempty"`
	ResolvedAt       string       `json:"resolved,omitempty"`
	ModifiedAt       string       `json:"modified,omitempty"`
	DetectedAt       string       `json:"detected,omitempty"`
	State            string       `json:"state,omitempty"`
	Title            string       `json:"title,omitempty"`
	Severity         string       `json:"severity,omitempty"`
	PostmortemId     string       `json:"postmortem_id,omitempty"`
	CustomerImpacted bool         `json:"customer_impacted,omitempty"`
}

type IncidentQueryOpts struct {
	Include    string
	PageSize   int64
	PageNumber int64
}

type Incident struct {
	Attributes Attributes `json:"attributes,omitempty"`
	ID         string     `json:"id,omitempty"`
}

type reqIncidents struct {
	Meta      Meta       `json:"meta,omitempty"`
	Incidents []Incident `json:"data,omitempty"`
}

type Meta struct {
	Pagination struct {
		Number     int `json:"number,omitempty"`
		NextNumber int `json:"next_number,omitempty"`
		Size       int `json:"size,omitempty"`
	}
}

func (client *Client) GetIncidentsWithOptions(opts IncidentQueryOpts) ([]Incident, error) {
	var out reqIncidents
	var query []string
	if len(opts.Include) > 0 {
		value := fmt.Sprintf("include=%s", opts.Include)
		query = append(query, value)
	}
	if opts.PageSize > 0 {
		value := fmt.Sprintf("page[size]=%d", opts.PageSize)
		query = append(query, value)
	}
	if opts.PageNumber > 0 {
		value := fmt.Sprintf("page[number]=%d", opts.PageNumber)
		query = append(query, value)
	}
	queryString, err := url.ParseQuery(strings.Join(query, "&"))
	if err != nil {
		return nil, err
	}
	err = client.doJsonRequest("GET", fmt.Sprintf("/v2/incidents?%v", queryString.Encode()), nil, &out)
	if err != nil {
		return nil, err
	}

	return out.Incidents, nil
}
