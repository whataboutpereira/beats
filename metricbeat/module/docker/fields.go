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
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "docker", asset.ModuleFieldsPri, AssetDocker); err != nil {
		panic(err)
	}
}

// AssetDocker returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/docker.
func AssetDocker() string {
	return "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"
}
