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

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("prombeat", "./vendor/github.com/elastic/beats/metricbeat/_meta/fields.common.yml", asset.BeatFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJyUUktu2zAU3OsUg+yjA2hRoPAVcoEXchwS5Ucln6zq9gVpOqiN1kW04+ObD2f0ih88FpgcY04ToF4DF5xuZ8tqil/V57Tg2wQAp5xUfKoDhLNnsBVyER/kPRA+QUIAL0wKPVbWecJYW6bO8YokkQsitXhTqXPMdgvsl39Vbd+bY8chn6GOuGKgThQfTCyitP2ma88D2CwskOCljskq6paxdCcc/UeRq6yWjf8y285ftHrDPnF7J1ZZLt5wFmsLa32udsop0bQpxv6nrhjnE3EuOWJ33rgHO7u0IkOgUdoZb87XT9peGaIcSFnxTqyFtZW6O6bOY0XlngIhGwnheHhO62DwFv7cfKEdEV+H/CVxDX+G/J8nt4BtNlu8/WQzvoddjooecsaLzeZlnn4HAAD///yQ6RY="
}
