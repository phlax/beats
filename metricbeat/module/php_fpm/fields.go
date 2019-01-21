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

package php_fpm

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "php_fpm", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
