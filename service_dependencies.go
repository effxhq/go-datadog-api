/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import (
	"encoding/json"
	"strconv"
)

type DependencyMap struct {
	Calls []string `json:"calls"`
}

func (client *Client) GetServiceMap(env string, startTime *int, endTime *int) (*map[string]DependencyMap{}, error) {
	uri := "/v1/service_dependencies?" + env
	dependencies := &map[string]DependencyMap{}

	if startTime != nil && endTime != nil {
		uri += "&start=" + strconv.Itoa(*startTime) + "&end=" + strconv.Itoa(*endTime)
	}

	err := client.doJsonRequest("GET", uri, nil, dependencies); err != nil {
		return dependencies, err
	}
	return dependencies, nil
}
