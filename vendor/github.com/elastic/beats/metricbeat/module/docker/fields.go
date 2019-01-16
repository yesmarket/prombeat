// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package docker

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "docker", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
